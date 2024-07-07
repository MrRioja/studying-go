package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Search servers by IP and servers"
	app.Version = "1.0.0"

	flags :=
		[]cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "google.com",
				Usage: "Host to search IPs",
			},
		}

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Usage:   "Search IPs by address on the internet",
			Aliases: []string{"s"},
			Action:  searchIps,
			Flags:   flags,
		},
		{
			Name:    "servers",
			Usage:   "Search servers names on the internet",
			Aliases: []string{"svs"},
			Action:  searchServers,
			Flags:   flags,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
