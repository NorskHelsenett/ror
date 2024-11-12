package pricesservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"
	pricesRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/pricesRepo"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

func GetAll(ctx context.Context) (*[]apicontracts.Price, error) {
	prices, err := pricesRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("could not get prices")
	}

	return prices, nil
}

func GetByProperty(ctx context.Context, property string, propertyValue string) (*[]apicontracts.Price, error) {
	prices, err := pricesRepo.GetByProperty(ctx, property, propertyValue)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.Price, error) {
	price, err := pricesRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return price, nil
}

func Create(ctx context.Context, priceInput *apicontracts.Price) (*apicontracts.Price, error) {
	exists, err := pricesRepo.FindOne(ctx, "machineclass", priceInput.MachineClass)
	if err != nil {
		return exists, fmt.Errorf("could not check if price exists: %v", err)
	}

	if exists != nil {
		return exists, fmt.Errorf("price already exists")
	}

	var mappedInput mongoTypes.MongoPrice
	err = mapping.Map(priceInput, &mappedInput)
	if err != nil {
		return nil, fmt.Errorf("could not map price from apitype to mongotype: %v", err)
	}

	createdPrice, err := pricesRepo.Create(ctx, &mappedInput)
	if err != nil {
		return nil, fmt.Errorf("could not create price: %v", err)
	}

	var mappedResult apicontracts.Price
	err = mapping.Map(createdPrice, &mappedResult)
	if err != nil {
		return nil, fmt.Errorf("could not map price from mongotype to apitype: %v", err)
	}

	return &mappedResult, nil
}

func Update(ctx context.Context, priceId string, priceInput *apicontracts.Price) (*apicontracts.Price, *apicontracts.Price, error) {
	var mongoPrice mongoTypes.MongoPrice
	err := mapping.Map(priceInput, &mongoPrice)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map from apitype to mongotype: %v", err)
	}

	newObject, oldObject, err := pricesRepo.Update(ctx, priceId, mongoPrice)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update price: %v", err)
	}

	var mappedNewObject apicontracts.Price
	err = mapping.Map(newObject, &mappedNewObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map updated price from apitype to mongotype: %v", err)
	}

	var mappedOldObject apicontracts.Price
	err = mapping.Map(oldObject, &mappedOldObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map original price from apitype to mongotype: %v", err)
	}

	return &mappedNewObject, &mappedOldObject, nil
}

func Delete(ctx context.Context, priceId string) (bool, *apicontracts.Price, error) {
	deleted, deletedPrice, err := pricesRepo.Delete(ctx, priceId)
	if err != nil {
		return false, nil, fmt.Errorf("could not delete price: %v", err)
	}

	var mappedDeletedPrice apicontracts.Price
	err = mapping.Map(deletedPrice, &mappedDeletedPrice)
	if err != nil {
		return false, nil, fmt.Errorf("could not map deleted price from mongotype to apitype: %v", err)
	}

	return deleted, &mappedDeletedPrice, nil
}
