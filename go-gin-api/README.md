# Go Gin API

This project is a simple API built using the Gin framework in Go. It serves as a demonstration of how to structure a Go application with a focus on separation of concerns.

## Project Structure

```
go-gin-api
├── cmd
│   └── main.go               # Entry point of the application
├── internal
│   ├── controllers           # Contains HTTP request handlers
│   │   └── example_controller.go
│   ├── models                # Contains data models
│   │   └── example_model.go
│   ├── routes                # Contains route definitions
│   │   └── router.go
│   └── services              # Contains business logic
│       └── example_service.go
├── go.mod                    # Module definition and dependencies
└── README.md                 # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-gin-api
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

4. **Access the API:**
   Open your browser or use a tool like Postman to access the API at `http://localhost:8080`.

## Usage Examples

- **Get Example:**
  ```
  GET /example
  ```

- **Create Example:**
  ```
  POST /example
  Content-Type: application/json

  {
      "name": "Example Name"
  }
  ```

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.