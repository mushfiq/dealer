package utils

import (
	"os"
	"io/ioutil"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/mailgun/mailgun-go"
)

func SendEmail(emailBody string, emailConfig *simplejson.Json){
	gun := mailgun.NewMailgun(emailConfig.Get("domain").MustString(), emailConfig.Get("api_key").MustString(), emailConfig.Get("public_key").MustString())
	for _, receiver_email := range emailConfig.Get("receipents").MustArray(){
		mail := mailgun.NewMessage(emailConfig.Get("sender_email").MustString(), emailConfig.Get("subject").MustString(), emailBody, receiver_email.(string))
		response, id, _ := gun.Send(mail)
		fmt.Printf("Response ID: %s\n", id)
		fmt.Printf("Message from server: %s\n", response)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func GetConfig(configType string) *simplejson.Json {
	fileName := "./config/"+configType+".json"
	file, _ := ioutil.ReadFile(fileName)
	config, err := simplejson.NewJson(file)
	CheckError(err)
	return config
}

