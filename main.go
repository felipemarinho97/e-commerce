package main

import (
	"context"

	"github.com/felipemarinho97/e-commerce/api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	api.Server(ctx)
}
