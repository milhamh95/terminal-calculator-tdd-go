package main

import (
	"flag"
	"log"
	"terminal-calculator-tdd-go/calculator"
	"terminal-calculator-tdd-go/input"
)

func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(
		engine.GetNumOperands(),
		engine.GetValidOperators(),
	)

	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
