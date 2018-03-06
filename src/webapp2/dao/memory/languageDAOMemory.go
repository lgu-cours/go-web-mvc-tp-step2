package dao_memory

import (
	"log"
	
	"webapp2/dao"
	"webapp2/entities"
)

// This type/struct stores no state, it’s just a collection of methods 
// To get an instance of this void struct : var dao = dao.StudentDAO{}
type LanguageDAOMemory struct {
}

// Check LanguageDAOMemory implements LanguageDAO
var _ dao.LanguageDAO = (*LanguageDAOMemory)(nil)

func (this *LanguageDAOMemory) values( m map[string]entities.Language ) []entities.Language  {
	var a = make([]entities.Language, len(m))
	i := 0
    for _, v := range m {
	    a[i] = v
    	i++
    }
	return a
}

func (this *LanguageDAOMemory) FindAll() []entities.Language {
	log.Print("DAO - FindAll() " )
	return this.values(LanguagesMap)
}

func (this *LanguageDAOMemory) Find(code string) *entities.Language {
	log.Printf("DAO - Find(%s) ", code )
	language := LanguagesMap[code]
	return &language

}
