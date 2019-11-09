package configs

//TODO: перевести в yaml
var Genres = [...]string{
	"боевик",
	"триллер",
	"фэнтези",
	"фантастика",
	"вестерн",
	"военный",
	"музыка",
	"драма",
	"криминал",
	"мелодрама",
	"приключения",
	"биография",
	"история",
	"детектив",
	"семейный",
	"комедия",
	"ужасы",
}

var Targets = [...]string{
	"film",
	"user",
	"person",
	"review",
}

var Roles = [...]string{
	"user",
	"admin",
	"actor",
	"director",
	"screenwriter",
	"producer",
	"compositor",
}

var FileTargetMap = map[string]string{
	"person": "persons_data.json",
	"user":   "users_data.json",
	"review": "reviews_data.json",
	"star":   "stars_data.json",
	"film":   "films_data.json",
	"like":   "likes_data.json",
}
