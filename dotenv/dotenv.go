package dotenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var envKeyValMap map[string]interface{} = make(map[string]interface{})

// Load (loads env file)
func Load() {

	dir, _ := os.Getwd()
	bytes, err := ioutil.ReadFile(filepath.Join(dir, ".env"))

	// If err is about .env not exist then handle it
	if os.IsNotExist(err) {
		fmt.Println(err)
	}

	// Unsplitted chunk of env key values
	unsplitedEnvVars := string(bytes)

	// Splitting the whole chunk of string into array by new line
	splitedByLinesEnvVars := strings.Split(unsplitedEnvVars, "\n")

	for _, eachEnvVarAfterSplitedByLine := range splitedByLinesEnvVars {
		// Removing any whitespace if exists
		trimmedEnvVar := strings.Trim(eachEnvVarAfterSplitedByLine, " ")
		trimmedEnvVar = strings.TrimSpace(eachEnvVarAfterSplitedByLine)
		// Now splitting each row of env key value pair by = sign
		eachEnvVarKeyValueSplitedByEqualSign := strings.Split(trimmedEnvVar, "=")

		if len(eachEnvVarKeyValueSplitedByEqualSign) > 1 {
			// Reserving the envKey and envValue in vars
			envKey, envValue := eachEnvVarKeyValueSplitedByEqualSign[0], eachEnvVarKeyValueSplitedByEqualSign[1]

			// Add the key value to the os env
			if len(os.Getenv(envKey)) == 0 {
				os.Setenv(envKey, envValue)
			}

			// Add the key value to the map if they are not there already
			if envKeyValMap[envKey] == nil {
				envKeyValMap[envKey] = envValue
			}
		}

	}
}

// Get (Gets the env variable)
func Get(key string) interface{} {
	// Check if that key actually exists or not
	if envKeyValMap[key] == nil {
		return nil
	}
	return os.Getenv(key)
}
