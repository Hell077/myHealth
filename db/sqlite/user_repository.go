package sqlite

import "gorm.io/gorm"

func (u *User) NewUser(db *gorm.DB, username, name string, tgID int64) error {
	user := User{
		TGID:     tgID,
		Username: username,
		Name:     name,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	*u = user
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

func (u *User) ExistsByTelegramID(db *gorm.DB, tgID int64) (bool, error) {
	var count int64
	if err := db.Model(&User{}).Where("tg_id = ?", tgID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
