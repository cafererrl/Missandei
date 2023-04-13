package main

import (
	"fmt"
	"os"
)

/*
	DEFINITIONS
*/
/*
astNode - An Abstract Syntax Tree guides the construction for the evaluator. Next to each node in the tree is an
environment, e, that holds the values of the variables. A top-down sweep will build an environment for each node, and
a bottom-up sweep (suitable for arithmetic expression) is performed to get the final value of the expression, at the
root node of the tree. Children are represented as a list of pointers to other AST nodes.
*/
type astNode struct {
	parent   *astNode // Pointer to parent node if exists. If null, then root.
	ttype    rune     // Token type, used to distinguish between vars and integers.
	termsym  bool     // Is it a terminal symbol (leaf node)?
	contents string   // Represent the contents as a string, even if it's an int.
	children []*astNode
}

// initTree - allocates a root node, initializes it, and advances queue to next Token.
func initTree(node *astNode, isTerminal bool) {
	node.termsym = isTerminal
	node.children = make([]*astNode, 0, 5) // allocate and init object of specified type, w/ return slice size and capacity
	node.ttype = tokenQueue[0].tokenTypeKeyId
	node.contents = tokenQueue[0].tokenSymbolValue
	advanceToken()
}

var globalRoot astNode

// runParser - the primary function call to activate parsing process
func runParser() {
	globalRoot = parseExpression()
	fmt.Println()
	fmt.Println(Info("Abstract Syntax Tree:"))
	startPrinting(globalRoot)
}

// startPrinting - the primary function call to begin printing the AST
func startPrinting(root astNode) {
	root.printTree(0)
}

// parseExpression -
func parseExpression() astNode {
	// Instantiate a local root.
	root := astNode{}
	queueHeadTokenType := tokenQueue[0].tokenTypeKeyId

	switch {
	// IDENTIFIER or INTEGER
	case queueHeadTokenType == lexMap["integer"] || queueHeadTokenType == lexMap["identifier"]:
		initTree(&root, true)
		return root
	// MINUS
	case queueHeadTokenType == lexMap["minus"]:
		initTree(&root, false) // Returns next token after initializing root
		checkToken(tokenQueue, "(", "Error in parseExp: lparen not found in case minus")
		child1 := parseExpression() // Pops tokenQueue
		checkToken(tokenQueue, ",", "Error in parseExp: comma not found in case minus")
		child2 := parseExpression() // Pops tokenQueue
		checkToken(tokenQueue, ")", "Error in parseExp: rparen not found in case minus")
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
	// ISZERO
	case queueHeadTokenType == lexMap["iszero"]:
		initTree(&root, false) // Returns next token after initializing root
		checkToken(tokenQueue, "(", "Error in parseExp: lparen not found in case iszero")
		child1 := parseExpression() // Pops tokenQueue
		checkToken(tokenQueue, ")", "Error in parseExp: rparen not found in case iszero")
		root.children = append(root.children, &child1)
	// IF
	case queueHeadTokenType == lexMap["if"]:
		initTree(&root, false)      // Returns next token after initializing root
		child1 := parseExpression() // Pops tokenQueue.
		checkToken(tokenQueue, "then", "Error in parseExp: then not found in case if")
		child2 := parseExpression() // Pops tokenQueue.
		checkToken(tokenQueue, "else", "Error in parseExp: else not found in case if")
		child3 := parseExpression() // Pops tokenQueue.
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)
	// LET
	case queueHeadTokenType == lexMap["let"]:
		initTree(&root, false)      // Returns next token after initializing root.
		child1 := parseExpression() // Pops tokenQueue.
		checkToken(tokenQueue, "=", "Error in parseExp: equals not found in case let")
		child2 := parseExpression() // Pops tokenQueue.
		checkToken(tokenQueue, "in", "Error in parseExp: in not found in case let")
		child3 := parseExpression() // Pops tokenQueue.
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)
	// UNKNOWN
	default:
		fmt.Println("Error in parseExp: unknown runParser.")
		os.Exit(1)
	}
	return root
}

/*
	HELPER FUNCTIONS
*/
// checkToken - Ensure Token is followed by appropriate symbol and advances input stream otherwise throws.
func checkToken(tokenList []Token, sym string, err string) {
	if tokenList[0].tokenSymbolValue == sym {
		advanceToken()
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

// advanceToken - Pop the next Token off the queue.
func advanceToken() {
	tokenQueue = tokenQueue[1:]
}

// startPrinting -
func (node astNode) printTree(indentLevel int) {
	output := ""
	for i := 0; i < indentLevel; i++ {
		output += "\t"
	}

	switch node.ttype {
	// IDENTIFIER
	case lexMap["identifier"]:
		fmt.Printf("%s", output)
		fmt.Println("VarExp(")
		fmt.Printf("%s", output)
		fmt.Print("\t")
		fmt.Println("\"" + node.contents + "\"")
		fmt.Printf("%s", output)
		fmt.Print(")")
	// INTEGER
	case lexMap["integer"]:
		fmt.Printf("%s", output)
		fmt.Println("ConstExp(")
		fmt.Printf("%s", output)
		fmt.Print("\t")
		fmt.Println(node.contents)
		fmt.Printf("%s", output)
		fmt.Print(")")
	// MINUS
	case lexMap["minus"]:
		fmt.Printf("%s", output)
		fmt.Println("DiffExp(")
		for idx, child := range node.children {
			child.printTree(indentLevel + 1)
			if idx != len(node.children)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Printf("%s", output)
		fmt.Print(")")
	// ISZERO
	case lexMap["iszero"]:
		fmt.Printf("%s", output)
		fmt.Println("IsZeroExp(")
		for idx, child := range node.children {
			child.printTree(indentLevel + 1)
			if idx != len(node.children)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Printf("%s", output)
		fmt.Print(")")
	// IF
	case lexMap["if"]:
		fmt.Printf("%s", output)
		fmt.Println("IfExp(")
		for idx, child := range node.children {
			child.printTree(indentLevel + 1)
			if idx != len(node.children)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Printf("%s", output)
		fmt.Print(")")
	// LET
	case lexMap["let"]:
		fmt.Printf("%s", output)
		fmt.Println("LetExp(")
		for i := 0; i < len(node.children); i++ {
			if i == 0 {
				fmt.Printf("%s", output)
				fmt.Print("\t")
				fmt.Print("\"" + node.children[i].contents + "\"")
			} else {
				node.children[i].printTree(indentLevel + 1)
			}
			if i != len(node.children)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Printf("%s", output)
		fmt.Print(")")
	// UNKNOWN
	default:
		fmt.Println("Error: Could not print Parse Tree.")
		os.Exit(1)
	}
}
