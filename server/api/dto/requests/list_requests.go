package requests

type CreateListReq struct {
	Name string `json:"name" binding:"required,min=6"`
}
