package main

import (
    "os"
    "net/http"
    restful "github.com/emicklei/go-restful"
    "github.com/iter8-tekton/pkg/endpoints"
)

func main() {
    // Set up routes
    wsContainer := restful.NewContainer()
    wsContainer.Router(restful.CurlyRouter{})
    r, _ := endpoints.NewResource()
    r.RegisterMonitorService(wsContainer)
    r.RegisterWeb(wsContainer)

    port := ":8080"
    portnum := os.Getenv("PORT")
    if portnum != "" {
        port = ":" + portnum
    }
    server := &http.Server{Addr: port, Handler: wsContainer}
    server.ListenAndServe()
}
