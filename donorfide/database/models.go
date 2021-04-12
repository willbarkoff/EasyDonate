package database

// Prefs is the table that holds system-wide settings
type Prefs struct {
	ID    uint `gorm:"autoIncrement"`
	Key   string
	Value string
}

// Users is the table that holds users
type Users struct {
	ID        uint `gorm:"autoIncrement"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	Level     uint8
}
