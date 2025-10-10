# PROJECT.md

## WebGen - Full-Stack Application Boilerplate Generator

### Purpose

WebGen is a schema-driven code generation framework that automates the creation of full-stack web applications with complete CRUD functionality. It eliminates repetitive boilerplate work by generating a production-ready application stack from simple YAML schema definitions, including backend services, API endpoints, database models, frontend admin interfaces, and automated GitHub repository deployment.

---

## Architecture

WebGen follows a **template-based code generation** architecture with two main components:

### 1. Generator CLI (`/generation`)
The code generation engine that reads YAML schemas and produces a complete application codebase using Go templates.

### 2. Base Application Template (`/app`)
A pre-configured full-stack application template that serves as the foundation for all generated projects. This template is copied, customized, and enhanced with schema-specific code.

---

## Technical Stack

### Backend
- **Language**: Go 1.21.4
- **API Protocol**: Connect-RPC (gRPC-compatible HTTP/2 protocol)
- **Web Framework**: `net/http` with `connect-go`
- **Database**: SQLite with `upper/db` ORM
- **Authentication**: JWT tokens (`dgrijalva/jwt-go`)
- **API Reflection**: gRPC reflection for service discovery
- **CORS**: Configurable cross-origin middleware

### Frontend
- **Framework**: React 18 with TypeScript
- **Styling**: TailwindCSS + DaisyUI
- **Build Tool**: esbuild
- **Routing**: React Router v6
- **API Client**: Bufbuild Connect-Web (type-safe RPC client)
- **UI Components**: React Bootstrap + custom admin components

### Infrastructure & Tooling
- **IaC**: Terraform (GitHub repository automation)
- **Code Generation**: Go templates + text/template
- **Protocol Buffers**: Bufbuild for schema compilation
- **Version Control**: Automated Git initialization and push

---

## Major Components

### 1. Schema Definition System
- **Location**: `generation/schema.yaml`
- **Purpose**: Declarative model definitions using YAML
- **Features**:
  - Field type definitions (int32, string, etc.)
  - Primary key and auto-increment configuration
  - Multi-model support

**Example Schema**:
```yaml
models:
  - name: Product
    fields:
      - name: id
        type: int32
        primary_key: true
        auto_increment: true
      - name: name
        type: string
      - name: amount
        type: int32
```

### 2. Code Generators (`generation/internal/generator/`)

#### Model Generator (`model.go`)
Generates Go structs with database tags for ORM mapping.

#### Database Generator (`database.go`)
Creates database initialization code and table schemas.

#### Service Generator (`service.go`)
Produces backend service layer with CRUD operations.

#### Proto Generator (`proto.go`)
Generates Protocol Buffer definitions for API contracts.

#### React Generator (`react.go`)
Creates:
- TypeScript service clients
- Admin CRUD pages with forms and data tables
- Routing configuration

#### Terraform Generator (`terraform.go`)
Produces Infrastructure as Code for GitHub repository creation.

### 3. Template System (`generation/templates/`)
- **Go Templates**: Backend models, services, database schemas
- **Proto Templates**: gRPC/Connect service definitions
- **React Templates**: TypeScript components and admin pages
- **Terraform Templates**: GitHub provider configuration

### 4. Base Application (`/app`)

#### Backend Structure
```
app/
├── main.go              # HTTP/2 server with Connect-RPC
├── pkg/
│   ├── database/        # SQLite initialization
│   ├── user/            # Example model and service
│   └── service/         # Service registry
└── gen/proto/           # Generated protobuf code
```

#### Frontend Structure
```
app/frontend/
├── src/
│   ├── pages/admin/     # Auto-generated admin pages
│   ├── services/        # API client services
│   └── pages/index.ts   # Route registry
├── generatePages.js     # Route generation script
└── tailwind.config.js   # Styling configuration
```

### 5. Project Initialization Pipeline (`generation/main.go`)

