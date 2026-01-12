package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Sample data structures for different resource types
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Created string `json:"created"`
}

var (
	users = []User{
		{ID: 0, Name: "Alice Johnson", Email: "alice@example.com", Role: "admin", Created: "2024-01-15"},
		{ID: 1, Name: "Bob Smith", Email: "bob@example.com", Role: "user", Created: "2024-02-20"},
		{ID: 2, Name: "Carol Wilson", Email: "carol@example.com", Role: "moderator", Created: "2024-03-10"},
	}
)

func userToJson(usr *User) string {
	j, _ := json.Marshal(usr)
	return string(j)
}

func setupResourcetemplates(s *server.MCPServer) {
	// Dynamic resource example - user profiles by ID
	template := mcp.NewResourceTemplate(
		"users://{id}/profile",
		"User Profile",
		mcp.WithTemplateDescription("Returns user profile information"),
		mcp.WithTemplateMIMEType("application/json"),
	)

	// Add template with its handler
	s.AddResourceTemplate(template, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from the URI using regex matching
		// The server automatically matches URIs to templates
		uri := strings.TrimLeft(request.Params.URI, "users://")
		userID, err := strconv.Atoi(strings.Split(uri, "/")[0])
		if err != nil {
			return nil, fmt.Errorf("%q : %q: %v", request.Params.URI, userID, err)
		}

		profile := userToJson(&users[userID])

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     profile,
			},
		}, nil
	})
}

func setupResources(s *server.MCPServer) {
	// Static resource example - exposing a README file
	for i := range 50 {
		resource := mcp.NewResource(
			"docs://readme"+strconv.Itoa(i),
			"Project README "+strconv.Itoa(i),
			mcp.WithResourceDescription("The project's README file"+strconv.Itoa(i)),
			mcp.WithMIMEType("text/markdown"),
		)

		// Add resource with its handler
		s.AddResource(resource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			content, err := os.ReadFile("README.md")
			if err != nil {
				return nil, err
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      "docs://readme" + strconv.Itoa(i),
					MIMEType: "text/markdown",
					Text:     string(content),
				},
			}, nil
		})
	}
}

func main() {
	// Create a new MCP server with resource capabilities
	s := server.NewMCPServer(
		"Resource Demo Server",
		"1.0.0",
		server.WithResourceCapabilities(true, true), // Support both subscribe and listChanged
	)

	setupResources(s)
	setupResourcetemplates(s)
	setupTools(s)

	// Start the stdio server
	fmt.Printf("Starting MCP Resource Server...\n")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// setupTools adds some utility tools for resource manipulation
func setupTools(s *server.MCPServer) {
	// Tool to create a new user (demonstrates data modification)
	createUserTool := mcp.NewTool("create_user",
		mcp.WithDescription("Create a new user in the system"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Full name of the user"),
		),
		mcp.WithString("email",
			mcp.Required(),
			mcp.Description("Email address of the user"),
		),
		mcp.WithString("role",
			mcp.Description("User role (admin, user, moderator)"),
			mcp.Enum("admin", "user", "moderator"),
		),
	)

	s.AddTool(createUserTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		var allOk = true
		name, ok := args["name"].(string)
		allOk = allOk && ok

		email, ok := args["email"].(string)
		allOk = allOk && ok

		role := args["role"].(string)
		allOk = allOk && ok

		if !allOk {
			return mcp.NewToolResultError("could not parse argument"), nil
		}

		// Create new user with next available ID
		newID := len(users) + 1
		newUser := User{
			ID:      newID,
			Name:    name,
			Email:   email,
			Role:    role,
			Created: time.Now().Format("2006-01-02"),
		}

		// Add to users slice
		users = append(users, newUser)

		// Send notification about the change (if clients are subscribed)
		// This would trigger a listChanged notification in a real implementation

		jsonData, err := json.MarshalIndent(newUser, "", "  ")
		if err != nil {
			return mcp.NewToolResultError("Failed to format user data"), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("User created successfully:\n%s", string(jsonData))), nil
	})
}
