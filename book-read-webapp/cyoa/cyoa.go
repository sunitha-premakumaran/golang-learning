package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

var defaultHTMLHandler = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose your own adventure</title>
    <meta name="description" content="This is an example of a meta description.">
  </head>
  <body>
		<h1>{{.Title}}</h1>
		{{range .Paragraph}}
		<p>{{.}}</p>
		{{end}}
		<div style="display:flex;justify-content: center">
			{{range .Options}}
			<a href="/{{.Arc}}">{{.Text}}</a>
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
	Arc  string `json:"arc_name"`
}

//NewHandler ..
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHTMLHandler))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
