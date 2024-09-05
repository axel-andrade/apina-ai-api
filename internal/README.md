# `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

Examples:

* https://github.com/hashicorp/terraform/tree/main/internal
* https://github.com/influxdata/influxdb/tree/master/internal
* https://github.com/perkeep/perkeep/tree/master/internal
* https://github.com/jaegertracing/jaeger/tree/main/internal
* https://github.com/moby/moby/tree/master/internal
* https://github.com/satellity/satellity/tree/main/internal

## `/internal/pkg`

Examples:

* https://github.com/hashicorp/waypoint/tree/main/internal/pkg

## Hexagonal Architecture Folder Structure

This project follows the principles of the hexagonal architecture, also known as ports and adapters. Below is an overview of the folder structure and how it relates to the hexagonal architecture:

### adapters

- primary: Contains adapters responsible for interacting with external systems, such as HTTP controllers, middlewares, presenters, and routes.
    - http: Handles HTTP-related concerns including handling routes, middleware, and controllers.
    - common: Contains common utilities and helpers for HTTP-related functionality.
    - controllers: Houses controllers responsible for handling HTTP requests and invoking corresponding use cases.
    - middlewares: Contains middleware functions for intercepting and modifying HTTP requests or responses.
    - presenters: Formats data from use cases into a presentable format for the HTTP response.
    - routes: Defines the routes and their corresponding handlers.
    - schemas: Contains JSON schema files for request validation.
    - server: Initializes and starts the HTTP server.

- secondary: Contains adapters responsible for interacting with secondary systems, such as databases or external APIs.
    - database: Contains database-related functionality, including mapper, models, and repositories.    
        - mongo: Handles MongoDB-specific functionality.
        - redis: Handles Redis-specific functionality.

### core

    - domain: Contains the core domain logic, including entities, value objects, errors, and constants.
        - constants: Houses application-wide constants.
        - errors: Defines custom error types.
        - value_objects: Contains reusable value objects used within the domain.

    - usecases: Contains application-specific use cases (business logic).
        Each use case typically consists of an interactor responsible for executing the use case logic and ports defining the input/output interfaces.

### infra

    - bootstrap: Contains initialization logic for setting up the application.
    - handlers: Contains infrastructure-related handlers, such as encrypter, JSON validator, and token manager handlers.
    - logger: Defines the logger used for logging application events.

This folder structure promotes separation of concerns and modularization, allowing for easier maintenance and scalability of the application. It ensures that business logic remains decoupled from external systems, facilitating testing and future changes.