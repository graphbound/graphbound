package domain

import "context"

type Quote struct {
	Value string `json:"value"`
}

type GetQuoteUseCase interface {
	Execute(ctx context.Context) (*Quote, error)
}
