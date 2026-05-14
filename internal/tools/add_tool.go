package tools

import "github.com/mark3labs/mcp-go/mcp"

func AddTool() mcp.Tool {

	return mcp.NewTool(
		"add",
		mcp.WithDescription("add two numbers"),

		mcp.WithNumber("a", mcp.Required()),
		mcp.WithNumber("b", mcp.Required()),
	)
}
