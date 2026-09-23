package rorcontext

import (
	"context"
	"reflect"
	"testing"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

func TestGetIdentityFromRorContext(t *testing.T) {
	t.Run("returns identity when present in context", func(t *testing.T) {
		expected, err := identitymodels.NewUserIdentity(identitymodels.AuthInfo{}, "test@example.com", "Test User", nil, nil)
		if err != nil {
			t.Fatalf("build user identity: %v", err)
		}
		ctx := context.WithValue(context.Background(), identitymodels.ContexIdentity, expected)

		got, err := GetIdentityFromRorContext(ctx)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		// DeepEqual covers the unexported state too, so the whole identity must
		// survive the round-trip, not just the fields a getter happens to expose.
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %+v, got %+v", expected, got)
		}
	})

	t.Run("returns cluster identity when present in context", func(t *testing.T) {
		expected, err := identitymodels.NewClusterIdentity(identitymodels.AuthInfo{}, "cluster-1", "uid-1")
		if err != nil {
			t.Fatalf("build cluster identity: %v", err)
		}
		ctx := context.WithValue(context.Background(), identitymodels.ContexIdentity, expected)

		got, err := GetIdentityFromRorContext(ctx)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %+v, got %+v", expected, got)
		}
	})

	t.Run("returns error when identity is missing", func(t *testing.T) {
		ctx := context.Background()

		got, err := GetIdentityFromRorContext(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !reflect.DeepEqual(got, identitymodels.Identity{}) {
			t.Fatalf("expected zero-value identity, got %+v", got)
		}
	})
}

func TestMustGetIdentityFromRorContext(t *testing.T) {
	t.Run("returns identity when present in context", func(t *testing.T) {
		expected, err := identitymodels.NewServiceIdentity(identitymodels.AuthInfo{}, "svc-1")
		if err != nil {
			t.Fatalf("build service identity: %v", err)
		}
		ctx := context.WithValue(context.Background(), identitymodels.ContexIdentity, expected)

		got := MustGetIdentityFromRorContext(ctx)
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %+v, got %+v", expected, got)
		}
	})

	t.Run("panics when identity is missing", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic, got nil")
			}
			if r != "Faild to get identity" {
				t.Fatalf("unexpected panic value: %v", r)
			}
		}()

		_ = MustGetIdentityFromRorContext(context.Background())
	})
}
