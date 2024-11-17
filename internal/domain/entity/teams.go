package entity

type Teams struct {
	Name      string `csv:"name" json:"name"`
	ShortName string `csv:"short_name" json:"short_name"`
	Slug      string `csv:"slug" json:"slug"`
}
