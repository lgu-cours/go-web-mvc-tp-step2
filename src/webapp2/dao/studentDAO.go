package dao

import (
	"webapp2/entities"
)

// DAO interface for Student
type StudentDAO interface {
   FindAll()     []entities.Student
   Find(id int)  *entities.Student
   Delete(id int)
   Create(student entities.Student)
   Update(student entities.Student)
}

