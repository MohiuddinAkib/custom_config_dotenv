package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MohiuddinAkib/my_first_goproj/dotenv"
)

var validJSONFiles = [4]string{"custom-environment-variables.json", "production.json", "development.json", "default.json"}

var goEnv = dotenv.Get("GO_ENV")

var fileContentMap = make(map[string]interface{})

// Load (load the config dir)
func Load() {
	// Set default go env
	if goEnv == nil {
		goEnv = "development"
	}
	// Currently working directory
	workingDir, _ := os.Getwd()
	configDir := path.Join(workingDir, "config-json")
	// Read the config dir
	files, err := ioutil.ReadDir(configDir)
	// Handle error if config folder does not exists
	if os.IsNotExist(err) {
		fmt.Println("Config dir doest not exist")
		return
	}

	for _, file := range files {
		// Get only the json files
		if !file.IsDir() && strings.Contains(file.Name(), ".json") {
			if contains(validJSONFiles, file.Name()) {
				// Reading the files
				bytes, _ := ioutil.ReadFile(path.Join(configDir, file.Name()))
				var mappedData map[string]interface{}
				// Decoding the json data
				json.Unmarshal(bytes, &mappedData)

				// Set the data and file name as key val pair
				fileContentMap[file.Name()] = mappedData
			}
		}

	}
}

// Get (Get config vars)
func Get(key string) interface{} {
	// Trim the key
	trimmedKey := strings.Trim(key, " ")
	trimmedKey = strings.TrimSpace(key)
	// Split the string with dot notion
	splittedByDot := strings.Split(trimmedKey, ".")
	return getPropVal(splittedByDot)
}

func getPropVal(keys []string) interface{} {
	var result interface{}
	for _, JSONFile := range validJSONFiles {
		value := fileContentMap[JSONFile]

		for index := 0; index < len(keys); index++ {
			if value != nil {
				value = value.(map[string]interface{})[keys[index]]
				result = value
			}
		}

		// Priority most to custom-environment-variables file
		if JSONFile == "custom-environment-variables.json" && result != nil {
			result = dotenv.Get(result.(string))
			break
		}

		envJSONFile := goEnv.(string) + ".json"

		// Priority secondly to environment specific json files
		if JSONFile == envJSONFile && result == nil {
			continue
		} else if JSONFile == envJSONFile && result != nil {
			break
		}
	}

	return result
}
