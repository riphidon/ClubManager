package server

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/pkg/errors"
	"github.com/riphidon/clubmanager/models"
)

type Page struct {
	Title       string
	Message     string
	BeltList    []string
	WeightList  []string
	AgeList     []string
	User        models.ClubUser
	Editable    models.ClubUser
	UserErr     string
	UserList    []*models.ClubUser
	ListByParam []*models.ClubUser
	UserProfile models.ClubUser
	Events      []*models.Event
	Event       *models.Event
	EventTitle  string
	Infos       []*models.Info
	Info        *models.Info
	LogState    bool
}

func ParseTemplate(path string) *template.Template {
	var tmpl = template.Must(template.New("").ParseGlob(path + "*.html"))
	return tmpl
}

func RenderPage(w http.ResponseWriter, path, template string, pageData interface{}) error {
	tmpl := ParseTemplate(path)
	bufTemplate := &bytes.Buffer{}
	err := tmpl.ExecuteTemplate(bufTemplate, template, pageData)
	if err != nil {
		return errors.Wrap(err, "unable to render template")
	} else {
		bufTemplate.WriteTo(w)
	}
	return nil
}
