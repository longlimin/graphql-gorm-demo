package db

import (
	"server_common/graphql-example/model"
)

// InsertPost db
func InsertPost(post *model.Post) error {
	return DB.Create(&post).Error
}

// RemovePostByID db
func RemovePostByID(id int) error {
	return DB.Delete(model.Post{}, "id=?", id).Error
}

// GetPostByID db
func GetPostByID(id int) (*model.Post, error) {
	var post model.Post;
	if err := DB.First(&post, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPostByIDAndUser db
func GetPostByIDAndUser(id, userID int) (*model.Post, error) {
	var post model.Post;
	if err := DB.First(&post, "id=? and user_id", id, userID).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPostsForUser db
func GetPostsForUser(userID int) ([]*model.Post, error) {
	var posts []*model.Post
	if err := DB.Where("user_id=?", userID).Find(&posts).Error; err != nil{
		return nil, err
	}
	return posts, nil
}

// PraisePost db
func PraisePost(postID int) error {
	return nil
}

// UnPraisePost db
func UnPraisePost(postID int) error {
	return nil
}

// UpdatePost db
func UpdatePost(postID int, title, body string) error {
	return nil
}
