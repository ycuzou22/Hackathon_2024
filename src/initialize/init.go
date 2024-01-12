package initialize

import (
	"encoding/json"
	"fmt"
	handle "hackathon/src/handlers"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var (
	oauthConf *oauth2.Config
	port      = ":8080"
)

func StartServ() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("http://localhost" + port + " 🚀")
	http.HandleFunc("/", handle.HandleMain)
	http.HandleFunc("/reverse", handle.HandleReverse)
	http.HandleFunc("/login", MicrosoftLogin)
	http.HandleFunc("/callback", MicrosoftCallback)
	http.HandleFunc("/warehouse", handle.ScrapWare)
	http.ListenAndServe(port, nil)
}

// init .env and OAuth2 configuration
func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	// Initialize OAuth2 configuration
	oauthConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     microsoft.AzureADEndpoint("common"),
		Scopes:       []string{"User.Read"},
	}
}

// func for Microsoft connexion

func MicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusFound)
}

func MicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := oauthConf.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Échec de l'échange de code", http.StatusInternalServerError)
		return
	}

	userDetails, err := getUserDetails(token.AccessToken)
	if err != nil {
		http.Error(w, "Erreur", http.StatusInternalServerError)
		return
	}

	fmt.Println("user : ", userDetails)

	http.Redirect(w, r, "/reverse", http.StatusFound)
}

func getUserDetails(accessToken string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userDetails map[string]interface{}
	err = json.Unmarshal(body, &userDetails)
	if err != nil {
		return nil, err
	}

	return userDetails, nil
}
