package main

import (
	"fmt"
	"os"
	"os/user"
	"bufio"
)

func main() {
	host, _ := os.Hostname()
	user, _ := user.Current()

	fmt.Fprintf(os.Stdout, "host: %s\n", host)
	fmt.Fprintf(os.Stdout, "user: %s\n", user.Username)
	fmt.Fprintf(os.Stdout, "---\n")
	fmt.Fprintf(os.Stdout, "Standard \033[0;32mOUT\033[0m write test\n")
	fmt.Fprintf(os.Stderr, "Srandard \033[0;31mERR\033[0m write test\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "Standard \033[0;36mIN\033[0m  read  test: %s\n", scanner.Text())
	}

}
