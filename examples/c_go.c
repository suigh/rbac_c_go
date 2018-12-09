#include "../rbac_c_go.h"
#include <stdio.h>
#include <string.h>

int main() {
    char * p;
    int rVal = 0;
    rVal = rbac_init();
    printf("The return value of rbac_init is %d. \n", rVal);

    
    {
        char name[] = "permAdmin";
        int lenName = strlen(name);
        GoString goName = {name, lenName};    

        char desc[] = "permpermission for admin";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_permission(goName, goDesc);  
    }

    {
        char name[] = "roleAdmin";
        int lenName = strlen(name);
        GoString goName = {name, lenName};

        char desc[] = "role for administrato";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc}; 

        rbac_create_role(goName, goDesc);  
    }

    {
        char name[] = "userAdmin";
        int lenName = strlen(name);
        GoString goName = {name, lenName};   

        char desc[] = "user for administrato";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_user(goName, goDesc);  
    }    

    p = rbac_list_users();
    printf("Result is %s. \n", p);

    p = rbac_list_roles();
    printf("Result is %s. \n", p);

    p = rbac_list_permissions();
    printf("Result is %s. \n", p);    

    {
        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};  

        char permission[] = "permAdmin";
        int lenPerm = strlen(permission);
        GoString goPerm = {permission, lenPerm};   

        p = rbac_bind_role_permission(goRole, goPerm);  
        printf("Result is %s. \n", p);
    } 

    {
        char user[] = "userAdmin";
        int lenUser = strlen(user);
        GoString goUser = {user, lenUser};   

        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};         

        p = rbac_bind_user_role(goUser, goRole);  
        printf("Result is %s. \n", p);
    } 

    /*rbac_bind_role_permission();
    rbac_bind_user_role();

    /*rbac_user_has_permission();
    
    rabc_list_roles_by_user();
    rabc_list_permissions_by_role();

    persist and load in Jason in GO.
    */


    rbac_uninit();

    return 0;
}
