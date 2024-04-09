package model

type Post struct {
	Id           int     `json:"id" DB:"id"`
	CategoryId   int     `json:"category_id" form:"category_id" DB:"category_id" binding:"required"`
	UserId       int     `json:"user_id" form:"user_id" DB:"user_id" binding:"required"`
	CategoryName string  `DB:"category_name"`
	UserName     string  `DB:"user_name"`
	Title        string  `json:"title" binding:"required" form:"title"`
	Description  string  `json:"description" binding:"required" form:"description"`
	ImagePath    *string `DB:"image_path" form:"image_path" json:"image_path"`
}
