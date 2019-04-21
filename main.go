package main

import (
	"flag"
	"fmt"
)

func main() {

	add := flag.Bool("a", false, "Implement append to the list: `urlshorten configure -a dogs -u www.dogs.com`")
	url := flag.String("u", "", "Implement append to the list: `urlshorten configure -a dogs -u www.dogs.com`")
	rem := flag.String("d", "", "Implement remove from the list: `urlshorten -d dogs`")
	lis := flag.Bool("l", false, "List redirections: `urlshorten -l`")
	port := flag.Int64("p", 8080, "Run HTTP server on a given port: `urlshorten run -p 8080`")
	help := flag.Bool("h", false, "Prints usage info: `urlshorten -h`")

	flag.Parse()

	//help flag
	if *help == true {
		flag.PrintDefaults()
	}

	fmt.Println(*add, *url, *rem, *lis, *port)
}
