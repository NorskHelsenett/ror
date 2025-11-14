# Setting up repositories

To ease the development on any core components of (ROR)[../../components/core/index.md] it is highly recommended to setup a Go workspace.
For more information on Go workspaces see (here)[https://go.dev/doc/tutorial/workspaces].

# Prerequisites

- Golang SDK (For debugging and changing) https://go.dev

# Short version

Make a directory to keep any related repositories to the ROR development. Avoid mixing projects not related to ROR out of this repository, as Go workspaces and react weirdly to that.
The directory name is not important but keep it sensible.

```bash .copy
mkdir ror-ws # short for ror-workspace
```

Copy each repository into this workspace folder.

```bash .copy
git clone git@github.com:NorskHelsenett/ror.git
```

```bash .copy
git clone git@github.com:NorskHelsenett/ror-api.git
```

```bash .copy
git clone git@github.com:NorskHelsenett/ror-web.git
```

## Creating the workspace

Create the workspace by using to `go work init` command on one of the repositories.

```bash .copy
go work init ./ror
```

Then add the other repositories.

```bash .copy
go work use  ./ror-api
```

```bash .copy
go work use  ./ror-web
```

Now you repositories are ready for you development.
