package models

type Player struct {
	PlayerId string `json:"playerId"`
	Player   string `json:"player"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
}
