package webserver

import (
	"log"
	"net/http"
)

type WebServer struct {
	Router        http.ServeMux
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func New(port string) *WebServer {
	return &WebServer{
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: port,
	}
}

func (w *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	w.Handlers[path] = handler
}

func (w *WebServer) Start() {
	for path, handler := range w.Handlers {
		w.Router.HandleFunc(path, handler)
	}

	log.Println("Starting web server...")
	if err := http.ListenAndServe(w.WebServerPort, &w.Router); err != nil {
		panic(err)
	}
}
