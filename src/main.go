package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	// require an argument
	if len(os.Args) == 1 {
		fmt.Printf("Please provide an expression to evaluate\n")
		os.Exit(1)
	}

	expr := os.Args[1]

	// create terraform manifest
	d1 := []byte(fmt.Sprintf("output value { value=%s }", expr))
	err := ioutil.WriteFile("temp.tf", d1, 0644)

	if err != nil {
		panic(err)
	}

	// run terraform command
	output, err := exec.Command("terraform", "apply").CombinedOutput()

	defer clean()

	if err != nil {
		fmt.Print(string(output))
		os.Exit(2)
	}

	stdout, err := exec.Command("terraform", "output", "value").CombinedOutput()

	if err != nil {
		fmt.Print(stdout)
		log.Fatal(err)
		panic(err)
	}

	fmt.Print(string(stdout))
}

func clean() {
	// cleanup directory
	os.Remove("temp.tf")
	os.Remove("terraform.tfstate")
	os.Remove("terraform.tfstate.backup")
}