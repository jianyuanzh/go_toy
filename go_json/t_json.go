package go_json

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"name"`
	ServerIp string `json:"ip"`
	ServerPort int `json:"port"`
}

func SerializeMap() string {
	server := make(map[string]interface{})
	server["name"] = "雪梨花旗"
	server["port"] = 8090
	server["ip"] = "192.168.1.110"

	b, err := json.Marshal(&server)
	if err != nil {
		fmt.Println("序列化失败:", err.Error())
		return ""
	}
	fmt.Println("Marsha map to json:", string(b))
	return string(b)
}

func DeserializeMap(jstr string) {
	server := make(map[string] interface{})
	err := json.Unmarshal([]byte(jstr), &server)
	if (err != nil) {
		fmt.Println("Unmarshal map failed:", err.Error())
		return
	}
	fmt.Println("序列化Map完成：", server)
}

func SerializeServerStruct() string {
	server  := Server{ServerName: "Master", ServerIp: "192.168.0.1", ServerPort: 8090}

	b, err := json.Marshal(&server)
	if err != nil {
		fmt.Println("序列化失败:", err.Error())
		return ""
	}

	fmt.Println("Marsha json: ", string(b))
	return string(b)
}

func DeserializeServer(str string) {
	server := new(Server)
	err := json.Unmarshal([]byte(str), &server)
	if err != nil {
		fmt.Println("Unmarshal error:" , err.Error())
		return
	}
	fmt.Println("Unmarshal struct,", server)
}