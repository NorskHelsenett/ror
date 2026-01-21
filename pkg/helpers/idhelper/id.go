package idhelper

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
)

// GetIdentifier returns a identifier from a name
//
// Parameters:
//
// - name: name to be turned into an identifier
func GetIdentifier(name string) string {
	idpostfix := stringhelper.RandomString(4, stringhelper.StringTypeClusterId)
	identifier := fmt.Sprintf("%s-%s", name, idpostfix)
	return identifier
}
