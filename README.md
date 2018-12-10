rbac_c_go 
======
rbac_c_go implements the RBAC with go language and shows how to use it with C language.

For the purposes of this package:

	* add/delete users, roles, permissions.
    * build/unbuild relationship between user and role, one user can have several roles. 
    * build/unbuild relationship between role and permission, one role can have several permissions.
    * user get its permissions by the assigned roles.
    * the above configuration can be saved/loaded.
	
Usage (only tested on Linux)
=====

Download the codes 
set GOPATH, and then run go get github.com/suigh/rbac_c_go

Run the sample
Enter the folder rbac_c_go, and run Make to build the go module, you will get the .h and .a file for the C language.
Enter the folder examples, and run Make to build the sample for C language, then you can run ./a to have a try.

The APIs
The two important APIs are:
rbac_init: it is used to initailize the environment and load the saved configuration, it needs to be called at the beginning of the usage, and only once.
rbac_uninit: it is used to save the configuration, it needs to be called at the ending of the usage.

List other APIs here, you can get the functions of them by name, they are thread safe.
rbac_bind_role_permission(role, perm string) *C.char
rbac_bind_user_role(user, role string) *C.char
rbac_create_permission(name, desc string) *C.char
rbac_create_role(name, desc string) *C.char
rbac_create_user(name, desc string) *C.char
rbac_delete_permission(name string) *C.char
rbac_delete_role(name string) *C.char
rbac_delete_user(name string) *C.char
rbac_list_permissions() *C.char
rbac_list_permissions_by_role(role string) *C.char
rbac_list_roles() *C.char
rbac_list_roles_by_user(user string) *C.char
rbac_list_users() *C.char
rbac_unbind_role_permission(role, perm string) *C.char
rbac_unbind_user_role(user, role string) *C.char
rbac_user_has_permission(user, perm string) *C.char

Authors
=======

 * Guang Han Sui <suiguanghan@gmail.com> 

Open Source - MIT Software License
==================================

See LICENSE.
