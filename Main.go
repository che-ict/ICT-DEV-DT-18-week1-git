package main

import "fmt"

func main() {
	// geeft alle even getallen weer tot en met 100

	fmt.Println("Alle even getallen tot en met 100: ")

	for i := 0; i <= 100; i += 2 {
		fmt.Printf("%v, ", i)
	}

	fmt.Println()

	// geeft alle oneven getallen weer tot 100
}
