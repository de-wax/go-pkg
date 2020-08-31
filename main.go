/*
Diese Datei dient lediglich dazu über Github Actions automatisierte Tests der Pakete laufen zu lassen.
This file is only used to run automated tests of the packages via Github Actions.
*/
package main

import (
	// Standard Packages
	"fmt"
	"log"

	// Self-Written Packages
	"github.com/de-wax/go-pkg/dewpoint"
)

func main() {
	// Test des Taupunkt-Pakets, das Ergebnis sollte 12.35 sein
	// Test of Dewpoint Package, Result should be 12.35
	DP, err := dewpoint.Calculate(21.5, 56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DP)
}
