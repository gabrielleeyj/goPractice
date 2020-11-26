package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	secretKey    string `json:"secret_key"`
	Name         string `json:"name,omitempty"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DontShowThis string `json:"-"`
}

func main() {
	// HTTP Client
	p := Person{
		secretKey:    "secret",
		Name:         "",
		Username:     "marcoshuck",
		Password:     "1234",
		DontShowThis: "value",
	}

	b, err := json.Marshal(p)
	if err != nil {
		panic("cannot marshal Person")
	}

	// HTTP Server

	fmt.Println("Output:", string(b))

	var result Person

	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("Bad request")
		return
	}

	fmt.Printf("%+v", result)

}
