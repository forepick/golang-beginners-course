package main

import (
	"flag"
	"fmt"
	"os"
)

type (
	Config struct {
		Username string
		Password string
		Threads  int
	}
)

func main() {
	config := &Config{}
	flag.StringVar(&config.Username,
		"username",
		"admin",
		"The user who invoke this process")

	flag.StringVar(&config.Password,
		"password",
		"",
		"Secret key")

	flag.IntVar(&config.Threads,
		"threads",
		2,
		"Number of threads")

	flag.Parse()

	fmt.Printf("num args: %d\n", flag.NArg())
	fmt.Printf("num flags: %d\n", flag.NFlag())

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Config object: %+v\n", config)
}
