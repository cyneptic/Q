package entities

type Message struct {
	DBModel
	Content  string   `json:"content"`
	Receiver Receiver `gorm:"type:VARCHAR(255)" json:"receiver"`
	Sender   string   `json:"sender"`
}

type DTOMessage struct {
	Content  string
	Receiver Receiver
	Sender   string
}
