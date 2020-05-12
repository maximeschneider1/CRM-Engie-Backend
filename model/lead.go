package model

type Lead struct {
	LeadID		int `json:"lead_id"`
	Name           string `json:"name"`
	Email string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Score          int    `json:"score"`
	FirstContact string
	City string
	Step int
	StepConverted string
	ContentDownloaded int
	TimeSpent int
	OpenedEmails int
	Profitability int
	WeeksSinceInactive int
}

type LeadHistory struct {
	LeadID int
	Type string
	Icon string
	Date string
	Comment string
}
