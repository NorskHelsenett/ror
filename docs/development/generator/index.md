# Code generator

# What is it?

The code generator uses (Go templating)[https://pkg.go.dev/html/template] to generate code based on a minimum required code and templates to

# Where is it used?

## v1

For v1 resources (which you should NOT be adding new features to) it uses these templates:

```.bash .copy
pkg/apicontracts/apiresourcecontracts/resource_models_generated.go.tmpl
pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go.tmpl
```

## v2

For v2 resources it uses these templates:

```.bash .copy
pkg/rorresources/fromstruct.go.tmpl
pkg/rorresources/resource.go.tmpl
pkg/rorresources/rorkubernetes/k8s.go.tmpl
pkg/rorresources/rorkubernetes/k8s_test.go.tmpl
pkg/rorresources/rorkubernetes/k8s_new.go.tmpl
pkg/rorresources/rorkubernetes/resource_interfaces.go.tmpl
pkg/rorresources/rorkubernetes/resourcemodels_methods.go.tmpl
pkg/rorresources/rortypes/resource_models_input_filter.go.tmpl
```

## TypeScript

TypeScript code is generated based on the Go struct data in the previous mentioned areas.
The resulting code can be found here:

```.bash .copy
typescript/models/resources.ts
```

# How do you use it?

On more details on how to use the generator, see (here)[./using-generator.md].
