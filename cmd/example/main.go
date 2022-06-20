package main

import (
	"errors"
	"flag"
	"io"
	"log"
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
		return errors.New("expression is not specified")
	}

	if *inputExp != "" && *inputFile != "" {
		return errors.New("only one of the input options can be used")
	}

	return nil
}

func main() {
	flag.Parse()
	inputFlagsErr := checkInputFlags(inputExpression, inputFile)
	if inputFlagsErr != nil {
		log.Fatalf("Wrong flags: %s", inputFlagsErr)
	}

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		inputF, openFileErr := os.Open(*inputFile)
		if openFileErr != nil {
			log.Fatalf("Cannot open specified file: %s", openFileErr)
		}
		reader = inputF
		defer inputF.Close()
	}

	if *outputFile == "" {
		writer = os.Stdout
	} else {
		createdFile, createFileErr := os.Create(*outputFile)
		if createFileErr != nil {
			log.Fatalf("Cannot create file: %s", createFileErr)
		}
		writer = createdFile
		defer createdFile.Close()
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	computeHandlerErr := handler.Compute()
	if computeHandlerErr != nil {
		log.Fatalf("Error during computing: %s", computeHandlerErr)
	}
}
