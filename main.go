package main

import (
	"os/user"
	"fmt"
	"github.com/reezaras/belt/belt-white/repl"
	"os"
)

func main(){
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s [Pretty Serious WhiteBelt Stage Interpreter Experience] \n",user.Username)
	repl.Start(os.Stdout,os.Stdout)
}
