package server

import (
	"context"
	"errors"
	"fmt"
	"strings"

	v1 "github.com/llm-operator/model-manager/api/v1"
	"github.com/llm-operator/model-manager/server/internal/store"
	"github.com/llmariner/common/pkg/id"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// ListModels lists models.
func (s *S) ListModels(
	ctx context.Context,
	req *v1.ListModelsRequest,
) (*v1.ListModelsResponse, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var modelProtos []*v1.Model
	// First include base models.
	bms, err := s.store.ListBaseModels(userInfo.TenantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list models: %s", err)
	}
	for _, m := range bms {
		modelProtos = append(modelProtos, baseToModelProto(m))
	}

	// Then add generated models owned by the project
	ms, err := s.store.ListModelsByProjectID(userInfo.ProjectID, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list models: %s", err)
	}

	for _, m := range ms {
		modelProtos = append(modelProtos, toModelProto(m))
	}

	return &v1.ListModelsResponse{
		Object: "list",
		Data:   modelProtos,
	}, nil
}

// GetModel gets a model.
func (s *S) GetModel(
	ctx context.Context,
	req *v1.GetModelRequest,
) (*v1.Model, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	// First check if it's a base model.
	bm, err := s.store.GetBaseModel(req.Id, userInfo.TenantID)
	if err == nil {
		return baseToModelProto(bm), nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.Internal, "get model: %s", err)
	}

	// Then check if it's a generated model.
	m, err := s.store.GetPublishedModelByModelIDAndProjectID(req.Id, userInfo.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "get model: %s", err)
	}
	return toModelProto(m), nil
}

// DeleteModel deletes a model.
func (s *S) DeleteModel(
	ctx context.Context,
	req *v1.DeleteModelRequest,
) (*v1.DeleteModelResponse, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	if _, err := s.store.GetBaseModel(req.Id, userInfo.TenantID); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "get model: %s", err)
		}
		// Do nothing
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "base model %q cannot be deleted", req.Id)
	}

	if err := s.store.DeleteModel(req.Id, userInfo.ProjectID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "delete model: %s", err)
	}
	return &v1.DeleteModelResponse{
		Id:      req.Id,
		Object:  "model",
		Deleted: true,
	}, nil
}

// ListBaseModels lists base models.
func (s *S) ListBaseModels(
	ctx context.Context,
	req *v1.ListBaseModelsRequest,
) (*v1.ListBaseModelsResponse, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	ms, err := s.store.ListBaseModels(userInfo.TenantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list base models: %s", err)
	}
	var modelProtos []*v1.BaseModel
	for _, m := range ms {
		modelProtos = append(modelProtos, toBaseModelProto(m))
	}
	return &v1.ListBaseModelsResponse{
		Object: "list",
		Data:   modelProtos,
	}, nil
}

// GetModel gets a model.
func (s *WS) GetModel(
	ctx context.Context,
	req *v1.GetModelRequest,
) (*v1.Model, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	// First check if it's a base model.
	bm, err := s.store.GetBaseModel(req.Id, clusterInfo.TenantID)
	if err == nil {
		return baseToModelProto(bm), nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.Internal, "get model: %s", err)
	}

	// Then check if it's a generated model.
	m, err := s.store.GetPublishedModelByModelIDAndTenantID(req.Id, clusterInfo.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "get model: %s", err)
	}
	return toModelProto(m), nil
}

// RegisterModel registers a model.
func (s *WS) RegisterModel(
	ctx context.Context,
	req *v1.RegisterModelRequest,
) (*v1.RegisterModelResponse, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.BaseModel == "" {
		return nil, status.Error(codes.InvalidArgument, "base_model is required")
	}
	if req.Suffix == "" &&
		req.Adapter == v1.AdapterType_ADAPTER_TYPE_UNSPECIFIED &&
		req.Quantization == v1.QuantizationType_QUANTIZATION_TYPE_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "one of suffix, adapter, or quantization is required")
	}
	if req.OrganizationId == "" {
		return nil, status.Error(codes.InvalidArgument, "organization_id is required")
	}
	if req.ProjectId == "" {
		return nil, status.Error(codes.InvalidArgument, "project_id is required")
	}

	id, err := s.genenerateModelID(req.BaseModel, req.Suffix)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "generate model ID: %s", err)
	}

	sc, err := s.store.GetStorageConfig(clusterInfo.TenantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get storage config: %s", err)
	}

	path := fmt.Sprintf("%s/%s/%s", sc.PathPrefix, clusterInfo.TenantID, id)
	_, err = s.store.CreateModel(store.ModelSpec{
		ModelID:        id,
		TenantID:       clusterInfo.TenantID,
		OrganizationID: req.OrganizationId,
		ProjectID:      req.ProjectId,
		IsPublished:    false,
		Path:           path,
		BaseModelID:    req.BaseModel,
		Adapter:        string(req.Adapter),
		Quantization:   string(req.Quantization),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create model: %s", err)
	}

	return &v1.RegisterModelResponse{
		Id:   id,
		Path: path,
	}, nil
}

