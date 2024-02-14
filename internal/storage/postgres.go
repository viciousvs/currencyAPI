package storage

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/viciousvs/currencyAPI/config"
)

var schema = `
create table if not exists currency(
    Id serial primary key,
    Date date not null,
    Time_stamp int,
    Base varchar(3),
    Rate varchar(3),
    Value numeric(20,16)
)
`

type PostgresDB struct {
	*sqlx.DB
}

func NewPostgresDB(cfg config.PostgresConfig) *PostgresDB {
	connUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password='%s' sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DatabaseName,
		cfg.Password)

	db, err := sqlx.Connect("postgres", connUrl)
	if err != nil {
		log.Fatalln("cannot connect to db ", err.Error())
	}
	db.MustExec(schema)
	fmt.Println("schema runned")
	return &PostgresDB{db}
}
