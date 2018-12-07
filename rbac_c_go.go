package main

import (
	"C"
	"fmt"
	"sort"

	"github.com/suigh/gorbac"
)

var gRbac *gorbac.RBAC

//export rbac_init
func rbac_init() int32 {
	gRbac = gorbac.New()

	rAdmin := gorbac.NewStdRole("Role_Admin")
	rView := gorbac.NewStdRole("Role_View")

	gRbac.Add(rAdmin)
	gRbac.Add(rView)

	pChange := gorbac.NewStdPermission("permission-change")
	pCheck := gorbac.NewStdPermission("permission-check")

	rAdmin.Assign(pChange)
	rAdmin.Assign(pCheck)

	rView.Assign(pCheck)

	return 0
}

//export rbac_list_users
func rbac_list_users() int32 {
	fmt.Println("rbac_list_users called.")

	return 0
}

//export rbac_list_roles
func rbac_list_roles() int32 {
	nameList := gRbac.GetRoleNameList()
	sort.Strings(nameList)

	for _, value := range nameList {
		fmt.Printf("rbac_list_roles: %s. \n", value)
	}

	return 0
}

//export rbac_list_permissions
func rbac_list_permissions() int32 {
	fmt.Println("rbac_list_permissions called.")

	return 0
}

func main() {

}
