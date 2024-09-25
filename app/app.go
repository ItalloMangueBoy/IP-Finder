package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gen() *cli.App {
	app := cli.NewApp()

	app.Name = "IP-Finder"
	app.HelpName = "ipf"
	app.Usage = "Busca por Ips e nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: FindIp,
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
