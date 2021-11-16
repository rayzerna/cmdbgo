package class

import (
	"bytes"
	"encoding/json"
)

// Json to dict list
func Json2ListMap(jsonData []byte) []map[string]interface{} {
	var listDataMap []map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &listDataMap)
	CheckError(err)
	return listDataMap
}

// Json to dict
func Json2Map(jsonData []byte) map[string]interface{} {
	var DataMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &DataMap)
	CheckError(err)
	return DataMap
}

// Dict list to json
func ListMap2Json(mapData []map[string]interface{}) []byte {
	dictDatajson, err := json.Marshal(mapData)
	CheckError(err)
	return dictDatajson
}

// Dict to json
func Map2Json(mapData map[string]interface{}) []byte {
	dictDatajson, err := json.Marshal(mapData)
	CheckError(err)
	return dictDatajson
}

// formated json
func FormatedJson(jsonData []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, jsonData, "", "\t")
	CheckError(err)
	return out.Bytes()
}
