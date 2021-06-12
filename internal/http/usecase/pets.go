package usecase

import (
	"fmt"
	"net/http"
	"simpleOpenapi/internal/http/gen"
	"sort"
	"sync"

	"github.com/labstack/echo/v4"
)

type PetUsecase struct {
	Pets   map[int64]gen.PetResponse
	NextId int64
	Lock   sync.Mutex
}

func NewPet() *PetUsecase {
	return &PetUsecase{
		Pets:   make(map[int64]gen.PetResponse),
		NextId: 1000,
	}
}

func (p *PetUsecase) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var result []gen.PetResponse

	// 順番通りに並べる(デフォルトAsc)
	isAsc := true
	if params.Order != nil && *params.Order == "desc" {
		isAsc = false
	}
	fmt.Println(isAsc)
	keys := make([]int64, 0, len(p.Pets))
	for k := range p.Pets {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if isAsc {
			return keys[int64(i)] < keys[int64(j)]
		}
		return keys[int64(j)] < keys[int64(i)]
	})

	fmt.Println(keys)

	for _, key := range keys {
		pet := p.Pets[key]
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

func (p *PetUsecase) AddPet(ctx echo.Context) error {
	req := new(gen.Pet)
	err := ctx.Bind(&req)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}
	// We now have a np, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	res := gen.PetResponse{
		Id: p.NextId,
		Pet: gen.Pet{
			Name: req.Name,
			Tag:  req.Tag,
		},
	}
	p.NextId = p.NextId + 1

	// Insert into map
	p.Pets[res.Id] = res

	// Now, we have to return the NewPet
	err = ctx.JSON(http.StatusCreated, res)
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

func (p *PetUsecase) FindPetById(ctx echo.Context, id int64) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[id]
	if !found {
		return sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", id))
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (p *PetUsecase) DeletePet(ctx echo.Context, id int64) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, found := p.Pets[id]
	if !found {
		return sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", id))
	}
	delete(p.Pets, id)
	return ctx.NoContent(http.StatusNoContent)
}
