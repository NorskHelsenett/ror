package apicontracts

import (
	"testing"
	"time"
)

func TestApiKey_IsExpired(t *testing.T) {
	type fields struct {
		Id          string
		Identifier  string
		DisplayName string
		Type        ApiKeyType
		ReadOnly    bool
		Expires     time.Time
		Created     time.Time
		LastUsed    time.Time
		Hash        string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test expired",
			fields: fields{
				Expires: time.Now().Add(-1 * time.Hour),
			},
			want: true,
		},
		{
			name: "Test not expired",
			fields: fields{
				Expires: time.Now().Add(1 * time.Hour),
			},
			want: false,
		},
		{
			name: "Test no expiry",
			fields: fields{
				Expires: time.Time{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := ApiKey{
				Id:          tt.fields.Id,
				Identifier:  tt.fields.Identifier,
				DisplayName: tt.fields.DisplayName,
				Type:        tt.fields.Type,
				ReadOnly:    tt.fields.ReadOnly,
				Expires:     tt.fields.Expires,
				Created:     tt.fields.Created,
				LastUsed:    tt.fields.LastUsed,
				Hash:        tt.fields.Hash,
			}
			if got := key.IsExpired(); got != tt.want {
				t.Errorf("ApiKey.IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
