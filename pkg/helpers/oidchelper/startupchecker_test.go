package oidchelper

import (
	"context"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
)

func TestOidcStartupChecker_NoIssuersPassesImmediately(t *testing.T) {
	c := &oidcStartupChecker{expected: 0, deadline: time.Now().Add(time.Minute)}
	checks := c.CheckHealth(context.Background())
	if len(checks) != 1 || checks[0].Status != rorhealth.StatusPass {
		t.Fatalf("expected pass with no expected issuers, got %+v", checks)
	}
}

func TestOidcStartupChecker_FailsWhileLoadingBeforeDeadline(t *testing.T) {
	c := &oidcStartupChecker{expected: 2, deadline: time.Now().Add(time.Minute)}
	checks := c.CheckHealth(context.Background())
	if len(checks) != 1 || checks[0].Status != rorhealth.StatusFail {
		t.Fatalf("expected fail while loading before deadline, got %+v", checks)
	}
}

func TestOidcStartupChecker_PassesWhenAllLoaded(t *testing.T) {
	c := &oidcStartupChecker{expected: 2, deadline: time.Now().Add(time.Minute)}
	c.markLoaded()
	c.markLoaded()
	checks := c.CheckHealth(context.Background())
	if len(checks) != 1 || checks[0].Status != rorhealth.StatusPass {
		t.Fatalf("expected pass when all issuers loaded, got %+v", checks)
	}
}

func TestOidcStartupChecker_PassesAfterDeadline(t *testing.T) {
	c := &oidcStartupChecker{expected: 2, deadline: time.Now().Add(-time.Second)}
	checks := c.CheckHealth(context.Background())
	if len(checks) != 1 || checks[0].Status != rorhealth.StatusPass {
		t.Fatalf("expected pass after deadline even if not all loaded, got %+v", checks)
	}
}
