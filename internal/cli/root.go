package cli

import (
	"ex.com/podman"
	"github.com/urfave/cli/v3"
)

func NewApp() *cli.Command {

	podmanClient := podman.New("http://localhost:9091")

	return &cli.Command{
		Name:  "pman",
		Usage: "Podman Environment Manager",

		Commands: []*cli.Command{
			Build(podmanClient),
		},
	}
}
