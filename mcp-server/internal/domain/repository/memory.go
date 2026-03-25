package repository

import (
	"context"

	"github.com/ukaaaai/sublimex/mcp-server/internal/domain/entity"
)

type IMemoryRepository interface {
	SaveMemory(ctx context.Context, text string) error

	GetMemory(ctx context.Context, id string) (*entity.Memory, error)

	SearchMemories(ctx context.Context, sentence string) ([]*entity.Memory, error)

	UpdateMemory(ctx context.Context, id string, text string) error

	DeleteMemory(ctx context.Context, id string) error
}
