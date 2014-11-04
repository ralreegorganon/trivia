package trivia

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateRouter(server *HTTPServer) (*mux.Router, error) {
	r := mux.NewRouter()
	m := map[string]map[string]HttpApiFunc{
		"GET": {
			"/api/trivia/random":      server.GetRandomTrivia,
			"/api/trivia":             server.GetAllTrivia,
			"/api/trivia/{id:[0-9]+}": server.GetTrivia,
		},
	}

	for method, routes := range m {
		for route, handler := range routes {
			localRoute := route
			localHandler := handler
			localMethod := method
			f := makeHttpHandler(localMethod, localRoute, localHandler)

			r.Path(localRoute).Methods(localMethod).HandlerFunc(f)
		}
	}

	return r, nil
}

func makeHttpHandler(localMethod string, localRoute string, handlerFunc HttpApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCorsHeaders(w, r)
		if err := handlerFunc(w, r, mux.Vars(r)); err != nil {
			httpError(w, err)
		}
	}
}

func writeCorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
}

type HttpApiFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string) error

type HTTPServer struct {
	DB *DB
}

func NewHTTPServer(db *DB) *HTTPServer {
	s := &HTTPServer{
		DB: db,
	}

	return s
}

func writeJSON(w http.ResponseWriter, code int, thing interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	val, err := json.Marshal(thing)
	w.Write(val)
	return err
}

func httpError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), statusCode)
	}
}

func (s *HTTPServer) GetRandomTrivia(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	trivia, err := s.DB.GetRandomTrivia()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, trivia)

	return nil
}

func (s *HTTPServer) GetTrivia(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	trivia, err := s.DB.GetTrivia(id)

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, emberify("trivium", trivia))

	return nil
}

func (s *HTTPServer) GetAllTrivia(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	trivia, err := s.DB.GetAllTrivia()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, emberify("trivia", trivia))

	return nil
}

func emberify(name string, thing interface{}) interface{} {
	return map[string]interface{}{name: thing}
}
