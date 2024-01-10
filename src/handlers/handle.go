package handle

import "net/http"

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="/login">Se connecter avec Microsoft</a>`))
}
func HandleProfile(w http.ResponseWriter, r *http.Request) {
	// Vous pouvez afficher les détails de l'utilisateur sur la page de profil
	w.Write([]byte("Page de profil de l'utilisateur"))
}
