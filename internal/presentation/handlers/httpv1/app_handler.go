package httpv1

type (
	AppHandler struct {
		AuthHandler
		FilmHandler
		ActorHandler
	}
)

var instance *AppHandler = nil

func NewAppHandler() *AppHandler {
	if instance != nil {
		return instance
	}

	instance = &AppHandler{
		AuthHandler:  NewAuthHandler(),
		FilmHandler:  NewFilmHandler(),
		ActorHandler: NewActorHandler(),
	}

	return instance
}
