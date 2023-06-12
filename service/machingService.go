package service

import (
	"machingMicroService/dbaccess"
	"machingMicroService/model"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MatchResponse struct {
	Matches []model.Match `json:"matches"`
}

func MakeMatch(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var matchCarrier model.Match
	err := c.BindJSON(&matchCarrier)
	if err != nil {
		log.Fatal("(RegisterUser) c.BindJSON", err)
	}

	fmt.Println(matchCarrier)

	query := `INSERT INTO matches (shelter_id, animal_id, adopter_id, picture, first_name, last_name, date_of_birth, gender, type, breed, shelter, address, postal_code, bio, status,adopter_email) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`
	res, err := db.Exec(query, matchCarrier.ShelterID, matchCarrier.AnimalID, matchCarrier.AdopterID, matchCarrier.Picture, matchCarrier.FirstName, matchCarrier.LastName, matchCarrier.DateOfBirth, matchCarrier.Gender, matchCarrier.Type, matchCarrier.Breed, matchCarrier.Shelter, matchCarrier.Address, matchCarrier.PostalCode, matchCarrier.Bio, matchCarrier.Status, matchCarrier.AdopterEmail)
	if err != nil {
		log.Fatal("(RegisterUser) db.Exec", err)
	}

	matchCarrier.ID, err = res.LastInsertId()
	if err != nil {
		log.Fatal("(CreateProduct) res.LastInsertId", err)
	}

	c.JSON(http.StatusOK, matchCarrier)
}

func GetMatches(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	query := "SELECT * FROM matches"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetProducts) db.Query", err)
	}

	matches := []model.Match{}
	for res.Next() {
		var match model.Match
		err := res.Scan(&match.ID, &match.ShelterID, &match.AnimalID, &match.AdopterID, &match.Picture, &match.FirstName, &match.LastName, &match.DateOfBirth, &match.Gender, &match.Type, &match.Breed, &match.Shelter, &match.Address, &match.PostalCode, &match.Bio, &match.Status)
		if err != nil {
			log.Fatal("(GetProducts) res.Scan", err)
		}
		matches = append(matches, match)
	}

	type MatchResponse struct {
		Matches []model.Match `json:"matches"`
	}

	// Wrap the users array within an object
	response := MatchResponse{Matches: matches}

	c.JSON(http.StatusOK, response)
}

func GetAdopterMatches(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	type RequestBody struct {
		AdopterID int64 `json:"adopter_id"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		return
	}

	adopterID := requestBody.AdopterID
	fmt.Println(adopterID)

	query := "SELECT * FROM matches WHERE adopter_id = ?"
	res, err := db.Query(query, adopterID)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetProducts) db.Query", err)
	}

	matches := []model.Match{}
	for res.Next() {
		var match model.Match
		err := res.Scan(&match.ID, &match.ShelterID, &match.AnimalID, &match.AdopterID, &match.Picture, &match.FirstName, &match.LastName, &match.DateOfBirth, &match.Gender, &match.Type, &match.Breed, &match.Shelter, &match.Address, &match.PostalCode, &match.Bio, &match.Status, &match.AdopterEmail)
		if err != nil {
			log.Fatal("(GetProducts) res.Scan", err)
		}
		matches = append(matches, match)
	}

	// Wrap the users array within an object
	response := MatchResponse{Matches: matches}

	c.JSON(http.StatusOK, response)
}

func GetShelterMatches(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	type RequestBody struct {
		ShelterID int64 `json:"shelter_id"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		return
	}

	shelterID := requestBody.ShelterID
	fmt.Println(shelterID)

	query := "SELECT * FROM matches WHERE shelter_id = ?"
	res, err := db.Query(query, shelterID)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetProducts) db.Query", err)
	}

	matches := []model.Match{}
	for res.Next() {
		var match model.Match
		err := res.Scan(&match.ID, &match.ShelterID, &match.AnimalID, &match.ShelterID, &match.Picture, &match.FirstName, &match.LastName, &match.DateOfBirth, &match.Gender, &match.Type, &match.Breed, &match.Shelter, &match.Address, &match.PostalCode, &match.Bio, &match.Status, &match.AdopterEmail)
		if err != nil {
			log.Fatal("(GetProducts) res.Scan", err)
		}
		matches = append(matches, match)
	}

	// Wrap the users array within an object
	response := MatchResponse{Matches: matches}

	c.JSON(http.StatusOK, response)
}

func UpdateMatchStatus(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	type RequestBody struct {
		MatchID int64  `json:"matchid"`
		Status  string `json: "status"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		return
	}

	query := "UPDATE matches SET Status = ? WHERE ID = ?"
	res, err := db.Query(query, requestBody.Status, requestBody.MatchID)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetProducts) db.Query", err)
	}

	c.JSON(http.StatusOK, "OK")
}
