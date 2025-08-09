# YOGO - REST API Client Generator

<img src="yogo-gopher.png" alt="YOGO" width="512" style="vertical-align:middle" />

[![Go Reference](https://pkg.go.dev/badge/github.com/luciancaetano/yogo.svg)](https://pkg.go.dev/github.com/luciancaetano/yogo)  
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue.svg)](https://golang.org/doc/go1.20)

YOGO is a **type-safe REST API client generator for Go**, inspired by [sqlc](https://github.com/kyleconroy/sqlc).  
Generate Go client code directly from your `yogo.yml` specification, making API consumption easy and error-proof.

---

## Features

- Generate strongly typed Go client code from API specs  
- Define endpoints, methods, request parameters, and response models  
- Automatic request building and response unmarshalling  
- CLI tool with commands for creating specs and generating code  

---

## Installation

```bash
go install github.com/luciancaetano/yogo/cmd/main@latest
````

---

## Usage

```bash
yogo [command]
```

Common commands:

| Command      | Description                        |
| ------------ | ---------------------------------- |
| `new`        | Create a new `yogo.yml` spec file  |
| `generate`   | Generate client code from the spec |
| `completion` | Generate shell autocomplete script |
| `help`       | Show help information              |

Example workflow:

```bash
yogo new
# Edit yogo.yml with your API spec
yogo generate
```

---

## Example `yogo.yml`

```yaml
name: ExampleAPI
version: v1
base_url: https://api.example.com

endpoints:
  - name: GetUser
    path: /users/{id}
    method: GET
    request:
      params:
        - name: id
          type: string
    response:
      type: User

models:
  User:
    id: string
    name: string
    email: string
```

---

## Contributing

Contributions, issues, and feature requests are welcome!
Feel free to open a pull request or issue.

---

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
