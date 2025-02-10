package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"fmt"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/connection"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

type ReportsRepository interface {
	GetReports(ctx context.Context, id string) (*model.HarvestReports, error)
	UpsertReports(ctx context.Context, dataSource model.HarvestReport) error
}

type ReportsRepositoryImpl struct {
	collection *mongo.Collection
}

var reportsRepository *ReportsRepositoryImpl

func InitReportsRepository() *ReportsRepositoryImpl {
	if reportsRepository == nil {
		reportsRepository = &ReportsRepositoryImpl{collection: connection.ReportsCollection()}
	}
	return reportsRepository
}


func (r *ReportsRepositoryImpl) GetReports(ctx context.Context, id string) (*model.HarvestReports, error) {
	if !isValidID(id) {
		return nil, fmt.Errorf("invalid id format")
	}
	
	filter := bson.D{{Key: "id", Value: id}}
	singleResult := r.collection.FindOne(ctx, filter)
	if err := singleResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	bytes, err := singleResult.Raw()
	if err != nil {
		return nil, err
	}

	var reports model.HarvestReports
	unmarshalError := bson.Unmarshal(bytes, &reports)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &reports, nil
}

func (r *ReportsRepositoryImpl) UpsertReports(ctx context.Context, report model.HarvestReport) error {
	if !isValidID(report.ID) {
		return fmt.Errorf("invalid id format")
	}
	
	filter := bson.D{{Key: "id", Value: report.ID}}
	singleResult := r.collection.FindOne(ctx, filter)
	if err := singleResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return r.createReports(ctx, report)
		} 
		return err
	}
	bytes, err := singleResult.Raw()
	if err != nil {
		return err
	} else {
		return r.updateReports(ctx, bytes, report)
	}
}

func (r *ReportsRepositoryImpl) createReports(ctx context.Context, report model.HarvestReport) error {
	reportsMap := make(map[string]model.HarvestReport)
	reportsMap[string(report.DataType)] = report
	reports := model.HarvestReports{
		ID:      report.ID,
		Reports: reportsMap,
	}
	_, err := r.collection.InsertOne(ctx, reports, nil)

	return err
}

func (r *ReportsRepositoryImpl) updateReports(ctx context.Context, dbReports bson.Raw, newReport model.HarvestReport) error {
	var updated model.HarvestReports
	err := bson.Unmarshal(dbReports, &updated)
	if err != nil {
		return err
	}

	updated.Reports[string(newReport.DataType)] = newReport

	filter := bson.D{{Key: "id", Value: newReport.ID}}
	result := r.collection.FindOneAndReplace(ctx, filter, updated, nil)

	return result.Err()
}
