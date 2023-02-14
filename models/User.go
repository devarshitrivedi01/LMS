package models

type User struct{
	UserID int
	Name string
	BirthDate string
	CardInitialTime string
	CardEndTime string
	BookIssued int
	BookCapacity int
	Role Role
}