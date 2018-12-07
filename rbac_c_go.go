package main

import (
	"C"
	"fmt"
)

var userMap map[string]string
var roleMap map[string]string
var permMap map[string]string

//export rbac_init
func rbac_init() int32 {

	userMap = make(map[string]string)
	roleMap = make(map[string]string)
	permMap = make(map[string]string)

	userMap["userAdmin"] = "user for administrator"
	userMap["userView"] = "user for viewer"

	roleMap["roleAdmin"] = "role for administrator"
	roleMap["roleView"] = "role for viewer"

	permMap["permChange"] = "permission for changing something"
	permMap["permCheck"] = "permission for viewing something"

	//sort.Strings(nameList)
	return 0
}

//export rbac_list_users
func rbac_list_users() int32 {

	var result []string
	for key, value := range userMap {
		str := key + ": " + value
		result = append(result, str)
	}

	for _, value := range result {
		fmt.Printf("%s. \n", value)
	}
	return 0
}

//export rbac_list_roles
func rbac_list_roles() int32 {

	var result []string
	for key, value := range roleMap {
		str := key + ": " + value
		result = append(result, str)
	}

	for _, value := range result {
		fmt.Printf("%s. \n", value)
	}

	return 0
}

//export rbac_list_permissions
func rbac_list_permissions() int32 {

	var result []string
	for key, value := range permMap {
		str := key + ": " + value
		result = append(result, str)
	}

	for _, value := range result {
		fmt.Printf("%s. \n", value)
	}
	return 0
}

func main() {

}
