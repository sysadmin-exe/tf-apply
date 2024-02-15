package printwithtimestamp

import (
	"fmt"
	"time"
)

func PrintWithTimestamp(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%v] %v\n", timestamp, message)
}
