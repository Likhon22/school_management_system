package contextkeys

type contextKey string

const (
	UserKey  contextKey = "userID"
	RoleKey  contextKey = "role"
	EmailKey contextKey = "email"
	UIdKey   contextKey = "uid"
)
