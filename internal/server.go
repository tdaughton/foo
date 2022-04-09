package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type FooNameDocument struct {
	Name string `json:"name"`
}

func (s *httpServer) handlePost(writer http.ResponseWriter, req *http.Request) {
	var fooReq FooNameDocument
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&fooReq)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	} else if fooReq.Name == "" {
		response := "Foo name is required."
		http.Error(writer, response, http.StatusBadRequest)
		return
	} else {
		foo := s.FooManager.insert(fooReq.Name)
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Transfer-Encoding", "chunked")
		json.NewEncoder(writer).Encode(foo)
	}
	return
}

func (s *httpServer) handleGet(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	foo, found := s.FooManager.retrieve(id)
	if found == false {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Transfer-Encoding", "chunked")
		json.NewEncoder(writer).Encode(foo)
	}
	return
}

func (s *httpServer) handleDelete(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	found := s.FooManager.delete(id)
	if found == false {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
	return
}

type httpServer struct {
	FooManager *FooManager
}

func NewHTTPServer(port string) *http.Server {
	server := &httpServer{
		FooManager: &FooManager{},
	}
	router := mux.NewRouter()
	router.HandleFunc("/foo/{id}", server.handleDelete).Methods("DELETE")
	router.HandleFunc("/foo/{id}", server.handleGet).Methods("GET")
	router.HandleFunc("/foo", server.handlePost).Methods("POST")
	return &http.Server{
		Addr:    port,
		Handler: router,
	}
}
