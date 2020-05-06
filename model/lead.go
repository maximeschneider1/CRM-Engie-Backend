package model



// Lead struct gather infos we have about the lead
type Lead struct {
	LeadID		int `json:"lead_id"`
	Name           string `json:"nom"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Score          int    `json:"score"`
	FirstContact string
	City string

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

type LeadTags struct {
	LeadID int
	TagID int
	TagContent string
	TagIcon string
}