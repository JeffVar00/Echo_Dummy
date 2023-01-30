package entity

// ENTITY, THIS IS THE STRUCTURE OF THE DATA THAT WILL BE RETURNED BY THE REPOSITORY
// SINCE WE ARE NOT USING A DATABASE, WE WILL USE A MODEL FOR NOW

type Player struct {
	PlayerId string `db:"playerId"`
	Player   string `db:"player"`
	Country  string `db:"country"`
	Currency string `db:"currency"`
}
