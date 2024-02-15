package applyapproval

import (
	"fmt"
)

// get permission to run by typing yes or no
func TfApplyApproval() bool {
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