// PublishModel publishes a model.
func (s *WS) PublishModel(
	ctx context.Context,
	req *v1.PublishModelRequest,
) (*v1.PublishModelResponse, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	if err := s.store.UpdateModel(req.Id, clusterInfo.TenantID, true); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "create model: %s", err)
	}
	return &v1.PublishModelResponse{}, nil
}

// GetModelPath gets a model path.
func (s *WS) GetModelPath(
	ctx context.Context,
	req *v1.GetModelPathRequest,
) (*v1.GetModelPathResponse, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	m, err := s.store.GetPublishedModelByModelIDAndTenantID(req.Id, clusterInfo.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "create model: %s", err)
	}
	return &v1.GetModelPathResponse{
		Path: m.Path,
	}, nil
}

// GetModelAttributes gets the model attributes.
func (s *WS) GetModelAttributes(
	ctx context.Context,
	req *v1.GetModelAttributesRequest,
) (*v1.ModelAttributes, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	m, err := s.store.GetPublishedModelByModelIDAndTenantID(req.Id, clusterInfo.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "create model: %s", err)
	}
	return &v1.ModelAttributes{
		Path:         m.Path,
		BaseModel:    m.BaseModelID,
		Adapter:      v1.AdapterType(v1.AdapterType_value[m.Adapter]),
		Quantization: v1.QuantizationType(v1.QuantizationType_value[m.Quantization]),
	}, nil
}

// CreateBaseModel creates a base model.
func (s *WS) CreateBaseModel(
	ctx context.Context,
	req *v1.CreateBaseModelRequest,
) (*v1.BaseModel, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.Path == "" {
		return nil, status.Error(codes.InvalidArgument, "path is required")
	}

	var formats []v1.ModelFormat
	if len(req.Formats) == 0 {
		// Fall back to GGUF for backward compatibility.
		// TODO(kenji): Make this to return an errror once all clients are updated.
		formats = append(formats, v1.ModelFormat_MODEL_FORMAT_GGUF)
	} else {
		formats = append(formats, req.Formats...)
	}

	for _, format := range formats {
		if format == v1.ModelFormat_MODEL_FORMAT_GGUF && req.GgufModelPath == "" {
			return nil, status.Error(codes.InvalidArgument, "gguf_model_path is required for the GGUF format")
		}
	}

	m, err := s.store.CreateBaseModel(req.Id, req.Path, formats, req.GgufModelPath, clusterInfo.TenantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create base model: %s", err)
	}

	return toBaseModelProto(m), nil
}

// GetBaseModelPath gets a model path.
func (s *WS) GetBaseModelPath(
	ctx context.Context,
	req *v1.GetBaseModelPathRequest,
) (*v1.GetBaseModelPathResponse, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	m, err := s.store.GetBaseModel(req.Id, clusterInfo.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "model %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "create model: %s", err)
	}

	formats, err := store.UnmarshalModelFormats(m.Formats)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unmarshal model formats: %s", err)
	}

	// Add GGUF as format if GGUF model path is set (for backward compatibility).
	if m.GGUFModelPath != "" && len(formats) == 0 {
		formats = append(formats, v1.ModelFormat_MODEL_FORMAT_GGUF)
	}

	return &v1.GetBaseModelPathResponse{
		Path:          m.Path,
		Formats:       formats,
		GgufModelPath: m.GGUFModelPath,
	}, nil
}

func (s *WS) genenerateModelID(baseModel, suffix string) (string, error) {
	const randomLength = 10
	// OpenAI uses ':" as a separator, but Ollama does not accept. Use '-' instead for now.
	// TODO(kenji): Revisit this.
	// Replace "/" with "-'. HuggingFace model contains "/", but that doesn't work for Ollama.
	base := fmt.Sprintf("ft:%s:%s-", strings.ReplaceAll(baseModel, "/", "-"), suffix)

	// Randomly create an ID and retry if it already exists.
	for {
		randomStr, err := id.GenerateID("", randomLength)
		if err != nil {
			return "", fmt.Errorf("generate ID: %s", err)
		}
		id := fmt.Sprintf("%s%s", base, randomStr)
		if _, err := s.store.GetModelByModelID(id); errors.Is(err, gorm.ErrRecordNotFound) {
			return id, nil
		}
	}
}

func toModelProto(m *store.Model) *v1.Model {
	return &v1.Model{
		Id:      m.ModelID,
		Object:  "model",
		Created: m.CreatedAt.UTC().Unix(),
		OwnedBy: "user",
	}
}

func baseToModelProto(m *store.BaseModel) *v1.Model {
	return &v1.Model{
		Id:      m.ModelID,
		Object:  "model",
		Created: m.CreatedAt.UTC().Unix(),
		OwnedBy: "system",
	}
}

func toBaseModelProto(m *store.BaseModel) *v1.BaseModel {
	return &v1.BaseModel{
		Id:      m.ModelID,
		Object:  "basemodel",
		Created: m.CreatedAt.UTC().Unix(),
	}
}
