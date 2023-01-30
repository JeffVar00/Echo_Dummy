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

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%d/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port, // 1433, remove if necesary (the string conection doesnt work)
		s.DB.Name)

	return sqlx.ConnectContext(ctx, "mssql", connectionString)

	//	package
	//	go get github.com/denisenkom/go-mssqldb

}
