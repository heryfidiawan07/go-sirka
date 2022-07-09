package request

type DisplayUser struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required"`
}
