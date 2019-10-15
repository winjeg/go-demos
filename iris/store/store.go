package store

type Site struct {
	Id      *int64  `json:"id"`
	Author  *string `json:"author"`
	Dir     *string `json:"dir"`
	Mapping *string `json:"mapping"`
	Created *int64  `json:"created"`
	Updated *int64  `json:"updated"`
}

type SiteStore interface {
	Store(author, dir, path string) error
	Exists(...string) bool
	GetAll() []Site
}
