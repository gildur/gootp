package main

import (
	"fmt"
	"github.com/gildur/gootp"
	"github.com/gildur/gootp/cmd"
)

func main() {
	seq, seed := cmd.ParseArguments()
	passphrase := cmd.ReadPassphrase()
	fmt.Println(gootp.OtpMd4(seq, seed, passphrase))
}
