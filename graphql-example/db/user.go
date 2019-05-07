package db

import (
	"errors"
	"server_common/graphql-example/model"
	"server_common/graphql-example/util"
)

// Login db
func Login(email string, pw string) (bool, error) {
	if !util.CheckEmail(email) {
		return false, errors.New("your email is unvalid")
	}
	if pw == "" {
		return false, errors.New("your password is empty")
	}
	var user model.User
	DB.First(&user, "email=?", email) //查询邮箱

	return true, nil
}

// InsertUser db
func InsertUser(user *model.User) error {
	return DB.Create(&model.User{Email: user.Email, PW:util.GetBcryptHash(user.PW), NickName:user.NickName}).Error
}

// CheckUserValid db
func CheckUserValid(id int, email string) (bool, error) {
	var user model.User
	err := DB.First(&user, "id=? & email=?", id, email).Error
	if err != nil || user.NickName != "" {
		return false, err
	}
	return true, nil
}

// GetUserByID db
func GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := DB.First(&user, id).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

// RemoveUserByID db
func RemoveUserByID(id int) error {
	return DB.Delete(model.User{}, "id=?", id).Error
}

// Follow db
func Follow(followerID, followeeID int) error {
	return DB.Create(&model.Follow{FolloweeID: followeeID, FollowerID: followerID}).Error
}

// Unfollow db
func Unfollow(followerID, followeeID int) error {
	return DB.Delete(model.Follow{}, "follower_id=? and followee_id=?", followerID, followeeID).Error
}

// GetFollowerByIDAndUser db
func GetFollowerByIDAndUser(id int, userID int) (*model.User, error) {

	var user model.User
	if err := DB.First(&user, "id=?", userID).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

// GetFollowersForUser db: get the people who you follow
func GetFollowersForUser(id int) ([]*model.User, error) {
	var users []*model.User
	if err := DB.Where("id=?", id).Find(&users).Error; err != nil{
		return nil, err
	}
	return users, nil

}

// GetFolloweeByIDAndUser db: get the people who follow you
func GetFolloweeByIDAndUser(id int, userID int) (*model.User, error) {
	var user model.User
	if err := DB.First(&user, "id=?", userID).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

// GetFolloweesForUser db
func GetFolloweesForUser(id int) ([]*model.User, error) {
	var users []*model.User
	if err := DB.Where("id=?", id).Find(&users).Error; err != nil{
		return nil, err
	}
	return users, nil
}
