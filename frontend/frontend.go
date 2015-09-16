package frontend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pdxjohnny/discovery/random"
)

var uuid = random.String(10)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2048)
	r.ParseForm()
	log.Println(r.Form)
	log.Println(r.PostForm)
	log.Println(r.MultipartForm)
	fmt.Fprintf(w, "Hello from %s\n", uuid)
}
