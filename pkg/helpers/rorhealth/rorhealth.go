package rorhealth

import (
	"context"

	"github.com/dotse/go-health"
)

// WrappedChecker is a wrapper that adapts the Checker interface to conform to health.Checker.
type WrappedChecker struct {
	Checker Checker
}

// CheckHealth implements the health.Checker interface by adding context support.
func (wc *WrappedChecker) CheckHealth(ctx context.Context) []health.Check {
	return wc.Checker.CheckHealth()
}

// WrapChecker wraps a Checker to conform to the health.Checker interface.
func WrapChecker(checker Checker) *WrappedChecker {
	return &WrappedChecker{Checker: checker}
}

type Checker interface {
	CheckHealth() (checks []health.Check)
}

// Deprecated: Use RegisterWithContext instead.
// Register registers a health checker with the given name.
// It wraps the provided Checker to conform to the health.Checker interface.
func Register(name string, checker Checker) {
	ctx := context.TODO()
	wrappedChecker := WrapChecker(checker)
	health.Register(ctx, name, wrappedChecker)
}

// RegisterWithContext registers a health checker with the given name and context.
// It wraps the provided Checker to conform to the health.Checker interface.
func RegisterWithContext(ctx context.Context, name string, checker health.Checker) health.Registered {
	return health.Register(ctx, name, checker)
}
