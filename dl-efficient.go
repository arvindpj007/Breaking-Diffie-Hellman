package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

var giantMap = make(map[string]string)
var babyMap = make(map[string]string)

//Funciton to throw an error when the input CLI has missing/wrong parameters
func missingParametersError() {

	fmt.Println("ERROR: Parameters missing!")
	fmt.Println("HELP:")
	fmt.Println("./dl-efficient <filename for input>")

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

//Function to return the value from the given input file
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

	return text
}

// Function to return the values from the input string
func getParameters(inputFile string) (*big.Int, *big.Int, *big.Int) {

	P := new(big.Int)
	G := new(big.Int)
	H := new(big.Int)

	fileText := getInputText(inputFile)
	fileText = fileText[2 : len(fileText)-2]
	fileTextSplit := strings.Split(fileText, ",")
	PText := fileTextSplit[0]
	GText := fileTextSplit[1]
	HText := fileTextSplit[2]

	P.SetString(PText, 10)
	G.SetString(GText, 10)
	H.SetString(HText, 10)

	return P, G, H
}

func getM(P *big.Int) *big.Int {

	// one := new(big.Int).SetInt64(1)
	M := new(big.Int)

	// M.Sub(P, one)
	M.Sqrt(P)

	return M
}

func setupGiantMapping(G, M, P *big.Int) {

	i := new(big.Int)
	Gim := new(big.Int)
	Gm := new(big.Int)
	one := new(big.Int).SetInt64(1)
	limit := new(big.Int)

	limit.Div(P, M)
	limit.Add(limit, one)

	Gm.Exp(G, M, P)
	for i.SetInt64(0); i.Cmp(limit) != 0; i.Add(i, one) {

		Gim.Exp(Gm, i, P)
		giantMap[i.String()] = Gim.String()
	}
}

func setupBabyMapping(G, M, P, H *big.Int) {

	i := new(big.Int)
	Gi := new(big.Int)
	HGi := new(big.Int)
	one := new(big.Int).SetInt64(1)
	limit := new(big.Int)

	limit.Div(P, M)
	limit.Add(limit, one)

	for i.SetInt64(0); i.Cmp(limit) != 0; i.Add(i, one) {

		Gi.Exp(G, i, P)
		HGi.Mul(Gi, H)
		HGi.Mod(HGi, P)
		babyMap[HGi.String()] = i.String()
	}

}

func babyGiant(G, H, M, P *big.Int) {

	X := new(big.Int).SetInt64(0)
	Gqm := new(big.Int)
	Q := new(big.Int)
	R := new(big.Int)
	limit := new(big.Int)
	one := new(big.Int).SetInt64(1)

	limit.Div(P, M)
	limit.Add(limit, one)

	for Q.SetInt64(0); Q.Cmp(limit) != 0; Q.Add(Q, one) {

		// fmt.Println(Q)
		// fmt.Println(inverseMap[Q.String()])

		Gqm.SetString(giantMap[Q.String()], 10)

		if r, ok := babyMap[Gqm.String()]; ok {

			R.SetString(r, 10)
			X.Mul(Q, M)
			X.Sub(X, R)

			fmt.Println(X)
			os.Exit(0)

		}
	}
}

//Main Function
func main() {

	inputFile := setupCLI()

	P, G, H := getParameters(inputFile)

	M := getM(P)

	setupGiantMapping(G, M, P)
	setupBabyMapping(G, M, P, H)

	// fmt.Println(H)
	babyGiant(G, H, M, P)

}
