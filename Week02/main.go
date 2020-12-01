package main

import (
	"fmt"
	"week02/model"

	"github.com/pkg/errors"
)

func main() {
	users, err := model.QueryUsers()
	if err != nil {
		fmt.Printf("oh, %+v happens %+v ", errors.Cause(err), err)
	}
	fmt.Printf("User data: %+v", users)
}
