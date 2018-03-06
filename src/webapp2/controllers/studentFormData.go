package controllers

import (
	"webapp2/entities"
)

type StudentFormData struct {
	CreationMode  bool
    Student       entities.Student 
    Languages     []entities.Language
}

