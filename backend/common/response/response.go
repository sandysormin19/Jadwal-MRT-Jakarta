package response
type APIResponse struct {
	Success bool        `json:"success"`
	Messege string      `json:"messege"`
	Data    interface{} `json:"data"`
}