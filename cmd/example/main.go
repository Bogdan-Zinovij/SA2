package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File with expression to compute")
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
	
	fmt.Println(*inputExpression)
	fmt.Println(*inputFile)
	fmt.Println(*outputFile)


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

	  //			handler := &lab2.ComputeHandler{
		//      	Input: {construct io.Reader according the command line parameters},
		//				Output: {construct io.Writer according the command line parameters},
		//      }
		//      err := handler.Compute()
		//			if err != nil {
		//			panic(err)	
		//}

}

