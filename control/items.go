package control

import (
	"cmdbgo/control/class"
	"fmt"
	"net/http"
)

// item method
func Item(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		resp := ItemList(request)
		fmt.Fprintf(writer, string(resp))
	}
}

// GET: List items
func ItemList(req *http.Request) []byte {
	query := req.URL.Query()
	id := query.Get("id")
	model := query.Get("model")
	itemsList := ListItem(model, id)

	returnData := class.RtnData{}
	if itemsList == nil {
		returnData.Code = "1"
		returnData.Msg = "Item(s) not found."
		return returnData.ToJson()
	}
	returnData = returnData.OK()
	fmt.Println(itemsList)
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
func ListItem(modelPath string, itemId string) []byte {
	itemFilePath := "data/items/" + modelPath
	itemsJson := class.ReadJson(itemFilePath)
	if itemId == "" {
		return itemsJson
	}
	itemsMap := class.Json2ListMap(itemsJson)
	for _, item := range itemsMap {
		if item["id"] == itemId {
			return class.Map2Json(item)
		}
	}
	return nil
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
