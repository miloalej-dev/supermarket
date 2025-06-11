# Supermarket

A simple supermarket management API built with Go and Chi, following Domain-Driven Design (DDD) and Hexagonal Architecture principles.

## Table of Contents
- [Overview](#overview)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
- [API Usage](#api-usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [References](#references)

## Overview
This project provides a RESTful API for managing supermarket products. It demonstrates clean architecture practices, separation of concerns, and testability.

**Key Features:**
- Product management (CRUD)
- In-memory and file-based repositories
- Modular, testable codebase

## Architecture
The project uses:
- **Domain-Driven Design (DDD):** Focuses on the core domain and domain logic.
- **Hexagonal Architecture (Ports & Adapters):** Decouples business logic from external systems (e.g., database, HTTP).

**Main Components:**
- `domain/`: Core business entities and logic
- `application/`: Use cases and service interfaces
- `infrastructure/`: Adapters for HTTP, database, etc.
- `cmd/api/`: Application entrypoint

## Getting Started
### Prerequisites
- Go 1.18+

### Installation & Running
1. Clone the repository:
   ```sh
   git clone https://github.com/miloalej-dev/supermarket
   cd supermarket
   ```
2. Run the API:
   ```sh
   go run cmd/api/main.go
   ```
   Or with Docker:
   ```sh
   docker build -t supermarket .
   docker run -p 8080:8080 supermarket
   ```

## API Usage
Example request to list products:
```http
GET /products HTTP/1.1
Host: localhost:8080/products
```

Example response:
```json
[
  {
    "id": 1,
   "name": "Apple",
    "quantity": 20,
    "code_value": "AS210D",
    "is_published": true,
    "expiration": "20/10/2025",
    "price": 12.65
  }
]
```

See [api/requests/products.http](api/requests/products.http) for more examples.

## Testing
To run tests:
```sh
go test ./...
```

## Contributing
Contributions are welcome! Please open issues or pull requests.

### Git Workflow
This project uses **git flow** for branching and release management. Please follow the git flow model when contributing.

We also use **gitmoji** for commit messages. Use appropriate emojis to describe your commits. See [gitmoji.dev](https://gitmoji.dev/) for the full list and usage guide.

## References
- [Gophercon 2018](https://www.youtube.com/watch?v=oL6JBUk6tj0)
- [Netflix Hexagonal architecture post](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)
- [Geek for Geeks Hexagonal architecture](https://www.geeksforgeeks.org/hexagonal-architecture-system-design/)
- [Hexagonal architecture post](https://medium.com/@edusalguero/arquitectura-hexagonal-5)