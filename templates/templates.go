/**
 * Created with IntelliJ IDEA.
 * User: syncim
 * Date: 11/14/12
 * Time: 12:59 PM
 * To change this template use File | Settings | File Templates.
 */
package templates

import (
	"net/http"
	"html/template"

)


type Entry struct {
	Title	  string
	Content	string
}
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p Entry) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


