package main

import "fmt"

var (
	StatusOnline     Status = "online"
	StatusOffline    Status = "offline"
	StatusOnVacation Status = "on-vacation"
	StatusAFK        Status = "afk"
)

type Status string

type User struct {
	Username string
	Password string
	Status   Status
}

func main() {
	u := User{
		Username: "marcoshuck",
		Password: "1234",
		Status:   StatusAFK,
	}

	fmt.Println("Username:", u.Username)
	fmt.Println("Password:", u.Password)
	fmt.Println("Status:", u.Status)
}
