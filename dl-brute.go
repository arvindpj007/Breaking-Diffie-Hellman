package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

//Funciton to throw an error when the input CLI has missing/wrong parameters
func missingParametersError() {

	fmt.Println("ERROR: Parameters missing!")
	fmt.Println("HELP:")
	fmt.Println("./dl-brute <filename for input>")

}

//Funciton to setup the CLI
func setupCLI() string {

	if len(os.Args) < 2 {

		missingParametersError()
		os.Exit(1)
	}

	inputFile := os.Args[1]
	return inputFile

}

//Function to get the binary value from the given input file and returns value
func getInputText(inputText string) string {

	file, err := os.Open(inputText)
	if err != nil {
		log.Fatal(err)
	}

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(dataBytes)

	// fmt.Println(binaryText)
	return text
}

//
func getParameters(inputFile string) (*big.Int, *big.Int, *big.Int) {

	P := new(big.Int)
	G := new(big.Int)
	H := new(big.Int)

	fileBobText := getInputText(inputFile)
	fileBobText = fileBobText[2 : len(fileBobText)-2]
	fileBobTextSplit := strings.Split(fileBobText, ",")
	PText := fileBobTextSplit[0]
	GText := fileBobTextSplit[1]
	HText := fileBobTextSplit[2]

	P.SetString(PText, 10)
	G.SetString(GText, 10)
	H.SetString(HText, 10)

	return P, G, H
}

func bruteForce(G, H, P *big.Int, i int64) {

	X := new(big.Int).SetInt64(2)
	one := new(big.Int).SetInt64(1)
	two := new(big.Int).SetInt64(2)
	HPrime := new(big.Int)
	startIter := new(big.Int)
	endIter := new(big.Int)
	start := new(big.Int)
	end := new(big.Int)

	switch i {
	case 1:
		startIter.SetInt64(1)
		endIter.SetInt64(30)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	case 2:
		startIter.SetInt64(30)
		endIter.SetInt64(34)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	case 3:
		startIter.SetInt64(34)
		endIter.SetInt64(36)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	case 4:
		startIter.SetInt64(36)
		endIter.SetInt64(37)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	case 5:
		startIter.SetInt64(37)
		endIter.SetInt64(38)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	case 6:
		startIter.SetInt64(38)
		endIter.SetInt64(39)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	case 7:
		startIter.SetInt64(39)
		endIter.SetInt64(40)
		start.Exp(two, startIter, nil)
		end.Exp(two, endIter, nil)
	}

	// fmt.Println("start: ", start)
	// fmt.Println("end: ", end)

	for X = start; X.Cmp(end) != 0; X.Add(X, one) {

		// Xiter.new(big.int).SetString("125217870160", 10)
		Xiter := new(big.Int).Mod(X, P)
		// if i == 4 {
		// 	fmt.Println("Xiter: ", Xiter)
		// }
		HPrime.Exp(G, Xiter, P)
		// fmt.Println("iterations: ", X)
		if HPrime.Cmp(H) == 0 {
			fmt.Println(X)
			os.Exit(0)
			// fmt.Println("answer found: ", time.Now())
		}

	}
}

//Main Function
func main() {

	inputFile := setupCLI()

	P, G, H := getParameters(inputFile)

	// fmt.Println("P,G,H: ", P, G, H)

	go bruteForce(G, H, P, 1)
	go bruteForce(G, H, P, 2)
	go bruteForce(G, H, P, 3)
	go bruteForce(G, H, P, 4)
	go bruteForce(G, H, P, 5)
	go bruteForce(G, H, P, 6)
	go bruteForce(G, H, P, 7)

	// fmt.Println("main started", time.Now())

	// two := new(big.Int).SetInt64(2)
	// startIter := new(big.Int)
	// endIter := new(big.Int)
	// start := new(big.Int)
	// end := new(big.Int)

	// startIter.SetInt64(37)
	// endIter.SetInt64(36)
	// start.Exp(two, startIter, nil)
	// end.Exp(two, endIter, nil)

	// fmt.Println("start ", start)
	// fmt.Println("end ", end)

	time.Sleep(300 * time.Minute)
	// fmt.Println("main terminated", time.Now())
}
