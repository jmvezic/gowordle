package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var allowedWords = []string{}

func main() {

	rand.Seed(time.Now().Unix())

	fmt.Println()
	fmt.Println("How to play Wordle: ")
	fmt.Println("1. The goal is to guess a 5-letter word")
	fmt.Println("2. You have 6 tries to guess it")
	fmt.Println("3. Character coloring indicates the following: ")
	fmt.Println("\t" + Yellow + "Yellow" + Reset + " means the character is somewhere in the word, but is not in the correct position")
	fmt.Println("\t" + Green + "Green" + Reset + " means the character is in the correct position")
	fmt.Println("\t" + "No color means the character is not in the word at all")
	fmt.Println()

	// all allowed words
	fileBytes, err := os.ReadFile("answers.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// file to slice
	correctAnswer := strings.Split(string(fileBytes), "\n")

	// all allowed words
	allowedBytes, err := os.ReadFile("allowed.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// file to slice
	allowedWords = strings.Split(string(allowedBytes), "\n")

	correctWordIndex := rand.Int() % len(correctAnswer)
	correctWord := correctAnswer[correctWordIndex]

	guessesSoFar := []string{}

	runGuess(correctWord, 6, guessesSoFar)

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeOnce(a string, list []string) []string {

	newList := []string{}

	alreadyFound := false

	for _, b := range list {
		if b == a {
			if !alreadyFound {
				alreadyFound = true
			} else {
				newList = append(newList, b)
			}
		} else {
			newList = append(newList, b)
		}
	}

	return newList
}

func runGuess(correctWord string, tries int, guessesSoFar []string) {

	fmt.Println()
	for _, g := range guessesSoFar {
		fmt.Println(g)
	}
	fmt.Println()

	if tries <= 0 {
		fmt.Println(Red + "Sorry, no more tries!" + Reset)
		fmt.Println("Correct word was " + Green + strings.ToUpper(correctWord) + Reset)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter guess (" + "Try " + strconv.Itoa(-(tries - 7)) + "/6): ")
	guess, _ := reader.ReadString('\n')
	// convert CRLF to LF
	guess = strings.Replace(guess, "\n", "", -1)

	if len(guess) != 5 {
		fmt.Println(Red + "Enter a 5-character guess, please." + Reset)
		runGuess(correctWord, tries, guessesSoFar)
		return
	}

	if !stringInSlice(guess, allowedWords) {
		fmt.Println(Red + "Illegal guess." + Reset)
		runGuess(correctWord, tries, guessesSoFar)
		return
	}

	charsGuess := strings.Split(guess, "")
	charsCorrect := strings.Split(correctWord, "")

	if guess == correctWord {
		fmt.Println("Correct! " + Green + strings.ToUpper(guess) + Reset)
		return
	}

	coloredGuess := []string{}
	correctGuess := []string{}
	correctLeftovers := []string{}

	for i := 0; i < len(charsGuess); i++ {

		char := string(charsGuess[i])

		if charsCorrect[i] == char {
			correctGuess = append(correctGuess, char)
			correctLeftovers = append(correctLeftovers, "")
		} else {
			correctGuess = append(correctGuess, "")
			correctLeftovers = append(correctLeftovers, charsCorrect[i])
		}

	}

	for i := 0; i < len(correctGuess); i++ {

		char := string(charsGuess[i])

		if correctGuess[i] != "" {
			coloredGuess = append(coloredGuess, Green+strings.ToUpper(char)+Reset)
		} else {
			if stringInSlice(char, correctLeftovers) {
				coloredGuess = append(coloredGuess, Yellow+strings.ToUpper(char)+Reset)
				correctLeftovers = removeOnce(char, correctLeftovers)
			} else {
				coloredGuess = append(coloredGuess, strings.ToUpper(char))
			}
		}

	}

	coloredGuessString := strings.Join(coloredGuess, "")
	guessesSoFar = append(guessesSoFar, coloredGuessString)
	runGuess(correctWord, tries-1, guessesSoFar)

}
