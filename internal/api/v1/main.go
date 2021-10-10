package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	m "server/internal/models"
	"time"
)

func Serve(oo *m.Operadoras) {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Authorization", "Content-Type", "Origin"}),
	)

	r := mux.NewRouter()
	r.Use(cors)
	r.HandleFunc("/api/v1",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			str := vars["string"]
			result := oo.Filter(str)
			e := json.NewEncoder(w)
			err := e.Encode(result)
			if err != nil {
				return
			}
		}).Methods("GET").Queries("string", "{string}")

	port := 8080
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler:        setHeaders(r),
	}
	log.Printf("Listening on: %d", port)
	log.Fatal(srv.ListenAndServe())
}

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Authorization, Content-Type, Origin")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
