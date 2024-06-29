package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Searches IPs and Sever Names on the Internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP Adresses on the web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIPs,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
