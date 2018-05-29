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
//批量入栏 和 单个 入栏相同
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
	m["fcn"] = "Register"
	m["args"] = make([]string, 1)
	m["peers"] = []string{"peer1"}

	var str string

	for i := 0; i < len(mx["PCList"].([]interface{})); i++ {
		switch mx["PCList"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PCList[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["PCList"].([]interface{})[i].(string) + `","batchNumber":"` + mx["PCNO"].(string) + `","kind":"` + mx["isType"].(string) + `","type":"` + mx["species"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Lairage","operator":"` + mx["CreatePerson"].(string) + `","createTime":"` + mx["createTime"].(string) + `"},`
	}

	batchOrSingleOperate("Register",str,args["header"]["Authorization"][0],result)
	return nil
}

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
//批量喂养  喂养
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
	var str string
	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","operator":"` + mx["Name"].(string) + `","feedName":"` + mx["Id"].(string) + `","feedTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Feed"}`
	}
	batchOrSingleOperate("Feed",str,args["header"]["Authorization"][0],result)
	return nil
}

//防疫 批量防疫
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
	var str string
	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","operator":"` + mx["OperatorName"].(string) + `","vaccineName":"` + mx["id"].(string) + `","VaccineTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Medication","vaccineType":"vaccineType","vaccineNumber":"vaccineNumber"}`

	}
	batchOrSingleOperate("Vaccine",str,args["header"]["Authorization"][0],result)
	return nil
}
//检疫 批量检疫
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
	var str string
	for i := 0; i < len(mx["EarTag"].([]interface{})); i++ {
		switch mx["EarTag"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "EarTag[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["EarTag"].([]interface{})[i].(string) + `","operator":"` + mx["OperatorName"].(string) + `","examConsequence":"` + mx["CheckResult"].(string) + `","examTime":"` + mx["CheckDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Prevention"}`
	}
	batchOrSingleOperate("Exam",str,args["header"]["Authorization"][0],result)
	log.Println(result)

	return nil
}

//救治 批量救治
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
	var str string
	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","operator":"` + mx["Name"].(string) + `","saveNumber":"` + mx["id"].(string) + `","saveName":"saveName","saveType":"` + mx["Treatment"].(string) + `","saveConsequence":"` + mx["InspectResult"].(string) + `","saveTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Save"}`

	}
	batchOrSingleOperate("Save",str,args["header"]["Authorization"][0],result)

	return nil
}

//灭失 批量灭失
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
	var str string
	for i := 0; i < len(mx["DeathObject"].([]interface{})); i++ {
		switch mx["DeathObject"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "DeathObject[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["DeathObject"].([]interface{})[i].(string) + `","operator":"` + mx["Name"].(string) + `","lostWay":"` + mx["TreatMethod"].(string) + `","lostReason":"` + mx["CauseDeath"].(string) + `","lostTime":"` + mx["SysDate"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `","operation":"Lost"}`

	}
	batchOrSingleOperate("Lost",str,args["header"]["Authorization"][0],result)

	return nil
}
//出栏
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
	var str string
	for i := 0; i < len(mx["PNO"].([]interface{})); i++ {
		switch mx["PNO"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "PNO[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["PNO"].([]interface{})[i].(string) + `","name":"` + mx["Name"].(string) + `","outputTime":"` + mx["SysDate"].(string) + `","operation":"Fattened","operator":"` + mx["CreatePerson"].(string) + `","mapPosition":"` + mx["TaskGps"].(string) + `"}`
	}
	batchOrSingleOperate("Output",str,args["header"]["Authorization"][0],result)

	return nil
}

func (this Remote)AssetButcher(args map[string]map[string][]string, result *Asset) error {
	var mx map[string]interface{}
	err := json.Unmarshal([]byte(args["body"]["b"][0]), &mx)
	if err != nil {
		log.Println(err)
	}

	_, ok := mx["operator"]
	if !ok {
		result.Message = "operator required"
		return nil
	}
	switch mx["operator"].(type) {
	case string:
	default:
		result.Message = "operator should be string"
		return nil
	}

	_, ok = mx["productIds"]
	if !ok {
		result.Message = "productIds required"
		return nil
	}
	switch mx["productIds"].(type) {
	case []interface{}:
	default:
		result.Message = "productIds should be []string"
		return nil
	}

	_, ok = mx["hookNo"]
	if !ok {
		result.Message = "hookNo required"
		return nil
	}
	switch mx["hookNo"].(type) {
	case string:
	default:
		result.Message = "hookNo should be string"
		return nil
	}

	_, ok = mx["Name"]
	if !ok {
		result.Message = "Name required"
		return nil
	}
	switch mx["operation"].(type) {
	case string:
	default:
		result.Message = "operation should be string"
		return nil
	}

	_, ok = mx["butcherTime"]
	if !ok {
		result.Message = "butcherTime required"
		return nil
	}
	switch mx["butcherTime"].(type) {
	case string:
	default:
		result.Message = "butcherTime should be string"
		return nil
	}

	_, ok = mx["mapPosition"]
	if !ok {
		result.Message = "mapPosition required"
		return nil
	}
	switch mx["mapPosition"].(type) {
	case string:
	default:
		result.Message = "mapPosition should be string"
		return nil
	}
	var str string
	for i := 0; i < len(mx["productIds"].([]interface{})); i++ {
		switch mx["productIds"].([]interface{})[i].(type) {
		case string:
		default:
			result.Message = "productIds[" + strconv.Itoa(i) + "] should be string"
			return nil
		}

		str += `{"productId":"` + mx["productIds"].([]interface{})[i].(string) + `","hookNo":"` + mx["hookNo"].(string) + `","butcherTime":"` + mx["butcherTime"].(string) + `","operation":"`+mx["operation"].(string)+`","operator":"` + mx["operator"].(string) + `","mapPosition":"` + mx["mapPosition"].(string) + `"}`
	}
	batchOrSingleOperate("Butcher",str,args["header"]["Authorization"][0],result)

	return nil
}
//批量或者单个操作
func batchOrSingleOperate(fcn string,str string,auth string ,result *Asset){
	m := make(map[string]interface{})
	m["fcn"] = fcn
	m["args"] = make([]string, 1)
	m["peers"] = []string{"peer1"}
	m["args"].([]string)[0] = "[" + str[0:len(str) - 1] + "]"

	log.Println(m)

	mJSON, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
		result.Message = "JSON Marshal error:" + err.Error()
		return
	}

	reader := bytes.NewReader(mJSON)

	request, err := http.NewRequest("POST", "http://" + model.CHAIN_CODE_DOMAIN + ":" + model.CHAIN_CODE_PORT + "/channels/mychannel/chaincodes/hlccc", reader)
	if err != nil {
		log.Println(err.Error())
		result.Message = err.Error()
		return
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("authorization", "Bearer " + auth)

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		result.Message = err.Error()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		result.Message = err.Error()
		return
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		log.Println(err.Error())
		result.Message = "JSON Unmarshal error:" + err.Error()
		return
	}

	log.Println(result.Message) // = append(result.Messages, mx["PCList"].([]interface{})[i].(string) + ":" + string(body))

	return
}
