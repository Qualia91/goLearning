package common

import (
	"encoding/json"
	"fmt"
)

func MessageToJson(message Message) string {

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Cant Serialise Message")
		return ""
	}

	return string(jsonMessage)
}
