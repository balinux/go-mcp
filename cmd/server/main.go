package main

import (
	"fmt"

	"go-mcp/internal/handlers"
	"go-mcp/internal/tools"

	"github.com/mark3labs/mcp-go/server"
)

func main() {

	s := server.NewMCPServer(
		"go mcp demo",
		"1.0.0",
	)

	// register tools
	s.AddTool(
		tools.HelloTool(),
		handlers.HelloHandler,
	)

	s.AddTool(
		tools.AddTool(),
		handlers.AddHandler,
	)

	s.AddTool(
		tools.SuratTool(),
		handlers.SuratHandler,
	)

	fmt.Println("Starting server")

	if err := server.ServeStdio(s); err != nil {
		fmt.Println(err)
	}
}
