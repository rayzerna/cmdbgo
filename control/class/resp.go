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

	mapResult := Json2Map(rtnJson)
	dictDataMap := Json2Map(dictDatajson)

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

	mapResult := Json2Map(rtnJson)
	listDataMap := Json2List(listDataJson)

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

	mapResult := Json2Map(rtnJson)
	listDataMap := Json2ListMap(listDataJson)

	mapResult["data"] = listDataMap

	jsonResp, err := json.Marshal(mapResult)
	CheckError(err)

	return jsonResp
}
