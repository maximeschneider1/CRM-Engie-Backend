package model

type HomeInfo struct {
	TotalClients int
	Todo int
	NewLeads int
	NewDocuments int
	PotentialValue int
}

type Todo struct {
	Id       int
	ClientID int
	Name     string
	Phone    string
	Category string
	Motif    string
}