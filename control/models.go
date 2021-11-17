package control

import (
	"cmdbgo/control/class"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// model method
func Model(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		resp := ModelList(request)
		fmt.Fprintf(writer, string(resp))
	case "POST":
		resp := ModelCreate(request)
		fmt.Fprintf(writer, string(resp))
	case "PUT":
		resp := ModelList(request)
		fmt.Fprintf(writer, string(resp))
	case "DELETE":
		resp := ModelList(request)
		fmt.Fprintf(writer, string(resp))
	default:
		fmt.Fprintf(writer, string([]byte{}))
	}
}

// GET: List models
func ModelList(req *http.Request) []byte {
	query := req.URL.Query()
	id := query.Get("id")
	modelPath := "data/models/"
	if id == "" {
		models, _ := ioutil.ReadDir(modelPath)
		var modelsList []string
		for _, mdl := range models {
			modelsList = append(modelsList, mdl.Name())
		}
		returnData := class.RtnData{}
		returnData.OK()
		return returnData.List(modelsList)
	} else {
		model := modelPath + id
		mdelJson := class.ReadJson(model)
		returnData := class.RtnData{}
		returnData.OK()
		mapModel := class.Json2ListMap(mdelJson)
		return returnData.DictList(mapModel)
	}
}

// POST: Create model
func ModelCreate(req *http.Request) []byte {
	decoder := json.NewDecoder(req.Body)
	var params map[string]interface{}
	decoder.Decode(&params)
	result := CreateModel(params)
	rtn := class.RtnData{}
	if result {
		resp := rtn.OK()
		return resp.ToJson()
	}
	rtn.Code = "1"
	rtn.Msg = "Create model failed."
	return rtn.ToJson()
}

// PUT: Update model
func ModelUpdate(req *http.Request) []byte {
	return []byte{}
}

// DELETE: Delete model
func ModelDelete(req *http.Request) []byte {
	return []byte{}
}

// Check weather model exists
func CheckModelExists(modelName string) bool {
	modelPath := "data/models/"
	models, _ := ioutil.ReadDir(modelPath)
	for _, mdl := range models {
		if modelName == mdl.Name() {
			return true
		}
	}
	return false
}

// Create model
func CreateModel(data map[string]interface{}) bool {
	name := fmt.Sprintf("%s", data["model_name"])
	modelFilePath := "data/models/" + name
	itemFilePath := "data/items/" + name
	ifExists := CheckModelExists(name)
	if ifExists {
		return false
	}
	data["id"] = class.UuidGen()
	class.WriteJson(modelFilePath, class.Map2Json(data))
	var newItem []map[string]interface{}
	newItemJson := class.ListMap2Json(newItem)
	newItemJson = class.FormatedJson(newItemJson)
	class.WriteJson(itemFilePath, newItemJson)
	return true
}
