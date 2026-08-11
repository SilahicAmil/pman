package cli

import "github.com/urfave/cli/v3"

func NewApp() *cli.Command {

	return &cli.Command{
		Name:  "pman",
		Usage: "Podman Environment Manager",

		Commands: []*cli.Command{
			Build(),
		},
	}
}
