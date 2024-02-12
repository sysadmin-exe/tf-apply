package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// declare variables as struct
// type userInput struct {
// 	resourceType  string
// 	resourceCount uint
// 	tfAction      string
// }

// var resourceDetails = make([]userInput, 0)

// Get input
func greetUsers() {
	fmt.Printf("You are now using TF-APPLY -  The best terraform orchestrator for non-tech users\n")
}

func getUserInput() (string, uint, string) {
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

func printWithTimestamp(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%v] %v\n", timestamp, message)
}

// run terraform plan and show output
func tfPlan(resourceType string, resourceCount uint) {
	var resourceCountAsStr string

	fmt.Printf("Planning to create %v instance of %v...\n", resourceCount, resourceType)

	resourceCountAsStr = strconv.Itoa(int(resourceCount))
	os.Setenv("TF_VAR_resource_count", resourceCountAsStr)
	os.Setenv("TF_VAR_resource_name", resourceType)

	// add logic for terraform plan
	// Run `terraform init`
	initCmd := exec.Command("terraform", "-chdir=terraform-resources", "init")
	initStdout, _ := initCmd.StdoutPipe()
	initStderr, _ := initCmd.StderrPipe()
	initCmd.Start()
	go func() {
		io.Copy(os.Stdout, initStdout)
	}()

	go func() {
		io.Copy(os.Stderr, initStderr)
	}()

	if err := initCmd.Wait(); err != nil {
		printWithTimestamp(fmt.Sprintf("Error running terraform init: %s", err))
		os.Exit(1)
	}
	printWithTimestamp("Terraform init successful")

	// Run `terraform plan`
	planCmd := exec.Command("terraform", "-chdir=terraform-resources", "plan")
	planStdout, _ := planCmd.StdoutPipe()
	planStderr, _ := planCmd.StderrPipe()
	planCmd.Start()

	go func() {
		io.Copy(os.Stdout, planStdout)
	}()

	go func() {
		io.Copy(os.Stderr, planStderr)
	}()

	if err := planCmd.Wait(); err != nil {
		printWithTimestamp(fmt.Sprintf("Error running terraform plan: %s", err))
		os.Exit(1)
	}
	printWithTimestamp("Terraform plan successful")
}

// run terraform plan, apply and show output
func tfApply(resourceType string, resourceCount uint) {
	var resourceCountAsStr string

	fmt.Printf("Planning to create %v instance of %v...\n", resourceCount, resourceType)

	resourceCountAsStr = strconv.Itoa(int(resourceCount))
	os.Setenv("TF_VAR_resource_count", resourceCountAsStr)
	os.Setenv("TF_VAR_resource_name", resourceType)

	// Run `terraform init`
	initCmd := exec.Command("terraform", "-chdir=terraform-resources", "init")
	initStdout, _ := initCmd.StdoutPipe()
	initStderr, _ := initCmd.StderrPipe()
	initCmd.Start()
	go func() {
		io.Copy(os.Stdout, initStdout)
	}()

	go func() {
		io.Copy(os.Stderr, initStderr)
	}()

	if err := initCmd.Wait(); err != nil {
		printWithTimestamp(fmt.Sprintf("Error running terraform init: %s", err))
		os.Exit(1)
	}
	printWithTimestamp("Terraform init successful")

	// Run `terraform plan`
	planCmd := exec.Command("terraform", "-chdir=terraform-resources", "plan")
	planStdout, _ := planCmd.StdoutPipe()
	planStderr, _ := planCmd.StderrPipe()
	planCmd.Start()

	go func() {
		io.Copy(os.Stdout, planStdout)
	}()

	go func() {
		io.Copy(os.Stderr, planStderr)
	}()

	if err := planCmd.Wait(); err != nil {
		printWithTimestamp(fmt.Sprintf("Error running terraform plan: %s", err))
		os.Exit(1)
	}
	printWithTimestamp("Terraform plan successful")

	applyChanges := tfApplyApproval()
	if applyChanges {
		fmt.Printf("Creating %v instance of %v...\n", resourceCount, resourceType)
		// Run `terraform apply`
		applyCmd := exec.Command("terraform", "-chdir=terraform-resources", "apply", "-auto-approve")
		applyStdout, _ := applyCmd.StdoutPipe()
		applyStderr, _ := applyCmd.StderrPipe()
		applyCmd.Start()

		go func() {
			io.Copy(os.Stdout, applyStdout)
		}()

		go func() {
			io.Copy(os.Stderr, applyStderr)
		}()

		if err := applyCmd.Wait(); err != nil {
			printWithTimestamp(fmt.Sprintf("Error running terraform apply: %s", err))
			os.Exit(1)
		}
		printWithTimestamp("Terraform apply successful")
	}
}

func main() {
	greetUsers()
	resourceType, resourceCount, tfAction := getUserInput()
	if tfAction == "plan" {
		tfPlan(resourceType, resourceCount)
	}
	if tfAction == "apply" {
		tfApply(resourceType, resourceCount)
	}
}
