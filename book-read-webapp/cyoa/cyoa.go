package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"text/template"
)

var defaultHTMLHandler = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
	<title>Choose your own adventure</title>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.2/css/bulma.min.css">
    <meta name="description" content="This is an example of a meta description.">
  </head>
  <body>
		<h1 class="subtitle is-1 has-text-centered">{{.Title}}</h1>
		{{range .Paragraph}}
		<p class="container">{{.}}</p>
		{{end}}
		<div class="field is-grouped is-grouped-centered is-grouped-multiline">
			{{range .Options}}
			<div class="control">
			<a class="button is-link" href="/{{.Arc}}">{{.Text}}</a>
			</div>
			{{end}}
		</div>
  </body>
</html>
`

// JSONStory ...
func JSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story ...
type Story map[string]Chapter

// Chapter ...
type Chapter struct {
	Title     string          `json:"title"`
	Paragraph []string        `json:"story"`
	Options   []ChapterOption `json:"options"`
}

// ChapterOption ..
type ChapterOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

//NewHandler ..
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	tpl := template.Must(template.New("").Parse(defaultHTMLHandler))
	err := tpl.Execute(w, h.s[path])
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
