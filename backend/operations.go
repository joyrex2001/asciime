package backend

import (
	"image"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// PostImage will handle the PostImage REST call, converting an image posted
// to this url. It will check if 'gradient' or 'width' query parameters are
// present, and use these accordingly (if not it will default to something
// sensible). The result is rendered using the 'result.tmpl' template in the
// templates folder.
func (f *handler) PostImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var width int

	// get query params
	q := r.URL.Query()
	grad := q.Get("gradient")
	if grad == "" {
		grad = ".-:+coeILCEOQB#@"
	}
	width_ := q.Get("width")
	if width_ != "" {
		width, err = strconv.Atoi(width_)
		if err != nil {
			f.Error(w, r, http.StatusInternalServerError, err)
			return
		}
	} else {
		width = 240
	}

	// retrieve form data
	r.ParseMultipartForm(5 << 20) // 5MB
	data, _, err := r.FormFile("file")
	if err != nil {
		f.Error(w, r, http.StatusInternalServerError, err)
		return
	}
	defer data.Close()

	img, _, err := image.Decode(data)
	if err != nil {
		f.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	// do magic
	art := Convert(img, width, grad)
	tmpl, err := f.GetTemplate("result.tmpl")
	if err != nil {
		f.Error(w, r, http.StatusInternalServerError, err)
		return
	}
	tmpl.Execute(w, &struct{ Art string }{Art: strings.Join(art, "<br>")})

	return
}
