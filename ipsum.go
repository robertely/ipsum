package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"syscall"
)

func main() {

	possibleSignals := []os.Signal{syscall.SIGABRT, syscall.SIGALRM, syscall.SIGBUS, syscall.SIGCHLD, syscall.SIGCONT, syscall.SIGEMT, syscall.SIGFPE, syscall.SIGHUP, syscall.SIGILL, syscall.SIGINFO, syscall.SIGINT, syscall.SIGIO, syscall.SIGIOT, syscall.SIGKILL, syscall.SIGPIPE, syscall.SIGPROF, syscall.SIGQUIT, syscall.SIGSEGV, syscall.SIGSTOP, syscall.SIGSYS, syscall.SIGTERM, syscall.SIGTRAP, syscall.SIGTSTP, syscall.SIGTTIN, syscall.SIGTTOU, syscall.SIGURG, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGVTALRM, syscall.SIGWINCH, syscall.SIGXCPU, syscall.SIGXFSZ, os.Interrupt}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, possibleSignals...)
	go func() {
		for sig := range signalChannel {
			fmt.Fprintf(os.Stdout, "Recived \033[0;35mSIG\033[0m %s\n", sig)
		}
	}()

	host, _ := os.Hostname()
	user, _ := user.Current()

	fmt.Fprintf(os.Stdout, "host: %s\n", host)
	fmt.Fprintf(os.Stdout, "user: %s\n", user.Username)
	fmt.Fprintf(os.Stdout, "---\n")
	fmt.Fprintf(os.Stdout, "Standard \033[0;32mOUT\033[0m write test\n")
	fmt.Fprintf(os.Stderr, "Srandard \033[0;31mERR\033[0m write test\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stdout, "Standard \033[0;36mIN\033[0m read test: %s\n", scanner.Text())
	}

}
