package dao_memory

import (
	"log"
	"sort"

	"webapp2/dao"
	"webapp2/entities"
)

// This type/struct stores no state, it’s just a collection of methods
// To get an instance of this void struct : var dao = dao.StudentDAO{}
type StudentDAOMemory struct {
}

// Check StudentDAOMemory implements StudentDAO
var _ dao.StudentDAO = (*StudentDAOMemory)(nil)

func (this *StudentDAOMemory) values(m map[int]entities.Student) []entities.Student {
	var a = make([]entities.Student, len(m))
	i := 0
	for _, v := range m {
		a[i] = v
		i++
	}
	this.sortById(a)
	return a
}

func (this *StudentDAOMemory) sortById(students []entities.Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Id < students[j].Id
	})
}

func (this *StudentDAOMemory) FindAll() []entities.Student {
	log.Print("DAO - FindAll() ")
	return this.values(StudentsMap)
}

func (this *StudentDAOMemory) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id)
	student := StudentsMap[id]
	return &student
}

func (this *StudentDAOMemory) Delete(id int) {
	log.Printf("DAO - Delete(%d) ", id)
	delete(StudentsMap, id)
}

func (this *StudentDAOMemory) Create(student entities.Student) {
	log.Printf("DAO - Create(%d) ", student.Id)
	StudentsMap[student.Id] = student
}

func (this *StudentDAOMemory) Update(student entities.Student) {
	log.Printf("DAO - Update(%d) ", student.Id)
	StudentsMap[student.Id] = student
}
