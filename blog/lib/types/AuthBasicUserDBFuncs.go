package types

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      AuthBasicUser-db-type-funcs
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Migrates (or creates) the table for AuthBasicUser
*/
func MigrateAuthBasicUserTable(db *gorm.DB) (err error) {
	// create own table
	err = db.Debug().AutoMigrate(&AuthBasicUser{}).Error

	// create index on uuid
	db.Model(&AuthBasicUser{}).Debug().AddUniqueIndex("idx_auth_basic_user_uuid", "uuid")

	// deal with foreign keys
	db.Model(&AuthBasicUser{}).Debug().AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	return
}

/*
GORM hook to ensure UUID is created
*/
func (T *AuthBasicUser) BeforeCreate(scope *gorm.Scope) error {
	T.UUID = uuid.New().String()
	return nil
}

/*
Creates a AuthBasicUser record.
*/
func CreateAuthBasicUser(db *gorm.DB, T *AuthBasicUser, UserUUID string) (err error) {

	// user-defined
	// it's not a view
	var User User
	err = db.
		Where("user_uuid = ?", UserUUID).
		First(User).Error
	if err != nil {
		return err
	}
	T.UserID = User.ID

	// other relations? (has-one, has-many, many-to-many)
	err = db.Create(T).Error

	return
}

/*
Find a AuthBasicUser record by ID
*/
func GetAuthBasicUserByID(db *gorm.DB, T *AuthBasicUser, UserUUID string) (err error) {

	err = db.
		Where("user_uuid = ?", UserUUID).
		First(T).Error

	return
}

/*
Find a AuthBasicUser record by UUID
*/
func GetAuthBasicUserByUUID(db *gorm.DB, UUID string) (T *AuthBasicUser, err error) {
	T = NewAuthBasicUser()
	err = db.
		Where("uuid = ?", UUID).
		First(T).Error
	return
}

/*
Updates a AuthBasicUser record
*/
func UpdateAuthBasicUser(db *gorm.DB, T *AuthBasicUser) (err error) {
	err = db.Update(T).Error
	return
}

/*
Soft deletes the record (i.e. it has a deleted at timestamp and is not returned in queries)
*/
func DeleteAuthBasicUser(db *gorm.DB, T *AuthBasicUser) (err error) {
	err = db.Delete(T).Error
	return
}

/*
Hard deletes the record, permenently deleting it from the DB
*/
func HardDeleteAuthBasicUser(db *gorm.DB, T *AuthBasicUser) (err error) {
	err = db.Unscoped().Delete(T).Error
	return
}

// HOFSTADTER_BELOW
