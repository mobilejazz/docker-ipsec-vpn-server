package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

type muxServer struct {
	router    *mux.Router
}

func main() {
	server := NewServer(mux.NewRouter())
	log.Fatal(server.ListenAndServe(":3000"))
}

func NewServer(r *mux.Router) *muxServer {
	server := muxServer{router: r}
	server.addRoutes()
	return &server
}

func (ms *muxServer) ListenAndServe(addr string) error {
	return http.ListenAndServe(":3000", ms.router)
}

func (ms *muxServer) addRoutes() {
	ms.router.HandleFunc("/adduser", ms.addUser).Methods(http.MethodPost)
	ms.router.HandleFunc("/lsusers", ms.lsUsers).Methods(http.MethodGet)
	ms.router.HandleFunc("/rmuser", ms.rmUser).Methods(http.MethodPost)

	// health check
	ms.router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	// Routing all the static paths
	ms.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public/"))))
}

type UserInfo struct {
	Name string
}

func (ms *muxServer) addUser(w http.ResponseWriter, req *http.Request) {
	var userInfo UserInfo
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&userInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userInfo.Name =="" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmd := exec.Command("/adduser.sh", userInfo.Name)
	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
	w.Write(stdout)
}

func (ms *muxServer) lsUsers(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("/lsusers.sh")
	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
	w.Write(stdout)
}

func (ms *muxServer) rmUser(w http.ResponseWriter, req *http.Request) {
	var userInfo UserInfo
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&userInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userInfo.Name =="" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmd := exec.Command("/rmuser.sh", userInfo.Name)
	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
	w.Write(stdout)
}

