package main

import (
	"time"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"strconv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	possibles := make(map[[2]string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	var w1, w2 string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		words := strings.Fields(line)
		for _, word := range words {
			possibles[[2]string{w1, w2}] = append(possibles[[2]string{w1, w2}], word)
			w1, w2 = w2, word
		}
	}
	possibles[[2]string{w1, w2}] = append(possibles[[2]string{w1, w2}], w2)
	possibles[[2]string{w2, ""}] = append(possibles[[2]string{w2, ""}], "")
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number of words to generate>")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number of words to generate:", os.Args[1])
		os.Exit(1)
	}
	var keys [][2]string
	for k := range possibles {
		if len(possibles[k]) > 1 {
			keys = append(keys, k)
		}
	}
	start := keys[rand.Intn(len(keys))]
	w1, w2 = start[0], start[1]
	output := []string{w1, w2}
	for i := 0; i < n; i++ {
		nextWords := possibles[[2]string{w1, w2}]
		word := nextWords[rand.Intn(len(nextWords))]
		output = append(output, word)
		w1, w2 = w2, word
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, strings.Join(output, " "))
	w.Flush()
}
