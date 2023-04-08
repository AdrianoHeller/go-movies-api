package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Fetch(url string) ([]byte, error) {

	apiKey := os.Getenv("API_KEY")

	baseUrl := "https://api.themoviedb.org/3"

	mountedUrl := fmt.Sprintf("%s%s?api_key=%s", baseUrl, url, apiKey)

	apiClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, mountedUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Skulldev")

	res, err := apiClient.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func SetJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
