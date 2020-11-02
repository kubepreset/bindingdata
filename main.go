package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func serveEnv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varName := vars["varName"]

	value := os.Getenv(varName)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(value))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	sbr := os.Getenv("SERVICE_BINDING_ROOT")
	w.Header().Set("Content-Type", "application/json")
	if sbr == "" {
		w.Write([]byte("{}\n"))
		err := errors.New("env var SERVICE_BINDING_ROOT not set")
		log.Println(err)
		return
	}

	fi, err := os.Stat(sbr)
	if err != nil {
		w.Write([]byte("{}\n"))
		err := errors.New("Error reading file: " + sbr)
		log.Println(err)
		return
	}
	if !fi.IsDir() {
		w.Write([]byte("{}\n"))
		err := errors.New("Not a directory: " + sbr)
		log.Println(err)
		return
	}

	files, err := ioutil.ReadDir(sbr)
	if err != nil {
		w.Write([]byte("{}\n"))
		err := errors.New("Not a directory: " + sbr)
		log.Println(err)
		return
	}

	content := make(map[string][]map[string]string)
	for _, f := range files {
		if f.IsDir() {
			content[f.Name()] = make([]map[string]string, 0)
			files2, err := ioutil.ReadDir(path.Join(sbr, f.Name()))
			if err != nil {
				continue
			}
			for _, g := range files2 {
				fileContent, err := ioutil.ReadFile(path.Join(sbr, f.Name(), g.Name()))
				if err != nil {
					continue
				}
				content[f.Name()] = append(content[f.Name()], map[string]string{g.Name(): string(fileContent)})
				if g.IsDir() {
					continue
				}

			}
		}
	}

	b, err := json.Marshal(content)
	if err != nil {
		w.Write([]byte("{}\n"))
		err := errors.New("Cannot marshal: " + fmt.Sprintf("%#v", content))
		log.Println(err)
		return
	}

	w.Write(b)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/env/{varName}", serveEnv).Methods("GET")
	r.HandleFunc("/files", serveFiles).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":7080")
}
