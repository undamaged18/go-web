package models

import (
	"context"
	"fmt"
)

type page struct {
	Nonce struct {
		Style string
		Script string
	}
}

func NewPage(ctx context.Context) *page {
	return &page{
		Nonce: struct {
			Style  string
			Script string
		}{
			fmt.Sprintf("%v", ctx.Value("styleNonce")),
			fmt.Sprintf("%v", ctx.Value("scriptNonce")),
		},
	}
}