# Go, gRPC, and React starter boilerplate.

## Features
- **SQL light backend**
- **Connect and wire framework**

## Getting Started

### Prerequisites
- Node.js
- Go (Golang)

### Setting Up the Project

1. **Set up project**  
    - Make sure you choose a new name for your project
      ```bash
      git clone https://github.com/n0remac/Go-gRPC-React-starter.git <new-project-name>
      ```
    - Change directory
      ```
      cd <new-project-name>
      ```
    - Use IDE to replace all instances of Go-gRPC-React-starter with `<new-project-name>`. In the package.json the new project name should be entirely lowercase.
2. **Set Up the Frontend (React)**
   - Ensure you are in the project's root directory.
   - Install dependencies:
     ```bash
     npm install
     ```
   - Start the React development server:
     ```bash
     npm run dev
     ```

3. **Set Up the Backend (Go)**
   - Ensure you are in the project's root directory.
   - Install Go dependencies:
     ```bash
     go mod tidy
     ```
   - Run the Go server:
     ```bash
     go run main.go
     ```
   - Generate gRPC code:
     ```bash
     go generate .
     ```


3. **Access the Site**
   - Open your browser and navigate to `http://localhost:8000` (or the port specified for your React app).


### Using the Project
  1. Addind a new protobuf
      - Add new .proto file in the proto folder
      - Whenever a new protobuf is added rerun `go generate .`
      - Add a new service and apiRoot.Handle() in the main function of main.go
      - Add a new service in frontend/src/service.ts