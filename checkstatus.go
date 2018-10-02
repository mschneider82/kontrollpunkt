package main

import "strings"

// CheckStatus defines the status of a check
type CheckStatus int

const (
	OK CheckStatus = iota + 1
	Warn
	Error
)

func (cs CheckStatus) String() string {
	if cs == Error {
		return "error"
	}
	if cs == Warn {
		return "warn"
	}
	return "ok"
}

// NewCheckStatus string status to type CheckStatus
func NewCheckStatus(s string) CheckStatus {
	if strings.ToLower(s) == "warn" {
		return Warn
	}
	if strings.ToLower(s) == "error" {
		return Error
	}
	return OK
}
