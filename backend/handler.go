package backend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func NewHandler() *handler {
	return &handler{}
}

type handler struct {
	once sync.Once
	mux  *httprouter.Router
}

func (f *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.once.Do(f.init)
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	f.mux.ServeHTTP(w, r)
}

func (f *handler) init() {
	f.mux = httprouter.New()
	f.mux.POST("/action/upload", f.PostImage)
	f.mux.GET("/public/*filepath", f.ServeFiles("public"))
	f.mux.GET("/", f.Redirect(307, "/public/"))
}

// ServeFiles will host http files based on a filestore as determined by the
// FileStore method for given folder.
func (f *handler) ServeFiles(folder string) httprouter.Handle {
	root := f.FileStore(folder)
	fileServer := http.FileServer(root)
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		req.URL.Path = ps.ByName("filepath")
		fileServer.ServeHTTP(w, req)
	}
}

// FileStore will return a http filesystem object for given folder.
func (f *handler) FileStore(folder string) http.FileSystem {
	return http.Dir("./frontend/" + folder)
}

// Redirect will redirect the user to given location.
func (f *handler) Redirect(status int, location string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, location, status)
		return
	}
}

// Error will return an error page based on the error.tmpl template.
func (f *handler) Error(w http.ResponseWriter, r *http.Request, code int, cerr error) {
	w.WriteHeader(code)
	log.Printf("ERROR - %d: %s", code, cerr)
	tmpl, err := f.GetTemplate("error.tmpl")
	if err != nil {
		return
	}
	tmpl.Execute(w, &struct {
		Code  int
		Error string
	}{
		Code:  code,
		Error: fmt.Sprintf("%s", cerr),
	})
}

// GetTemplate will return a template instance for given file. This file should
// be present in the frontend/templates folder.
func (f *handler) GetTemplate(file string) (*template.Template, error) {
	d, err := ioutil.ReadFile("./frontend/templates/" + file)
	if err != nil {
		return nil, err
	}
	return template.New(file).Parse(string(d))
}
