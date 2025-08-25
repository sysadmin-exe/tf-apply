package main

import (
	"tf-apply/actions/apply"
	"tf-apply/actions/plan"
	"tf-apply/helpers/getuserinput"
	"tf-apply/helpers/greetuser"
)

func main() {
	greetuser.GreetUser()
	application, resourceType, resourceCount, tfAction := getuserinput.GetUserInput()
	debugEnabled := true
	if tfAction == "plan" {
		plan.TfPlan(application, resourceType, resourceCount, debugEnabled)
	}
	if tfAction == "apply" {
		apply.TfApply(application, resourceType, resourceCount, debugEnabled)
	}
}
