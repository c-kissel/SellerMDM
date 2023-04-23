// Package specs provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package specs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/chi/v5"
)

const (
	KeycloakScopes = "Keycloak.Scopes"
)

// EditSellerRequest defines model for EditSellerRequest.
type EditSellerRequest struct {
	City *string             `json:"city,omitempty"`
	Id   *openapi_types.UUID `json:"id,omitempty"`
	Inn  *string             `json:"inn,omitempty"`
	Logo *string             `json:"logo,omitempty"`
	Memo *string             `json:"memo,omitempty"`
	Name *string             `json:"name,omitempty"`
	Ogrn *string             `json:"ogrn,omitempty"`
	Site *string             `json:"site,omitempty"`
	Yml  *string             `json:"yml,omitempty"`
}

// NewSellerRequest Seller Master Data
type NewSellerRequest struct {
	City *string `json:"city,omitempty"`
	Inn  *string `json:"inn,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Name *string `json:"name,omitempty"`
	Ogrn *string `json:"ogrn,omitempty"`
	Site *string `json:"site,omitempty"`
	Yml  *string `json:"yml,omitempty"`
}

// SellerResponse defines model for SellerResponse.
type SellerResponse struct {
	City    *string             `json:"city,omitempty"`
	Created *string             `json:"created,omitempty"`
	Id      *openapi_types.UUID `json:"id,omitempty"`
	Inn     *string             `json:"inn,omitempty"`
	Logo    *string             `json:"logo,omitempty"`
	Memo    *string             `json:"memo,omitempty"`
	Name    *string             `json:"name,omitempty"`
	Ogrn    *string             `json:"ogrn,omitempty"`
	Site    *string             `json:"site,omitempty"`
	Updated *string             `json:"updated,omitempty"`
	Yml     *string             `json:"yml,omitempty"`
}

// GetSellersByNameParams defines parameters for GetSellersByName.
type GetSellersByNameParams struct {
	// Name name of seller
	Name string `form:"name" json:"name"`
}

// PostNewSellerJSONRequestBody defines body for PostNewSeller for application/json ContentType.
type PostNewSellerJSONRequestBody = NewSellerRequest

// PutSellerJSONRequestBody defines body for PutSeller for application/json ContentType.
type PutSellerJSONRequestBody = EditSellerRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /v1/sellers)
	PostNewSeller(w http.ResponseWriter, r *http.Request)

	// (GET /v1/sellers/all)
	GetSellersAll(w http.ResponseWriter, r *http.Request)

	// (DELETE /v1/sellers/id/{id})
	DeleteSeller(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (GET /v1/sellers/id/{id})
	GetSeller(w http.ResponseWriter, r *http.Request, id string)

	// (PUT /v1/sellers/id/{id})
	PutSeller(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (GET /v1/sellers/search)
	GetSellersByName(w http.ResponseWriter, r *http.Request, params GetSellersByNameParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostNewSeller operation middleware
func (siw *ServerInterfaceWrapper) PostNewSeller(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, KeycloakScopes, []string{"sellers:write"})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNewSeller(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSellersAll operation middleware
func (siw *ServerInterfaceWrapper) GetSellersAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, KeycloakScopes, []string{"sellers:read"})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSellersAll(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteSeller operation middleware
func (siw *ServerInterfaceWrapper) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, KeycloakScopes, []string{"sellers:delete"})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteSeller(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSeller operation middleware
func (siw *ServerInterfaceWrapper) GetSeller(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, KeycloakScopes, []string{"sellers:read"})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSeller(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutSeller operation middleware
func (siw *ServerInterfaceWrapper) PutSeller(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, KeycloakScopes, []string{"sellers:write"})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutSeller(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSellersByName operation middleware
func (siw *ServerInterfaceWrapper) GetSellersByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, KeycloakScopes, []string{"sellers:read"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSellersByNameParams

	// ------------- Required query parameter "name" -------------

	if paramValue := r.URL.Query().Get("name"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "name"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "name", r.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSellersByName(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/sellers", wrapper.PostNewSeller)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/sellers/all", wrapper.GetSellersAll)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/sellers/id/{id}", wrapper.DeleteSeller)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/sellers/id/{id}", wrapper.GetSeller)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/v1/sellers/id/{id}", wrapper.PutSeller)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/sellers/search", wrapper.GetSellersByName)
	})

	return r
}
