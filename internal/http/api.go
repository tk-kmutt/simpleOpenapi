package api

import (
	"simpleOpenapi/internal/http/gen"
	"simpleOpenapi/internal/http/usecase"

	"github.com/labstack/echo/v4"
)

type Api struct {
	pet *usecase.Pet
}

func NewApi() *Api {
	return &Api{pet: usecase.NewPet()}
}

var _ gen.ServerInterface = (*Api)(nil)

func (p *Api) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	return p.pet.FindPets(ctx, params)
}
func (p *Api) AddPet(ctx echo.Context) error {
	return p.pet.AddPet(ctx)
}

func (p *Api) FindPetById(ctx echo.Context, id int64) error {
	return p.pet.FindPetById(ctx, id)
}

func (p *Api) DeletePet(ctx echo.Context, id int64) error {
	return p.pet.DeletePet(ctx, id)
}
