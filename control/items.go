package control

import (
	"cmdbgo/control/class"
	"fmt"
)

// Check weather item exists in items of model
func CheckExists(modelPath string, itemName string) bool {
	items := class.ReadJson(modelPath)
	itemsMap := class.Json2Map(items)
	for _, i := range itemsMap {
		if itemName == fmt.Sprintf("%s", i["username"]) {
			return true
		}
	}
	return false
}

// List models items
func ListItem(modelPath string, itemId string) {

}

// Create model's item
func CreateItem(modelPath string, data map[string]interface{}) bool {
	modelFilePath := "data/models/" + modelPath
	username := fmt.Sprintf("%s", data["username"])
	ifExists := CheckExists(modelFilePath, username)
	if ifExists {
		return false
	} else {
		return true
	}
}

// Update models item
func UpdateItem(modelPath string, data map[string]interface{}) {

}

// Delete models item
func DeleteItem(modelPath string, itemId string) {

}
