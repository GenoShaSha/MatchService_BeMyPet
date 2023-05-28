package service

import (
	"machingMicroService/dbaccess"
	"machingMicroService/model"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeMatch(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var matchCarrier model.Match
	err := c.BindJSON(&matchCarrier)
	if err != nil {
		log.Fatal("(RegisterUser) c.BindJSON", err)
	}

	query := `INSERT INTO matches (user1id, user2id) VALUES (?, ?)`
	res, err := db.Exec(query, matchCarrier.User1ID, matchCarrier.User2ID)
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
		err := res.Scan(&match.ID, &match.User1ID, &match.User2ID)
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

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, "updated")
	return
}

func Get1(c *gin.Context) {
	c.JSON(http.StatusOK, "YESSSSSS")
	return
}

func Get3(c *gin.Context) {
	c.JSON(http.StatusOK, "okayyy")
	return
}
