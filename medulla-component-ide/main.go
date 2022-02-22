package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		args = append(args, "info")
	}

	switch args[1] {
	case "version":
		fmt.Println(`

medulla-cli Version 1.01

`)
	default:
		fmt.Println(`medulla-cli [command] [option]	

version							Prints version
setpath							Set path to medulla repo
installmedulla						Installs medulla		
						
				`)
	}

}
