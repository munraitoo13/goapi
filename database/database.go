package database

type Client struct {
	Id    string
	Email string
	Token string
	Name  string
	Age   uint32
}

// to export a variable always use uppercase
var Clients = map[string]Client{
	"1": {
		Id:    "1",
		Email: "michael.smith@example.com",
		Token: "michael2024!",
		Name:  "Michael Smith",
		Age:   35,
	},
	"2": {
		Id:    "2",
		Email: "sarah.johnson@example.com",
		Token: "s4rahjohn$on",
		Name:  "Sarah Johnson",
		Age:   28,
	},
}
