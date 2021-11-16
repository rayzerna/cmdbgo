package class

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

// Transform to json
func (rtn *RtnData) ToJson() []byte {
	rtnJson, err := json.Marshal(rtn)
	CheckError(err)
	return rtnJson
}

// Set Dict
func (rtn *RtnData) Dict(dictData map[string]interface{}) []byte {
	rtnJson, err := json.Marshal(rtn)
	CheckError(err)
	dictDatajson, err := json.Marshal(dictData)
	CheckError(err)

	var mapResult map[string]interface{}
	var dictDataMap map[string]interface{}

	err = json.Unmarshal([]byte(rtnJson), &mapResult)
	CheckError(err)
	err = json.Unmarshal([]byte(dictDatajson), &dictDataMap)
	CheckError(err)

	mapResult["data"] = dictDataMap

	jsonResp, err := json.Marshal(mapResult)
	CheckError(err)

	return jsonResp
}

// Set List
func (rtn *RtnData) List(listData []string) []byte {
	rtnJson, err := json.Marshal(rtn)
	CheckError(err)
	listDataJson, err := json.Marshal(listData)
	CheckError(err)

	var mapResult map[string]interface{}
	var listDataMap []interface{}

	err = json.Unmarshal([]byte(rtnJson), &mapResult)
	CheckError(err)
	err = json.Unmarshal([]byte(listDataJson), &listDataMap)
	CheckError(err)

	mapResult["data"] = listDataMap

	jsonResp, err := json.Marshal(mapResult)
	CheckError(err)

	return jsonResp
}

// Set Dict-List
func (rtn *RtnData) DictList(listData []map[string]interface{}) []byte {
	rtnJson, err := json.Marshal(rtn)
	CheckError(err)
	listDataJson, err := json.Marshal(listData)
	CheckError(err)

	var mapResult map[string]interface{}
	var listDataMap []map[string]interface{}

	err = json.Unmarshal([]byte(rtnJson), &mapResult)
	CheckError(err)
	err = json.Unmarshal([]byte(listDataJson), &listDataMap)
	CheckError(err)

	mapResult["data"] = listDataMap

	jsonResp, err := json.Marshal(mapResult)
	CheckError(err)

	return jsonResp
}
