package main

import (
	"context"
	"github.com/xamust/ekom-test/config"
	"github.com/xamust/ekom-test/internal/application"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	application.RunWithCtx(ctx, cfg)
}
