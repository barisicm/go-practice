package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// jsonResultValid1 := json.RawMessage(`
	// {
	// 	"result": [
	// 		[
	// 			{
	// 				"id": 1,
	// 				"name":"test",
	// 				"last_name":"test",
	// 				"name2":"test",
	// 				"icon":"female",
	// 				"uuid":"1109e309-59c0-488b-8603-6d92fc996b92",
	// 				"created_at":"2020-09-22 14:42:07"
	// 			},
	// 			{
	// 				"id": 1,
	// 				"name":"test",
	// 				"last_name":"test",
	// 				"name2":"test",
	// 				"icon":"female",
	// 				"uuid":"1109e309-59c0-488b-8603-6d92fc996b92",
	// 				"created_at":"2020-09-22 14:42:07"
	// 			}
	// 		],
	// 		[
	// 			{
	// 				"mac":"98:fa:9b:84:a2:31",
	// 				"test":null,
	// 				"uuid":"c6db5055-7e65-4e5f-9833-a830e4e17763",
	// 				"created_at":"2020-09-23 12:58:42",
	// 				"ipv4":"192.168.1.182",
	// 				"hostname":"mbarisic",
	// 				"port":"eth1.0",
	// 				"conn_type":"wired"
	// 			}
	// 		]
	// 	]
	// }
	// `)

	jsonResValid2 := json.RawMessage(`
	[[[1,"test","test","test","female","1109e309-59c0-488b-8603-6d92fc996b92","2020-09-22 14:42:07"],[3,"subskrajber1","","","","def22de9-89e9-44d7-bdf8-e2210a75b769","2020-09-23 10:45:30"],[4,"Im","Ne","a@f.c","male","70e93218-daee-490f-b44f-911fa91a6d78","2020-09-09 12:09:47"],[5,"ysd","erw","a.b@gh","female","f5bd4077-2342-4e89-933c-c8478170e639","2020-09-09 11:52:43"]],[["98:fa:9b:84:a2:31",null,"c6db5055-7e65-4e5f-9833-a830e4e17763","2020-09-21 14:48:23","192.168.1.182","","mbarisic","","eth1.0","wired"],["10:10:10:01:01:01","1109e309-59c0-488b-8603-6d92fc996b92","01f7cd40-c896-450a-895a-f7d9effb5cb6","2020-09-23 11:50:39","192.168.1.50","","","","",""],["f8:e6:1a:c8:34:01",null,"b2cf4fef-a476-42b3-b3f3-af5621a92716","2020-09-07 11:05:43","192.168.1.179","","Galaxy-S8","","wl0","wireless_50"],["56:36:35:66:36:53",null,"3d26c6b6-3af6-46db-b646-ab6bc93de7aa","2020-09-15 13:57:24","","","","","",""]],[]]`)

	var responses [][][]interface{}

	err := json.Unmarshal(jsonResValid2, &responses)
	if err != nil {
		fmt.Printf("things went to shit")
	}

}
