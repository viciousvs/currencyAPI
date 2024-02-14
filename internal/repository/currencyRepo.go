package repository

import (
	"context"
	"fmt"

	"github.com/viciousvs/currencyAPI/internal/model"
	"github.com/viciousvs/currencyAPI/internal/storage"
)

// 43
type CurrencyRepo interface {
	GetLastRates(ctx context.Context, limit int) (model.Currencies, error)
	GetCurrencyByName(ctx context.Context, name string) (model.Currency, error)
	UpdateCurrencies(ctx context.Context, c model.Currencies) (int, error)
}

type repoPostgres struct {
	db *storage.PostgresDB
}

func NewCurrencyRepo(db *storage.PostgresDB) CurrencyRepo {
	return repoPostgres{db: db}
}

func (r repoPostgres) GetLastRates(ctx context.Context, limit int) (model.Currencies, error) {
	res := model.Currencies{Rates: make(map[string]float64)}
	stmt := fmt.Sprintf("select * from currency limit %d", limit)
	rows, err := r.db.QueryxContext(ctx, stmt)
	if err != nil {
		return model.Currencies{}, err
	}
	for rows.Next() {
		m := model.Currency{}
		rows.StructScan(&m)
		res.Date = m.Date
		res.Base = m.Base
		res.Timestamp = m.TimeStamp
		res.Rates[m.Rate] = m.Value
	}
	return res, nil
}
func (r repoPostgres) GetCurrencyByName(ctx context.Context, name string) (model.Currency, error) {
	res := model.Currency{}
	fmt.Println(name)
	err := r.db.Get(&res, "SELECT * FROM currency where rate=$1 order by id DESC LIMIT 1", name)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (r repoPostgres) UpdateCurrencies(ctx context.Context, c model.Currencies) (int, error) {
	stmt := "INSERT INTO currency (date, time_stamp, base, rate, value) VALUES ($1, $2, $3, $4, $5)"
	count := 0
	for key, val := range c.Rates {
		_, err := r.db.ExecContext(ctx, stmt, c.Date, c.Timestamp, c.Base, key, val)
		if err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}
