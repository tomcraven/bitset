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

func outputBitset(b bitset.Bitset) {
	for i := uint(0); i < b.Size(); i++ {
		if b.Get(i) {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Printf("\n")
}

func main() {
	b := bitset.Create(8)
	outputBitset(b) // 00000000

	b.Set(1)
	b.Set(2)
	b.Set(7)
	outputBitset(b) // 01100001

	b.Clear(2)
	outputBitset(b) // 01000001
}
```
