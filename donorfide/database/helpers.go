package database

import "gorm.io/gorm"

//GetPref returns the preference with the given string, or the empty string if it doesn't exist.
func GetPref(db *gorm.DB, key string) string {
	pref := Pref{}
	db.First(&pref, "`key` = ?", key)
	return pref.Value
}

//GetUserInfo gets the information for a specific user, or returns nil if the user doesn't exist.
func GetUserInfo(db *gorm.DB, id int) User {
	user := User{}
	db.First(&user, "id = ?", id)
	return user
}
