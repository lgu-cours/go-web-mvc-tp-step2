package dao_bolt

import (
	"log"
	"strconv"
	"encoding/json"
	
	"webapp2/dao"
	"webapp2/entities"
)

// This type/struct stores no state, it’s just a collection of methods
// To get an instance of this void struct : var dao = dao.StudentDAO{}
type StudentDAOBolt struct {
	bucketName string 
}

// Check studentDAOBolt implements StudentDAO
var _ dao.StudentDAO = (*StudentDAOBolt)(nil)

func NewStudentDAOBolt() StudentDAOBolt {
	// Creates a DAO using the "students" bucket
	return StudentDAOBolt{"students"}
}

func (this *StudentDAOBolt) toStudent(value string) entities.Student {
	student := entities.Student{}
	err := json.Unmarshal([]byte(value), &student)
	if ( err != nil ) {
		panic(err)
	}	
	return student
}

func (this *StudentDAOBolt) FindAll() []entities.Student {
	log.Print("DAO - FindAll() ")
	students := make([]entities.Student,0)
	values := dbGetAll(this.bucketName)
	for _, v := range values {
		student := this.toStudent(v)
		students = append(students, student)
	}
	return students
}

func (this *StudentDAOBolt) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id)
	key := strconv.Itoa(id)
	value := dbGet(this.bucketName, key)
	if value != "" {
		student := this.toStudent(value)
		return &student
	} else {
		return nil 
	}
}

func (this *StudentDAOBolt) Delete(id int) {
	log.Printf("DAO - Delete(%d) ", id)
	key := strconv.Itoa(id)
	dbDelete(this.bucketName, key)
}

func (this *StudentDAOBolt) Create(student entities.Student) {
	log.Printf("DAO - Create(%d) ", student.Id)
	this.Save(student)
}

func (this *StudentDAOBolt) Update(student entities.Student) {
	log.Printf("DAO - Update(%d) ", student.Id)
	this.Save(student)
}

func (this *StudentDAOBolt) Save(student entities.Student) {
	log.Printf("DAO - Save(%d) ", student.Id)
	key := strconv.Itoa(student.Id)
	// value : student converted to JSON
	value, err := json.Marshal(student)
	if ( err != nil ) {
		panic(err)
	}
	dbPut(this.bucketName, key, string(value[:]))
}
