package app

import (
	"fmt"

	"github.com/urfave/cli"
)

func Gen() *cli.App {
	app := cli.NewApp()

	app.Name = "IP-Finder"
	app.HelpName = "ipf"
	app.Usage = "Busca por Ips e nomes de servidor na internet"

	fmt.Println(app)

	return app
}
