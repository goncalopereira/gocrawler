package data

import (
	"net/url"
)

//Request representation
type Request struct {
	Path   string
	Host   string
	IP     string
	Scheme string
	Depth  int
	Key    string
}

//URL returns the clean URL for the page
func (r *Request) URL() string {
	return r.Scheme + "://" + r.Host + r.Path
}

//MakeRequestFromURL requests obtained from Server
func MakeRequestFromURL(newURL string) (result Request, ok bool) {
	u, _ := url.Parse(newURL)
	if u.Host == "" {
		//err nil doesn't work here? test available
		var niLRequest Request
		return niLRequest, false
	}

	return Request{Host: u.Host, Path: u.Path, Scheme: u.Scheme, Depth: 0}, true
}

//Response struct for all links from a response
type Response struct {
	Request Request
	Links   chan string
}

//ProcessedResponse struct for all next Requests processed from a response
type ProcessedResponse struct {
	Requests chan Request
}
