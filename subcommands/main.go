package main

import (
	"flag"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("todo", "0.1.0")
	c.Args = os.Args[1:]

	var judge = true
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{judge: judge}, nil
		},
	}
	code, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(code)
}

// サブコマンドをインタフェースとして定義
type AddCommand struct {
	judge bool
}

func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

func (c *AddCommand) Run(args []string) int {
	log.Println("Call Run ... ")
	log.Println(c.judge)

	var treat string
	flags := flag.NewFlagSet("add", flag.ContinueOnError)
	flags.StringVar(&treat, "t", "FILE", "treat type")
	if err := flags.Parse(args); err != nil {
		return 1
	}

	log.Println(treat)
	return 0
}
