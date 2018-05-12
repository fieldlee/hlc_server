// lcserver
package server

import (
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
	"strconv"

	"model"
)

type Asset struct {
	Success		bool	`json:"success"`
	Payloads	[]string	`json:"payloads"`
	Timestamp	int64	`json:"timestamp"`
	Message		string	`json:"message"`
	Messages	[]string	`json:"messages"`
}

func (this Remote)AssetRegister(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["PCList"]
	if !ok {
		result.Message = "PCList required"
		return nil
	}
	switch mx["PCList"].(type) {
	case []interface{}:
	default:
		result.Message = "PCList should be []string"
		return nil
	}

	_, ok = mx["CreatePerson"]
	if !ok {
		result.Message = "CreatePerson required"
		return nil
	}
	switch mx["CreatePerson"].(type) {
	case string:
	default:
		result.Message = "CreatePerson should be string"
		return nil
	}

	_, ok = mx["username"]
	if !ok {
		result.Message = "PCList required"
		return nil
	}
	switch mx["username"].(type) {
	case string:
	default:
		result.Message = "username should be string"
		return nil
	}

	_, ok = mx["PCNO"]
	if !ok {
		result.Message = "PCNO required"
		return nil
	}
	switch mx["PCNO"].(type) {
	case string:
	default:
		result.Message = "PCNO should be string"
		return nil
	}

	_, ok = mx["isType"]
	if !ok {
		result.Message = "isType required"
		return nil
	}
	switch mx["isType"].(type) {
	case string:
	default:
		result.Message = "isType should be string"
		return nil
	}

	_, ok = mx["species"]
	if !ok {
		result.Message = "species required"
		return nil
	}
	switch mx["species"].(type) {
	case string:
	default:
		result.Message = "species should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	_, ok = mx["createTime"]
	if !ok {
		result.Message = "createTime required"
		return nil
	}
	switch mx["createTime"].(type) {
	case string:
	default:
		result.Message = "createTime should be string"
		return nil
	}

	m := make(map[string]interface{})
	m["fcn"] = "BatchRegister"
	m["args"] = make([]string, 2)
	m["args"].([]string)[0] = m["fcn"].(string)
	m["peers"] = []string{"peer0.creator.com"}

	var str string

	for i := 0; i < len(mx["PCList"].([]interface{})); i++ {
		switch mx["PCList"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PCList[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["PCList"].([]interface{})[i].(string) + `","username":"` + mx["username"].(string) + `","inModule":"` + mx["PCNO"].(string) + `","kind":"` + mx["isType"].(string) + `","type":"` + mx["species"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Lairage","operator":"` + mx["CreatePerson"].(string) + `","createTime":"` + mx["createTime"].(string) + `"},`
	}

	m["args"].([]string)[1] = "[" + str[0:len(str) - 1] + "]"

	log.Println(m)

	mJSON, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
		result.Message = "JSON Marshal error:" + err.Error()
		return nil
	}

	reader := bytes.NewReader(mJSON)

	request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
	if err != nil {
		log.Println(err.Error())
		result.Message = err.Error()
		return nil
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		result.Message = err.Error()
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		result.Message = err.Error()
		return nil
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		log.Println(err.Error())
		result.Message = "JSON Unmarshal error:" + err.Error()
		return nil
	}

	log.Println(result.Message) // = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + ":" + string(body))

	return nil
}

/*
func (this Remote)AssetRegister(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["PCList"]
	if !ok {
		result.Message = "PCList required"
		return nil
	}
	switch mx["PCList"].(type) {
	case []interface{}:
	default:
		result.Message = "PCList should be []string"
		return nil
	}

	_, ok = mx["CreatePerson"]
	if !ok {
		result.Message = "CreatePerson required"
		return nil
	}
	switch mx["CreatePerson"].(type) {
	case string:
	default:
		result.Message = "CreatePerson should be string"
		return nil
	}

	_, ok = mx["username"]
	if !ok {
		result.Message = "PCList required"
		return nil
	}
	switch mx["username"].(type) {
	case string:
	default:
		result.Message = "username should be string"
		return nil
	}

	_, ok = mx["PCNO"]
	if !ok {
		result.Message = "PCNO required"
		return nil
	}
	switch mx["PCNO"].(type) {
	case string:
	default:
		result.Message = "PCNO should be string"
		return nil
	}

	_, ok = mx["isType"]
	if !ok {
		result.Message = "isType required"
		return nil
	}
	switch mx["isType"].(type) {
	case string:
	default:
		result.Message = "isType should be string"
		return nil
	}

	_, ok = mx["species"]
	if !ok {
		result.Message = "species required"
		return nil
	}
	switch mx["species"].(type) {
	case string:
	default:
		result.Message = "species should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	_, ok = mx["createTime"]
	if !ok {
		result.Message = "createTime required"
		return nil
	}
	switch mx["createTime"].(type) {
	case string:
	default:
		result.Message = "createTime should be string"
		return nil
	}


	for i := 0; i < len(mx["PCList"].([]interface{})); i++ {
		switch mx["PCList"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PCList[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["PCList"].([]interface{})[i].(string) + `","username":"` + mx["username"].(string) + `","inModule":"` + mx["PCNO"].(string) + `","kind":"` + mx["isType"].(string) + `","type":"` + mx["species"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Lairage","operator":"` + mx["CreatePerson"].(string) + `","createTime":"` + mx["createTime"].(string) + `"}`

		m := make(map[string]interface{})
		m["fcn"] = "Register"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + ":" + string(body))
	}

	return nil
}
*/

func (this Remote)AssetQueryDetail(args map[string]map[string][]string, result *Asset) error {
	m := make(map[string]interface{})
	m["fcn"] = "QueryProductDetail"
	m["args"] = make([]string, 2)
	m["args"].([]string)[0] = m["fcn"].(string)
	m["args"].([]string)[1] = args["body"]["b"][0]
	m["peers"] = []string{"peer0.creator.com"}

	mJSON, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
	}

	reader := bytes.NewReader(mJSON)

	request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/query/channels/mychannel/chaincodes/hlccc", reader)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		log.Println(err.Error())
		result.Message = "500:服务器内部错误:" + err.Error() + string(body)
		return err
	}

	return nil
}

func (this Remote)AssetFeed(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["PNO"]
	if !ok {
		result.Message = "PNO required"
		return nil
	}
	switch mx["PNO"].(type) {
	case []interface{}:
	default:
		result.Message = "PNO should be []string"
		return nil
	}

	_, ok = mx["Name"]
	if !ok {
		result.Message = "Name required"
		return nil
	}
	switch mx["Name"].(type) {
	case string:
	default:
		result.Message = "Name should be string"
		return nil
	}

	_, ok = mx["Id"]
	if !ok {
		result.Message = "Id required"
		return nil
	}
	switch mx["Id"].(type) {
	case string:
	default:
		result.Message = "Id should be string"
		return nil
	}

	_, ok = mx["SysDate"]
	if !ok {
		result.Message = "SysDate required"
		return nil
	}
	switch mx["Id"].(type) {
	case string:
	default:
		result.Message = "SysDate should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","feeder":"` + mx["Name"].(string) + `","feedId":"` + mx["Id"].(string) + `","feedTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Feed"}`

		m := make(map[string]interface{})
		m["fcn"] = "ChangeProduct"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"].([]string)[1])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + string(body))

	}
	log.Println(result)

	return nil
}

func (this Remote)AssetMedication(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["PNO"]
	if !ok {
		result.Message = "PNO required"
		return nil
	}
	switch mx["PNO"].(type) {
	case []interface{}:
	default:
		result.Message = "PNO should be []string"
		return nil
	}

	_, ok = mx["OperatorName"]
	if !ok {
		result.Message = "OperatorName required"
		return nil
	}
	switch mx["OperatorName"].(type) {
	case string:
	default:
		result.Message = "OperatorName should be string"
		return nil
	}

	_, ok = mx["SysDate"]
	if !ok {
		result.Message = "SysDate required"
		return nil
	}
	switch mx["SysDate"].(type) {
	case string:
	default:
		result.Message = "SysDate should be string"
		return nil
	}

	_, ok = mx["id"]
	if !ok {
		result.Message = "id required"
		return nil
	}
	switch mx["id"].(type) {
	case string:
	default:
		result.Message = "id should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","medder":"` + mx["OperatorName"].(string) + `","medId":"` + mx["id"].(string) + `","medicationTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Prevention"}`

		m := make(map[string]interface{})
		m["fcn"] = "ChangeProduct"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"].([]string)[1])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + string(body))

	}
	log.Println(result)

	return nil
}

