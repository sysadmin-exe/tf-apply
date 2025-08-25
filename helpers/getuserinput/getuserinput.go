package getuserinput

import (
	"fmt"
)

func GetUserInput() (application string, resourceType string, resourceCount uint, tfAction string) {
	fmt.Println("Application name: ")
	fmt.Scan(&application)

	fmt.Println("What resource do you want to create: ")
	fmt.Scan(&resourceType)

	fmt.Println("How many of this resource do you want to create: ")
	fmt.Scan(&resourceCount)

	fmt.Println("Plan or Apply: ")
	fmt.Scan(&tfAction)

	return application, resourceType, resourceCount, tfAction
}
