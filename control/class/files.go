package class

import (
	"io/ioutil"
)

// Write json to file
func WriteJson(filePath string, data []byte) {
	err := ioutil.WriteFile(filePath, data, 0777)
	CheckError(err)
}

// Read json from file
func ReadJson(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	CheckError(err)
	return data
}
