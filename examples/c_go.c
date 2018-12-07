#include "../rbac_c_go.h"
#include <stdio.h>
#include <string.h>

int main() {
    char cvalue[] = "From C.";
    int length = strlen(cvalue);
    int rVal = 0;
    GoString value = {cvalue, length};//go中的字符串类型在c中为GoString

    rVal = rbac_init();
    printf("The return value of rbac_init is %d. \n", rVal);


    rbac_list_users();
    rbac_list_roles();
    rbac_list_permissions();
    
    /*rbac_create_permission();  
    rbac_create_role();
    rbac_create_user();  

    rbac_assign_permission_to_role();
    rabc_assign_role_to_user();

    rbac_user_has_permission();*/

    return 0;
}
