package students

var params = map[string]string{
	"first_name": "first_name",
	"last_name":  "last_name",
	"class":      "class",
}
var allowedSortFields = map[string]bool{
	"first_name": true,
	"last_name":  true,
	"class":      true,
	"created_at": true,
}

var allowedFields = map[string]bool{
	"first_name": true,
	"last_name":  true,
	"email":      true,
	"class":      true,
}
