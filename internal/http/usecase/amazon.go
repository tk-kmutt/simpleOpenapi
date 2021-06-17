package usecase

import (
	"net/http"
	"simpleOpenapi/internal/database/repository"
	"simpleOpenapi/internal/http/gen"
	"time"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type AmazonUsecase struct {
	*gorm.DB
}

func NewAmazon(db *gorm.DB) *AmazonUsecase {
	return &AmazonUsecase{db}
}

func (p *AmazonUsecase) FindAmazons(ctx echo.Context, params gen.FindAmazonsParams) error {
	//TODO implement
	return ctx.NoContent(http.StatusNoContent)
}

func (p *AmazonUsecase) CreateAmazon(ctx echo.Context) error {
	req := new(gen.Amazon)
	err := ctx.Bind(&req)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for NewAmazon")
	}

	// Create
	now := time.Now()
	tx := p.Create(&repository.AmazonItems{
		Asin:      req.Asin,
		Name:      req.Name,
		Maker:     req.Maker,
		Price:     req.Price,
		Url:       req.Url,
		Comment:   req.Comment,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if tx.Error != nil {
		return sendError(ctx, http.StatusBadRequest, tx.Error.Error())
	}
	return ctx.JSON(http.StatusCreated, req)
}

func (p *AmazonUsecase) FindAmazonByKey(ctx echo.Context, key gen.Key) error {
	m := new(repository.AmazonItems)
	if tx := p.First(m, "asin = ?", key); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}
	res := &gen.Amazon{
		Name:    m.Name,
		Maker:   m.Maker,
		Price:   m.Price,
		Comment: m.Comment,
		Url:     m.Url,
		Asin:    m.Asin,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (p *AmazonUsecase) DeleteAmazon(ctx echo.Context, key gen.Key) error {
	m := new(repository.AmazonItems)
	if tx := p.Unscoped().First(m, "asin = ?", key); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}
	if tx := p.Unscoped().Delete(m, "asin = ?", key); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (p *AmazonUsecase) ActiveAmazon(ctx echo.Context, key gen.Key) error {
	m := new(repository.AmazonItems)
	if tx := p.Unscoped().Model(m).Where("asin = ?", key).Update("is_delete", repository.NOT_DELETE); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (p *AmazonUsecase) InactiveAmazon(ctx echo.Context, key gen.Key) error {
	m := new(repository.AmazonItems)

	m.IsDelete = repository.DELETE
	if tx := p.Model(m).Where("asin = ?", key).Updates(m); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (p *AmazonUsecase) UpdateAmazon(ctx echo.Context, key gen.Key) error {
	req := new(gen.AmazonUpdate)
	err := ctx.Bind(&req)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for NewAmazon")
	}

	now := time.Now()
	m := &repository.AmazonItems{
		Name:      req.Name,
		Maker:     req.Maker,
		Price:     req.Price,
		Comment:   req.Comment,
		Url:       req.Url,
		UpdatedAt: now,
	}
	tx := p.Model(m).
		Where("asin = ?", key).
		Updates(m)
	if tx.Error != nil {
		return sendError(ctx, http.StatusBadRequest, tx.Error.Error())
	}
	res := &gen.Amazon{
		Name:    m.Name,
		Maker:   m.Maker,
		Price:   m.Price,
		Comment: m.Comment,
		Url:     m.Url,
		Asin:    key.String(),
	}

	return ctx.JSON(http.StatusOK, res)
}

func (p *AmazonUsecase) PatchAmazon(ctx echo.Context, key gen.Key) error {
	req := new(gen.AmazonPatch)
	err := ctx.Bind(&req)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for NewAmazon")
	}

	m := new(repository.AmazonItems)
	if tx := p.Model(m).Find(m, "asin = ?", key); tx.Error != nil {
		return sendError(ctx, http.StatusBadRequest, tx.Error.Error())
	}

	if req.Name != nil {
		m.Name = *req.Name
	}
	if req.Maker != nil {
		m.Maker = *req.Maker
	}
	if req.Price != nil {
		m.Price = *req.Price
	}
	if req.Comment != nil {
		m.Comment = req.Comment
	}
	if req.Url != nil {
		m.Url = *req.Url
	}

	tx := p.Model(m).Where("asin = ?", key).Updates(*m)
	if tx.Error != nil {
		return sendError(ctx, http.StatusBadRequest, tx.Error.Error())
	}

	return ctx.JSON(http.StatusOK, &gen.Amazon{
		Name:    m.Name,
		Maker:   m.Maker,
		Price:   m.Price,
		Comment: m.Comment,
		Url:     m.Url,
		Asin:    m.Asin,
	})
}
