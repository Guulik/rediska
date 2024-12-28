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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

}

func interactiveMode() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()

		args, err := extractArgs(scanner.Text())
		if err != nil {
			continue
		}
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

func extractArgs(input string) ([]string, error) {
	if input == "exit" {
		os.Exit(0)
	}

	args := strings.Fields(input)
	if len(args) == 0 {
		fmt.Println("empty command")
		return nil, fmt.Errorf("empty command")
	}
	fmt.Println("!!DEBUG!! args: ", args)

	return args, nil
}
