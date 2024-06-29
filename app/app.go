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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		}}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Searches IP Adresses on the web",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "servers",
			Usage:  "Searches Server Names on the web",
			Flags:  flags,
			Action: searchServers,
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

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servers := range servers {
		fmt.Println(servers.Host)
	}
}
