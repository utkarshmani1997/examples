package types

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      Post-db-type-funcs
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Migrates (or creates) the table for Post
*/
func MigratePostTable(db *gorm.DB) (err error) {
	// create own table
	err = db.Debug().AutoMigrate(&Post{}).Error

	// create index on uuid
	db.Model(&Post{}).Debug().AddUniqueIndex("idx_post_uuid", "uuid")

	// deal with foreign keys
	db.Model(&Post{}).Debug().AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	return
}

/*
GORM hook to ensure UUID is created
*/
func (T *Post) BeforeCreate(scope *gorm.Scope) error {
	T.UUID = uuid.New().String()
	return nil
}

/*
Creates a Post record.
*/
func CreatePost(db *gorm.DB, T *Post, UserUUID string) (err error) {

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
	// other relations? (has-one, has-many, many-to-many)
	// other relations? (has-one, has-many, many-to-many)
	err = db.Create(T).Error

	return
}

/*
Find a Post record by ID
*/
func GetPostByID(db *gorm.DB, T *Post, UserUUID string) (err error) {

	err = db.
		Where("user_uuid = ?", UserUUID).
		First(T).Error

	return
}

/*
Find a Post record by UUID
*/
func GetPostByUUID(db *gorm.DB, UUID string) (T *Post, err error) {
	T = NewPost()
	err = db.
		Where("uuid = ?", UUID).
		First(T).Error
	return
}

/*
Updates a Post record
*/
func UpdatePost(db *gorm.DB, T *Post) (err error) {
	err = db.Update(T).Error
	return
}

/*
Soft deletes the record (i.e. it has a deleted at timestamp and is not returned in queries)
*/
func DeletePost(db *gorm.DB, T *Post) (err error) {
	err = db.Delete(T).Error
	return
}

/*
Hard deletes the record, permenently deleting it from the DB
*/
func HardDeletePost(db *gorm.DB, T *Post) (err error) {
	err = db.Unscoped().Delete(T).Error
	return
}

// HOFSTADTER_BELOW
