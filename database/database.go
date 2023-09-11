package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hugolesta/go-inventory-api/internal/settings"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) { 
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=True",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)
	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
