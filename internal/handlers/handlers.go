package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/GiovannyHdz/bookings/internal/config"
	models2 "github.com/GiovannyHdz/bookings/internal/models"
	render2 "github.com/GiovannyHdz/bookings/internal/render"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config2.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config2.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handles
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render2.RenderTemplate(w, r, "home.page.tmpl", &models2.TemplateDate{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render2.RenderTemplate(w, r, "about.page.tmpl", &models2.TemplateDate{
		StringMap: stringMap,
	})
}

// Reservation renders the make reservation page and displays from
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render2.RenderTemplate(w, r, "make-reservation.page.tmpl", &models2.TemplateDate{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render2.RenderTemplate(w, r, "generals.page.tmpl", &models2.TemplateDate{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render2.RenderTemplate(w, r, "majors.page.tmpl", &models2.TemplateDate{})
}

// Availability renders the availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render2.RenderTemplate(w, r, "search-availability.page.tmpl", &models2.TemplateDate{})
}

// PostAvailability make post request
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	response := fmt.Sprintf("start date is %s and end date is %s", start, end)

	w.Write([]byte(response))
}

type jsonResponse struct {
	Ok      bool `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handler request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok: true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp,"", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render2.RenderTemplate(w, r, "contact.page.tmpl", &models2.TemplateDate{})
}
