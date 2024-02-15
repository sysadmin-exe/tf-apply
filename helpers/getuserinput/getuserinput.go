package getuserinput

import (
	"fmt"
)

func GetUserInput() (string, uint, string) {
	var resourceType string
	var resourceCount uint
	var tfAction string

	fmt.Println("What resource do you want to create: ")
	fmt.Scan(&resourceType)

	fmt.Println("How many of this resource do you want to create: ")
	fmt.Scan(&resourceCount)

	fmt.Println("Plan or Apply: ")
	fmt.Scan(&tfAction)

	return resourceType, resourceCount, tfAction
}
