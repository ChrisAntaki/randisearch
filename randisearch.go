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

	settings := getSettings()

	for settings.loop {
		seedRandomGenerator()
		delay := settings.delay*rand.Float64()*1000 + 300
		duration := time.Duration(delay) * time.Millisecond
		time.Sleep(duration)

		searchSomethingSomewhere()
	}
}

type Settings struct {
	delay float64
	loop  bool
}

func getSettings() Settings {
	delay := flag.Float64("delay", 10, "Maximum delay between looping searches, in seconds.")
	loop := flag.Bool("loop", false, "Continuously search.")

	flag.Parse()

	return Settings{*delay, *loop}
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
	for {
		lines := strings.Split(string(content), "\n")
		seedRandomGenerator()
		index := rand.Intn(len(lines))
		line := lines[index]
		if len(line) > 0 {
			return line
		}
	}
}

func seedRandomGenerator() {
	rand.Seed(time.Now().Unix())
}

func dealbreaker(err error) {
	if err != nil {
		panic(err)
	}
}
