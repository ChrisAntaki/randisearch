package main

import (
	"fmt"
	"github.com/toqueteos/webbrowser"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Create a URL.
	engine := chooseRandomLineFrom("engines.txt")
	query := strings.Replace(chooseRandomLineFrom("queries.txt"), " ", "+", -1)
	url := strings.Replace(engine+query, "\r", "", -1)

	// Let's give a heads up.
	fmt.Println("Visiting", url)

	// Open the URL.
	webbrowser.Open(url)
}

func chooseRandomLineFrom(filename string) string {
	content, err := ioutil.ReadFile(filename)
	dealbreaker(err)

	lines := strings.Split(string(content), "\n")

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(lines))

	return lines[index]
}

func dealbreaker(err error) {
	if err != nil {
		panic(err)
	}
}
