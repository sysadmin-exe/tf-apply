package main

import (
	"fmt"
	"os"
	"tf-apply/actions/apply"
	"tf-apply/actions/plan"
	"tf-apply/helpers/greetuser"
	"tf-apply/helpers/parser"
)

func main() {
	var tfAction string
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "plan" || arg == "apply" {
			tfAction = arg
		} else {
			println("Error: Invalid argument. Valid arguments are 'plan' or 'apply'.")
			os.Exit(1)
		}
	} else {
		tfAction = "plan"
	}
	greetuser.GreetUser()
	resources_list_string := parser.ToTerraform("inventory.yaml")
	fmt.Printf("Parsed %d resources from inventory\n", len(resources_list_string))
	fmt.Printf("Resources: %+v\n", resources_list_string)
	debugEnabled := true
	if tfAction == "plan" {
		plan.TfPlan(resources_list_string, debugEnabled)
	}
	if tfAction == "apply" {
		apply.TfApply(resources_list_string, debugEnabled)
	}
}
