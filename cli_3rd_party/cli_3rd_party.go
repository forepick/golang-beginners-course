package main

import (
	"flag"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.App{
		Name:     "go-course-cli",
		Version:  "6.1.3.770",
	}


	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "./config.yaml",
		},
		cli.StringFlag{
			Name:  "overrides",
			Value: "./overrides.yaml",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "version",
			Usage: "Print the version",
			Action: func(c *cli.Context) error {
				fmt.Printf(app.Version)
				return nil
			},
		},
		{
			Name:  "say_hello",
			Usage: "flush entire cache",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Value: "Moyshe Zuchmir",
					Usage: "name of the good guy",
				},
			},
			Action: SayHello,
		},
		{
			Name: "sc_test",
			Subcommands: []cli.Command{
				{
					Name:   "sub1",
					Action: SubCommand1,
				},
			},
		},

	}

	app.EnableBashCompletion = true

	flag.Parse()

	app.Run(os.Args)
}
func SayHello(c *cli.Context) error {
	if name := c.String("name"); name != "" {
		fmt.Printf("Hello %s\n", name)
	} else {
		fmt.Println("Hello man!")
	}

	return nil
}

func SubCommand1(c *cli.Context) error {

	fmt.Println("This is sub command 1")
	return nil
}
