package cmd

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"syscall"
)

func ParseArguments() (int, string) {
	if len(os.Args) != 3 {
		printUsageAndExit()
	}

	seq, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		printUsageAndExit()
	}

	return int(seq), os.Args[2]
}

func ReadPassphrase() string {
	var passphrase string
	if terminal.IsTerminal(syscall.Stdin) {
		fmt.Print("Passphrase: ")
		passphraseBytes, _ := terminal.ReadPassword(syscall.Stdin)
		passphrase = string(passphraseBytes)
	} else {
		fmt.Scanln(&passphrase)
	}
	return passphrase
}

func printUsageAndExit() {
	fmt.Printf("Usage: %s seq seed\n", os.Args[0])
	os.Exit(1)
}
