package guid_test

import (
	"fmt"

	"github.com/beevik/guid"
)

// Parse a string containing a guid.
func ExampleParseString() {
	g, err := guid.ParseString("0e545c9c-f942-4988-4ab0-145274cfaded")
	if err != nil {
		fmt.Printf("Guid: %v\n", g)
	}
}
