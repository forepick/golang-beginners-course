package main

import (
	"flag"
	"os"
)

type (
	Config struct {
		Name string
	}
)

func main(){
	config := &Config{}
	flag.StringVar(&config.Name, "name", "my-job","What's the name of the job to execute")

	help := flag.Bool("help", false, "To get this help")
	h := flag.Bool("h", false, "To get this help")

	flag.Parse()

	if flag.NArg() == 0 || *help || *h{
		flag.Usage()
		os.Exit(1)
	}


}
