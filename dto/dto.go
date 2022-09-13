package dto

type Status uint16

type Message struct {
	Ip        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
	Status    Status `json:"status"`
}

type Error struct {
	Ip    string `json:"ip"`
	Error string `json:"error"`
}

type Flap struct {
	// Ip     string `json:"ip"`
	Start  int64 `json:"start"`
	Amount int64 `json:"amount"`
}
