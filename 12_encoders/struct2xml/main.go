package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	secretKey    string `xml:"secret_key"`
	Name         string `xml:"name"`
	Username     string `xml:"username"`
	Password     string `xml:"password"`
	DontShowThis string `xml:"-"`
}

func main() {
	p := Person{
		secretKey:    "secret",
		Name:         "",
		Username:     "marcoshuck",
		Password:     "1234",
		DontShowThis: "value",
	}

	b, err := xml.Marshal(p)
	if err != nil {
		panic("cannot marshal Person")
	}

	fmt.Println("Output:", string(b))

	var result Person

	err = xml.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("Bad request")
		return
	}

	fmt.Printf("%+v", result)

}
