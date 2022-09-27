package goalRenderTemplate

import (
	"html/template"
	"net/http"

	"github.com/Hari-Kiri/goalApplicationSettingsLoader"
	"github.com/Hari-Kiri/temboLog"
)

// HTML renderer
func Process(htmlTemplates *template.Template, responseWriter http.ResponseWriter, template string,
	appSettings *goalApplicationSettingsLoader.ApplicationSettings, request *http.Request) {
	// Open HTML page
	error := htmlTemplates.ExecuteTemplate(responseWriter, template+".html", appSettings)
	temboLog.InfoLogging("Serving", template, "page, requested from", request.RemoteAddr, "with request url path [",
		request.URL.Path, "]")
	// Failed when opening HTML page
	if error != nil {
		// Give 500 response code
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
	}
}
