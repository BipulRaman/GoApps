# GoApps

## Getting Started with Go and GoGIN

This guide will help you set up your environment for Go and the GoGIN framework, build applications, and run them.

### Prerequisites

1. **Windows Package Manager (`winget`)**
   Ensure you have `winget` installed on your Windows machine. It comes pre-installed on Windows 10 (version 1809 and later) and Windows 11.

2. **Visual Studio Code**
   Download and install Visual Studio Code if you don’t already have it:

   ```bash
   winget install -e --id Microsoft.VisualStudioCode
   ```

3. **Go Programming Language**
   Install Go using `winget`:

   ```bash
   winget install -e --id GoLang.Go
   ```

4. **Git**
   Install Git to manage version control:

   ```bash
   winget install -e --id Git.Git
   ```

---

### Setting Up the Environment

1. **Verify Go Installation**
   After installing Go, verify the installation:

   ```bash
   go version
   ```

   This should display the installed Go version.

2. **Set Up Go Workspace**
   Ensure your Go workspace is set up correctly. By default, Go uses the `GOPATH` environment variable. You can check it with:

   ```bash
   go env GOPATH
   ```

   If needed, create a directory for your Go workspace and set it in your environment variables.

3. **Install GoGIN Framework**
   Install the GoGIN framework using the following command:

   ```bash
   go get -u github.com/gin-gonic/gin
   ```

---

### Building and Running Go Applications

1. **Navigate to Your Project**
   Open a terminal and navigate to your project directory:

   ```bash
   cd c:\Code\GitHub\GoApps\go-gin-api
   ```

2. **Install Dependencies**
   Use `go mod tidy` to download and install all required dependencies:

   ```bash
   go mod tidy
   ```

3. **Run the Application**
   Navigate to the `cmd` folder and run the application:

   ```bash
   go run ./cmd/main.go
   ```

---

### Building the Application

To create a build of your application:

1. **Build for Local Environment**
   Run the following command to build the application for your local environment:

   ```bash
   go build -o bin/app.exe ./cmd
   ```

   The binary will be available in the `bin` folder.

2. **Build for Windows Deployment**
   To build a Windows-specific binary:

   ```bash
   GOOS=windows GOARCH=amd64 go build -o bin/app.exe ./cmd
   ```

---

### Running the Built Application

After building the application, you can run it using:

```bash
bin\app.exe
```

---

### Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [GoGIN Documentation](https://gin-gonic.com/docs/)
- [Visual Studio Code Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)