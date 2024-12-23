package database

type Client struct {
	Id       uint32
	Email    string
	Password string
	Name     string
	Age      uint32
}

// to export a variable always use uppercase
var Clients = map[string]Client{
	"1": {
		Id:       1,
		Email:    "michael.smith@example.com",
		Password: "michael2024!",
		Name:     "Michael Smith",
		Age:      35,
	},
	"2": {
		Id:       2,
		Email:    "sarah.johnson@example.com",
		Password: "s4rahjohn$on",
		Name:     "Sarah Johnson",
		Age:      28,
	},
	"3": {
		Id:       3,
		Email:    "william.brown@example.com",
		Password: "Willi@mbrown123",
		Name:     "William Brown",
		Age:      42,
	},
	"4": {
		Id:       4,
		Email:    "emily.davis@example.com",
		Password: "EmilyD@vis2023",
		Name:     "Emily Davis",
		Age:      30,
	},
	"5": {
		Id:       5,
		Email:    "james.miller@example.com",
		Password: "J@m3smiller",
		Name:     "James Miller",
		Age:      50,
	},
	"6": {
		Id:       6,
		Email:    "olivia.wilson@example.com",
		Password: "Olivia2024!",
		Name:     "Olivia Wilson",
		Age:      26,
	},
	"7": {
		Id:       7,
		Email:    "daniel.moore@example.com",
		Password: "D@nielmoore42",
		Name:     "Daniel Moore",
		Age:      38,
	},
	"8": {
		Id:       8,
		Email:    "isabella.taylor@example.com",
		Password: "Isab3llaT@ylor",
		Name:     "Isabella Taylor",
		Age:      24,
	},
	"9": {
		Id:       9,
		Email:    "jack.anderson@example.com",
		Password: "JackAnd3rson!",
		Name:     "Jack Anderson",
		Age:      33,
	},
	"10": {
		Id:       10,
		Email:    "sophia.thomas@example.com",
		Password: "SophiaT2024$",
		Name:     "Sophia Thomas",
		Age:      29,
	},
}
