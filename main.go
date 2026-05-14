package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// membuat mcp server
	s := server.NewMCPServer(
		"go mcp demo", "1.0.0",
	)

	// register tool
	helloTool := mcp.NewTool(
		"hello",
		mcp.WithDescription("Say hello to user"),
		mcp.WithString(
			"name",
			mcp.Required(),
			mcp.Description("username"),
		),
	)

	// register addtool
	addTool := mcp.NewTool(
		"add",
		mcp.WithDescription(" add two number"),
		mcp.WithNumber("a", mcp.Required()),
		mcp.WithNumber("b", mcp.Required()),
	)

	// surat tool
	suratTool := mcp.NewTool(
		"get_surat",
		mcp.WithDescription("get surat al quran"),
		mcp.WithNumber("nomor", mcp.Required()),
	)

	// handler tool
	s.AddTool(helloTool, helloHandler)
	s.AddTool(addTool, addHandler)
	s.AddTool(suratTool, suratHandler)

	fmt.Println("Starting server")

	// jalankan STDIO Server
	if err := server.ServeStdio(s); err != nil {
		fmt.Println(err)
	}
}

func helloHandler(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	args := request.Params.Arguments.(map[string]any)
	name := args["name"].(string)

	result := fmt.Sprintf("Hello %s 🚀", name)

	return mcp.NewToolResultText(result), nil
}

func addHandler(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments.(map[string]any)

	a := args["a"].(float64)
	b := args["b"].(float64)

	result := fmt.Sprintf("Result: %f", a+b)

	return mcp.NewToolResultText(result), nil
}

type Ayat struct {
	Ar    string `json:"ar"`
	Id    string `json:"id"`
	Tr    string `json:"tr"`
	Nomor string `json:"nomor"`
}

func suratHandler(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments.(map[string]any)

	nomor := args["nomor"].(float64)

	url := fmt.Sprintf("https://api.npoint.io/99c279bb173a6e28359c/surat/%d", int(nomor))

	resp, err := http.Get(url)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var ayat []Ayat
	err = json.Unmarshal(body, &ayat)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var result string
	for _, a := range ayat {
		result += fmt.Sprintf(
			"%s. %s\n%s\n\n",
			a.Nomor,
			a.Ar,
			a.Id,
		)
	}

	return mcp.NewToolResultText(result), nil
}
