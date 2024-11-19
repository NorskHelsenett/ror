package apikeysservice

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/spf13/viper"
)

func Test_mustGetApikeySalt(t *testing.T) {
	tests := []struct {
		name string
		salt string
		want string
	}{
		{
			name: "Test salt",
			salt: "salt",
			want: "salt",
		},
	}
	for _, tt := range tests {
		viper.Set(configconsts.API_KEY_SALT, tt.salt)
		t.Run(tt.name, func(t *testing.T) {
			if got := mustGetApikeySalt(); got != tt.want {
				t.Errorf("mustGetApikeySalt() = %v, want %v", got, tt.want)
			}
		})
	}
	// test panic
	viper.Set(configconsts.API_KEY_SALT, "")
	defer func() { _ = recover() }()
	mustGetApikeySalt()
	t.Errorf("did not panic")

}
