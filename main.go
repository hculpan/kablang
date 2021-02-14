package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hculpan/kablang/executor"
	"github.com/hculpan/kablang/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect arguments")
		printHelp()
		return
	}

	lines, err := readInputFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	parser := parser.NewParser()
	program, errs := parser.Parse(lines)
	if len(errs) > 0 {
		fmt.Println("** We have errors")
		for _, e := range errs {
			fmt.Println("    ", e)
		}
		return
	}

	if program == nil {
		fmt.Println("We have an invalid program node")
		return
	}

	fmt.Println(program.AsString(""))

	fmt.Println("Starting execution")

	ex := executor.NewExecutor()
	ex.Execute(program)

	fmt.Println("Program processed successfully")
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("        kablang <input filename>")
	fmt.Println()
}

func readInputFile(filename string) ([]string, error) {
	fmt.Printf("Reading %s...", filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error")
		return nil, err
	}
	defer f.Close()

	result := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result = append(result, strings.Trim(scanner.Text(), " "))
	}

	fmt.Printf("%d lines read, done\n", len(result))

	return result, nil
}
