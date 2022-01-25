package response

type MessageDescription struct {
	Message Message `json:"message"`
}
type Message struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
