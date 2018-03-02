package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

type Hello interface {
	Hello(name string) string
}

type Base struct {
	Type string `json:"type"`
}

type Loader struct {
	impl Hello
}

func (s Loader) Hello(name string) string {
	return s.impl.Hello(name)
}

func (b *Loader) UnmarshalJSON(raw []byte) error {
	typ := gjson.Get(string(raw), "type")
	switch typ.Str {
	case "prefix":
		pre := Prefix{}
		err := json.Unmarshal(raw, &pre)
		if err != nil {
			return err
		}
		b.impl = pre
		return nil
	case "suffix":
		suf := Suffix{}
		err := json.Unmarshal(raw, &suf)
		if err != nil {
			return err
		}
		b.impl = suf
		return nil
	}
	return errors.New("failed to unmarshal, unknown type")
}

type Prefix struct {
	Base
	Prefix string `json:"prefix"`
}

func (p Prefix) Hello(name string) string {
	return fmt.Sprintf("Hello %s %s", p.Prefix, name)
}

type Suffix struct {
	Base
	Suffix string `json:"suffix"`
}

func (s Suffix) Hello(name string) string {
	return fmt.Sprintf("Hello %s %s", name, s.Suffix)
}

func main() {
	x := []Loader{}
	err := json.Unmarshal([]byte(conf), &x)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(x)
	for _, y := range x {
		fmt.Println(y.Hello("Alan"))
	}
}

const conf = `[{
	"type": "prefix",
	"prefix": "Sir"
},
{
	"type": "suffix",
	"suffix": "Braithwaite"
}]`
