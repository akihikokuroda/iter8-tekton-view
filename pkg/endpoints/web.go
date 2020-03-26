package endpoints

import (
        "fmt"
	"os"
	"net/http"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restful "github.com/emicklei/go-restful"
)

// RegisterWeb registers the web
func (r Resource) RegisterWeb(container *restful.Container) {
	fs := http.FileServer(http.Dir("/var/run/ko"))
	container.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        fmt.Println(r)
		fmt.Println(os.Stat("/var/run/ko/index.html"))
		fs.ServeHTTP(w, r)
	}))
}
