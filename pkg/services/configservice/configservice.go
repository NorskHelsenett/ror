package configservice

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"
)

func ConfigService(resourceConfig *ResourceConfig, ctx context.Context) (*ResourceConfig, error) {

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	id := identity.GetId()

	var filtered []ResourceConfigData

	for _, data := range r.Spec.Data {
		if data.Filter != string(identity.Type) {
			continue
		}
		tmpl, err := template.New("value").Delims("{{", "}}").Parse(data.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template %q: %w", data.Value, err)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, map[string]string{"id": id}); err != nil {
			return nil, fmt.Errorf("failed to execute template %q: %w", data.Value, err)
		}
		data.Value = buf.String()

		var noe map[string]interface{}

		switch data.Source {
		 case "vault":
			// get from vault
			// for now we just set it to a dummy value
		default:
			// if no source is specified, we just use the value from the template
			
		


		filtered = append(filtered, data)
	}
	r.Spec.Data = filtered

	return r, nil
}

