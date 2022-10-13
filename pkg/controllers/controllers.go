package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/lordscoba/create-id-card-golang/pkg/utils"
	"github.com/lordscoba/create-id-card-golang/pkg/models"
)


var NewCard models.Card

func GetCard(w http.ResponseWriter, r *http.Request){
	newCards:=models.GetAllCards()
	res, _ :=json.Marshal(newCards)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCardById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	cardId := vars["cardId"]
	ID, err:= strconv.ParseInt(cardId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	cardDetails, _:= models.GetCardById(ID)
	res, _ := json.Marshal(cardDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCard(w http.ResponseWriter, r *http.Request){
	CreateCard := &models.Card{}
	utils.ParseBody(r, CreateCard)
	b:= CreateCard.CreateCard()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCard(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	cardId := vars["cardId"]
	ID, err := strconv.ParseInt(cardId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	card := models.DeleteCard(ID)
	res, _ := json.Marshal(card)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCard(w http.ResponseWriter, r *http.Request){
	var updateCard = &models.Card{}
	utils.ParseBody(r, updateCard)
	vars := mux.Vars(r)
	cardId := vars["cardId"]
	ID, err := strconv.ParseInt(cardId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	cardDetails, db:=models.GetCardById(ID)
	if updateCard.Name != ""{
		cardDetails.Name = updateCard.Name
	}
	if updateCard.YearOfBirth != 0{
		cardDetails.YearOfBirth = updateCard.YearOfBirth
	}
  if updateCard.State != ""{
		cardDetails.State = updateCard.State
	}
  if updateCard.Country != ""{
		cardDetails.Country = updateCard.Country
	}
  if updateCard.Height != ""{
		cardDetails.Height = updateCard.Height
	}
  if updateCard.Weight != ""{
		cardDetails.Weight = updateCard.Weight
	}

  if updateCard.About != ""{
		cardDetails.About = updateCard.About
	}

	db.Save(&cardDetails)
	res, _ := json.Marshal(cardDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
