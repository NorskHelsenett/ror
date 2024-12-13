# Getting started with ROR Client

## Prerequisites
    
- A ROR resourcedef - To create this see docs/ror-generator-client/docs/generator/getting-started.md

### Optional


## Getting started

### Configuration variables

First to implement it we need to add the rorclient to the project

```
go get "github.com/NorskHelsenett/ror/pkg/clients/rorclient"
```

And then add configuration variables for the following parameters:

The Url of the ROR instance you're communicating with
``BaseUrl``

The auth proivder from "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
which is either:
httpauthprovider.NewNoAuthProvider()
or an actual AuthProvider

TODO

``AuthProvider``

The current version of the client, for example 0.0.1:<commit>
``Version``

The name of the client, for example vpshereAgent
``Role``
