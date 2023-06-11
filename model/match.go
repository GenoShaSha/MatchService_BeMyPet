package model

type Match struct {
	ID          int64  `json:"id"`
	ShelterID   int64  `json:"shelterId"`
	AnimalID    int64  `json:"animalId"`
	AdopterID   int64  `json:"adopterId"`
	Picture     string `json:"picture"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
	Breed       string `json:"breed"`
	Shelter     string `json:"shelter"`
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	Bio         string `json:"bio"`
	Status      string `json:"status"`
}
