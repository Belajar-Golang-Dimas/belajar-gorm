package main

import (
	"fmt"
	"log"

	"github.com/Belajar-Golang-Dimas/belajar-gorm.git/app"
	"github.com/Belajar-Golang-Dimas/belajar-gorm.git/repo"
)

func main() {
	db := app.NewDB()
	// isCreated, err := repo.GormCreateDataUser(db, &models.User{
	// 	ID:       2,
	// 	Username: "dimasfebriyanto",
	// 	Password: "aa11bb22",
	// })
	// if err != nil {
	// 	log.Fatal("Failed created user", err.Error())
	// }

	// if isCreated {
	allData, err := repo.GormQuerySelect(db)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(allData)

	// }

}
