package main

import (
	"ecaterminal/internal/domain/appcontext"
	"ecaterminal/internal/domain/ecaterminal"
	"ecaterminal/internal/infrastructure/database"
	"ecaterminal/internal/infrastructure/worker"
	"fmt"

	"github.com/go-skynet/go-llama.cpp"
)

var l *llama.LLama

func main() {

	ctx := appcontext.NewBackground()

	ecaterminal, err := setupecaterminal()
	if err != nil {
		fmt.Println("Failed to setup Ecaterminal")
		panic(err)
	}
	l = worker.LoadAiModel()
	ecaterminal.Screen(ctx, l)
}

func setupecaterminal() (ecaterminal.UseCases, error) {

	dynamodb, err := setupDynamoDB()
	if err != nil {
		return nil, err
	}

	memdbInput := &ecaterminal.Input{
		Repository: dynamodb,
	}
	ecatrom2000UseCases := ecaterminal.New(memdbInput)
	return ecatrom2000UseCases, nil
}

func setupDynamoDB() (ecaterminal.Repository, error) {
	// env := environment.GetInstance()
	// if !env.DEFAULT_PERSISTENT {
	// 	return database.NewMemoryDatabase(), nil
	// }

	// awsRegion := env.AWS_REGION
	// awsEndpoint := env.DYNAMO_AWS_ENDPOINT
	// table := env.DYNAMO_TABLE_NAME
	// cfg, err := aws.EndpointResolverWithOptionsFunc(awsEndpoint, awsRegion)
	// if err != nil {
	// 	return nil, err
	// }
	// return database.NewDynamoDB(cfg, table), nil

	return database.NewMemoryDatabase(), nil
}
