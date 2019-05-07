package model

import (
	"time"
)
// User model
type User struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email      string    `json:"email"`
	NickName   string    `json:"nickname"`
	PW         string    `json:"pw"`
}

// Post model
type Post struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	PraiseNum  int       `json:"praise_num"`
	CommentNum int       `json:"comment_num"`
}


// Auth model
type Auth struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Token string `json:"token"`
	User  User   `json:"user"`
}

// Comment model
type Comment struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID     int       `json:"user_id"`
	PostID     int       `json:"post_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
}

// Follow model
type Follow struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FollowerID int       `json:"follower_id"`
	FolloweeID int       `json:"followee_id"`
}