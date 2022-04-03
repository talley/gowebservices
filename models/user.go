package models

type User struct {
	ID int32
	Name string
	Location string
}
 var(
 	users []*User
 	NextID=1
 )