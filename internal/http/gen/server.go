// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /amazons)
	FindAmazons(ctx echo.Context, params FindAmazonsParams) error

	// (POST /amazons)
	CreateAmazon(ctx echo.Context) error

	// (PATCH /amazons/active/{key})
	ActiveAmazon(ctx echo.Context, key Key) error

	// (PATCH /amazons/inactive/{key})
	InactiveAmazon(ctx echo.Context, key Key) error

	// (DELETE /amazons/{key})
	DeleteAmazon(ctx echo.Context, key Key) error

	// (GET /amazons/{key})
	FindAmazonByKey(ctx echo.Context, key Key) error

	// (PATCH /amazons/{key})
	PatchAmazon(ctx echo.Context, key Key) error

	// (PUT /amazons/{key})
	UpdateAmazon(ctx echo.Context, key Key) error

	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error

	// (POST /pets)
	AddPet(ctx echo.Context) error

	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id ID) error

	// (GET /pets/{id})
	FindPetById(ctx echo.Context, id ID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindAmazons converts echo context to params.
func (w *ServerInterfaceWrapper) FindAmazons(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindAmazonsParams
	// ------------- Optional query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, false, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindAmazons(ctx, params)
	return err
}

// CreateAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAmazon(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateAmazon(ctx)
	return err
}

// ActiveAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) ActiveAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ActiveAmazon(ctx, key)
	return err
}

// InactiveAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) InactiveAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.InactiveAmazon(ctx, key)
	return err
}

// DeleteAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteAmazon(ctx, key)
	return err
}

// FindAmazonByKey converts echo context to params.
func (w *ServerInterfaceWrapper) FindAmazonByKey(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindAmazonByKey(ctx, key)
	return err
}

// PatchAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) PatchAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchAmazon(ctx, key)
	return err
}

// UpdateAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateAmazon(ctx, key)
	return err
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, false, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id ID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetById converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id ID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetById(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/amazons", wrapper.FindAmazons)
	router.POST(baseURL+"/amazons", wrapper.CreateAmazon)
	router.PATCH(baseURL+"/amazons/active/:key", wrapper.ActiveAmazon)
	router.PATCH(baseURL+"/amazons/inactive/:key", wrapper.InactiveAmazon)
	router.DELETE(baseURL+"/amazons/:key", wrapper.DeleteAmazon)
	router.GET(baseURL+"/amazons/:key", wrapper.FindAmazonByKey)
	router.PATCH(baseURL+"/amazons/:key", wrapper.PatchAmazon)
	router.PUT(baseURL+"/amazons/:key", wrapper.UpdateAmazon)
	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
