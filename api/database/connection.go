package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
)

var Manager *Queries

func Connect() {
	ctx := context.Background()
	connection, err := pgx.Connect(ctx, fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s", os.Getenv("POSTGRES_USER"), os.Getenv(("POSTGRES_PASSWORD")), os.Getenv("POSTGRES_DB"), os.Getenv(("POSTGRES_HOST")), os.Getenv("POSTGRES_PORT")))
	if err != nil {
		panic(err)
	}

	Manager = New(connection)
}

func MockConnection() pgxmock.PgxConnIface {
	mock, err := pgxmock.NewConn()
	if err != nil {
		panic(err)
	}
	defer mock.Close(context.Background())

	Manager = New(mock)

	return mock
}
