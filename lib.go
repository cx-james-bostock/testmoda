package testmoda

import (
	"fmt"
	"log"
	"go.uber.org/zap"
)

func Foo() {
	fmt.Println("foo!")
}

func Baz() {
	fmt.Println("baz!")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	defer logger.Sync()
	sugar.Debug("This is a debug message")
}
