package control

import (
	"cmdbgo/control/class"
	"fmt"
	"net/http"
	"strings"
)

// item method
func Item(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		resp := ItemList(request)
		fmt.Fprintf(writer, string(resp))
	case "POST":
		resp := ItemList(request)
		fmt.Fprintf(writer, string(resp))
	case "PUT":
		resp := ItemList(request)
		fmt.Fprintf(writer, string(resp))
	case "DELETE":
		resp := ItemList(request)
		fmt.Fprintf(writer, string(resp))
	}
}

// GET: List items
func ItemList(req *http.Request) []byte {
	query := req.URL.Query()
	id := query.Get("id")
	model := query.Get("model")
	listDetail := query.Get("showDetail")
	showDetail := true
	if listDetail == "" {
		showDetail = false
	}
	itemsList := ListItem(model, id, showDetail)

	returnData := class.RtnData{}
	if itemsList == nil {
		returnData.Code = "1"
		returnData.Msg = "Item(s) not found."
		return returnData.ToJson()
	}
	returnData = returnData.OK()
	if id != "" {
		return returnData.Dict(class.Json2Map(itemsList))
	}
	return returnData.DictList(class.Json2ListMap(itemsList))
}

// Check weather item exists in items of model
func CheckItemExists(itemPath string, itemName string) bool {
	items := class.ReadJson(itemPath)
	itemsMap := class.Json2ListMap(items)
	for _, i := range itemsMap {
		if itemName == fmt.Sprintf("%s", i["username"]) {
			return true
		}
	}
	return false
}

// List models items
func ListItem(modelPath string, itemId string, relatedDetail bool) []byte {
	itemFilePath := "data/items/" + modelPath
	itemsJson := class.ReadJson(itemFilePath)
	var result []byte
	if itemId == "" {
		result = itemsJson
		// 查询详情，展示关联关系第一层
		if relatedDetail {
			result = RelatedItemReplace(itemsJson, modelPath)
		}
	}
	itemsMap := class.Json2ListMap(itemsJson)
	for _, item := range itemsMap {
		if item["id"] == itemId {
			result = class.Map2Json(item)
			break
		}
	}
	return result
}

// Create model's item
func CreateItem(modelPath string, data map[string]interface{}) bool {
	modelFilePath := "data/models/" + modelPath
	itemFilePath := "data/items/" + modelPath
	name := fmt.Sprintf("%s", data["name"])
	ifExists := CheckItemExists(itemFilePath, name)
	if ifExists {
		return false
	}

	modelField := class.ReadJson(modelFilePath)
	modelMap := class.Json2Map(modelField)
	for key, _ := range modelMap {
		if key == "id" {
			modelMap["id"] = class.UuidGen()
		} else {
			modelMap[key] = fmt.Sprintf("%s", data[key])
		}
	}
	itemJson := class.ReadJson(itemFilePath)
	itemListMap := class.Json2ListMap(itemJson)
	itemListMap = append(itemListMap, modelMap)
	itemListJson := class.ListMap2Json(itemListMap)
	itemListJson = class.FormatedJson(itemListJson)
	class.WriteJson(itemFilePath, itemListJson)
	return true
}

// Update models item
func UpdateItem(modelPath string, data map[string]interface{}) {

}

// Delete models item
func DeleteItem(modelPath string, itemId string) {

}

// ====== Custom functions place there =======
// Key value formated of a dict list
func KeyValueFormat(key string, dictListData []map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, item := range dictListData {
		getKey := item[key].(string)
		result[getKey] = item
	}
	return result
}

// Replace the 1st layer of items to detail
func RelatedItemReplace(itemJson []byte, modelName string) []byte {
	itemMap := class.Json2ListMap(itemJson)
	modelFilePath := "data/models/" + modelName
	modelJson := class.ReadJson(modelFilePath)
	modelMap := class.Json2Map(modelJson)
	// 循环取数，替换原列表
	for key, item := range modelMap {
		strItem := item.(string)
		if strings.HasPrefix(strItem, "Refer") {
			splitString := strings.Split(strItem, ":")
			refModelName := splitString[1]
			refItemFilePath := "data/items/" + refModelName
			refItemJson := class.ReadJson(refItemFilePath)
			refItemKVMapKey := "id"
			// 转换成KV格式的MAP，便于取值
			refItemKVMap := KeyValueFormat(refItemKVMapKey, class.Json2ListMap(refItemJson))
			var refItemMap []map[string]interface{}
			for seq, i := range itemMap {
				refItemArray := i[key].([]interface{})
				for _, j := range refItemArray {
					refItemId := j.(string)
					refItem := refItemKVMap[refItemId].(map[string]interface{})
					refItemMap = append(refItemMap, refItem)
				}
				itemMap[seq][key] = refItemMap
			}
		}
	}
	return class.ListMap2Json(itemMap)
}
