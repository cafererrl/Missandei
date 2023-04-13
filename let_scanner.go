package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
	GLOBALS & REFERENCE
*/

// Grammar - the BNF grammar of the LetLang.
// var grammar = [6][]string{
//	{"E", "->", "Number"},
//	{"E", "->", "minus", "(", "E", ",", "E", ")"},
//	{"E", "->", "iszero", "(", "E", ")"},
//	{"E", "->", "if", "E", "then", "E", "else", "E"},
//	{"E", "->", "Identifier"},
//	{"E", "->", "let", "Identifier", "=", "E", "in", "E"},
//}

// fileContents - The program from the file.
var fileContents string

// Token - Define <int32 key, string tokenSymbolValue> Token type.
type Token struct {
	tokenTypeKeyId   rune
	tokenSymbolValue string
}

// tokenQueue - Declare global for containing input string tokens.
var tokenQueue []Token

// lexeme - The lexical unit being constructed to form/match a Token in the language.
var lexeme string

// lexMap - Map/Dictionary of Tokens that belong in LetLang.
var lexMap = map[string]rune{
	"(":          1,
	")":          2,
	",":          3,
	"minus":      4,
	"=":          5,
	"iszero":     6,
	"if":         7,
	"then":       8,
	"else":       9,
	"let":        10,
	"in":         11,
	"identifier": 12,
	"integer":    13,
	"unknown":    -1,
}

/*
	IMPLEMENT SCANNER AND PRINT TOKEN QUEUE.
*/

// promptForFileInput - Takes in the file name from user if file is found, and sets file (program) contents as a string.
func promptForFileInput() {
	// Taking input from user
	fmt.Println(Prompt("Enter the name of your Let program (with NO file extension suffix): "))
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	filename := input + ".txt"

	// Keep asking user for correct file name if Path Error.
	if fileExists(filename) {
		file, err := os.ReadFile(filename)
		check(err)
		fileContents = string(file)
		//println()
		//print(fileContents)
	} else {
		fmt.Println(Error("Path error. Check file name and try again."))
		promptForFileInput()
	}
}

// scan - Returns a queue of Tokens from string in fileContents
func scan() []Token {
	tokenize(fileContents)
	printQueue()
	return tokenQueue
}

// tokenize - Builds Tokens from program string.
func tokenize(contents string) {
	// The Token key (symbol) that will be generated to add to the queue.
	//lexeme = ""
	// Iterate through every character in the program file content string and track its index.
	for pos, char := range contents {
		lexeme += string(char) // Build the Token.
		switch {
		// Absorb whitespace.
		case isWhiteSpace(lexeme):
			clearLexeme()
		// If the lexeme being built is a known Token in the Let Language, add it to the queue.
		case checkLexemeExistsAsToken(lexeme):
			tokenQueue = append(tokenQueue, Token{lexMap[lexeme], lexeme})
			clearLexeme()
		// If the Token is intact, identify it, otherwise continue building the lexeme. Requires lookahead.
		default:
			if pos == len(contents)-1 {
				if isNumeric(lexeme) { // Int
					tokenQueue = append(tokenQueue, Token{13, lexeme})
					clearLexeme()
				} else { // Id
					tokenQueue = append(tokenQueue, Token{12, lexeme})
					clearLexeme()
				}
				continue
			}
			// Lookahead next char in content string as UTF-8 tokenSymbolValue
			lookahead := string([]rune(contents)[pos+1])
			switch {
			case isWhiteSpace(lookahead):
				if isNumeric(lexeme) { // Integer
					tokenQueue = append(tokenQueue, Token{13, lexeme})
					clearLexeme()
				} else { // Identifier
					tokenQueue = append(tokenQueue, Token{12, lexeme})
					clearLexeme()
				}
			case isIdentifier(lookahead):
				if isNumeric(lexeme) { // Integer
					tokenQueue = append(tokenQueue, Token{13, lexeme})
					clearLexeme()
				} else { // Identifier
					tokenQueue = append(tokenQueue, Token{12, lexeme})
					clearLexeme()
				}
			default:
				continue
			}
		}
	}
}

/*
	HELPER FUNCTIONS
*/

// isNumeric - Tests if string is numeric.
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// isWhiteSpace - Used to determine if chars in lexeme (string) is whitespace.
func isWhiteSpace(str string) bool {
	if unicode.IsSpace(rune(str[0])) {
		return true
	}
	return false
}

// isIdentifier - Used to determine if char in lexeme (string) is identifier.
func isIdentifier(str string) bool {
	// Type-coercing by checking total length of string...
	isPunct := unicode.IsPunct(rune(str[0]))
	isSymbol := unicode.IsSymbol(rune(str[0]))
	if isSymbol || isPunct {
		return true
	} else {
		return false
	}
}

// clearLexeme - Reset the lexeme being built.
func clearLexeme() {
	lexeme = ""
}

// checkLexemeExistsAsToken - Reference the map of existing Tokens in Let Language.
func checkLexemeExistsAsToken(lex string) bool {
	_, isPresent := lexMap[lex]
	return isPresent
}

// addToQueue - Add Token to queue based on its type. Refactor switch/case redundancy.
//func addToQueue

// printQueue - Display Tokens in queue.
func printQueue() {
	fmt.Println()
	fmt.Println(Info("Queue Contents:"))
	for idx, element := range tokenQueue {
		fmt.Println(idx, ".\t", element)
	}
}
