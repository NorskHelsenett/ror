# NHN-ROR-API

Webapi made with Golang and Gin webapi framework

# Prerequisites

- Golang 1.20.x https://go.dev

# Get started

Bash commands is from `<repo root>/src/backend/ror-api/`

Download dependencies:

```bash
go get ./...
```

Start webapi

```bash
go run main.go
```

Or
Start the `Debug ROR-Api` debugger config from VS Code

# Generate swagger docs:

Foreach endpoint function, you must add comments for it to show in generated openapi spec

ex:

```go

// @Summary 	Create cluster
// @Schemes
// @Description Create a cluster
// @Tags 		cluster
// @Accept 		application/json
// @Produce 	application/json
// @Success 	200 {object} responses.ClusterResponse
// @Failure 	403  {string}  Forbidden
// @Failure 	401  {string}  Unauthorized
// @Failure 	500  {string}  Failure message
// @Router		/v1/cluster [post]
// @Security	ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		...
	}
}

```

[Examples of annotations](https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html)

To generate new swagger you need to install a cmd called `swag` (https://github.com/swaggo/swag):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

(and remember to set `<userprofile>\go\bin` in PATH to terminal)

And run this command from `ror-api` root:

```bash
 swag init -g cmd/api/main.go --parseDependency --output cmd/api/docs
```

the folder `docs` and `docs\swagger.json` and `docs\swagger.yaml` is updated/created
