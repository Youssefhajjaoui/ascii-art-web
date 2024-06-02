package asciiweb

import (
	"net/http"
)

func Web_Proc(w http.ResponseWriter, r *http.Request) {
	var tmplFile string
	path := "../pkg/AsciiWeb/"
	if r.URL.Path == "/" {
		tmplFile = path + "index.html"
	} else {
		tmplFile = path + "index-err.html"
	}
	http.ServeFile(w, r, tmplFile)
}

// var obj outil.Outils

// obj.Args = append(obj.Args, r.FormValue("input"))
// obj.Args = append(obj.Args, r.FormValue("banners"))

// fmt.Println(obj.WordToChange)
// fmt.Println(obj.Banner)

// ascii.Process_Text(&obj)
