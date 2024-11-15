package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"rediska/config"
	"rediska/internal/cli"
	"strings"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	cli.InitCobra(address)

	fmt.Println("Hello! Its redis cli :>")
	go interactiveMode()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}

func interactiveMode() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		fmt.Println("!!DEBUG!! args: ", args)
		strings.ToUpper(args[0])

		cmd, _, err := cli.RootCmd.Find(args)
		if err != nil {
			fmt.Println("Unknown command:", args[0])
			continue
		}
		cli.RootCmd.SetArgs(args)
		cmd.Execute()
	}
}
