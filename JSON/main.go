package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User to store user
type User struct {
	Name     string `json:"Name"`
	Password string `json:"-"`
	Handle   string `json:"handle"`
}

func main() {
	users := []User{
		{"Rahul Sachan", "abc123", "rahul4u4i"},
		{"Isha Singh", "a1b2b3", "isha4u4i"},
	}
	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println("There was an error.")
	}
	fmt.Println(string(bs))

	var data []User
	err = json.Unmarshal(bs, &data)
	if err != nil {
		fmt.Println("There was an error.")
	}
	fmt.Println(data)

	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("There was an error.")
	}
}
