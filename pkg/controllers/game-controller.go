package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kelechiigbe/rock-game/pkg/models"
	"github.com/kelechiigbe/rock-game/pkg/utils"
)

var NewChoice models.Choice

func GetAllChoices(w http.ResponseWriter, r *http.Request) {
	allChoices := models.GetAllChoices()
	res, _ := json.Marshal(allChoices)
	setupCORS(&w, r)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetChoiceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	choice_id := vars["choice_id"]
	id, err := strconv.ParseInt(choice_id, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}
	result := models.GetChoiceById(int(id))
	res, _ := json.Marshal(result)
	setupCORS(&w, r)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRandomChoice(w http.ResponseWriter, r *http.Request) {
	result := models.GetChoice()
	res, _ := json.Marshal(result)
	setupCORS(&w, r)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func PlayGame(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	playerChoice := &models.PlayerChoice{}
	utils.ParseBody(r, playerChoice)
	result := models.PlayGame(*playerChoice)
	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
}
