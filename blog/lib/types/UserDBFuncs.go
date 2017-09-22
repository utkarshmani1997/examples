package types

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      User-db-type-funcs
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Migrates (or creates) the table for User
*/
func MigrateUserTable(db *gorm.DB) (err error) {
	// create own table
	err = db.Debug().AutoMigrate(&User{}).Error

	// create index on uuid
	db.Model(&User{}).Debug().AddUniqueIndex("idx_user_uuid", "uuid")

	// deal with foreign keys

	return
}

/*
GORM hook to ensure UUID is created
*/
func (T *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

/*
Creates a User record.
*/
func CreateUser(db *gorm.DB, T *User) (err error) {

	err = db.Create(T).Error

	return
}

/*
Find a User record by ID
*/
func GetUserByID(db *gorm.DB, T *User) (err error) {

	err = db.
		First(T).Error

	return
}

/*
Find a User record by UUID
*/
func GetUserByUUID(db *gorm.DB, UUID string) (T *User, err error) {
	T = NewUser()
	err = db.
		Where("uuid = ?", UUID).
		First(T).Error
	return
}

/*
Updates a User record
*/
func UpdateUser(db *gorm.DB, T *User) (err error) {
	err = db.Update(T).Error
	return
}

/*
Soft deletes the record (i.e. it has a deleted at timestamp and is not returned in queries)
*/
func DeleteUser(db *gorm.DB, T *User) (err error) {
	err = db.Delete(T).Error
	return
}

/*
Hard deletes the record, permenently deleting it from the DB
*/
func HardDeleteUser(db *gorm.DB, T *User) (err error) {
	err = db.Unscoped().Delete(T).Error
	return
}

// HOFSTADTER_BELOW
