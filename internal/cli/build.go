package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func Build() *cli.Command {

	return &cli.Command{
		Name:  "build",
		Usage: "Builds podman images from designated YAML file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "pman.yaml",
				Usage:   "File used to build images",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			file := c.String("file")

			fmt.Println("using file ", file)

			// Debugging
			file = "../" + file
			actFile, err := os.ReadFile(file)

			if err != nil {
				log.Fatal(err)
				return err
			}

			fmt.Println("actual yaml file contents ", string(actFile))
			return nil
		},
	}
}
