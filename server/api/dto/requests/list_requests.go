package requests

type CreateListRequest struct {
	Name string `json:"name"`
}

type UpdateListRequest struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
