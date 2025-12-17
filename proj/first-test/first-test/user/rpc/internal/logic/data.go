package logic

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "炭治郎",
		Phone: "13000000001",
	},
	"2": {
		Id:    "2",
		Name:  "香奈乎",
		Phone: "15000000001",
	},
}
