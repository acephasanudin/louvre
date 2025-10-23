package usecase

import (
	"context"
	"example/service/internal/domain/example/model"
	"example/service/internal/domain/example/repository"

	"github.com/google/uuid"
)

// ExampleUseCase defines the business logic interface for examples
type ExampleUseCase interface {
	CreateExample(ctx context.Context, req *CreateExampleRequest) (*model.Example, error)
	GetExample(ctx context.Context, id uuid.UUID) (*model.Example, error)
	GetExamples(ctx context.Context, req *GetExamplesRequest) (*GetExamplesResponse, error)
	UpdateExample(ctx context.Context, req *UpdateExampleRequest) (*model.Example, error)
	DeleteExample(ctx context.Context, id uuid.UUID) error
}

// CreateExampleRequest represents the request to create an example
type CreateExampleRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"max=1000"`
}

// UpdateExampleRequest represents the request to update an example
type UpdateExampleRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        *string   `json:"name,omitempty" validate:"omitempty,min=1,max=255"`
	Description *string   `json:"description,omitempty" validate:"omitempty,max=1000"`
	Status      *string   `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}

// GetExamplesRequest represents the request to get examples
type GetExamplesRequest struct {
	Limit  int `json:"limit" validate:"min=1,max=100"`
	Offset int `json:"offset" validate:"min=0"`
}

// GetExamplesResponse represents the response for getting examples
type GetExamplesResponse struct {
	Examples []*model.Example `json:"examples"`
	Total    int              `json:"total"`
}

// exampleUseCase implements the ExampleUseCase interface
type exampleUseCase struct {
	repo repository.ExampleRepository
}

// NewExampleUseCase creates a new example use case
func NewExampleUseCase(repo repository.ExampleRepository) ExampleUseCase {
	return &exampleUseCase{
		repo: repo,
	}
}

// CreateExample creates a new example
func (uc *exampleUseCase) CreateExample(ctx context.Context, req *CreateExampleRequest) (*model.Example, error) {
	example := &model.Example{
		Name:        req.Name,
		Description: req.Description,
		Status:      "active",
	}

	if err := uc.repo.Create(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}

// GetExample retrieves an example by ID
func (uc *exampleUseCase) GetExample(ctx context.Context, id uuid.UUID) (*model.Example, error) {
	return uc.repo.GetByID(ctx, id)
}

// GetExamples retrieves a list of examples
func (uc *exampleUseCase) GetExamples(ctx context.Context, req *GetExamplesRequest) (*GetExamplesResponse, error) {
	examples, err := uc.repo.GetAll(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &GetExamplesResponse{
		Examples: examples,
		Total:    len(examples), // In a real implementation, you'd get the total count separately
	}, nil
}

// UpdateExample updates an existing example
func (uc *exampleUseCase) UpdateExample(ctx context.Context, req *UpdateExampleRequest) (*model.Example, error) {
	example, err := uc.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		example.Name = *req.Name
	}
	if req.Description != nil {
		example.Description = *req.Description
	}
	if req.Status != nil {
		example.Status = *req.Status
	}

	if err := uc.repo.Update(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}

// DeleteExample deletes an example by ID
func (uc *exampleUseCase) DeleteExample(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}