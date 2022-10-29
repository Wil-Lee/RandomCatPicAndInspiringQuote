package main

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// muss noch an index html weitergegeben werden
var maxSavedCombos int = 10

const (
	USER_DOES_NOT_EXIST = "User does not exist."
	MAX_COMBOS_SAVED    = "Max amount of combos are saved already."
	COMBO_ALREADY_SAVED = "Combo is saved already."
	COMBO_NOT_STORED    = "Combo was not safed yet."
)

type User struct {
	gorm.Model
	UserName   string
	Password   string
	CombosID   uint
	ComboLinks ComboLinks
}

type ComboLinks struct {
	gorm.Model
	ComboLinks []string
}

var db *gorm.DB

func initializeDataBase() {
	var dataBase, error = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if error != nil {
		panic("Couldn't connect to database.")
	}
	db = dataBase
}

// returns user by userName or nil if user does not exist
func getUser(userName string) *User {
	var user User
	if db.Where("name = ?", userName).First(&user).Error != nil {
		return nil
	}
	return &user
}

// returns true if successfull, false if User already exists or an error occured
func addUser(userName string, password string) bool {
	var user User
	if db.Where("name = ?", userName).First(&user).Error == nil {
		return false
	}

	var combos ComboLinks
	comboErr := db.Create(&combos)
	user = User{UserName: userName, Password: password, CombosID: combos.ID, ComboLinks: combos}
	userErr := db.Create(&user)

	return (comboErr.Error == nil) && (userErr.Error == nil)
}

// returns true if successfull, false if not successfull
func deleteUser(user User) bool {
	return db.Delete(&user).Error == nil
}

// Saves a combo for an user.
func addCombo(user User, comboLink string) error {
	cLinks := user.ComboLinks
	if db.First(&cLinks).Error != nil {
		return errors.New(USER_DOES_NOT_EXIST)
	}
	if len(cLinks.ComboLinks) >= maxSavedCombos {
		return errors.New(MAX_COMBOS_SAVED)
	}
	for _, cLink := range cLinks.ComboLinks {
		if cLink == comboLink {
			return errors.New(COMBO_ALREADY_SAVED)
		}
	}

	cLinks.ComboLinks = append(cLinks.ComboLinks, comboLink)
	db.Save(&cLinks)
	return nil
}

// Deletes a combo from an user.
func deleteCombo(user User, comboLink string) error {
	cLinks := user.ComboLinks
	if db.First(&cLinks).Error != nil {
		return errors.New(USER_DOES_NOT_EXIST)
	}

	for i, cLink := range cLinks.ComboLinks {
		if cLink == comboLink {
			// delete entry of comboLink
			cLinks.ComboLinks = append(cLinks.ComboLinks[:i], cLinks.ComboLinks[i+1:]...)
			db.Save(&cLinks)
			return nil
		}
	}
	return errors.New(COMBO_NOT_STORED)
}
