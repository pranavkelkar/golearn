package brego

import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}

func Sum(a int, b int) int {
	return  a + b
}
