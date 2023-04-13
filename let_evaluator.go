package main

import (
	"fmt"
	"strconv"
)

/*
The evaluator maintains an environment that consists of keys (variables) and values (their values).
Entry into a let statement adds a new binding to the environment.
Exit from a let statement pops the environment. We can nest environments such that the variable declared with a first
expression is referenced within a second expression, so that environment holds effect. These will be maintained using
a stack and lookup table.
*/

// Binding - represents a var/value pair.
type Binding struct {
	varname string
	value   rune
}

func runEval() {
	var e []Binding
	result := evaluate(globalRoot, e)
	fmt.Println()
	fmt.Println()
	fmt.Print(Info("RESULT: "))
	fmt.Println(result)
}

// evaluate - Compute the expression result of the operation on variables within environments from parsed nodes.
func evaluate(localRoot astNode, e []Binding) rune {
	switch localRoot.ttype {
	case lexMap["let"]:
		newBindingList := []Binding{{
			localRoot.children[0].contents,
			evaluate(*localRoot.children[1], e)}}
		e = append(newBindingList, e...)
		return evaluate(*localRoot.children[2], e)
	case lexMap["identifier"]:
		for _, itemPair := range e {
			if itemPair.varname == localRoot.contents {
				return itemPair.value
			}
		}
	case lexMap["integer"]:
		integer, _ := strconv.Atoi(localRoot.contents)
		return rune(integer)
	case lexMap["minus"]:
		exp1 := evaluate(*localRoot.children[0], e)
		exp2 := evaluate(*localRoot.children[1], e)
		return exp1 - exp2
	case lexMap["iszero"]:
		exp1 := evaluate(*localRoot.children[0], e)
		if exp1 == 0 {
			return 1
		} else {
			return 0
		}
	case lexMap["if"]:
		exp1 := evaluate(*localRoot.children[0], e)
		if exp1 == 1 {
			return evaluate(*localRoot.children[1], e)
		} else if exp1 == 0 {
			return evaluate(*localRoot.children[2], e)
		}
	}
	// Get the final value of the expression at the root node of the tree.
	return e[0].value
}
