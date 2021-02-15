package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hculpan/kablang/ast"
	"github.com/hculpan/kablang/executor"
	"github.com/hculpan/kablang/parser"
)

var outputAST bool = false
var outputSymbols bool = false

var inputFilename string
var inputFilenameBase string

func main() {
	if !processCommandLine() {
		return
	}

	fmt.Println("Kab Interpreter v0.1")
	lines, err := readInputFile(inputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputFilenameBase = inputFilename[:len(inputFilename)-(len(filepath.Ext(inputFilename)))]

	parser := parser.NewParser()
	program, errs := parser.Parse(lines)

	if len(errs) > 0 {
		fmt.Println("Errors reported:")
		for _, e := range errs {
			fmt.Println("    ", e)
		}
		return
	}

	if program == nil {
		fmt.Println("We have an invalid program node")
		return
	}

	if outputAST {
		outputASTToFile(inputFilenameBase, program)
	}

	if outputSymbols {
		outputSymbolsToFile(inputFilenameBase, program)
	}

	ex := executor.NewExecutor()
	ex.Execute(program)
}

func processCommandLine() bool {
	flag.BoolVar(&outputAST, "a", false, "Output AST")
	flag.BoolVar(&outputSymbols, "s", false, "Output symbols")
	flag.Parse()
	if inputFilename = flag.Arg(0); len(flag.Args()) != 1 || inputFilename == "" {
		fmt.Println("Error: Incorrect arguments")
		printHelp()
		return false
	}

	return true
}

func printHelp() {
	fmt.Println("  Usage:")
	fmt.Println("        kablang <options> <input-filename>")
	fmt.Println()
	fmt.Println("  Options:")
	fmt.Println("        -a    Output AST")
	fmt.Println("        -s    Output symbols")
	fmt.Println()
}

func readInputFile(filename string) ([]string, error) {
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

	return result, nil
}

func outputASTToFile(filenameBase string, program *ast.Program) {
	file, err := os.Create(filenameBase + ".kab-ast")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(program.AsString("") + "\n")
	if err != nil {
		panic(err)
	}
	writer.Flush()
}

func outputSymbolsToFile(filenameBase string, program *ast.Program) {
	file, err := os.Create(filenameBase + ".kab-symbols")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("Symbols:\n")
	for _, v := range program.BlockNode.Symbols {
		_, err = writer.WriteString(v.AsString("  ") + "\n")
	}
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
