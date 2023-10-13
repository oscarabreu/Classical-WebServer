package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Form struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var forms = make(map[int]Form)
var idCounter int
var formsMutex sync.Mutex

func getAllFormsHandler(w http.ResponseWriter, r *http.Request) {
	formsMutex.Lock()
	defer formsMutex.Unlock()

	var formList []Form
	for _, form := range forms {
		formList = append(formList, form)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formList)
}

func getFormHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if form, exists := forms[id]; exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(form)
	} else {
		http.Error(w, "Form not found", http.StatusNotFound)
	}
}

func createFormHandler(w http.ResponseWriter, r *http.Request) {
	var newForm Form
	if err := json.NewDecoder(r.Body).Decode(&newForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	formsMutex.Lock()
	idCounter++
	newForm.ID = idCounter
	forms[idCounter] = newForm
	formsMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newForm)
}

func updateFormHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var updatedForm Form
	if err := json.NewDecoder(r.Body).Decode(&updatedForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	formsMutex.Lock()
	defer formsMutex.Unlock()

	if _, exists := forms[id]; exists {
		updatedForm.ID = id
		forms[id] = updatedForm
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedForm)
	} else {
		http.Error(w, "Form not found", http.StatusNotFound)
	}
}

func deleteFormHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	formsMutex.Lock()
	defer formsMutex.Unlock()

	if _, exists := forms[id]; exists {
		delete(forms, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Form not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/api/forms", getAllFormsHandler)
	http.HandleFunc("/api/form", getFormHandler)
	http.HandleFunc("/api/form/create", createFormHandler)
	http.HandleFunc("/api/form/update", updateFormHandler)
	http.HandleFunc("/api/form/delete", deleteFormHandler)

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
