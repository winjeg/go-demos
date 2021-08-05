package casbin

import (
	"fmt"
	"github.com/casbin/casbin"
	"testing"
)

// TestTenants test tenants
func TestTenants(t *testing.T) {
	e := casbin.NewEnforcer("model.conf", "policy.csv")
	roles := e.GetAllRoles()
	fmt.Println(roles)
	subs := e.GetAllSubjects()
	fmt.Println(subs)
	objs := e.GetAllObjects()
	fmt.Println(objs)
	acts := e.GetAllActions()
	fmt.Println(acts)
	users, _:= e.GetUsersForRole("admin")
	fmt.Println(users)
	//e.RemoveNamedGroupingPolicy()
}