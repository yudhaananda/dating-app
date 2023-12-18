package filter

type PremiumFeatureFilter struct {
	Id        int       `db:"id" json:"id" form:"id"`
    Name string `db:"name" json:"name" form:"name"`
    Flag string `db:"flag" json:"flag" form:"flag"`
    
}
