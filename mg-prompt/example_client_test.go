package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Example test file demonstrating how the MCP resources would be accessed
// This is not a functional test but shows the expected request/response patterns

func TestResourceExamples(t *testing.T) {
	// Example 1: Static resource request/response
	fmt.Println("=== Example 1: Static Resource (system://info) ===")
	
	// Request that would be sent to the server
	systemInfoRequest := `{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "resources/read",
		"params": {
			"uri": "system://info"
		}
	}`
	
	// Expected response format
	systemInfoResponse := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": {
			"contents": [
				{
					"uri": "system://info",
					"name": "System Information",
					"mimeType": "application/json",
					"text": "{\n  \"server_name\": \"Resource Demo Server\",\n  \"version\": \"1.0.0\",\n  \"timestamp\": \"2024-01-15T10:30:00Z\",\n  \"uptime\": \"running\",\n  \"resources\": {\n    \"users\": 3,\n    \"projects\": 3\n  }\n}"
				}
			]
		}
	}`
	
	fmt.Printf("Request:\n%s\n\n", formatJSON(systemInfoRequest))
	fmt.Printf("Response:\n%s\n\n", formatJSON(systemInfoResponse))

	// Example 2: Resource template request/response
	fmt.Println("=== Example 2: Resource Template (users://1) ===")
	
	userRequest := `{
		"jsonrpc": "2.0",
		"id": 2,
		"method": "resources/read",
		"params": {
			"uri": "users://1"
		}
	}`
	
	userResponse := `{
		"jsonrpc": "2.0",
		"id": 2,
		"result": {
			"contents": [
				{
					"uri": "users://1",
					"name": "User Alice Johnson",
					"mimeType": "application/json",
					"text": "{\n  \"id\": 1,\n  \"name\": \"Alice Johnson\",\n  \"email\": \"alice@example.com\",\n  \"role\": \"admin\",\n  \"created\": \"2024-01-15\"\n}"
				}
			]
		}
	}`
	
	fmt.Printf("Request:\n%s\n\n", formatJSON(userRequest))
	fmt.Printf("Response:\n%s\n\n", formatJSON(userResponse))

	// Example 3: Resource listing request/response
	fmt.Println("=== Example 3: Resource Listing ===")
	
	listRequest := `{
		"jsonrpc": "2.0",
		"id": 3,
		"method": "resources/list"
	}`
	
	listResponse := `{
		"jsonrpc": "2.0",
		"id": 3,
		"result": {
			"resources": [
				{
					"uri": "system://info",
					"name": "System Information",
					"title": "System Information",
					"description": "Current system information and server status",
					"mimeType": "application/json",
					"annotations": {
						"audience": ["user", "assistant"],
						"priority": 0.8,
						"lastModified": "2024-01-15T10:30:00Z"
					}
				},
				{
					"uri": "docs://readme",
					"name": "README Documentation",
					"title": "README Documentation",
					"description": "Server documentation and usage instructions",
					"mimeType": "text/markdown",
					"annotations": {
						"audience": ["user"],
						"priority": 0.9
					}
				}
			]
		}
	}`
	
	fmt.Printf("Request:\n%s\n\n", formatJSON(listRequest))
	fmt.Printf("Response:\n%s\n\n", formatJSON(listResponse))

	// Example 4: Resource templates listing
	fmt.Println("=== Example 4: Resource Templates Listing ===")
	
	templatesRequest := `{
		"jsonrpc": "2.0",
		"id": 4,
		"method": "resources/templates/list"
	}`
	
	templatesResponse := `{
		"jsonrpc": "2.0",
		"id": 4,
		"result": {
			"resourceTemplates": [
				{
					"uriTemplate": "users://{id}",
					"name": "User Information",
					"title": "User Information",
					"description": "Individual user information by ID",
					"mimeType": "application/json",
					"annotations": {
						"audience": ["user", "assistant"],
						"priority": 0.7
					}
				},
				{
					"uriTemplate": "projects://{id}",
					"name": "Project Information", 
					"title": "Project Information",
					"description": "Individual project information by ID",
					"mimeType": "application/json",
					"annotations": {
						"audience": ["user", "assistant"],
						"priority": 0.7
					}
				},
				{
					"uriTemplate": "data://{collection}",
					"name": "Data Collections",
					"title": "Data Collections",
					"description": "Access to data collections (users, projects)",
					"mimeType": "application/json",
					"annotations": {
						"audience": ["user", "assistant"],
						"priority": 0.6
					}
				},
				{
					"uriTemplate": "file://{path}",
					"name": "File System Access",
					"title": "File System Access",
					"description": "Access to files in the current directory",
					"mimeType": "text/plain",
					"annotations": {
						"audience": ["user", "assistant"],
						"priority": 0.5
					}
				}
			]
		}
	}`
	
	fmt.Printf("Request:\n%s\n\n", formatJSON(templatesRequest))
	fmt.Printf("Response:\n%s\n\n", formatJSON(templatesResponse))

	// Example 5: Tool call example
	fmt.Println("=== Example 5: Tool Call (create_user) ===")
	
	toolRequest := `{
		"jsonrpc": "2.0",
		"id": 5,
		"method": "tools/call",
		"params": {
			"name": "create_user",
			"arguments": {
				"name": "David Brown",
				"email": "david@example.com",
				"role": "user"
			}
		}
	}`
	
	toolResponse := `{
		"jsonrpc": "2.0",
		"id": 5,
		"result": {
			"content": [
				{
					"type": "text",
					"text": "User created successfully:\n{\n  \"id\": 4,\n  \"name\": \"David Brown\",\n  \"email\": \"david@example.com\",\n  \"role\": \"user\",\n  \"created\": \"2024-01-15\"\n}"
				}
			]
		}
	}`
	
	fmt.Printf("Request:\n%s\n\n", formatJSON(toolRequest))
	fmt.Printf("Response:\n%s\n\n", formatJSON(toolResponse))

	// Example 6: Error response
	fmt.Println("=== Example 6: Error Response (resource not found) ===")
	
	errorRequest := `{
		"jsonrpc": "2.0",
		"id": 6,
		"method": "resources/read",
		"params": {
			"uri": "users://999"
		}
	}`
	
	errorResponse := `{
		"jsonrpc": "2.0",
		"id": 6,
		"error": {
			"code": -32002,
			"message": "Resource not found",
			"data": {
				"uri": "users://999",
				"details": "user not found: 999"
			}
		}
	}`
	
	fmt.Printf("Request:\n%s\n\n", formatJSON(errorRequest))
	fmt.Printf("Response:\n%s\n\n", formatJSON(errorResponse))
}

// Helper function to format JSON for better readability
func formatJSON(jsonStr string) string {
	var obj interface{}
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return jsonStr // Return original if parsing fails
	}
	
	formatted, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return jsonStr // Return original if formatting fails
	}
	
	return string(formatted)
}