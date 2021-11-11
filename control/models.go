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
	}
}

// GET
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
		var mapModel []map[string]interface{}
		err := json.Unmarshal([]byte(mdelJson), &mapModel)
		class.CheckError(err)
		return returnData.DictList(mapModel)
	}
}
