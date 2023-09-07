package repo

import (
	"github.com/Belajar-Golang-Dimas/belajar-gorm.git/models"
	"gorm.io/gorm"
)

func GormQuerySelect(db *gorm.DB) (*models.Bookmarks, error) {
	var bookmark models.Bookmarks
	err := db.Select("bookmarks.title, users.username").Joins("JOIN users ON bookmarks.user_id = users.id").Find(&bookmark).Where("users.id = 1").Error
	if err != nil {
		return nil, err
	}

	return &bookmark, nil
}

func GormCreateDataUser(db *gorm.DB, user *models.User) (bool, error) {
	err := db.Create(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
