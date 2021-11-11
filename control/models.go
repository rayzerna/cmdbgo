package control

import (
	"cmdbgo/control/class"
	"fmt"
	"io/ioutil"
	"net/http"
)

// model method
func Model(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		resp := ModelList()
		fmt.Fprintf(writer, string(resp))
	}
}

// GET
func ModelList() []byte {
	models, _ := ioutil.ReadDir("data/models/")
	var modelsList []string
	for _, mdl := range models {
		modelsList = append(modelsList, mdl.Name())
	}
	returnData := class.RtnData{}
	returnData.OK()
	return returnData.List(modelsList)
}
