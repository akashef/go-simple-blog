package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
}

type Post struct {
	gorm.Model
	Title  string
	Body   string
	UserID uint
}

type Comment struct {
	gorm.Model
	Body   string
	PostID uint
	UserID uint
}