func (this Remote)AssetPrevention(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	log.Println(mx)
	_, ok := mx["EarTag"]
	if !ok {
		result.Message = "EarTag required"
		return nil
	}
	switch mx["EarTag"].(type) {
	case []interface{}:
	default:
		result.Message = "EarTag should be []string"
		return nil
	}

	_, ok = mx["OperatorName"]
	if !ok {
		result.Message = "OperatorName required"
		return nil
	}
	switch mx["OperatorName"].(type) {
	case string:
	default:
		result.Message = "OperatorName should be string"
		return nil
	}

	_, ok = mx["Immunion"]
	if !ok {
		result.Message = "Immunion required"
		return nil
	}
	switch mx["Immunion"].(type) {
	case string:
	default:
		result.Message = "Immunion should be string"
		return nil
	}

	_, ok = mx["CheckDate"]
	if !ok {
		result.Message = "CheckDate required"
		return nil
	}
	switch mx["CheckDate"].(type) {
	case string:
	default:
		result.Message = "CheckDate should be string"
		return nil
	}

	_, ok = mx["CheckResult"]
	if !ok {
		result.Message = "CheckResult required"
		return nil
	}
	switch mx["CheckResult"].(type) {
	case string:
	default:
		result.Message = "CheckResult should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	for i := 0; i < len(mx["EarTag"].([]interface{})); i++ {
		switch mx["EarTag"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "EarTag[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["EarTag"].([]interface{})[i].(string) + `","preventer":"` + mx["OperatorName"].(string) + `","preventName":"` + mx["Immunion"].(string) + `","preventResult":"` + mx["CheckResult"].(string) + `","preventionTime":"` + mx["CheckDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Medication"}`

		m := make(map[string]interface{})
		m["fcn"] = "ChangeProduct"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"].([]string)[1])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["EarTag"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["EarTag"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["EarTag"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["EarTag"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["EarTag"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["EarTag"].([]interface{})[i].(string) + ":" + string(body))

	}
	log.Println(result)

	return nil
}

func (this Remote)AssetSave(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["PNO"]
	if !ok {
		result.Message = "PNO required"
		return nil
	}
	switch mx["PNO"].(type) {
	case []interface{}:
	default:
		result.Message = "PNO should be []string"
		return nil
	}

	_, ok = mx["Name"]
	if !ok {
		result.Message = "Name required"
		return nil
	}

	switch mx["Name"].(type) {
	case string:
	default:
		result.Message = "Name should be string"
		return nil
	}

	_, ok = mx["InspectResult"]
	if !ok {
		result.Message = "InspectResult required"
		return nil
	}
	switch mx["InspectResult"].(type) {
	case string:
	default:
		result.Message = "InspectResult should be string"
		return nil
	}

	_, ok = mx["SysDate"]
	if !ok {
		result.Message = "SysDate required"
		return nil
	}
	switch mx["SysDate"].(type) {
	case string:
	default:
		result.Message = "SysDate should be string"
		return nil
	}

	_, ok = mx["Treatment"]
	if !ok {
		result.Message = "Treatment required"
		return nil
	}
	switch mx["Treatment"].(type) {
	case string:
	default:
		result.Message = "Treatment should be string"
		return nil
	}

	_, ok = mx["id"]
	if !ok {
		result.Message = "id required"
		return nil
	}
	switch mx["id"].(type) {
	case string:
	default:
		result.Message = "id should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","inspector":"` + mx["Name"].(string) + `","inspectId":"` + mx["id"].(string) + `","treatment":"` + mx["Treatment"].(string) + `","inspectResult":"` + mx["InspectResult"].(string) + `","saveTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Save"}`

		m := make(map[string]interface{})
		m["fcn"] = "ChangeProduct"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"].([]string)[1])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + string(body))

	}
	log.Println(result)

	return nil
}

func (this Remote)AssetLost(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["DeathObject"]
	if !ok {
		result.Message = "DeathObject required"
		return nil
	}
	switch mx["DeathObject"].(type) {
	case []interface{}:
	default:
		result.Message = "DeathObject should be []string"
		return nil
	}

	_, ok = mx["Name"]
	if !ok {
		result.Message = "Name required"
		return nil
	}
	switch mx["Name"].(type) {
	case string:
	default:
		result.Message = "Name should be string"
		return nil
	}

	_, ok = mx["CauseDeath"]
	if !ok {
		result.Message = "CauseDeath required"
		return nil
	}
	switch mx["CauseDeath"].(type) {
	case string:
	default:
		result.Message = "CauseDeath should be string"
		return nil
	}

	_, ok = mx["SysDate"]
	if !ok {
		result.Message = "SysDate required"
		return nil
	}
	switch mx["SysDate"].(type) {
	case string:
	default:
		result.Message = "SysDate should be string"
		return nil
	}

	_, ok = mx["TreatMethod"]
	if !ok {
		result.Message = "TreatMethod required"
		return nil
	}
	switch mx["TreatMethod"].(type) {
	case string:
	default:
		result.Message = "TreatMethod should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	for i := 0; i < len(mx["DeathObject"].([]interface{})); i++ {
		switch mx["DeathObject"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "DeathObject[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["DeathObject"].([]interface{})[i].(string) + `","loser":"` + mx["Name"].(string) + `","lostTreat":"` + mx["TreatMethod"].(string) + `","lostCause":"` + mx["CauseDeath"].(string) + `","lostTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Lost"}`

		m := make(map[string]interface{})
		m["fcn"] = "ChangeProduct"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"].([]string)[1])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["DeathObject"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["DeathObject"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["DeathObject"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["DeathObject"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["DeathObject"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["DeathObject"].([]interface{})[i].(string) + ":" + string(body))

	}
	log.Println(result)

	return nil
}

