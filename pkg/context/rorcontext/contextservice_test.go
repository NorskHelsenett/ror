package rorcontext

import (
	"context"
	"testing"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

func TestGetIdentityFromRorContext(t *testing.T) {
	t.Run("returns identity when present in context", func(t *testing.T) {
		expected := identitymodels.Identity{
			Type: identitymodels.IdentityTypeUser,
			User: &identitymodels.User{Email: "test@example.com"},
		}
		ctx := context.WithValue(context.Background(), identitymodels.ContexIdentity, expected)

		got, err := GetIdentityFromRorContext(ctx)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if got.Type != expected.Type {
			t.Fatalf("expected type %q, got %q", expected.Type, got.Type)
		}
		if got.User == nil || got.User.Email != expected.User.Email {
			t.Fatalf("expected user email %q, got %+v", expected.User.Email, got.User)
		}
	})

	t.Run("returns cluster identity when present in context", func(t *testing.T) {
		expected := identitymodels.Identity{
			Type:            identitymodels.IdentityTypeCluster,
			ClusterIdentity: &identitymodels.ServiceIdentity{Id: "cluster-1"},
		}
		ctx := context.WithValue(context.Background(), identitymodels.ContexIdentity, expected)

		got, err := GetIdentityFromRorContext(ctx)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if got.Type != expected.Type {
			t.Fatalf("expected type %q, got %q", expected.Type, got.Type)
		}
		if got.ClusterIdentity == nil || got.ClusterIdentity.Id != expected.ClusterIdentity.Id {
			t.Fatalf("expected cluster id %q, got %+v", expected.ClusterIdentity.Id, got.ClusterIdentity)
		}
	})

	t.Run("returns error when identity is missing", func(t *testing.T) {
		ctx := context.Background()

		got, err := GetIdentityFromRorContext(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if got != (identitymodels.Identity{}) {
			t.Fatalf("expected zero-value identity, got %+v", got)
		}
	})
}

func TestMustGetIdentityFromRorContext(t *testing.T) {
	t.Run("returns identity when present in context", func(t *testing.T) {
		expected := identitymodels.Identity{
			Type:            identitymodels.IdentityTypeService,
			ServiceIdentity: &identitymodels.ServiceIdentity{Id: "svc-1"},
		}
		ctx := context.WithValue(context.Background(), identitymodels.ContexIdentity, expected)

		got := MustGetIdentityFromRorContext(ctx)
		if got.Type != expected.Type {
			t.Fatalf("expected type %q, got %q", expected.Type, got.Type)
		}
		if got.ServiceIdentity == nil || got.ServiceIdentity.Id != expected.ServiceIdentity.Id {
			t.Fatalf("expected service id %q, got %+v", expected.ServiceIdentity.Id, got.ServiceIdentity)
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
