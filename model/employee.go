package model

type HomeInfo struct {
	TotalClients int
	Todo int
	NewLeads int
	NewDocuments int
	PotentialValue int
}

type Todo struct {
	Name string
	Id int
	Telephone string
	Category string
	Motif string
}