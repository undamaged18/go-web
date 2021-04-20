package auth

import (
	"context"
	"fmt"
)

func HasAuth(ctx context.Context) bool {
	token := ctx.Value("auth_token")
	if token != nil {
		fmt.Println(token.(string))
		return true
	}
	return false
}
