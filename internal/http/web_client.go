package http

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//WebClient that runs through external requests to crawl
func WebClient(parser BodyParser, webRequests <-chan data.Request, responses chan<- data.Response, shutdown <-chan bool) error {
	for {
		select {
		case r, ok := <-webRequests:

			if ok {
				response, err := GetURL(r, parser)

				if err != nil {
					continue
				}

				responses <- response
			}
		case <-shutdown:
			return nil
		}

	}
}

//GetURL sends all links found in page back
//http.Client misbehaves and might create too many goroutines
func GetURL(r data.Request, parser BodyParser) (data.Response, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	req, err := http.NewRequest("GET", r.URL(), nil)
	if err != nil {
		return data.Response{}, err
	}

	//https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	client := &http.Client{

		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},

			Dial: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		}}

	resp, err := client.Do(req.WithContext(ctx))

	if err != nil {
		return data.Response{}, err
	}
	defer resp.Body.Close()

	response := parser.Parse(resp.Body, r)

	return response, nil
}
