/**
 * Created with IntelliJ IDEA.
 * User: syncim
 * Date: 11/14/12
 * Time: 12:57 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import ("github.com/gorilla/mux"
	"net/http"
	"github.com/thesyncim/mvc/controllers"

)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{teste}/{id}", controllers.HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8090", r)

}

