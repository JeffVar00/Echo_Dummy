package database

import (
	"Echo_Dummy/settings"
	"context"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	//   sqlserver://username:password@host/instance?param1=value1&param2=value2

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		s.DB.Host,
		s.DB.User,
		s.DB.Password,
		s.DB.Port,
		s.DB.Name)

	return sqlx.ConnectContext(ctx, "mssql", connectionString)

	//    package
	//    go get github.com/denisenkom/go-mssqldb

}
