package entities
import 	(
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
)
type Movie struct {
	Id string `validate:required`
	Title string `validate:"required"`
	Genre [] string `validate:"required,min=1"`
	Description string `validate:"required"`
	Actors [] string `validate:"required,min=1"`
	Rating float32 `validate:"required,max=5"`
}

func (movie *Movie) SetId () {
	movie.Id = uuid.New().String()
}

func (movie *Movie) ValidateMovie() string {
	validate := validator.New()
	err := validate.Struct(movie)
	if err != nil {
		return err.Error()
	}
	return "valid"
}

