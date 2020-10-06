// This example shows loc2map's usage as a Go package
package main

import (
	"github.com/forrest321/loc2map"
	"log"
)

func main() {
	err := loc2map.Convert(30.330557, -86.164910, "grayton.png")
	if err != nil {
		log.Fatal(err)
	}
}
