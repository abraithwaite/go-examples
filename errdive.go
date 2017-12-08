package main

import (
	"log"

	"github.com/pkg/errors"
)

func Dive(err error) error {
	const limit = 64
	ret := err
	for i := 0; i < limit; i++ {
		// if error is a Causer and that Cause()
		// has itself a stack trace, dive.
		if cs, ok := ret.(interface {
			Cause() error
		}); ok {
			if _, ok = cs.Cause().(interface {
				StackTrace() errors.StackTrace
			}); ok {
				ret = cs.Cause()
			} else {
				break
			}
		} else {
			break
		}
	}
	return ret
}

func nah() error {
	return errors.New("whaaat are those?!?")
}

func blah() error {
	err := nah()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func hah() error {
	err := blah()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func main() {
	err := hah()
	err = Dive(err)
	log.Printf("%+v", err)
}
