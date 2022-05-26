package repositories

type UrlDictRepository interface {
	Find() error
	Create() error
	Update() error
	Delete() error
}
