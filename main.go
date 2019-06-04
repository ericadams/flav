package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"os"
	"time"
)

const Port = 3000

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/flava", ServeFlava)

	http.ListenAndServe(fmt.Sprintf(":%d", Port), r)
}

func getLogger() *log.Entry {
	log.SetHandler(json.New(os.Stdout))
	return log.WithField("service", "flava").
		WithField("port", Port)
}

func ServeFlava(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./img/Flava_Flav.jpg")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer file.Close()

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)

}

func ServerTime(w http.ResponseWriter, r *http.Request) {

	time.Now().UTC()

}
