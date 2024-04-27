package model

type PostResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Link        string `json:"link"`
	Category    int    `json:"category"`
}
