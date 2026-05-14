package handlers

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func HelloHandler(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	args := request.Params.Arguments.(map[string]any)

	name := args["name"].(string)

	result := fmt.Sprintf(
		"Hello %s 🚀",
		name,
	)

	return mcp.NewToolResultText(result), nil
}
