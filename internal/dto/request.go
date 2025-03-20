package dto

type ListQuery struct {
	Page     int  `form:"page"`
	Limit    int  `form:"limit"`
	Simplify bool `form:"simplify"`
}
