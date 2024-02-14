package usecase

import (
	"context"

	"github.com/viciousvs/currencyAPI/internal/model"
	"github.com/viciousvs/currencyAPI/internal/repository"
)

const LIMIT = 43

type CurrencyUseCase interface {
	GetLastRates(ctx context.Context) (model.Currencies, error)
	GetCurrencyByName(ctx context.Context, name string) (model.Currency, error)
	UpdateCurrencies(ctx context.Context, c model.Currencies) error
}

type currencyUseCase struct {
	repo repository.CurrencyRepo
}

func NewCurrencyUseCase(repo repository.CurrencyRepo) CurrencyUseCase {
	return currencyUseCase{repo: repo}
}

func (cu currencyUseCase) GetLastRates(ctx context.Context) (model.Currencies, error) {
	return cu.repo.GetLastRates(ctx, LIMIT)
}
func (cu currencyUseCase) GetCurrencyByName(ctx context.Context, name string) (model.Currency, error) {
	return cu.repo.GetCurrencyByName(ctx, name)
}
func (cu currencyUseCase) UpdateCurrencies(ctx context.Context, c model.Currencies) error {
	_, err := cu.repo.UpdateCurrencies(ctx, c)
	if err != nil {
		return err
	}
	return nil
}
