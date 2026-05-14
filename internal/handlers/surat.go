package handlers

import (
	"context"
	"fmt"

	"go-mcp/internal/services"

	"github.com/mark3labs/mcp-go/mcp"
)

func SuratHandler(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	args := request.Params.Arguments.(map[string]any)

	nomor := int(args["nomor"].(float64))

	data, err := services.GetSurat(nomor)
	if err != nil {
		return mcp.NewToolResultError(
			err.Error(),
		), nil
	}

	var result string

	for _, a := range data {

		result += fmt.Sprintf(
			"%s. %s\n%s\n\n",
			a.Nomor,
			a.Ar,
			a.Id,
		)
	}

	return mcp.NewToolResultText(result), nil
}
