package plan

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"tf-apply/helpers/printwithtimestamp"
)

// run terraform plan and show output
func TfPlan(resourcesList string, debugEnabled bool) {
	// Set environment variables for Terraform
	os.Setenv("TF_VAR_resources_list", resourcesList)

	// add logic for terraform plan
	// Run `terraform init`
	initCmd := exec.Command("terraform", "-chdir=terraform-resources/modules", "init")
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
	planCmd := exec.Command("terraform", "-chdir=terraform-resources/modules", "plan")
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
}
