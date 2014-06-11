package main

import (
	"flag"
	"github.com/toqueteos/webbrowser"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Search.
	searchSomethingSomewhere()

	// Loop, if requested.
	loop := flag.Bool("loop", false, "Continuously search")
	flag.Parse()
	for *loop {
		time.Sleep(10 * time.Second)
		searchSomethingSomewhere()
	}
}

func searchSomethingSomewhere() {
	// Create a URL.
	engine := chooseRandomLineFrom("engines.txt")
	query := strings.Replace(chooseRandomLineFrom("queries.txt"), " ", "+", -1)
	url := strings.Replace(engine+query, "\r", "", -1)

	// Open the URL.
	webbrowser.Open(url)
}

func chooseRandomLineFrom(filename string) string {
	// Read file.
	content, err := ioutil.ReadFile(filename)
	dealbreaker(err)

	// Pick a random line.
	lines := strings.Split(string(content), "\n")
	rollTheDice()
	index := rand.Intn(len(lines))
	line := lines[index]

	return line
}

func rollTheDice() {
	rand.Seed(time.Now().Unix())
}

func dealbreaker(err error) {
	// Oh jeez!
	if err != nil {
		panic(err)
	}
}
