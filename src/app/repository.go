package app

type LinkRepository struct {
	links []Link
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{}
}

func (repo *LinkRepository) Add(l *Link) {
	repo.links = append(repo.links, *l)
}
