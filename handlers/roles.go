package handlers

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

var RolePermissions = map[string][]string{
	RoleAdmin: {"register", "mint", "transfer", "balance", "accountID", "initialize"},
	RoleUser:  {"transfer", "balance", "accountID"},
}

func HasPermission(role, permission string) bool {
	permissions, ok := RolePermissions[role]
	if !ok {
		return false
	}

	for _, p := range permissions {
		if p == permission {
			return true
		}
	}

	return false
}