package rorhealth

import (
	"context"
	"net/http"

	"github.com/dotse/go-health"
)

const (
	// StatusPass is ‘pass’.
	StatusPass Status = iota
	// StatusWarn is ‘warn’.
	StatusWarn Status = iota
	// StatusFail is ‘fail’.
	StatusFail Status = iota
)

// WrappedChecker is a wrapper that adapts the Checker interface to conform to health.Checker.
type WrappedChecker struct {
	Checker CheckerWithoutContext
}

// CheckHealth implements the health.Checker interface by adding context support.
func (wc *WrappedChecker) CheckHealth(ctx context.Context) []Check {
	return wc.Checker.CheckHealthWithoutContext()
}

// WrapChecker wraps a Checker to conform to the health.Checker interface.
func WrapChecker(checker CheckerWithoutContext) *WrappedChecker {
	return &WrappedChecker{Checker: checker}
}

type CheckerWithoutContext interface {
	CheckHealthWithoutContext() (checks []Check)
}
type Checker interface {
	CheckHealth(ctx context.Context) (checks []Check)
}

type Check = health.Check

type Status = health.Status

// Deprecated: Use Register instead.
// RegisterWithoutContext registers a health checker with the given name.
// It wraps the provided Checker to conform to the health.Checker interface.
func RegisterWithoutContext(name string, checker CheckerWithoutContext) {
	ctx := context.TODO()
	wrappedChecker := WrapChecker(checker)
	health.Register(ctx, name, wrappedChecker)
}

// Register registers a health checker with the given name and context.
// It wraps the provided Checker to conform to the health.Checker interface.
func Register(ctx context.Context, name string, checker Checker) health.Registered {
	return health.Register(ctx, name, checker)
}

func GetHttpHandler(w http.ResponseWriter, req *http.Request) {
	health.HandleHTTP(w, req)
}
