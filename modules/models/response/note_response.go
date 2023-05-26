package response

// Models for request, and response
type NoteResponse struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
