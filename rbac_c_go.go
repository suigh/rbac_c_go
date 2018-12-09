package main

import (
	"C"
	"fmt"
	"sort"
	"sync"
)

var userMap map[string]string
var roleMap map[string]string
var permMap map[string]string

var userRoleMap map[string]map[string]string
var rolePermMap map[string]map[string]string

var initialized = 0 /*0 not inited, 1 inited, 2 uninited*/
var mutex sync.RWMutex

var resultStr string

/*used for initialization, must be called once and only once*/
//export rbac_init
func rbac_init() int32 {

	mutex.Lock()
	defer mutex.Unlock()

	if initialized != 0 {
		return -1
	}

	initialized = 1

	userMap = make(map[string]string)
	roleMap = make(map[string]string)
	permMap = make(map[string]string)
	userRoleMap = make(map[string]map[string]string)
	rolePermMap = make(map[string]map[string]string)

	/*do something to load the change*/
	return 0
}

//export rbac_uninit
func rbac_uninit() int32 {

	mutex.Lock()
	defer mutex.Unlock()

	if initialized != 1 {
		return -1
	}

	initialized = 2

	/*do something to save the change*/

	return 0
}

//export rbac_create_permission
func rbac_create_permission(name, desc string) int32 {

	mutex.Lock()
	defer mutex.Unlock()

	permMap[name] = desc

	return 0
}

//export rbac_create_role
func rbac_create_role(name, desc string) int32 {

	mutex.Lock()
	defer mutex.Unlock()

	roleMap[name] = desc

	return 0
}

//export rbac_create_user
func rbac_create_user(name, desc string) int32 {

	mutex.Lock()
	defer mutex.Unlock()

	userMap[name] = desc

	return 0
}

//export rbac_list_users
func rbac_list_users() *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = ""
	var resultArray []string
	for key, value := range userMap {
		str := key + ": " + value
		resultStr = resultStr + "," + str
		resultArray = append(resultArray, str)

	}

	//print_str_array(resultArray)
	//fmt.Printf("%s. \n", resultStr)
	return C.CString(resultStr)
}

//export rbac_list_roles
func rbac_list_roles() *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = ""
	var resultArray []string
	for key, value := range roleMap {
		str := key + ": " + value
		resultStr = resultStr + "," + str
		resultArray = append(resultArray, str)

	}

	//print_str_array(resultArray)
	//fmt.Printf("%s. \n", resultStr)
	return C.CString(resultStr)
}

//export rbac_list_permissions
func rbac_list_permissions() *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = ""
	var resultArray []string
	for key, value := range permMap {
		str := key + ": " + value
		resultStr = resultStr + "," + str
		resultArray = append(resultArray, str)

	}

	//print_str_array(resultArray)
	//fmt.Printf("%s. \n", resultStr)
	return C.CString(resultStr)
}

//export rbac_bind_role_permission
func rbac_bind_role_permission(role, perm string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := roleMap[role]; !exist {
		resultStr = "Role doesn't exist"
		C.CString(resultStr)
	}

	if _, exist := permMap[perm]; !exist {
		resultStr = "permission doesn't exist"
		C.CString(resultStr)
	}

	if _, exist := rolePermMap[role][perm]; exist {
		resultStr = "The relationship already exist"
		C.CString(resultStr)
	}

	add_kv_to_2D_map(rolePermMap, role, perm, "")

	return C.CString(resultStr)
}

//export rbac_bind_user_role
func rbac_bind_user_role(user, role string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := userMap[user]; !exist {
		resultStr = "User doesn't exist"
		C.CString(resultStr)
	}

	if _, exist := roleMap[role]; !exist {
		resultStr = "Role doesn't exist"
		C.CString(resultStr)
	}

	if _, exist := userRoleMap[user][role]; exist {
		resultStr = "The relationship already exist"
		C.CString(resultStr)
	}

	add_kv_to_2D_map(userRoleMap, user, role, "")

	return C.CString(resultStr)
}

func print_str_array(array []string) {
	sort.Strings(array)
	for _, value := range array {
		fmt.Printf("%s. \n", value)
	}
}

/*return 1 if already exist, return 0 for success*/
func add_kv_to_2D_map(map2D map[string]map[string]string, key1 string, key2 string, value string) int32 {

	if _, exist := map2D[key1]; exist {
		if _, exist1 := map2D[key1][key2]; exist1 {
			return 1
		}
		map2D[key1][key2] = value
	} else {
		c := make(map[string]string)
		c[key2] = value
		map2D[key1] = c
	}

	return 0
}

func print_2D_map(map2D map[string]map[string]string) {
	for keyB, valB := range userRoleMap {
		for subKeyB, subValB := range valB {
			fmt.Printf("KeyL1:%s -> KeyL2:%s -> value:%s.\n", keyB, subKeyB, subValB)
		}
	}
}

func main() {

}
