package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/Bogdan-Zinovij/SA2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	//Only one of the input options can be used

	outputFile = flag.String("o", "", "File to write the result to")
	//If outputFile is not specified, print the result to standard output

	reader io.Reader
	writer io.Writer
)

func checkInputFlags(inputExp *string, inputFile *string) error {
	if *inputExp == "" && *inputFile == "" {
		return errors.New("expression is not set")
	}

	if *inputExp != "" && *inputFile != "" {
		return errors.New("only one of the input options can be used")
	}

	return nil
}

func getFileOrCreateIfNotExist(filename string) (os.File, error) {
	file, err := os.Open(*outputFile)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(*outputFile)
	}
	return *file, err
}

func main() {
	flag.Parse()
	err := checkInputFlags(inputExpression, inputFile)
	if err != nil {
		panic(err)
	}

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		inputF, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		reader = inputF
		defer inputF.Close()
	}

	if *outputFile == "" {
		writer = os.Stdout
	} else {
		outputF, err := getFileOrCreateIfNotExist(*outputFile)
		if err != nil {
			panic(err)
		}
		writer = &outputF
		defer outputF.Close()
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	e := handler.Compute()
	if e != nil {
		panic(e)
	}

}
