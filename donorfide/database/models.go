package database

// Pref is the table that holds system-wide settings
type Pref struct {
	ID    uint   `gorm:"autoIncrement" json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// User is the table that holds users
type User struct {
	ID        uint   `gorm:"autoIncrement,primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Level     uint8  `json:"level"`
}

type Donation struct {
	ID            uint   `gorm:"autoIncrement,primaryKey" json:"id"`
	Email         string `json:"email"`
	PaymentIntent string `json:"payment_intent"`
	Currency      string `json:"currency"`
	Amount        int64  `json:"amount"`
	Status        string `json:"status"`
}
