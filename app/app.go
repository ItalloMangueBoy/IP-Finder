package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

var flags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "host",
		Value: "github.com",
	},
}

func Gen() *cli.App {
	app := cli.NewApp()

	app.Name = "IP-Finder"
	app.HelpName = "ipf"
	app.Usage = "Busca por Ips e nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: FindIp,
		},

		{
			Name:   "server",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: FindServer,
		},
	}

	return app
}

func FindIp(context *cli.Context) {
	host := context.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func FindServer(context *cli.Context) {
	host := context.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
