// lcserver
package server

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"log"

	"model"
)

type LoginResult struct {
	Code	string	`json:"code"`
	Success	bool	`json:"success"`
	Message	string	`json:"message"`
	Token	string	`json:"token"`
	Exp	int64	`json:"exp"`
}

func (this Remote) Login(args []byte, result *LoginResult) error {
	reader := bytes.NewReader(args)
	log.Println("http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/users")
	request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/users", reader)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	//var r interface{}
	//err = global.Call("Remote.SessionInsertion", body, &r)
	//if err != nil {
	//	log.Println(err.Error())
	//	return err
	//}

	err = json.Unmarshal(body, result)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
