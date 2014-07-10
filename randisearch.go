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
		time.Sleep(time.Duration(settings.delay) * time.Second)
		searchSomethingSomewhere()
	}
}

type Settings struct {
	delay int
	loop  bool
}

func getSettings() Settings {
	delay := flag.Int("delay", 10, "Delay between searches, when looping.")
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
	if err != nil {
		panic(err)
	}
}
