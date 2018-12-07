package main

import (
	"C"
	"fmt"
	"sort"
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
	permMap["permView"] = "permission for viewing something"

	//sort.Strings(nameList)
	return 0
}

//export rbac_create_permission
func rbac_create_permission(name, desc string) int32 {

	permMap[name] = desc

	return 0
}

//export rbac_create_role
func rbac_create_role(name, desc string) int32 {

	roleMap[name] = desc

	return 0
}

//export rbac_create_user
func rbac_create_user(name, desc string) int32 {

	userMap[name] = desc

	return 0
}

//export rbac_list_users
func rbac_list_users() int32 {

	var result []string
	for key, value := range userMap {
		str := key + ": " + value
		result = append(result, str)
	}

	print_str_array(result)

	return 0
}

//export rbac_list_roles
func rbac_list_roles() int32 {

	var result []string
	for key, value := range roleMap {
		str := key + ": " + value
		result = append(result, str)
	}

	print_str_array(result)

	return 0
}

//export rbac_list_permissions
func rbac_list_permissions() int32 {

	var result []string
	for key, value := range permMap {
		str := key + ": " + value
		result = append(result, str)
	}

	print_str_array(result)

	return 0
}

func print_str_array(array []string) {
	sort.Strings(array)
	for _, value := range array {
		fmt.Printf("%s. \n", value)
	}
}

func main() {

}
