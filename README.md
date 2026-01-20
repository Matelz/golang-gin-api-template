# 🥤 Golang GIN API Template

This is a template for building RESTful APIs using the [Gin web framework](https://github.com/gin-gonic/gin) in Go. It provides a structured starting point for developing scalable and maintainable APIs. This template is designed to be deployed on a [Lambda](https://aws.amazon.com/lambda/) environment.

## 📑 How to use

1. Clone the repository:
   ```bash
   git clone https://github.com/Matelz/golang-gin-api-template.git
   cd golang-gin-api-template
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run ./cmd/app/main.go
   ```

4. Test the API endpoints using a tool like Postman or cURL.

## 🛠️ Project Structure

```bash
golang-gin-api-template/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   └── account/
│       ├── account.go/
│       ├── handler.go
│       ├── service.go
│       └── repository.go
│   └── user/
│       ├── user.go/
│       ├── handler.go
│       ├── service.go
│       └── repository.go
│   └── shared/
└── go.mod
```

- `cmd/app/main.go`: The entry point of the application.
- `internal/account/`: Contains the account module with its handler, service, and repository.
- `internal/user/`: Contains the user module with its handler, service, and repository.
- `internal/shared/`: Contains shared utilities and helper functions.

## 📄 License

This project is licensed under the GNU General Public License v3.0 See the [LICENSE](LICENSE) file for details.

## 🤝 Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.