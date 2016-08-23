# Bitset

A simple bitset implementation in Go.

## Installation

```
go get github.com/tomcraven/bitset
```

## Example

```go
package main

import (
	"fmt"

	"github.com/tomcraven/bitset"
)

func settingAndClearing() {
	fmt.Println("  ** Setting and clearing **")

	b := bitset.Create(8)
	b.Output() // 00000000

	b.Set(1, 2, 7)
	b.Output() // 01100001

	b.Clear(2)
	b.Output() // 01000001

	b.SetAll()
	b.Output() // 11111111

	b.Clear(1, 2, 3, 4, 5)
	b.Output() // 10000011

	b.ClearAll()
	b.Output() // 00000000

	fmt.Println()
}

func slicing() {
	fmt.Println("  ** Slicing **")

	b := bitset.Create(8)
	b.Output() // 00000000

	slice := b.Slice(0, 4)
	b.Set(0, 1, 4, 5)
	b.Output()     // 11001100
	slice.Output() // 1100

	slice.SetAll()
	b.Output()     // 11111100
	slice.Output() // 1111

	slicedSlice := slice.Slice(2, 4)
	b.ClearAll()
	b.Output()           // 00000000
	slice.Output()       // 0000
	slicedSlice.Output() // 00

	slicedSlice.Set(0, 1)
	slicedSlice.Output() // 11
	slice.Output()       // 0011
	b.Output()           // 00110000

	fmt.Println()
}

func building() {
	fmt.Println("  ** Building **")
	b := bitset.Create(16)
	var output uint8

	b.BuildUint8(&output)
	b.Output()          // 0000000000000000
	fmt.Println(output) // 0

	b.Set(0, 1, 2, 3, 4, 5, 6, 7)
	b.Output() // 1111111100000000
	b.BuildUint8(&output)
	fmt.Println(output) // 255

	slice := b.Slice(4, 12)
	slice.Output() // 11110000
	slice.BuildUint8(&output)
	fmt.Println(output) // 15
}

func main() {
	settingAndClearing()
	slicing()
	building()

	/*
		Program output:

		** Setting and clearing **
		00000000
		01100001
		01000001
		11111111
		10000011
		00000000

		  ** Slicing **
		00000000
		11001100
		1100
		11111100
		1111
		00000000
		0000
		00
		11
		0011
		00110000

		  ** Building **
		0000000000000000
		0
		1111111100000000
		255
		11110000
		15
	*/
}

```
