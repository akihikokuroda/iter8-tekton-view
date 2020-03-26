package endpoints

import (
        "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restful "github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ExperimentStatus struct {
	Name                 string `json:"name"`
	Baseline             string `json:"baseline"`
	Candidate            string `json:"candidate"`
	CurrentIteration     int64 `json:"currentiteration"`
	Message              string `json:"message"`
	BaselineTraffic      int64 `json:"baselinetraffic"`
	CandidateTraffic     int64 `json:"candidatetraffic"`
	Conclusions          []string `json:"conclusions"`
}

// RegisterMonitorbService registers the monitor service
func (r Resource) RegisterMonitorService(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/monitor").Consumes(restful.MIME_JSON, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_JSON)
	ws.Route(ws.GET("/{experiment-id}").To(r.checkStatus))

	container.Add(ws)
}
func (r Resource) checkStatus(request *restful.Request, response *restful.Response) {
        fmt.Println("Hello, world.")
        id := request.PathParameter("experiment-id")
	get, err := r.experimentClient.Namespace("kabanero").Get( id, metav1.GetOptions{})
	if err != nil {
            fmt.Println(err)
	}
	name, found, err :=unstructured.NestedString(get.Object, "spec", "targetService", "name")
	if found && err == nil {
	    fmt.Printf("Name: %v\n", name)
	}
	baseline, found, err :=unstructured.NestedString(get.Object, "spec", "targetService", "baseline")
	if found && err == nil {
	    fmt.Printf("Baseline: %v\n", baseline)
	}
	candidate, found, err :=unstructured.NestedString(get.Object, "spec", "targetService", "candidate")
	if found && err == nil {
	    fmt.Printf("Candidate: %v\n", candidate)
	}
	currentIteration, found, err :=unstructured.NestedInt64(get.Object, "status", "currentIteration")
	if found && err == nil {
	    fmt.Printf("Current Iteration: %v\n", currentIteration)
	}
	message, found, err :=unstructured.NestedString(get.Object, "status", "message")
	if found && err == nil {
	    fmt.Printf("Message: %v\n", message)
	}
	baselineTraffic, found, err :=unstructured.NestedInt64(get.Object, "status", "trafficSplitPercentage", "baseline")
	if found && err == nil {
	    fmt.Printf("Baseline Traffic: %v\n", baselineTraffic)
	}
	candidateTraffic, found, err :=unstructured.NestedInt64(get.Object, "status", "trafficSplitPercentage", "candidate")
	if found && err == nil {
	    fmt.Printf("Candidate Traffic: %v\n", candidateTraffic)
	}
	conclusions, found, err :=unstructured.NestedStringSlice(get.Object, "status", "assessment", "conclusions")
	if found && err == nil {
	    for _, conclude := range conclusions {
	        fmt.Printf("Conclude: %v\n", conclude)
	    }
	}
        status := ExperimentStatus{
	    Name:             name,
	    Baseline:         baseline,
	    Candidate:        candidate,
	    CurrentIteration: currentIteration,
	    Message:          message,
	    BaselineTraffic:  baselineTraffic,
	    CandidateTraffic: candidateTraffic,
	    Conclusions:      conclusions,
        }
	
	response.WriteEntity(status)
}