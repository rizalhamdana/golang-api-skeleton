package main

import (
	"flag"
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"
)

func main() {
	isLocalDevelopment := flag.Bool("local", false, "=(true/false)")
	flag.Parse()
	if *isLocalDevelopment {
		err := configEnv.Load(".env")
		if err != nil {
			fmt.Println(".env is not loaded properly")
			os.Exit(2)
		}
	}

	service := MakeHandler()
	service.HttpServerMain()
}
