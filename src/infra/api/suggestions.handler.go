package api

import (
	"encoding/json"
	"ifood-backend-test/src/application/services"
	"ifood-backend-test/src/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func MakeSuggestionHandler(r *chi.Mux) {

	r.Get("/api/latitude/{lat}/longitude/{long}", func(w http.ResponseWriter, r *http.Request) {
		
		config := config.LoadConfig()

		// lat := r.URL.Query().Get("lat")
		// long := r.URL.Query().Get("long")
		
		lat := chi.URLParam(r, "lat")
    long := chi.URLParam(r, "long")
		
		svc := services.NewWeatherService(config.WeatherUrlApi, config.WeatherApiKey)

		resp, httpError, err := svc.FindByLocation(lat, long)

		if err != nil {
      log.Println("Error occurred:", err)
			return
		}

		if httpError != nil {
			w.Header().Set("Content-type", "application/json")
    	w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}

		w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	})

	r.Get("/api/city/{city}", func(w http.ResponseWriter, r *http.Request) {
		config := config.LoadConfig()

		city := chi.URLParam(r, "city")

		svc := services.NewWeatherService(config.WeatherUrlApi, config.WeatherApiKey)

		resp, httpError, err := svc.FindByCity(city)

		if err != nil {
      log.Println("Error occurred:", err)
			return
		}

		if httpError != nil {
			w.Header().Set("Content-type", "application/json")
    	w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}

		w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	})

	r.Get("/api/music", func(w http.ResponseWriter, r *http.Request) {
		config := config.LoadConfig()

		svc := services.NewSpotifyService("https://accounts.spotify.com/api/token", config.SporifyClientId, config.SporifyClientSecret)

		resp, httpError := svc.GetAccessToken();

		if httpError != nil {
			w.Header().Set("Content-type", "application/json")
    	w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}

		var tokenReponse services.TokenResponse
		err := json.Unmarshal([]byte(resp), &tokenReponse)
		
		if err != nil {
			w.Header().Set("Content-type", "application/json")
    	w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		svc.AcessToken = tokenReponse.AccessToken
		
		tracks, httpError := svc.GetMusic("")

		if httpError != nil || err != nil{
			w.Header().Set("Content-type", "application/json")
    	w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}

		w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)
		w.Write([]byte(tracks))
	})
	
}