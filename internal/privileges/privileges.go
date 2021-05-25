package privileges

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

// type Priveleges interface {
// 	CreateUser()
// 	UpdateUser()
// 	GetAllUsers()
// 	GetUserById()
// }

type PrivelegesList = interface{}{
	"CreateUser":  true,
	"DeleteUser":  true,
	"GetAllUsers": true,
}

type PrivelegesStruct struct {
	Priveleges PrivelegesList
}

type CompanyDirector struct {
	AllUsers
	AllPrivelegs
	AllGroupsCreated
	AllModels
}

type Groups struct {
	Name string
	Director.id string
	PrivelegesList map[string]bool
	Company.id
	user.ids
}

type Priveleges struct {
	Id int
	PrivName string
	State bool
}

