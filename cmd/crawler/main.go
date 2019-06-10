package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	crawler "github.com/goncalocool/coolcoolcool/internal/crawler"
	rules "github.com/goncalocool/coolcoolcool/internal/crawler/rules"
	data "github.com/goncalocool/coolcoolcool/internal/data"
	env "github.com/goncalocool/coolcoolcool/internal/env"
	chttp "github.com/goncalocool/coolcoolcool/internal/http"
	obs "github.com/goncalocool/coolcoolcool/internal/observability"
	sitemap "github.com/goncalocool/coolcoolcool/internal/sitemap"
	storage "github.com/goncalocool/coolcoolcool/internal/storage"
)

//main setups up I/O, env, stop, default service
func main() {
	log.SetOutput(os.Stdout)

	sTerm := make(chan os.Signal, 1)

	signal.Notify(sTerm, syscall.SIGTERM) //docker stop, ctrl+c
	signal.Notify(sTerm, syscall.SIGINT)

	externalRequestsToFilter := make(chan data.Request, env.MaxRequestsToBeFilteredQueueSize())

	go chttp.Server(env.Port(), externalRequestsToFilter)

	for {
		select {
		case <-sTerm:
			//Will try to finish crawl if Crawler in process for a clean shutdown
			return
		case req, ok := <-externalRequestsToFilter:
			//Only engages with one domain to crawl at a time
			if !ok {
				continue
			}
			//start a crawler!
			webRequests := make(chan data.Request, env.MaxWebClientRequestsQueueSize())
			webResponses := make(chan data.Response, env.MaxWebClientResponsesQueueSize())

			webShutdowns := [](chan bool){}
			for i := 1; i <= env.NumberWebClients(); i++ {
				shutdown := make(chan bool)
				go chttp.WebClient(
					chttp.BasicBodyParser{MaxLinks: env.MaxLinksPerPage()},
					webRequests,
					webResponses, shutdown)

				webShutdowns = append(webShutdowns, shutdown)
			}

			storageLinks := storage.BasicInMemoryDB{
				DB: make(map[string]data.Request, env.MaxURLSStored())}
			storageURLs := storage.BasicInMemoryDB{
				DB: make(map[string]data.Request, env.MaxURLSStored())}

			followRules := []rules.FollowRule{
				rules.RelativeLinkRule{},
				rules.ValidURLRule{},
				rules.DifferentDomainRule{},
				rules.MaxDepthRule{MaxDepth: env.MaxDepth()},
				rules.WriteKeyRule{},
				rules.WriteLinkRule{CurrentStorage: storageLinks},
				rules.WriteURLRule{CurrentStorage: storageURLs}}

			reqCommands := []crawler.RequestCommand{
				//Could try to check if page already exists but no need to block on read
				//it's checked on write
				crawler.WebClientCommand{Requests: webRequests}}

			filterRequests := make(chan data.Request, env.MaxRequestsToBeFilteredQueueSize())
			processedResponses := make(chan data.ProcessedResponse, env.MaxWebClientResponsesQueueSize())

			c := crawler.Crawler{
				FilterRequests:     filterRequests,
				WebResponses:       webResponses,
				ProcessedResponses: processedResponses,
				Commands:           reqCommands,
				FollowRules:        followRules,
				MaxLinks:           env.MaxLinksPerPage(),
				Shutdown:           make(chan bool, 1)}

			allMonitorReqs := []obs.Reqs{{Name: "filterRequests", Chan: filterRequests},
				{Name: "webRequests", Chan: webRequests}}
			allMonitorRess := []obs.Ress{{Name: "webResponses", Chan: webResponses}}
			allMonitorProcs := []obs.Procs{{Name: "processedResponses", Chan: processedResponses}}
			go obs.Monitor(allMonitorReqs, allMonitorRess, allMonitorProcs, env.MonitorDelay(), c.Shutdown)

			filterRequests <- req
			start := time.Now()
			log.Println("Req: ", req.URL())
			c.Do()

			close(filterRequests)
			close(processedResponses)
			for _, s := range webShutdowns {
				s <- true
				close(s)
			}

			siteMap := sitemap.BasicSiteMap{LinksDB: storageLinks, URLDB: storageURLs}
			siteMap.Output()
			log.Println(req.URL(), "Crawler: ", time.Now().Sub(start))
		default:
		}
	}
}
