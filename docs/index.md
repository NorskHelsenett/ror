<img src="images/ror.svg" width="200" height="100" alt="ror logo">

# Release Operate Report

ROR is a multi cluster management tool aimed to streamline the process of creating kubernetes clusters on any provider or architecture.
ROR is developed by Norsk Helsenett SF but we move to remove any references and internal logic related to the organization. Our aim is to release the project as an open source project.

## API driven

ROR is API driven with provided web and cli clients. ROR leverages the Kubernetes API definition extending it with additional metadata and resources.

## Provider agnostic but extensible

ROR is aimed to be provider/cloud agnostic but stil extensible to provide provider specific functionality by the use of microservices.

## Distributed model

ROR relies on a distributed model using agents in each cluster. This model ensures that ROR can't be used to access the cluster directly.

## Development values

- Support simple primitives first then extend support if needed.
- Collect only needed data, scope the datamodel to suit our need.

## Features

For coming features, checkout the roadmap (here)[./roadmap/index.md]
