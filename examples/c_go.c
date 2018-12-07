#include "../rbac_c_go.h"
#include <stdio.h>
#include <string.h>

int main() {
    int rVal = 0;
    rVal = rbac_init();
    printf("The return value of rbac_init is %d. \n", rVal);

    rbac_list_users();
    rbac_list_roles();
    rbac_list_permissions();
    
    {
        char name[] = "permGetLog";
        int lenName = strlen(name);
        GoString goName = {name, lenName};    

        char desc[] = "permpermission for getting logs";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_permission(goName, goDesc);  
    }

    {
        char name[] = "roleLog";
        int lenName = strlen(name);
        GoString goName = {name, lenName};

        char desc[] = "role for logs";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc}; 

        rbac_create_role(goName, goDesc);  
    }

    {
        char name[] = "userLog";
        int lenName = strlen(name);
        GoString goName = {name, lenName};   

        char desc[] = "user for logs";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_user(goName, goDesc);  
    }    

    printf("\n");
    
    rbac_list_users();
    rbac_list_roles();
    rbac_list_permissions();

    /*rbac_assign_permission_to_role();
    rabc_assign_role_to_user();

    rbac_user_has_permission();*/

    return 0;
}
