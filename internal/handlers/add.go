package handlers

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func AddHandler(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	args := request.Params.Arguments.(map[string]any)

	a := args["a"].(float64)
	b := args["b"].(float64)

	result := fmt.Sprintf(
		"Result: %f",
		a+b,
	)

	return mcp.NewToolResultText(result), nil
}
