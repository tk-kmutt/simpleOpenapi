package api

import (
	"simpleOpenapi/internal/http/gen"
	"simpleOpenapi/internal/http/usecase"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Api struct {
	pet    *usecase.PetUsecase
	amazon *usecase.AmazonUsecase
}

func (p *Api) PatchAmazon(ctx echo.Context, key gen.Key) error {
	return p.amazon.PatchAmazon(ctx, key)
}

func (p *Api) UpdateAmazon(ctx echo.Context, key gen.Key) error {
	return p.amazon.UpdateAmazon(ctx, key)
}

func (p *Api) ActiveAmazon(ctx echo.Context, key gen.Key) error {
	return p.amazon.ActiveAmazon(ctx, key)
}

func (p *Api) InactiveAmazon(ctx echo.Context, key gen.Key) error {
	return p.amazon.InactiveAmazon(ctx, key)
}

func (p *Api) FindAmazons(ctx echo.Context, params gen.FindAmazonsParams) error {
	panic("implement me")
}

func (p *Api) CreateAmazon(ctx echo.Context) error {
	return p.amazon.CreateAmazon(ctx)
}

func (p *Api) DeleteAmazon(ctx echo.Context, key gen.Key) error {
	return p.amazon.DeleteAmazon(ctx, key)
}

func (p *Api) FindAmazonByKey(ctx echo.Context, key gen.Key) error {
	return p.amazon.FindAmazonByKey(ctx, key)
}

func NewApi(db *gorm.DB) *Api {
	return &Api{
		pet:    usecase.NewPet(),
		amazon: usecase.NewAmazon(db),
	}
}

func (p *Api) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	return p.pet.FindPets(ctx, params)
}
func (p *Api) AddPet(ctx echo.Context) error {
	return p.pet.AddPet(ctx)
}

func (p *Api) FindPetById(ctx echo.Context, id gen.ID) error {
	return p.pet.FindPetById(ctx, id.Int64())
}

func (p *Api) DeletePet(ctx echo.Context, id gen.ID) error {
	return p.pet.DeletePet(ctx, id.Int64())
}

var _ gen.ServerInterface = (*Api)(nil)
