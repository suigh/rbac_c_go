#include "../rbac_c_go.h"
#include <stdio.h>
#include <string.h>

int main() {
    char * p;
    int rVal = 0;
    rVal = rbac_init();
    printf("The return value of rbac_init is %d. \n", rVal);

    
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

    p = rbac_list_users();
    printf("Result is %s. \n", p);

    p = rbac_list_roles();
    printf("Result is %s. \n", p);

    p = rbac_list_permissions();
    printf("Result is %s. \n", p);    

    /*rbac_assign_permission_to_role();
    rabc_assign_role_to_user();

    rbac_user_has_permission();*/

    rbac_uninit();

    return 0;
}
