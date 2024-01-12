package handle

import (
	"bytes"
	"fmt"
	scraperUtils "hackathon/src/scraper"
	"net/http"
	"path/filepath"
	"text/template"

	"golang.org/x/net/html"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	tmplFiles, err := filepath.Glob(fmt.Sprintf("./templates/%s", tmpl))
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(tmplFiles) == 0 {
		fmt.Println("Template not found")
		return
	}
	ts, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := new(bytes.Buffer)
	if err := ts.Execute(buffer, p); err != nil {
		fmt.Println(err)
		return
	}
	buffer.WriteTo(w)
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound)
}
func HandleHome(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html", nil)
}
func HandleReverse(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "reverse_home.html", nil)
}

func ScrapWare(w http.ResponseWriter, r *http.Request) {
	url := "https://www.warehouse-nantes.fr/event"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de la page", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du HTML", http.StatusInternalServerError)
		return
	}
	var results []scraperUtils.Result

	targetClasses := []string{"col-md-3", "col-12", "g-2"}
	for _, class := range targetClasses {
		scraperUtils.FindElementsByClass(doc, class, &results)
	}

	RenderTemplate(w, "warehouse.html", scraperUtils.PageData{Results: results})
}
