package apply

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"tf-apply/helpers/applyapproval"
	"tf-apply/helpers/printwithtimestamp"
)

// declare variables as struct
// type userInput struct {
// 	resourceType  string
// 	resourceCount uint
// 	tfAction      string
// }

// var resourceDetails = make([]userInput, 0)

// run terraform plan, apply and show output
func TfApply(resourceType string, resourceCount uint, debugEnabled bool) {
	var resourceCountAsStr string

	printwithtimestamp.PrintWithTimestamp(fmt.Sprintf("Planning to create %v instance of %v...\n", resourceCount, resourceType))

	resourceCountAsStr = strconv.Itoa(int(resourceCount))
	os.Setenv("TF_VAR_resource_count", resourceCountAsStr)
	os.Setenv("TF_VAR_resource_name", resourceType)

	// Run `terraform init`
	initCmd := exec.Command("terraform", "-chdir=terraform-resources", "init")
	initStdout, _ := initCmd.StdoutPipe()
	initStderr, _ := initCmd.StderrPipe()
	initCmd.Start()
	if debugEnabled {
		go func() {
			io.Copy(os.Stdout, initStdout)
		}()
	}

	go func() {
		io.Copy(os.Stderr, initStderr)
	}()

	if err := initCmd.Wait(); err != nil {
		printwithtimestamp.PrintWithTimestamp(fmt.Sprintf("Error running terraform init: %s", err))
		os.Exit(1)
	}
	printwithtimestamp.PrintWithTimestamp("Terraform init successful")

	// Run `terraform plan`
	planCmd := exec.Command("terraform", "-chdir=terraform-resources", "plan")
	planStdout, _ := planCmd.StdoutPipe()
	planStderr, _ := planCmd.StderrPipe()
	planCmd.Start()

	if debugEnabled {
		go func() {
			io.Copy(os.Stdout, planStdout)
		}()
	}

	go func() {
		io.Copy(os.Stderr, planStderr)
	}()

	if err := planCmd.Wait(); err != nil {
		printwithtimestamp.PrintWithTimestamp(fmt.Sprintf("Error running terraform plan: %s", err))
		os.Exit(1)
	}
	printwithtimestamp.PrintWithTimestamp("Terraform plan successful")

	applyChanges := applyapproval.TfApplyApproval()
	if applyChanges {
		printwithtimestamp.PrintWithTimestamp(fmt.Sprintf("Creating %v instance of %v...\n", resourceCount, resourceType))

		// Run `terraform apply`
		applyCmd := exec.Command("terraform", "-chdir=terraform-resources", "apply", "-auto-approve")
		applyStdout, _ := applyCmd.StdoutPipe()
		applyStderr, _ := applyCmd.StderrPipe()
		applyCmd.Start()

		if debugEnabled {
			go func() {
				io.Copy(os.Stdout, applyStdout)
			}()
		}
		go func() {
			io.Copy(os.Stderr, applyStderr)
		}()

		if err := applyCmd.Wait(); err != nil {
			printwithtimestamp.PrintWithTimestamp(fmt.Sprintf("Error running terraform apply: %s", err))
			os.Exit(1)
		}
		printwithtimestamp.PrintWithTimestamp("Terraform apply successful")
	}
}
