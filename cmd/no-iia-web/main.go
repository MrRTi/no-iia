package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/MrRTi/no-iia/pkg/files"
	"github.com/MrRTi/no-iia/pkg/translit"
)

type Body struct {
	Name     string
	Translit string
}

type Page struct {
	Title string
	Body  *Body
}

// --- Templates preload

func path(filename string) string {
	pwd, _ := os.Getwd()
	return pwd + "/cmd/no-iia-web/" + filename
}

var templates = template.Must(template.ParseFiles(path("root.html"), path("name.html")))

// ---

var validPath = regexp.MustCompile("^/(all|female|male)/$")

// --- Load names

func loadNames(filename string) []string {
	filePath := "./ru-pnames-list/lists/" + filename
	rows, err := files.Read(filePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Names loaded (%s)\n", filename)
	return rows
}

func loadFemaleNames() []string {
	return loadNames("female_names_rus.txt")
}

func loadMaleNames() []string {
	return loadNames("male_names_rus.txt")
}

var femaleNames = loadFemaleNames()
var maleNames = loadMaleNames()
var allNames = append(femaleNames, maleNames...)

// ---

func randomNumber(max int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	return r.Intn(max)
}

func randomName(names *[]string) *string {
	idx := randomNumber(len(*names)) - 1
	return &(*names)[idx]
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Visiting %s", r.URL)

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func rootViewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "root", nil)
}

func namesPage(names *[]string) *Page {
	name := randomName(names)
	log.Printf("Showing %s", *name)

	body := &Body{
		Name:     *name,
		Translit: translit.Transliterate(*name),
	}

	return &Page{Title: *name, Body: body}
}

func namesViewHandler(w http.ResponseWriter, r *http.Request, names *[]string) {
	p := namesPage(names)
	renderTemplate(w, "name", p)
}

func maleNamesHandler(w http.ResponseWriter, r *http.Request) {
	namesViewHandler(w, r, &maleNames)
}

func femaleNamesHandler(w http.ResponseWriter, r *http.Request) {
	namesViewHandler(w, r, &femaleNames)
}

func allNamesHandler(w http.ResponseWriter, r *http.Request) {
	namesViewHandler(w, r, &allNames)
}

// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	body := r.FormValue("body")
// 	p := &Page{Title: title, Body: []byte(body)}
// 	err := p.save()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

func main() {
	http.HandleFunc("/", rootViewHandler)
	http.HandleFunc("/female/", makeHandler(femaleNamesHandler))
	http.HandleFunc("/male/", makeHandler(maleNamesHandler))
	http.HandleFunc("/all/", makeHandler(allNamesHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
