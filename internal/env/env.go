package env

import (
	"os"
	"strconv"
)

//GetEnv returns env variable or default
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

//GetEnvInt returns env variable or default as Int
func GetEnvInt(key string, defaultValue int) int {
	envString := GetEnv(key, strconv.Itoa(defaultValue))
	v, err := strconv.Atoi(envString)

	if err != nil {
		return defaultValue
	}

	return v
}

//GetEnv default Port setup
func Port() string {
	return GetEnv("PORT", "8080")
}

//MaxDepth number of links to follow from original link
func MaxDepth() int {
	return GetEnvInt("MAX_DEPTH", 100)
}

//MaxLinksPerPage Max number of links to retrieve in a single page before returning
func MaxLinksPerPage() int {
	return GetEnvInt("MAX_LINKS", 500)
}

//MaxURLSStored for the BasicInMemoryDB we have a fixed max number of URLs
func MaxURLSStored() int {
	return GetEnvInt("MAX_URLS", 1000000)
}

//MaxWebClientRequestsQueueSize Max number of requests queued for all web clients (post filter)
func MaxWebClientRequestsQueueSize() int {
	return GetEnvInt("MAX_WEB_CLIENT_REQUESTS", 10000)
}

//MaxWebClientResponsesQueueSize Max number of responses queued from all web clients (pre filter)
func MaxWebClientResponsesQueueSize() int {
	return GetEnvInt("MAX_WEB_CLIENT_RESPONSES", 10000)
}

//MaxRequestsToBeFilteredQueueSize Max number of requests to filter before pushed to all web clients
func MaxRequestsToBeFilteredQueueSize() int {
	return GetEnvInt("MAX_FILTER_REQUESTS", 1000)
}

//NumberWebClients parallel web clients querying domain
func NumberWebClients() int {
	return GetEnvInt("WEB_CLIENTS", 15)
}

//MonitorDelay seconds between displaying monitoring
func MonitorDelay() int {
	return GetEnvInt("MONITOR_DELAY", 2)
}
