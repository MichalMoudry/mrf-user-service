<p align="center">
    <img src="./assets/mrf_logo.svg" alt="Microservices Reference Framework logo" />
</p>


# User service
A repository with a service for managing users. This repository is part of [Microservices Reference Framework](https://github.com/MichalMoudry/microservices-reference-framework "Link to Microservices Reference Framework").

[![Build and test project](https://github.com/MichalMoudry/mrf-user-service/actions/workflows/go.yml/badge.svg)](https://github.com/MichalMoudry/mrf-user-service/actions/workflows/go.yml)
[![Deploy to Azure](https://github.com/MichalMoudry/mrf-user-service/actions/workflows/deploy.yml/badge.svg)](https://github.com/MichalMoudry/mrf-user-service/actions/workflows/deploy.yml)

## Project structure
- **/src** - Folder with all the source/test code related to this service.
    - /cmd - Folder with service's entrypoints.
    - /internal
        - /transport
        - /service
        - /config
- **/.dapr** - Folder with Dapr components for local deployment.
- **/.azure** - Folder containing definitions for Azure resources.
- **/assets** - Folder with repository's assets that are not part of the source code.

### Service architecture
This section describes architecture of this particular service and not the entire system.

**Note**: Arrows in the diagram below display direction of dependencies between layers. This project utilises Inversion of Control pattern for many component, layers including.

```mermaid
---
title: "Layers of the workflow service"
---
classDiagram
    class transport["Transport layer"]
    click transport href "https://github.com/MichalMoudry/mrf-workflow-service/tree/main/transport" "Go to transport layer package"
    class service["Service layer"]
    transport <-- service

    note for transport "validates requests\nparses request content\ncontains HTTP middleware\nrequest & response contracts"
    note for service "contains business logic\nhandles commiting transactions\nhandles rolling back of transactions\n..."
```

**Note**: This is a stateless service, so there is no persistance layer/database.

## Getting started

## Used technologies
