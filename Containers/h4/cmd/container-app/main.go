package main

import (
	"container-app/config"
	"log"
	"net/http"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal("failed while reading config", err)
	}
	app := createApp(c)
	log.Printf("server is listening on %s", c.Address)
	err = app.ListenAndServe()
	if err != nil {
		log.Fatal("failed while listening server", err)
	}
}

// create hello world server from standrad library
func createApp(config *config.Config) *http.Server {
	helloWorldHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	server := &http.Server{
		Addr:    config.Address,
		Handler: helloWorldHandler,
	}

	return server
}
