package main

import (
	"C"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

//used to store users, roles, and permissions
var userMap map[string]string
var roleMap map[string]string
var permMap map[string]string

//used to store the relationshiops for user->role and role->permission
var userRoleMap map[string]map[string]string
var rolePermMap map[string]map[string]string

var initialized = 0 /*0 not inited, 1 inited, 2 uninited*/
var mutex sync.RWMutex

var resultStr string

func loadJSON(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}

func saveJSON(filename string, v interface{}) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(v)
}

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

	if err := loadJSON("user.json", &userMap); err != nil {
		//log.Fatal(err)
	}

	if err := loadJSON("role.json", &roleMap); err != nil {
		//log.Fatal(err)
	}

	if err := loadJSON("perm.json", &permMap); err != nil {
		//log.Fatal(err)
	}

	if err := loadJSON("userrole.json", &userRoleMap); err != nil {
		//log.Fatal(err)
	}
	//print2DMap(userRoleMap)

	if err := loadJSON("roleperm.json", &rolePermMap); err != nil {
		//log.Fatal(err)
	}
	//print2DMap(rolePermMap)

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

	if err := saveJSON("user.json", &userMap); err != nil {
		log.Fatal(err)
	}

	if err := saveJSON("role.json", &roleMap); err != nil {
		log.Fatal(err)
	}

	if err := saveJSON("perm.json", &permMap); err != nil {
		log.Fatal(err)
	}

	if err := saveJSON("userrole.json", &userRoleMap); err != nil {
		log.Fatal(err)
	}

	if err := saveJSON("roleperm.json", &rolePermMap); err != nil {
		log.Fatal(err)
	}

	return 0
}

//export rbac_create_permission
func rbac_create_permission(name, desc string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := permMap[name]; exist {
		resultStr = "permission already exist"
		return C.CString(resultStr)
	}

	permMap[name] = desc

	return C.CString(resultStr)
}

//export rbac_delete_permission
func rbac_delete_permission(name string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := permMap[name]; !exist {
		resultStr = "permission doesn't exist"
		return C.CString(resultStr)
	}

	delete(permMap, name)

	for _, valL1 := range rolePermMap {
		for keyL2 := range valL1 {
			if keyL2 == name {
				delete(valL1, name)
			}
		}
	}

	return C.CString(resultStr)
}

//export rbac_create_role
func rbac_create_role(name, desc string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := roleMap[name]; exist {
		resultStr = "role already exist"
		return C.CString(resultStr)
	}

	roleMap[name] = desc

	return C.CString(resultStr)
}

//export rbac_delete_role
func rbac_delete_role(name string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := roleMap[name]; !exist {
		resultStr = "role doesn't exist"
		return C.CString(resultStr)
	}

	delete(roleMap, name)

	for keyL1 := range rolePermMap {
		if keyL1 == name {
			delete(rolePermMap, name)
		}
	}

	return C.CString(resultStr)
}

//export rbac_create_user
func rbac_create_user(name, desc string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := userMap[name]; exist {
		resultStr = "user already exist"
		return C.CString(resultStr)
	}

	userMap[name] = desc

	return C.CString(resultStr)
}

//export rbac_delete_user
func rbac_delete_user(name string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := userMap[name]; !exist {
		resultStr = "user doesn't exist"
		return C.CString(resultStr)
	}

	delete(userMap, name)

	for keyL1 := range userRoleMap {
		if keyL1 == name {
			delete(userRoleMap, name)
		}
	}

	return C.CString(resultStr)
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

	//printStrArray(resultArray)
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

	//printStrArray(resultArray)
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

	//printStrArray(resultArray)
	//fmt.Printf("%s. \n", resultStr)
	return C.CString(resultStr)
}

//export rabc_list_roles_by_user
func rabc_list_roles_by_user(user string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = ""

	for keyL1, valL1 := range userRoleMap {
		for keyL2 := range valL1 {
			if user == keyL1 {
				resultStr = resultStr + "," + keyL2
			}
		}
	}

	return C.CString(resultStr)
}

//export rabc_list_permissions_by_role
func rabc_list_permissions_by_role(role string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = ""

	for keyL1, valL1 := range rolePermMap {
		for keyL2 := range valL1 {
			if role == keyL1 {
				resultStr = resultStr + "," + keyL2
			}
		}
	}

	return C.CString(resultStr)
}

//export rbac_user_has_permission
func rbac_user_has_permission(user, perm string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "NOK"

	for keyURL1, valURL1 := range userRoleMap {
		if user == keyURL1 {
			for keyURL2 := range valURL1 {
				for keyRPL1, valRPL1 := range rolePermMap {
					for keyRPL2 := range valRPL1 {
						if keyURL2 == keyRPL1 && keyRPL2 == perm {
							resultStr = "OK"
							return C.CString(resultStr)
						}
					}
				}
			}

		}
	}

	return C.CString(resultStr)
}

//export rbac_bind_role_permission
func rbac_bind_role_permission(role, perm string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := roleMap[role]; !exist {
		resultStr = "Role doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := permMap[perm]; !exist {
		resultStr = "permission doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := rolePermMap[role][perm]; exist {
		resultStr = "The relationship already exist"
		return C.CString(resultStr)
	}

	addKVTo2DMap(rolePermMap, role, perm, "")

	return C.CString(resultStr)
}

//export rbac_unbind_role_permission
func rbac_unbind_role_permission(role, perm string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := roleMap[role]; !exist {
		resultStr = "Role doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := permMap[perm]; !exist {
		resultStr = "permission doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := rolePermMap[role][perm]; !exist {
		resultStr = "The relationship doesn't exist"
		return C.CString(resultStr)
	}

	for keyL1, valL1 := range rolePermMap {
		for keyL2 := range valL1 {
			if keyL1 == role && keyL2 == perm {
				delete(valL1, perm)
			}
		}
	}

	return C.CString(resultStr)
}

//export rbac_unbind_user_role
func rbac_unbind_user_role(user, role string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := userMap[user]; !exist {
		resultStr = "User doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := roleMap[role]; !exist {
		resultStr = "Role doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := userRoleMap[user][role]; !exist {
		resultStr = "The relationship doesn't exist"
		return C.CString(resultStr)
	}

	for keyL1, valL1 := range userRoleMap {
		for keyL2 := range valL1 {
			if keyL1 == user && keyL2 == role {
				delete(valL1, role)
			}
		}
	}

	return C.CString(resultStr)
}

//export rbac_bind_user_role
func rbac_bind_user_role(user, role string) *C.char {

	mutex.Lock()
	defer mutex.Unlock()

	resultStr = "OK"

	if _, exist := userMap[user]; !exist {
		resultStr = "User doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := roleMap[role]; !exist {
		resultStr = "Role doesn't exist"
		return C.CString(resultStr)
	}

	if _, exist := userRoleMap[user][role]; exist {
		resultStr = "The relationship already exist"
		return C.CString(resultStr)
	}

	addKVTo2DMap(userRoleMap, user, role, "")

	return C.CString(resultStr)
}

func printStrArray(array []string) {
	sort.Strings(array)
	for _, value := range array {
		fmt.Printf("%s. \n", value)
	}
}

/*return 1 if already exist, return 0 for success*/
func addKVTo2DMap(map2D map[string]map[string]string, key1 string, key2 string, value string) int32 {

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

func print2DMap(map2D map[string]map[string]string) {
	for keyL1, valL1 := range map2D {
		for keyL2, valL2 := range valL1 {
			fmt.Printf("KeyL1:%s -> KeyL2:%s -> value:%s.\n", keyL1, keyL2, valL2)
		}
	}
}

func main() {

}
