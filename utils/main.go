package utils

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func GetKeywords() []interface{} {
	file, err := ioutil.ReadFile("./keywords.json")
	CheckError(err)
	js, err := simplejson.NewJson(file)
	CheckError(err)
	keywords := js.Get("keywords").MustArray()
	return keywords
}
