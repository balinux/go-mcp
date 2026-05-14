package tools

import "github.com/mark3labs/mcp-go/mcp"

func SuratTool() mcp.Tool {

	return mcp.NewTool(
		"get_surat",
		mcp.WithDescription("get surat al quran"),

		mcp.WithNumber(
			"nomor",
			mcp.Required(),
		),
	)
}
