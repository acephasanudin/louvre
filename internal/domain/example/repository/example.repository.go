package repository

import (
	"context"
	"example/service/internal/domain/example/model"

	"github.com/google/uuid"
)

// ExampleRepository defines the interface for example data operations
type ExampleRepository interface {
	Create(ctx context.Context, example *model.Example) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Example, error)
	GetAll(ctx context.Context, limit, offset int) ([]*model.Example, error)
	Update(ctx context.Context, example *model.Example) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByName(ctx context.Context, name string) (*model.Example, error)
}