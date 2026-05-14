package tools

import "github.com/mark3labs/mcp-go/mcp"

func HelloTool() mcp.Tool {
	return mcp.NewTool(
		"hello",
		mcp.WithDescription("Say hello to user"),
		mcp.WithString(
			"name",
			mcp.Required(),
			mcp.Description("username"),
		),
	)
}
