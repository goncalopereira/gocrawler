package http

import (
	"net/http"

	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//NewRequestHandler Parses and sends requests to Crawler
type NewRequestHandler struct {
	FilterRequests chan<- data.Request
}

//Post Parses Post and sends to Crawler queue
func (nrh NewRequestHandler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			return
		}

		req, ok := data.MakeRequestFromURL(r.Form.Get("url"))

		if ok {
			nrh.FilterRequests <- req
		}
	}
}

//Server starts new HTTP server that receives external requests
//sends request to be filtered on Crawler
func Server(port string, filterRequests chan<- data.Request) error {
	handler := http.NewServeMux()
	handler.HandleFunc("/", NewRequestHandler{FilterRequests: filterRequests}.Post)
	return http.ListenAndServe(":"+port, handler)
}
