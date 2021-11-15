package class

import "encoding/json"

// Json to dict list
func Json2Map(jsonData []byte) []map[string]interface{} {
	var listDataMap []map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &listDataMap)
	CheckError(err)
	return listDataMap
}
