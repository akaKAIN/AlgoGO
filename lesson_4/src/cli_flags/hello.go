package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello"
	app.Usage = "Print hello dude"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "body",
			Usage: "Saying hello to  dude, whom names ...",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hellow %s!\n", name)
		return nil
	}
	app.Run(os.Args)

}
