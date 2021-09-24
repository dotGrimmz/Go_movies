package entities
import 	"github.com/google/uuid"

type Movie struct {
	Id string
	Title string
	Genre [] string
	Description string
	Actors [] string
	Rating float32
}

func (movie *Movie) SetId () {
	movie.Id = uuid.New().String()
}