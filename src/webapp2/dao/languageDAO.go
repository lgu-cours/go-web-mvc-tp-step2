package dao

import (
	"webapp2/entities"
)

// DAO interface for Language
type LanguageDAO interface {
	
	FindAll()         []entities.Language 
	Find(code string) *entities.Language

}
