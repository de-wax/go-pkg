// Only for testinf my Packages
package main

import (
	// Standard Packages
	"fmt"
	"log"

	// Self-Written Packages
	"github.com/de-wax/go-pkg/dewpoint"
)

func main() {
	// Test of Dewpoint Package, Result should be 12.35
	DP, err := dewpoint.Calculate(21.5, 56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DP)
}
