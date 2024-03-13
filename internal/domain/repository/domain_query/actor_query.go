package domainQuery

type ActorRepositoryQuery struct {
	CurrentPage    int
	PageCount      int
	WithConnection []string
}

func NewActorRepositoryQuery() *FilmRepositoryQuery {
	return &FilmRepositoryQuery{
		CurrentPage:    1,
		PageCount:      10,
		WithConnection: make([]string, 0, 4),
	}
}
