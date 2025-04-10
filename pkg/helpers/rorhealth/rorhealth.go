package rorhealth

import (
	"context"

	newhealth "github.com/dotse/go-health"
)

// WrappedChecker is a wrapper that adapts the Checker interface to conform to newhealth.Checker.
type WrappedChecker struct {
	Checker Checker
}

// CheckHealth implements the newhealth.Checker interface by adding context support.
func (wc *WrappedChecker) CheckHealth(ctx context.Context) []newhealth.Check {
	return wc.Checker.CheckHealth()
}

// WrapChecker wraps a Checker to conform to the newhealth.Checker interface.
func WrapChecker(checker Checker) *WrappedChecker {
	return &WrappedChecker{Checker: checker}
}

type Checker interface {
	CheckHealth() (checks []newhealth.Check)
}

func Register(name string, checker Checker) {
	ctx := context.TODO()
	wrappedChecker := WrapChecker(checker)
	newhealth.Register(ctx, name, wrappedChecker)
}

func RegisterWithContext(ctx context.Context, name string, checker newhealth.Checker) newhealth.Registered {
	return newhealth.Register(ctx, name, checker)
}
