package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-mcp/internal/models"
)

func GetSurat(
	nomor int,
) ([]models.Ayat, error) {

	url := fmt.Sprintf(
		"https://api.npoint.io/99c279bb173a6e28359c/surat/%d",
		nomor,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ayat []models.Ayat

	err = json.Unmarshal(body, &ayat)
	if err != nil {
		return nil, err
	}

	return ayat, nil
}
