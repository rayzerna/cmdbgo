package control

// Check weather item exists in items of model
func CheckExists(modelPath string, itemName string) {

}

// List models items
func ListItem(modelPath string, itemId string) {

}

// Create model's item
func CreateItem(modelPath string, data map[string]interface{}) {
	modelFilePath := "data/models/" + modelPath
	CheckExists(modelFilePath, data.Get("username"))
}

// Update models item
func UpdateItem(modelPath string, data map[string]interface{}) {

}

// Delete models item
func DeleteItem(modelPath string, itemId string) {

}
