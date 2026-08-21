package credshelper

import (
	"errors"
	"testing"
)

type fakeForceRenew struct {
	forced  int
	renewed int
	err     error
}

func (f *fakeForceRenew) GetCredentials() (string, string) { return "user", "pass" }
func (f *fakeForceRenew) CheckAndRenew() bool              { f.renewed++; return false }
func (f *fakeForceRenew) ForceRenew() error                { f.forced++; return f.err }

type fakeNoForce struct {
	renewed int
}

func (f *fakeNoForce) GetCredentials() (string, string) { return "user", "pass" }
func (f *fakeNoForce) CheckAndRenew() bool              { f.renewed++; return true }

// TestForceRenewDelegates verifies the wrapper calls the wrapped helper's
// ForceRenew when it implements ForceRenewer.
func TestForceRenewDelegates(t *testing.T) {
	inner := &fakeForceRenew{}
	w := WrapSimpleCredsHelperWithRenew(inner)

	if err := w.ForceRenew(); err != nil {
		t.Fatalf("ForceRenew returned error: %v", err)
	}
	if inner.forced != 1 {
		t.Errorf("wrapped ForceRenew called %d times, want 1", inner.forced)
	}
	if inner.renewed != 0 {
		t.Errorf("CheckAndRenew called %d times, want 0", inner.renewed)
	}
}

// TestForceRenewPropagatesError verifies the wrapper propagates the wrapped
// helper's ForceRenew error.
func TestForceRenewPropagatesError(t *testing.T) {
	wantErr := errors.New("vault down")
	inner := &fakeForceRenew{err: wantErr}
	w := WrapSimpleCredsHelperWithRenew(inner)

	if err := w.ForceRenew(); !errors.Is(err, wantErr) {
		t.Errorf("ForceRenew error = %v, want %v", err, wantErr)
	}
}

// TestForceRenewFallsBackToCheckAndRenew verifies that when the wrapped helper
// does not implement ForceRenewer, ForceRenew falls back to CheckAndRenew.
func TestForceRenewFallsBackToCheckAndRenew(t *testing.T) {
	inner := &fakeNoForce{}
	w := WrapSimpleCredsHelperWithRenew(inner)

	if err := w.ForceRenew(); err != nil {
		t.Fatalf("ForceRenew returned error: %v", err)
	}
	if inner.renewed != 1 {
		t.Errorf("CheckAndRenew called %d times, want 1", inner.renewed)
	}
}
