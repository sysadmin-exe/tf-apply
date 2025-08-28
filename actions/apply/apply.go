package apply

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"tf-apply/actions/plan"
	"tf-apply/helpers/applyapproval"
	"tf-apply/helpers/printwithtimestamp"
)

// run terraform plan, apply and show output
func TfApply(resourcesList string, debugEnabled bool) {
	printwithtimestamp.PrintWithTimestamp("Planning to create resources\n")

	// Set environment variables for Terraform
	os.Setenv("TF_VAR_resources_list", resourcesList)

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

	// Reuse the plan logic from actions/plan
	plan.TfPlan(resourcesList, debugEnabled)

	applyChanges := applyapproval.TfApplyApproval()
	if applyChanges {
		printwithtimestamp.PrintWithTimestamp("Creating Resources\n")

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
