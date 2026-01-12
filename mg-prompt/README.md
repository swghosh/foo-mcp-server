# MCP Resource Server in Go

A comprehensive Model Context Protocol (MCP) server implementation in Go that demonstrates how to expose various types of resources using the `github.com/mark3labs/mcp-go` library.

## Overview

This MCP server showcases the implementation of MCP resources according to the [MCP specification](https://modelcontextprotocol.io/specification/2025-06-18/server/resources). It demonstrates:

- **Static Resources**: Fixed URI resources like system information and documentation
- **Dynamic Resources**: Template-based resources with parameterized URIs
- **Resource Annotations**: Metadata for audience targeting, priority, and modification tracking
- **Tools Integration**: Additional functionality for resource manipulation
- **Multiple Content Types**: JSON, Markdown, and plain text resources

## Features

### Static Resources
- `system://info` - Real-time system information and server status
- `docs://readme` - Server documentation and usage instructions

### Dynamic Resource Templates
- `users://{id}` - Individual user information by ID
- `projects://{id}` - Individual project information by ID  
- `data://{collection}` - Access to data collections (users, projects)
- `file://{path}` - Secure file system access (current directory only)

### Tools
- `create_user` - Create new users in the system
- `list_resources` - List all available resources with descriptions

### Capabilities
- ✅ Resource listing with pagination support
- ✅ Resource reading with proper content types
- ✅ Resource templates with URI parameter extraction
- ✅ Resource annotations (audience, priority, lastModified)
- ✅ Error handling with proper MCP error codes
- ✅ Security validation for file access
- ✅ Multiple MIME types (JSON, Markdown, Go source, etc.)

## Installation

1. Ensure you have Go 1.21 or later installed
2. Clone or download this code
3. Install dependencies:

```bash
go mod tidy
```

## Usage

### Running the Server

```bash
go run main.go
```

The server runs using the stdio transport, which is the standard way MCP servers communicate with clients.

### Example Resource Access

Once connected through an MCP client, you can access resources like:

```
# Get system information
resources/read: system://info

# Get user with ID 1
resources/read: users://1

# Get all projects
resources/read: data://projects

# Read the Go source file
resources/read: file://main.go

# Read the README
resources/read: docs://readme
```

### Example Tool Usage

```
# Create a new user
tools/call: create_user
{
  "name": "John Doe",
  "email": "john@example.com", 
  "role": "user"
}

# List available resources
tools/call: list_resources
```

## Resource Types Demonstrated

### 1. Static Resources
Resources with fixed URIs that provide consistent data:
```go
systemResource := mcp.NewResource(
    "system://info",
    "System Information",
    mcp.WithResourceDescription("Current system information and server status"),
    mcp.WithMIMEType("application/json"),
)
```

### 2. Resource Templates
Dynamic resources using URI templates with parameters:
```go
userTemplate := mcp.NewResourceTemplate(
    "users://{id}",
    "User Information", 
    mcp.WithTemplateDescription("Individual user information by ID"),
    mcp.WithTemplateMIMEType("application/json"),
)
```

### 3. Resource Annotations
Metadata to help clients understand resource usage:
```go
mcp.WithResourceAnnotations(mcp.ResourceAnnotations{
    Audience:     []string{"user", "assistant"},
    Priority:     0.8,
    LastModified: time.Now().Format(time.RFC3339),
})
```

## Security Features

- **File Access Control**: Only allows access to files in the current directory and subdirectories
- **Path Validation**: Prevents directory traversal attacks (../, absolute paths)
- **Input Validation**: Proper validation of all URI parameters and tool inputs
- **Error Handling**: Secure error messages that don't leak sensitive information

## Sample Data

The server includes sample data for demonstration:

**Users:**
- Alice Johnson (admin)
- Bob Smith (user)  
- Carol Wilson (moderator)

**Projects:**
- MCP Resource Server (active)
- Data Analytics Tool (completed)
- Mobile App (in-progress)

## Code Structure

- `main.go` - Main server implementation
- `setupStaticResources()` - Static resource definitions
- `setupResourceTemplates()` - Dynamic resource template definitions
- `setupTools()` - Tool implementations for resource manipulation

## MCP Compliance

This server implements the MCP specification features:

- ✅ Resource listing (`resources/list`)
- ✅ Resource reading (`resources/read`) 
- ✅ Resource templates (`resources/templates/list`)
- ✅ Proper error handling with MCP error codes
- ✅ Multiple content types (text, JSON, binary support)
- ✅ Resource annotations
- ✅ URI scheme compliance

## Extending the Server

To add new resources:

1. **Static Resource**: Use `s.AddResource()` with a fixed URI
2. **Dynamic Resource**: Use `s.AddResourceTemplate()` with URI templates
3. **New Tools**: Use `s.AddTool()` for additional functionality

Example of adding a new static resource:
```go
newResource := mcp.NewResource(
    "custom://data",
    "Custom Data",
    mcp.WithResourceDescription("My custom data resource"),
    mcp.WithMIMEType("application/json"),
)

s.AddResource(newResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
    // Your resource implementation here
    return []mcp.ResourceContents{
        mcp.TextResourceContents{
            URI:      "custom://data",
            MIMEType: "application/json",
            Text:     `{"message": "Hello from custom resource"}`,
        },
    }, nil
})
```

## Requirements

- Go 1.21+
- github.com/mark3labs/mcp-go v0.8.0+

## License

This is a demonstration/example implementation. Feel free to use it as a starting point for your own MCP servers.