func (this Remote)AssetFattened(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["CreatePerson"]
	if !ok {
		result.Message = "CreatePerson required"
		return nil
	}
	switch mx["CreatePerson"].(type) {
	case string:
	default:
		result.Message = "CreatePerson should be string"
		return nil
	}

	_, ok = mx["PNO"]
	if !ok {
		result.Message = "PNO required"
		return nil
	}
	switch mx["PNO"].(type) {
	case []interface{}:
	default:
		result.Message = "PNO should be []string"
		return nil
	}

	_, ok = mx["CLPCNO"]
	if !ok {
		result.Message = "CLPCNO required"
		return nil
	}
	switch mx["CLPCNO"].(type) {
	case string:
	default:
		result.Message = "CLPCNO should be string"
		return nil
	}

	_, ok = mx["Name"]
	if !ok {
		result.Message = "Name required"
		return nil
	}
	switch mx["Name"].(type) {
	case string:
	default:
		result.Message = "Name should be string"
		return nil
	}

	_, ok = mx["SysDate"]
	if !ok {
		result.Message = "SysDate required"
		return nil
	}
	switch mx["SysDate"].(type) {
	case string:
	default:
		result.Message = "SysDate should be string"
		return nil
	}

	_, ok = mx["TaskGps"]
	if !ok {
		result.Message = "TaskGps required"
		return nil
	}
	switch mx["TaskGps"].(type) {
	case string:
	default:
		result.Message = "TaskGps should be string"
		return nil
	}

	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str := `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","name":"` + mx["Name"].(string) + `","outModule":"` + mx["CLPCNO"].(string) + `","FattenedTime":"` + mx["SysDate"].(string) + `","operation":"Out-Fence","operator":"` + mx["CreatePerson"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `"}`

		m := make(map[string]interface{})
		m["fcn"] = "ChangeProduct"
		m["args"] = make([]string, 2)
		m["args"].([]string)[0] = m["fcn"].(string)
		m["args"].([]string)[1] = str
		m["peers"] = []string{"peer0.creator.com"}

		log.Println(m["args"].([]string)[1])

		mJSON, err := json.Marshal(m)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		reader := bytes.NewReader(mJSON)

		request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		request.Header.Set("Content-Type", "application/json;charset=utf-8")
		request.Header.Set("authorization", "Bearer " + args["header"]["Authorization"][0])

		client := http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + err.Error())
			continue
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			log.Println(err.Error())
			result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + " JSON parse error:" + err.Error())
			continue
		}

		result.Messages = append(result.Messages, mx["PNO"].([]interface{})[i].(string) + ":" + string(body))

	}
	log.Println(result)

	return nil
}
