package usecase

import (
	"fmt"
	"net/http"
	"simpleOpenapi/internal/http/gen"
	"sync"

	"github.com/labstack/echo/v4"
)

type Pet struct {
	Pets   map[int64]gen.Pet
	NextId int64
	Lock   sync.Mutex
}

func NewPet() *Pet {
	return &Pet{
		Pets:   make(map[int64]gen.Pet),
		NextId: 1000,
	}
}

func (p *Pet) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var result []gen.Pet

	for _, pet := range p.Pets {
		if params.Tags != nil {
			// If we have tags,  filter pets by tag
			for _, t := range *params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			// Add all pets if we're not filtering
			result = append(result, pet)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}
	return ctx.JSON(http.StatusOK, result)
}

func (p *Pet) AddPet(ctx echo.Context) error {
	// We expect a NewPet object in the request body.
	var newPet gen.NewPet
	err := ctx.Bind(&newPet)
	if err != nil {
		return sendPetstoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}
	// We now have a pet, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	var pet gen.Pet
	pet.Name = newPet.Name
	pet.Tag = newPet.Tag
	pet.Id = p.NextId
	p.NextId = p.NextId + 1

	// Insert into map
	p.Pets[pet.Id] = pet

	// Now, we have to return the NewPet
	err = ctx.JSON(http.StatusCreated, pet)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}

	// Return no error. This refers to the handler. Even if we return an HTTP
	// error, but everything else is working properly, tell Echo that we serviced
	// the error. We should only return errors from Echo handlers if the actual
	// servicing of the error on the infrastructure level failed. Returning an
	// HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
	// return nil.
	return nil
}

func (p *Pet) FindPetById(ctx echo.Context, id int64) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[id]
	if !found {
		return sendPetstoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", id))
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (p *Pet) DeletePet(ctx echo.Context, id int64) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, found := p.Pets[id]
	if !found {
		return sendPetstoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", id))
	}
	delete(p.Pets, id)
	return ctx.NoContent(http.StatusNoContent)
}
