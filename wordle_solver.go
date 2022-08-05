// Wordle Solver

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getPossibleWords(currentWord string,
                      targetWordLength int,
                      validLetters []string,
                      knownPositions map[int]string,
                      dictionary map[string]bool) []string {
	var possibleWords []string

	currentLength := len(currentWord)
	if currentLength < targetWordLength {
		// Keep trying to append letters to find a word
		var words []string
		// First, check if the current position is known
		nextLetter, ok := knownPositions[currentLength]
		if ok {
			currentWord += nextLetter
			words = getPossibleWords(
				currentWord,
				targetWordLength,
				validLetters,
				knownPositions,
				dictionary,
			)
			possibleWords = append(possibleWords, words...)
		} else {
			for _, letter := range validLetters {
				newWord := currentWord + letter
				words = getPossibleWords(
					newWord,
					targetWordLength,
					validLetters,
					knownPositions,
					dictionary,
				)
				possibleWords = append(possibleWords, words...)
			}
		}
	} else {
		// Check if a valid word has been found
		_, ok := dictionary[currentWord]
		if ok {
			possibleWords = append(possibleWords, currentWord)		
		}
	}

	return possibleWords
}


func main() {
	// ------------
	// SET UP FLAGS
	// ------------

	debug        := flag.Bool("debug", false,
		"Print the dictionary map to ensure " + 
		"that the text file was read correctly.")
	verbose      := flag.Bool("verbose", false,
		"Show how the user's arguments were interpreted.")
	wordFilePath := flag.String("file", "freebsd_words.txt",
		"User-specified text file from which to read in English-language words.")
	wordLength   := flag.Int("length", 5,
		"The length of the word to be found.")
	excludeArg   := flag.String("exclude", "",
		"List of letters that are known to not be in the word. " +
		"Separate multiple with a comma. For example: -exclude m,s,e")
	knownArg     := flag.String("known", "",
		"List of known positions (zero-indexed) and letters. " +
		"Separate multiple with a comma. For example: -known 0m,1o,2u")
	saveToTxt    := flag.Bool("txt", false,
		"Write the possible solutions to a .txt file.")
	flag.Parse()

	// -----------------------
	// VALIDATE USER ARGUMENTS
	// -----------------------

	if *wordLength <= 0 {
		fmt.Printf("Error: Word length must be greater than 0.\n")
		os.Exit(1)
	}

	// Create map of letters to exclude
	excludedLetters := make(map[string]bool)
	excludeArgs := strings.Split(*excludeArg, ",")
	for _, arg := range excludeArgs {
		if len(arg) == 1 {
			matched, err := regexp.Match(`\w`, []byte(arg))
			if err != nil {
				fmt.Printf("Error when using regexp " +
					"on letters to exclude: %v\n", err)
			} else if matched {
				letter := strings.ToLower(string(arg[0]))
				excludedLetters[letter] = true
			}
		}
	}
	if len(excludedLetters) > 0 && *verbose {
		fmt.Println("Letters to exclude:")
		for letter, shouldExclude := range excludedLetters {
			if shouldExclude {
				fmt.Println(letter)
			}
		}
		fmt.Println()
	}
	
	// Use map of excluded letters to create array of valid letters
	var validLetters []string
	for i := 0; i < 26; i++ {
		letter := string(rune('a' + i))
		_, ok := excludedLetters[letter]
		if !ok {
			validLetters = append(validLetters, letter)
		}
	}
	if len(validLetters) == 0 {
		fmt.Println("Error: All letters of the alphabet have been excluded.")
		os.Exit(1)
	}
	if len(validLetters) < 26 && *verbose {
		fmt.Printf("Valid letters:\n%v\n\n", validLetters)
	}

	// Create map of known positions
	knownPositions := make(map[int]string)
	knownArgs := strings.Split(*knownArg, ",")
	if len(knownArgs) > *wordLength {
		fmt.Printf("Error: Number of known positions (%d) is greater than " +
			"the word length (%d).\n", len(knownArgs), *wordLength)
		os.Exit(1)
	}
	for _, arg := range knownArgs {
		if len(arg) > 1 {
			matched, err := regexp.Match(`\d\w`, []byte(arg))
			if err != nil {
				fmt.Printf("Error when using regexp on " +
					"-known arg \"%s\": %v\n", arg, err)
			} else if matched {
				position, err := strconv.Atoi(string(arg[0]))
				if err != nil {
					fmt.Printf("Error when converting first character in " +
						"-known arg \"%s\" to int: %v\n", arg, err)
				} else if position >= *wordLength {
					fmt.Printf("Error: Position in -known arg \"%s\" exceeds " +
						"word length (%d).\n", arg, *wordLength)
				} else {
					letter := strings.ToLower(string(arg[1]))
					if excludedLetters[letter] == true {
						fmt.Printf("Letter \"%s\" already in " +
							"excluded letters list.\n", letter)
					} else {
						knownPositions[position] = letter
					}
				}
			}
		}
	}
	if len(knownPositions) > 0 && *verbose {
		wordSoFar := ""
		for i := 0; i < *wordLength; i++ {
			letter, ok := knownPositions[i]
			if ok {
				wordSoFar += letter
			} else {
				wordSoFar += "_"
			}
		}
		fmt.Printf("Current progress:\n%s\n\n", wordSoFar)
	}

	// --------------------
	// BUILD THE DICTIONARY
	// --------------------

	// Read in words from text file
	if *verbose {
		fmt.Printf("Reading in words from: %s\n", *wordFilePath)
	}
	txtFile, err := os.Open(*wordFilePath)
	if err != nil {
		fmt.Printf("Error when opening \"%s\": %v\n", *wordFilePath, err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(txtFile)
	wordMap := make(map[string]bool)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if len(word) > 0 {
			wordMap[word] = true
		}
	}
	if err = scanner.Err(); err != nil{
		fmt.Printf("Error when creating new Scanner: %v\n", err)
		os.Exit(1)
	}
	txtFile.Close()

	// Check if the resulting map is valid
	if len(wordMap) == 0 {
		fmt.Println("Error: No words were found in the text file.")
		os.Exit(1)
	}
	if *verbose {
		fmt.Printf("Dictionary contains %d words.\n\n", len(wordMap))
	}
	if *debug {
		fmt.Println("Displaying the contents of the map:")
		for word, _ := range wordMap {
			fmt.Println(word)
		}
		fmt.Println()
	}

	// -----------------------
	// FIND ALL POSSIBLE WORDS
	// -----------------------

	possibleWords := getPossibleWords(
		"",
		*wordLength,
		validLetters,
		knownPositions,
		wordMap,
	)

	fmt.Println("Possible solutions:")
	for _, word := range possibleWords {
		fmt.Println(word)
	}

	if *saveToTxt {
		txtFilename := "results.txt"
		if _, err = os.Stat(txtFilename); err == nil {
			err = os.Remove(txtFilename)
			if err != nil {
				fmt.Printf("Error when attempting to remove existing file " +
					"\"%s\": %v\n", txtFilename, err)
				os.Exit(1)
			}
		}
		txtFile, err := os.OpenFile(txtFilename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Error when opening \"%s\": %v\n", txtFilename, err)
			os.Exit(1)
		}

		txtFile.WriteString("Possible solutions:\n")
		for _, word := range possibleWords {
			txtFile.WriteString(word + "\n")
		}

		txtFile.Close()
	}
}

