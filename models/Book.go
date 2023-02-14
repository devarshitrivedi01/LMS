package models

type Book struct{
	BookID int
	Name string
	Author string
	Issued bool
	ISBN string
	AssocUserId int
}