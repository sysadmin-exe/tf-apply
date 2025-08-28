package main

import (
	"fmt"
	"os"
	"tf-apply/actions/apply"
	"tf-apply/actions/plan"
	"tf-apply/helpers/greetuser"
	"tf-apply/helpers/parser"
)

func tfActionArg() string {
	// Helper function to parse command line arguments
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "plan" || arg == "apply" {
			return arg
		}
	}
	return ""
}

func inventoryFile() string {
	// Helper function to get the inventory file argument
	if len(os.Args) > 2 {
		filePath := os.Args[2]
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Error: Inventory file '%s' does not exist.\n", filePath)
			os.Exit(1)
		}
		return filePath
	}
	return ""
}

func main() {
	tfAction := tfActionArg()
	file := inventoryFile()
	if tfAction != "plan" && tfAction != "apply" {
		println("Error: Invalid argument. Valid arguments are 'plan' or 'apply'.")
		return
	}
	greetuser.GreetUser()
	resources_list_string := parser.ToTerraform(file)
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
