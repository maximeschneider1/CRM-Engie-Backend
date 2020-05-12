package model

type TagClient struct {
	TagID int
	Name string
}
type LeadTags struct {
	LeadID int
	TagID int
	TagContent string
	TagIcon string
}
