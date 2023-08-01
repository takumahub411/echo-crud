package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `json:"title"`
	LeadingRole string `json:"leading_role"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
}

type RequestMovie struct {
	Title       string `json:"title"`
	LeadingRole string `json:"leading_role"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
}
