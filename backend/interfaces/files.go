package interfaces

type File struct {
	ID        string `json:"id"`
	FileName  string `json:"fileName"`
	Url       string `json:"url"`
	UpdatedAt int64  `json:"updatedAt"`
}
