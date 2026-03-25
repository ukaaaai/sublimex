package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ukaaaai/sublimex/mcp-server/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewClient(ctx context.Context, cfg *config.Config) (*bun.DB, error) {
	pgCfg := cfg.PostgreSQL
	databaseUrl := fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s&sslmode=disable",
		pgCfg.Host, pgCfg.Port, pgCfg.Name, pgCfg.User, pgCfg.Password)

	sqldb, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())

	return db, nil
}