**Workflow**:
1. **Copy Template**: Duplicates base application
2. **Name Substitution**: Replaces "CodeGen" with project name
3. **Schema Processing**: Loads YAML and generates code
4. **Terraform Execution**: Creates GitHub repository
5. **Git Automation**: Initializes repo, commits, and pushes code

---

## Unique Features

### 1. End-to-End Automation
Unlike scaffolding tools that only generate initial code, WebGen creates a complete deployment pipeline from schema to live GitHub repository.

### 2. Connect-RPC Integration
Modern gRPC-compatible protocol that works over HTTP/1.1 and HTTP/2, providing type-safe APIs without the complexity of traditional gRPC.

### 3. Schema-Driven Development
Single source of truth (YAML schema) automatically propagates changes through:
- Database schemas
- Go models
- Proto definitions
- TypeScript types
- React UI components

### 4. Auto-Generated Admin Interface
Every model defined in the schema automatically receives a fully functional admin page with:
- Data tables with CRUD operations
- Form validation
- Type-safe API calls
- Responsive design

### 5. Infrastructure as Code
Terraform integration automates:
- GitHub repository creation
- Initial commit and push
- Repository configuration

### 6. Type Safety Across Stack
- Backend: Go static typing
- API: Protocol Buffers
- Frontend: TypeScript
- Database: ORM with type mapping

---

## Development Workflow

### Creating a New Project

1. **Define Schema** (`generation/schema.yaml`):
```yaml
models:
  - name: Todo
    fields:
      - name: id
        type: int32
        primary_key: true
      - name: title
        type: string
      - name: completed
        type: bool
```

2. **Configure Project** (`generation/main.go`):
```go
var projectName string = "MyTodoApp"
```

3. **Run Generator**:
```bash
cd generation
go run main.go
```

4. **Generated Output**:
- Complete Go backend with Todo model and service
- Proto definitions for Todo API
- React admin page for managing todos
- GitHub repository created and code pushed

---

## Integration Points

### Protocol Buffer Workflow
1. Generator creates `.proto` files
2. `buf generate` compiles to Go and TypeScript
3. Backend implements service interfaces
4. Frontend consumes typed clients

### Database Layer
- SQLite for development simplicity
- `upper/db` provides database-agnostic ORM
- Auto-migration from schema definitions

### API Layer
- Connect-RPC handlers registered in `main.go`
- CORS middleware for local development
- gRPC reflection for debugging tools

---

## Use Cases

- **Rapid Prototyping**: Transform ideas into working apps in minutes
- **Internal Tools**: Generate admin dashboards for data management
- **API Development**: Quickly scaffold microservices with type-safe contracts
- **Learning Platform**: Study full-stack architecture with working examples
- **Hackathons**: Bootstrap projects with production-ready foundations

---

## Future Enhancements

Potential areas for expansion:
- Multiple database support (PostgreSQL, MySQL)
- Authentication/authorization code generation
- API documentation generation (OpenAPI/Swagger)
- Docker and Kubernetes deployment templates
- GraphQL endpoint generation
- Real-time features (WebSockets, Server-Sent Events)
- Testing scaffolds (unit tests, integration tests)

---

## Repository Structure

```
WebGen/
├── app/                    # Base application template
│   ├── frontend/          # React + TypeScript frontend
│   ├── pkg/               # Go backend packages
│   └── main.go            # Server entry point
├── generation/            # Code generation CLI
│   ├── internal/
│   │   └── generator/    # Generator implementations
│   ├── templates/        # Go template files
│   ├── schema.yaml       # Example schema
│   └── main.go           # Generator entry point
└── README.md
```

---

## Technical Philosophy

WebGen embodies **convention over configuration** and **DRY (Don't Repeat Yourself)** principles. By codifying common full-stack patterns into templates, it allows developers to focus on business logic rather than boilerplate. The schema-first approach ensures consistency across all layers of the application while maintaining flexibility for customization after generation.
