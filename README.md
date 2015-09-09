
# go-number-scrambler [![Build Status](https://travis-ci.org/delaemon/go-number-scrambler.svg?branch=master)](https://travis-ci.org/delaemon/go-number-scrambler) [![GoCover](http://gocover.io/_badge/github.com/delaemon/go-number-scrambler)](http://gocover.io/github.com/delaemon/go-number-scrambler) [![GoDoc](https://godoc.org/github.com/delaemon/go-number-scrambler?status.png)](https://godoc.org/github.com/delaemon/go-number-scrambler)

##Description
Very simple library to generate reversible integer that is scrambled using the odd and inverse.

##Installation
This package can be installed with the go get command:
```sh
go get github.com/delaemon/go-number-scrambler
```

##Usage
```
package main

import (
	"github.com/delaemon/go-number-scrambler"
)

func main() {
	salt, inverse := scramble.GenerateSalt()
	in := 123
	scrambled, _ := scramble.Scrabmble(in, salt, inverse)
	out, _ := scramble.Scrabmble(scrambled, salt, inverse)
	if in == out {
		// success
	}
}
```
