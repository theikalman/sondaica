package main

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/spf13/viper"
	"github.com/theikalman/sondaica/internal/config"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var conf config.Config
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unable to unmarshal configuration file: %w", err))
	}

	fmt.Println("successfully unmarshal config")
	fmt.Printf("%+v\n\n", conf)

	docModel := mustReadOpenAPISpecFile(conf.RestAPI.OpenAPISpec)

	mustGenerateServerCode(docModel)

	// TODO Generate client code
}

func mustReadOpenAPISpecFile(file string) *libopenapi.DocumentModel[v3.Document] {
	petstore, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("unable to read open-api-specs file: %w", err))
	}

	document, err := libopenapi.NewDocument(petstore)
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	docModel, errors := document.BuildV3Model()
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}

		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
	}

	return docModel
}

func mustGenerateServerCode(docModel *libopenapi.DocumentModel[v3.Document]) {
	// TODO Generate server code

	for pathItem := docModel.Model.Paths.PathItems.First(); pathItem != nil; pathItem = pathItem.Next() {
		pathItemValue := pathItem.Value()

		fmt.Printf("pathItem Key: %+v\n", pathItem.Key())
		if pathItemValue.Put != nil {
			fmt.Printf("pathItem PUT operation Id: %+v\n", pathItemValue.Put.OperationId)
		}

		if pathItemValue.Get != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Put != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Post != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Delete != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Options != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Head != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Patch != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		if pathItemValue.Trace != nil {
			// TODO Check if the handler name and operation id mapping is exists!
		}

		fmt.Println("")

		// TODO Generate list of handlers we need to create
	}
}
