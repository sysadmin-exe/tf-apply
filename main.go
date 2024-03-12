package main

import (
	"tf-apply/actions/apply"
	"tf-apply/actions/plan"
	"tf-apply/helpers/getuserinput"
	"tf-apply/helpers/greetuser"
)

func main() {
	greetuser.GreetUser()
	resourceType, resourceCount, tfAction := getuserinput.GetUserInput()
	debugEnabled := false
	if tfAction == "plan" {
		plan.TfPlan(resourceType, resourceCount, debugEnabled)
	}
	if tfAction == "apply" {
		apply.TfApply(resourceType, resourceCount, debugEnabled)
	}
}
