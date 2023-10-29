package main

import (
	"fmt"
	"os"

	"github.com/shimizu/calculator"
	"github.com/shimizu/calculator/auth"
	"github.com/shimizu/calculator/internal/validation"
)

func main() {
	if !auth.Authenticate(os.Args[1], os.Args[2]) {
		fmt.Println("Authentication failed")
		os.Exit(1)
	}
	a, err := validation.ValidateInput(os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := validation.ValidateInput(os.Args[4])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(calculator.Add(a, b))
}
