package main

import (
	"errors"
	"fmt"
)

// what satisfies the error interface
type what struct {
	w string
}

func (w what) Error() string {
	return "what:" + w.w
}

// foo wraps an error and returns what, but it's also an error
func foo(orig error) what {
	return what{w: "foo:" + orig.Error()}
}

func blah(orig error) (err error) {
	return errors.New("blah:" + orig.Error())
}

// chain chains error wrapping funcs
func chain(wraps ...func(error) error) func(error) error {
	return func(e error) error {
		for _, w := range wraps {
			e = w(e)
		}
		return e
	}
}

type errWrap func(error) error

type funcer func(error) what

// force foo to be seen as a "middleware"
func coerce(b funcer) errWrap {
	return func(e error) error {
		return error(b(e))
	}
}

func main() {
	call := chain(coerce(foo), blah)
	v := call(errors.New("base"))
	fmt.Println(v)
}
