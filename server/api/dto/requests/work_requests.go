package requests

type WorkRequest struct {
	Address string `json:"address" binding:"required"`
	PortMin int    `json:"port_min" binding:"required,min=1"`
	PortMax int    `json:"port_max" binding:"required,max=65535"`
}
