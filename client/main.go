package main

import (
	"email-center/center"
)

func main() {
	result, _ := center.CreateEstimate()
	result.AuditAllEmailItems()
}
