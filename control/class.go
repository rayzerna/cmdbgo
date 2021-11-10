package control

import "encoding/json"

type RtnData struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// OK response
func (rtn *RtnData) OK() RtnData {
	rtn.Code = "0"
	rtn.Msg = "success"
	return *rtn
}

// Set Dict
func (rtn *RtnData) Dict(dictData map[string]interface{}) []byte {
	rtnJson, err := json.Marshal(rtn)
	if err != nil {
		panic(err)
	}
	dictDatajson, err := json.Marshal(dictData)
	if err != nil {
		panic(err)
	}

	var mapResult map[string]interface{}
	var dictDataMap map[string]interface{}

	err = json.Unmarshal([]byte(rtnJson), &mapResult)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(dictDatajson), &dictDataMap)
	if err != nil {
		panic(err)
	}

	mapResult["data"] = dictDataMap

	jsonResp, err := json.Marshal(mapResult)
	if err != nil {
		panic(err)
	}

	return jsonResp
}

// Set List
func (rtn *RtnData) List(listData []string) []byte {
	rtnJson, err := json.Marshal(rtn)
	if err != nil {
		panic(err)
	}
	listDataJson, err := json.Marshal(listData)
	if err != nil {
		panic(err)
	}

	var mapResult map[string]interface{}
	var listDataMap []interface{}

	err = json.Unmarshal([]byte(rtnJson), &mapResult)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(listDataJson), &listDataMap)
	if err != nil {
		panic(err)
	}

	mapResult["data"] = listDataMap

	jsonResp, err := json.Marshal(mapResult)
	if err != nil {
		panic(err)
	}

	return jsonResp
}
