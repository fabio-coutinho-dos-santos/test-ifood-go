package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	http_errors "ifood-backend-test/src/application/http-errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type TokenResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
}

type SpotifyApiResponse struct {
    Items []Item `json:"items"`
}

type Item struct {
    Track Track `json:"track"`
}

type Track struct {
    Name string `json:"name"`
}

type TracksResponse struct {
    Tracks []string `json:"tracks"`
}

type ISpotifyService interface {
	GetAccessToken(clientId string, clientSecret string) (string, *http_errors.HttpError)
}

type SpotifyService struct {
	TokenUrl string
	ClientId string
	ClientSecret string
	AcessToken string
}

func NewSpotifyService(tokenUrl string, clientId string, clientSecret string) *SpotifyService {
	return &SpotifyService{
		TokenUrl: tokenUrl,
		ClientId: clientId,
		ClientSecret: clientSecret,
	}
}

func (svc *SpotifyService) GetAccessToken() (string, *http_errors.HttpError) {
	authToken := base64.StdEncoding.EncodeToString([]byte(svc.ClientId + ":" + svc.ClientSecret))

	req, _ := http.NewRequest("POST", svc.TokenUrl, strings.NewReader("grant_type=client_credentials"))
	req.Header.Add("Authorization", "Basic "+ authToken)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(resp.Body)

	if err != nil {
		return "", http_errors.NewHttpError(resp.StatusCode, string(body))
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", http_errors.NewHttpError(resp.StatusCode, string(body))
	}

	return string(body), nil
}

func (svc *SpotifyService) GetMusic(category string) ([]byte, *http_errors.HttpError) {

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/37i9dQZF1DX7GWlXStIq5M/tracks", strings.NewReader(""))
	req.Header.Add("Authorization", "Bearer " + svc.AcessToken)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return nil, http_errors.NewHttpError(http.StatusInternalServerError, "Intenal Server Error")
	}

	resp, err := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(resp.Body)

	log.Println(body)

	if err != nil {
		return nil, http_errors.NewHttpError(resp.StatusCode, string(body))
	}

	defer resp.Body.Close()

	var spotifyApiResponse SpotifyApiResponse
	err = json.Unmarshal([]byte(body), &spotifyApiResponse)

	if err != nil {
		return nil, http_errors.NewHttpError(resp.StatusCode, string(body))
	}

	var trackNames []string
	for _, item := range spotifyApiResponse.Items {
			trackNames = append(trackNames, item.Track.Name)
	}

  tracksResponse := TracksResponse{Tracks: trackNames}

	result, err := json.MarshalIndent(tracksResponse, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
	}

	return result, nil
}