package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"path/filepath"

	mv1 "github.com/llm-operator/model-manager/api/v1"
	"github.com/llm-operator/model-manager/loader/internal/config"
	"github.com/llm-operator/model-manager/loader/internal/loader"
	"github.com/llm-operator/model-manager/loader/internal/s3"
	"github.com/llm-operator/rbac-manager/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const flagConfig = "config"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString(flagConfig)
		if err != nil {
			return err
		}

		c, err := config.Parse(path)
		if err != nil {
			return err
		}

		if err := c.Validate(); err != nil {
			return err
		}

		if err := run(cmd.Context(), &c); err != nil {
			return err
		}
		return nil
	},
}

func run(ctx context.Context, c *config.Config) error {
	s3c := c.ObjectStore.S3
	d, err := newModelDownloader(c)
	if err != nil {
		return err
	}

	var mclient loader.ModelClient
	if c.Debug.Standalone {
		mclient = loader.NewFakeModelClient()
	} else {
		var option grpc.DialOption
		if c.Worker.TLS.Enable {
			option = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
		} else {
			option = grpc.WithTransportCredentials(insecure.NewCredentials())
		}

		conn, err := grpc.Dial(c.ModelManagerServerWorkerServiceAddr, option)
		if err != nil {
			return err
		}
		mc := mv1.NewModelsWorkerServiceClient(conn)
		if err := createStorageClass(ctx, mc, s3c.PathPrefix); err != nil {
			return err
		}
		mclient = mc
	}

	s := loader.New(
		c.BaseModels,
		filepath.Join(s3c.PathPrefix, s3c.BaseModelPathPrefix),
		d,
		newS3Client(c),
		mclient,
	)

	if c.RunOnce {
		return s.LoadBaseModels(ctx)
	}

	return s.Run(ctx, c.ModelLoadInterval)
}

func createStorageClass(ctx context.Context, mclient mv1.ModelsWorkerServiceClient, pathPrefix string) error {
	ctx = auth.AppendWorkerAuthorization(ctx)

	_, err := mclient.GetStorageConfig(ctx, &mv1.GetStorageConfigRequest{})
	if err == nil {
		return nil
	}

	if s, ok := status.FromError(err); ok && s.Code() != codes.NotFound {
		return err
	}

	log.Printf("Creating a storage class with path prefix %q", pathPrefix)
	_, err = mclient.CreateStorageConfig(ctx, &mv1.CreateStorageConfigRequest{
		PathPrefix: pathPrefix,
	})
	return err
}

func newModelDownloader(c *config.Config) (loader.ModelDownloader, error) {
	switch c.Downloader.Kind {
	case config.DownloaderKindS3:
		s3Client := s3.NewClient(s3.NewOptions{
			EndpointURL: c.Downloader.S3.EndpointURL,
			Region:      c.Downloader.S3.Region,
			Bucket:      c.Downloader.S3.Bucket,
			// Use anonymous credentials as the S3 bucket is public and we don't want to use the credential that is
			// used to upload the model.
			UseAnonymousCredentials: true,
		})
		return loader.NewS3Downloader(s3Client, c.Downloader.S3.PathPrefix), nil
	case config.DownloaderKindHuggingFace:
		return loader.NewHuggingFaceDownloader(c.Downloader.HuggingFace.CacheDir), nil
	default:
		return nil, fmt.Errorf("unknown downloader kind: %s", c.Downloader.Kind)
	}
}

func newS3Client(c *config.Config) loader.S3Client {
	return s3.NewClient(s3.NewOptions{
		EndpointURL: c.ObjectStore.S3.EndpointURL,
		Region:      c.ObjectStore.S3.Region,
		Bucket:      c.ObjectStore.S3.Bucket,
	})

}

func init() {
	runCmd.Flags().StringP(flagConfig, "c", "", "Configuration file path")
	_ = runCmd.MarkFlagRequired(flagConfig)
}
