package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"ex.com/config"
	"ex.com/podman"
	"github.com/urfave/cli/v3"
)

type App struct {
	Podman *podman.HTTPClient
}

func Build(client *podman.HTTPClient) *cli.Command {

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

			// So pass the file tot eh YAML Parrser
			parsedStruct, err := config.Parse(actFile)

			if err != nil {
				log.Fatal(err)
				return err
			}

			for name, service := range parsedStruct.Services {
				fmt.Println("name", name)
				fmt.Println("service", service)
				fmt.Println("context", service.Build.Context)
				req := podman.BuildRequest{
					Name:    name,
					Image:   service.Image,
					Context: service.Build.Context,
				}

				err := client.Build(req)

				if err != nil {
					log.Fatal(err)
					return err
				}
			}

			// fmt.Println("actual yaml file contents ", string(actFile))
			return nil
		},
	}
}
