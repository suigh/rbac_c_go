#include "../rbac_c_go.h"
#include <stdio.h>
#include <string.h>

int main() {
    char * p;
    int rVal = 0;
    rVal = rbac_init();

    /* {
        char name[] = "permAdmin";
        int lenName = strlen(name);
        GoString goName = {name, lenName};    

        char desc[] = "permpermission for admin";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_permission(goName, goDesc);  
    }

    {
        char name[] = "permView";
        int lenName = strlen(name);
        GoString goName = {name, lenName};    

        char desc[] = "permpermission for view";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_permission(goName, goDesc);  
    }    

    {
        char name[] = "roleAdmin";
        int lenName = strlen(name);
        GoString goName = {name, lenName};

        char desc[] = "role for admin";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc}; 

        rbac_create_role(goName, goDesc);  
    }

    {
        char name[] = "roleView";
        int lenName = strlen(name);
        GoString goName = {name, lenName};

        char desc[] = "role for view";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc}; 

        rbac_create_role(goName, goDesc);  
    }    

    {
        char name[] = "userAdmin";
        int lenName = strlen(name);
        GoString goName = {name, lenName};   

        char desc[] = "user for admin";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_user(goName, goDesc);  
    }    

    {
        char name[] = "userView";
        int lenName = strlen(name);
        GoString goName = {name, lenName};   

        char desc[] = "user for view";
        int lenDesc = strlen(desc);
        GoString goDesc = {desc, lenDesc};  

        rbac_create_user(goName, goDesc);  
    } */

    p = rbac_list_users();
    printf("Result is %s. \n", p);

    p = rbac_list_roles();
    printf("Result is %s. \n", p);

    p = rbac_list_permissions();
    printf("Result is %s. \n", p);    

    /*{
        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};  

        char permission[] = "permAdmin";
        int lenPerm = strlen(permission);
        GoString goPerm = {permission, lenPerm};   

        p = rbac_bind_role_permission(goRole, goPerm);  
        //printf("Result is %s. \n", p);
    }     

    {
        char role[] = "roleView";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};  

        char permission[] = "permView";
        int lenPerm = strlen(permission);
        GoString goPerm = {permission, lenPerm};   

        p = rbac_bind_role_permission(goRole, goPerm);  
        //printf("Result is %s. \n", p);
    }    

    {
        char user[] = "userAdmin";
        int lenUser = strlen(user);
        GoString goUser = {user, lenUser};   

        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};         

        p = rbac_bind_user_role(goUser, goRole);  
        //printf("Result is %s. \n", p);
    } 

    {
        char user[] = "userAdmin";
        int lenUser = strlen(user);
        GoString goUser = {user, lenUser};   

        char role[] = "roleView";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};         

        p = rbac_bind_user_role(goUser, goRole);  
        //printf("Result is %s. \n", p);
    }   */ 

    {
        char user[] = "userAdmin";
        int lenUser = strlen(user);
        GoString goUser = {user, lenUser};          

        p = rbac_list_roles_by_user(goUser);  
        printf("Result is %s. \n", p);
    }

    {
        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};           

        p = rbac_list_permissions_by_role(goRole);  
        printf("Result is %s. \n", p);
    }

    {
        char user[] = "userAdmin";
        int lenUser = strlen(user);
        GoString goUser = {user, lenUser};          

        char permission[] = "permAdmin";
        int lenPerm = strlen(permission);
        GoString goPerm = {permission, lenPerm};

        p = rbac_user_has_permission(goUser, goPerm);  
        printf("Result is %s. \n", p);
    } 

    /*
    {
        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};  

        char permission[] = "permAdmin";
        int lenPerm = strlen(permission);
        GoString goPerm = {permission, lenPerm};   

        p = rbac_unbind_role_permission(goRole, goPerm);  
        //printf("Result is %s. \n", p);
    }

    {
        char user[] = "userAdmin";
        int lenUser = strlen(user);
        GoString goUser = {user, lenUser};   

        char role[] = "roleAdmin";
        int lenRole = strlen(role);
        GoString goRole = {role, lenRole};         

        p = rbac_unbind_user_role(goUser, goRole);  
        //printf("Result is %s. \n", p);
    } 
    

    {
        char permission[] = "permAdmin";
        int lenPerm = strlen(permission);
        GoString goPerm = {permission, lenPerm};   

        p = rbac_delete_permission(goPerm);  
    }      
    */

    rbac_uninit();

    return 0;
}
