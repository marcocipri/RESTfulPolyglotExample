/*
 * Swagger Petstore
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	RESTfulPolyglotEampleGoServer "github.com/marcocipri/RESTfulPolyglotExample/go/servermock"
)

func main() {

	t := reflect.TypeOf(MyPetsAPIService{}) // the Singer type
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&MyPetsAPIService{}) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}

	log.Printf("Server started")

	DefaultAPIService := &MyPetsAPIService{}
	DefaultAPIController := RESTfulPolyglotEampleGoServer.NewDefaultApiController(DefaultAPIService)

	router := RESTfulPolyglotEampleGoServer.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}

type MyPetsAPIService struct {
	RESTfulPolyglotEampleGoServer.DefaultApiService
}

// FindPets -
func (s *MyPetsAPIService) FindPets(ctx context.Context, tags []string, limit int32) (RESTfulPolyglotEampleGoServer.ImplResponse, error) {

	cat := RESTfulPolyglotEampleGoServer.Pet{Name: "cat", Tag: "felis", Id: 1}
	dog := RESTfulPolyglotEampleGoServer.Pet{Name: "dog", Tag: "canis", Id: 2}

	pets := []RESTfulPolyglotEampleGoServer.Pet{cat, dog}
	return RESTfulPolyglotEampleGoServer.Response(http.StatusOK, pets), nil

}
