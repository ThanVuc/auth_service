package helper

import (
	"auth_service/global"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

type ResourceItem struct {
	Resource string   `json:"resource"`
	Actions  []string `json:"actions"`
}

var resourceList []ResourceItem = []ResourceItem{}

func AddResource(resource string, actions []string) {
	resourceList = append(resourceList, ResourceItem{
		Resource: resource,
		Actions:  actions,
	})
}

func GetResources(fileName string) []ResourceItem {
	filePath := "./backup/" + fileName + ".json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []ResourceItem{}
	}

	file, err := os.Open(filePath)
	if err != nil {
		return []ResourceItem{}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	resourceList = []ResourceItem{}
	if err := decoder.Decode(resourceList); err != nil {
		return []ResourceItem{}
	}
	return resourceList
}

func WriteToJsonFile(fileName string) {
	logger := global.Logger
	filePath := "./backup/" + fileName + ".json"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logger.ErrorString("Failed to open file for writing", zap.Error(err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(resourceList); err != nil {
		logger.ErrorString("Failed to encode resource list to JSON", zap.Error(err))
	} else {
		logger.InfoString("Resource list written to JSON file successfully", zap.String("file_path", filePath))
	}
}
