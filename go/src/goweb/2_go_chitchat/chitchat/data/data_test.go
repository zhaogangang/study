package data

import (
	"fmt"
	"testing"
)

//test data
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func setup() {
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll(0)
}

func Test_createUUID(t *testing.T) {
	fmt.Println(createUUID())
	fmt.Println(createUUID())
	fmt.Println(createUUID())
}
