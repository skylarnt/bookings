package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/skylarnt/bookings/internal/config"
	"github.com/skylarnt/bookings/internal/models"
	"github.com/skylarnt/bookings/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) MajorsSuits(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "search-availablity.page.html", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.Form.Get("start")
	endDate := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", startDate, endDate)))

}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Got here"))
	// render.RenderTemplate(w, "make-reservations.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
