package main

import (
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("Hello")
}
