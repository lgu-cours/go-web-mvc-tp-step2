package dao_memory

import (
	"webapp2/entities"
)

var LanguagesMap = map[string]entities.Language {
	"JA" : { "JA", "Java" },
	"JS" : { "JS", "JavaScript" },
	"PH" : { "PH", "PHP" },
	"PY" : { "PY", "Python" },
	"GO" : { "GO", "Go lang" } }
