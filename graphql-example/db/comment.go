package db

import (
	"server_common/graphql-example/model"

)

// InsertComment db
func InsertComment(comment *model.Comment) error {
	return DB.Create(&comment).Error


}

// RemoveCommentByID db
func RemoveCommentByID(id int) error {
	return nil
}

// GetCommentByIDAndPost db
func GetCommentByIDAndPost(id int, postID int) (*model.Comment, error) {
	var comment model.Comment
	if err := DB.First(&comment, "id=? and post_id=?", id, postID).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentsForPost db
func GetCommentsForPost(id int) ([]*model.Comment, error) {
	var comment []*model.Comment
	if err := DB.Where("post_id=?", id).Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
