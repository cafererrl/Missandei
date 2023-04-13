package main

// Main driver for the Let Language Interpreter
func main() {
	// Call the Scanner and init program file data.
	promptForFileInput()

	// Scan the user program file.
	scan()

	// Create a parse tree and print the Abstract Syntax Tree generated from the scanned Tokens.
	runParser()

	// Evaluate the parse tree.
	runEval()

}
