package obs

import (
	"time"

	"github.com/goncalopereira/gocrawler/internal/data"
)

type Reqs struct {
	Name string
	Chan chan data.Request
}
type Ress struct {
	Name string
	Chan chan data.Response
}

type Procs struct {
	Name string
	Chan chan data.ProcessedResponse
}

func Monitor(requesters []Reqs, responsers []Ress, processeds []Procs, monitorDelay int, shutdownCrawler chan bool) {

	for {
		var state = true
		time.Sleep(time.Duration(monitorDelay) * time.Second)

		for _, reqs := range requesters {
			//log.Printf("Req %s cap %s len %s", reqs.Name, strconv.Itoa(cap(reqs.Chan)), strconv.Itoa(len(reqs.Chan)))
			if len(reqs.Chan) != 0 {
				state = false
			}
		}
		for _, res := range responsers {
			//log.Printf("Res %s cap %s len %s", res.Name, strconv.Itoa(cap(res.Chan)), strconv.Itoa(len(res.Chan)))
			if len(res.Chan) != 0 {
				state = false
			}
		}

		for _, pres := range processeds {
			//	log.Printf("Res %s cap %s len %s", pres.Name, strconv.Itoa(cap(pres.Chan)), strconv.Itoa(len(pres.Chan)))
			if len(pres.Chan) != 0 {
				state = false
			}
		}

		//log.Println("GoRoutines: ", runtime.NumGoroutine())
		//pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

		if state {
			//poll channels now and then for size
			shutdownCrawler <- true
			return //close Monitor
		}
	}
}
