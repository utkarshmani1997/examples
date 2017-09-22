package types

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      Comment-db-type-funcs
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Migrates (or creates) the table for Comment
*/
func MigrateCommentTable(db *gorm.DB) (err error) {
	// create own table
	err = db.Debug().AutoMigrate(&Comment{}).Error

	// create index on uuid
	db.Model(&Comment{}).Debug().AddUniqueIndex("idx_comment_uuid", "uuid")

	// deal with foreign keys
	db.Model(&Comment{}).Debug().AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")

	db.Model(&Comment{}).Debug().AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	return
}

/*
GORM hook to ensure UUID is created
*/
func (T *Comment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

/*
Creates a Comment record.
*/
func CreateComment(db *gorm.DB, T *Comment, PostID uint, UserID uint) (err error) {

	T.PostID = PostID
	T.UserID = UserID
	err = db.Create(T).Error

	return
}

/*
Find a Comment record by ID
*/
func GetCommentByID(db *gorm.DB, T *Comment, PostID uint, UserID uint) (err error) {

	err = db.
		Where("post_id = ?", PostID).
		Where("user_id = ?", UserID).
		First(T).Error

	return
}

/*
Find a Comment record by UUID
*/
func GetCommentByUUID(db *gorm.DB, UUID string) (T *Comment, err error) {
	T = NewComment()
	err = db.
		Where("uuid = ?", UUID).
		First(T).Error
	return
}

/*
Updates a Comment record
*/
func UpdateComment(db *gorm.DB, T *Comment) (err error) {
	err = db.Update(T).Error
	return
}

/*
Soft deletes the record (i.e. it has a deleted at timestamp and is not returned in queries)
*/
func DeleteComment(db *gorm.DB, T *Comment) (err error) {
	err = db.Delete(T).Error
	return
}

/*
Hard deletes the record, permenently deleting it from the DB
*/
func HardDeleteComment(db *gorm.DB, T *Comment) (err error) {
	err = db.Unscoped().Delete(T).Error
	return
}

// HOFSTADTER_BELOW
