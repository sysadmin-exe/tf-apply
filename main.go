package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// declare variables as struct
// type userInput struct {
// 	resourceType  string
// 	resourceCount uint
// }

// var resourceDetails = make([]userInput, 0)

// Get input
func greetUsers() {
	fmt.Printf("You are now using TF-APPLY -  The best terraform orchestrator for non-tech users\n")
}

func getUserInput() (string, uint) {
	var resourceType string
	var resourceCount uint

	fmt.Println("What resource do you want to create: ")
	fmt.Scan(&resourceType)

	fmt.Println("How many of this resource do you want to create: ")
	fmt.Scan(&resourceCount)

	return resourceType, resourceCount
}

// get permission to run by typing yes or no
func tfApplyApproval() bool {
	var applyApproval string
	var applyChanges bool
	fmt.Println("Are you sure you want to apply this change: ")
	fmt.Scan(&applyApproval)
	if applyApproval == "yes" {
		applyChanges = true
	} else {
		applyChanges = false
	}
	return applyChanges
}

// run terraform plan and apply and show use output
func tfRun() {
	var resourceCountAsStr string

	resourceType, resourceCount := getUserInput()
	fmt.Printf("Planning to create %v instance of %v...\n", resourceCount, resourceType)

	resourceCountAsStr = strconv.Itoa(int(resourceCount))
	os.Setenv("TF_VAR_resource_count", resourceCountAsStr)
	os.Setenv("TF_VAR_resource_name", resourceType)

	// add logic for terraform plan
	// Run `terraform init`
	initCmd := exec.Command("terraform", "-chdir=terraform-resources", "init")
	output, err := initCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running terraform init:", err)
		fmt.Println("Output:", string(output))
		return
	}
	fmt.Println("Terraform init successful")

	// Run `terraform plan`
	planCmd := exec.Command("terraform", "-chdir=terraform-resources", "plan")
	output, err = planCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running terraform plan:", err)
		fmt.Println("Output:", string(output))
		return
	}
	fmt.Println("Terraform plan successful")

	applyChanges := tfApplyApproval()
	if applyChanges {
		fmt.Printf("Creating %v instance of %v...\n", resourceCount, resourceType)
		// add logic for terraform apply
		// Run `terraform apply`
		applyCmd := exec.Command("terraform", "-chdir=terraform-resources", "apply", "-auto-approve")
		output, err = applyCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error running terraform apply:", err)
			fmt.Println("Output:", string(output))
			return
		}
		fmt.Println("Terraform apply successful")
	} else {
		fmt.Printf("Will not %v instance of %v. Exiting...\n", resourceCount, resourceType)
	}

}

// give results and output of resource created
// func tfApplyOutput() {

// }

func main() {
	greetUsers()
	tfRun()
}
