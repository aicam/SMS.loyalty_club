package sms

import (
	"encoding/json"
	"net/http"
)

var HEADERS = map[string]string{
	"Content-Type": "application/json",
}
//sms server address
var SMSTokenApi = "https://api.sms.ir/users/v1/Token/GetToken"
// just for special sms api
var AddPhoneNumberToGroup := "https://api.sms.ir/users/v1/Contacts/AddContacts"
func getToken() string {
	var result map[string]string
	client := &http.Client{}
	req, _ := http.NewRequest("POST", SMSTokenApi, nil)
	for headerName, headerValue := range HEADERS {
		req.Header.Add(headerName, headerValue)
	}
	res, _ := client.Do(req)
	_ = json.NewDecoder(res.Body).Decode(&result)
	tokenKey := result["TokenKey"]
	return tokenKey
}

func addContact(phonenumber string) {
	client := &http.Client{}

}

func SendSMS(phonenumber string) {
	tokenKey := getToken()

}