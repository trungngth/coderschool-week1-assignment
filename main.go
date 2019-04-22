package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type flagStorage struct {
	addURL    *string
	stringURL *string
	removeURL *string
	listURL   *bool
	port      *int64
	help      *bool
	run       *bool
	configure *bool
}

var serialData map[string]string

func main() {

	fs := flagStorage{}
	initFlag(&fs)
	readYAML()

	if *fs.help {
		flag.PrintDefaults()
	}
	//List the redirects list
	if *fs.listURL {
		printYAML()
	}
	//Add url to the list
	if *fs.configure {
		serialData[*fs.addURL] = *fs.stringURL
		writeYAML()
	}
	//Remove url from the list
	if *fs.removeURL != "" {
		delete(serialData, *fs.removeURL)
		writeYAML()
	}
	if *fs.run {
		http.HandleFunc("/", initServer)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *fs.port), nil))
	}

}

func initServer(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path[1:]
	urlMatch, match := serialData[urlPath]
	if match {
		http.Redirect(w, r, urlMatch, http.StatusMovedPermanently)
	}
	http.Error(w, "Not in the redirect list", http.StatusNotFound)
}

func initFlag(fs *flagStorage) {
	fs.addURL = flag.String("a", "", "Implement append to the list: `urlshorten configure -a dogs -u www.dogs.com`")
	fs.stringURL = flag.String("u", "", "Implement append to the list: `urlshorten configure -a dogs -u www.dogs.com`")
	fs.removeURL = flag.String("d", "", "Implement remove from the list: `urlshorten -d dogs`")
	fs.listURL = flag.Bool("l", false, "List redirections: `urlshorten -l`")
	fs.port = flag.Int64("p", 8080, "Run HTTP server on a given port: `urlshorten run -p 8080`")
	fs.help = flag.Bool("h", false, "Prints usage info: `urlshorten -h`")
	fs.configure = flag.Bool("c", false, "Configure the app")
	fs.run = flag.Bool("r", false, "Run the app")

	flag.Parse()
}

func readYAML() {
	bytes, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &serialData)
	if err != nil {
		log.Fatal(err)
	}
}

func writeYAML() {
	bytes, err := yaml.Marshal(&serialData)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./config.yaml", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func printYAML() {
	for k, v := range serialData {
		fmt.Printf("%v : %v\n", k, v)
	}
}
