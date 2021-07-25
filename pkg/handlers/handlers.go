package handlers

import (
	"net/http"

	"github.com/alvinashiatey/app5/pkg/config"
	"github.com/alvinashiatey/app5/pkg/models"
	"github.com/alvinashiatey/app5/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hi from Alvin"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Paynin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "paynin.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Kakra(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "kakra.page.tmpl", &models.TemplateData{})
}
