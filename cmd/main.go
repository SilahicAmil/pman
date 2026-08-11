package main

import (
	"context"
	"log"
	"os"

	"ex.com/internal/cli"
)

func main() {

	if err := cli.NewApp().Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
