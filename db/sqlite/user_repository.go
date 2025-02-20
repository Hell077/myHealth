package sqlite

import "gorm.io/gorm"

func (u *User) NewUser(db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Create(&User{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *User) DeleteUser(db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Delete(&User{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
