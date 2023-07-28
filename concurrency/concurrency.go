package concurrency

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printMe(text string, secs int) {
	defer wg.Done()
	duration := time.Duration(secs) * time.Second
	time.Sleep(duration)
	fmt.Printf("Finally... %s after %d seconds\n", text, secs)
}

func SimplePrintWithTimeout(times int) {
	start := time.Now()
	fmt.Println("*** concurrency ****")
	wg.Add(times)
	supposedTime := 0
	for i := 0; i < times; i++ {
		thisTime := rand.Intn(5)
		supposedTime = supposedTime + thisTime
		go printMe(fmt.Sprintf("#%d", i+1), thisTime)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Function took %v and not %d seconds", elapsed, supposedTime)
}

func WordCount(textFilePath string) {
	start := time.Now()
	fmt.Println("*** concurrency ****")

	file, err := os.Open(textFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines = []string{""}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	wg.Add(len(lines))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var wordCounts = []map[string]int{}
	for _, line := range lines {
		wordsChan := make(chan map[string]int)
		go countWords(line, wordsChan)
		words := <-wordsChan
		wordCounts = append(wordCounts, words)
	}
	wg.Wait()

	var finalWordCount = map[string]int{}
	for _, wordCount := range wordCounts {
		for word, value := range wordCount {
			count, ok := finalWordCount[word]
			if ok {
				finalWordCount[word] = count + value
			} else {
				finalWordCount[word] = value
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Function took %v seconds\n", elapsed)
	printWordCount(finalWordCount)
}

func printWordCount(wordCount map[string]int) {
	fmt.Println("Word count:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(text string, wordsChan chan map[string]int) {
	defer wg.Done()
	var tokensCount = map[string]int{}
	tokens := strings.Split(strings.ToLower(text), " ")
	for _, token := range tokens {
		count, ok := tokensCount[token]
		if ok {
			tokensCount[token] = count + 1
		} else {
			tokensCount[token] = 1
		}
	}
	wordsChan <- tokensCount
}
