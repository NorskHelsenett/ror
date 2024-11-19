package auditlogs

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	auditlogrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/auditlogRepo"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.AuditLog], error) {
	auditLogs, totalCount, err := auditlogrepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error when getting auditlogs by filter from repo: %v", err)
	}

	var results []apicontracts.AuditLog
	for _, v := range auditLogs {
		var auditLog apicontracts.AuditLog
		err := mapping.Map(v, &auditLog)
		if err != nil {
			return nil, fmt.Errorf("could not map from mongotype to apitype: %v", err)
		}
		results = append(results, auditLog)
	}

	paginatedResult := apicontracts.PaginatedResult[apicontracts.AuditLog]{}

	paginatedResult.Data = results
	paginatedResult.DataCount = int64(len(results))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.AuditLog, error) {
	mongoAuditLog, err := auditlogrepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get auditlog from repo: %v", err)
	}

	auditLog := apicontracts.AuditLog{}
	err = mapping.Map(mongoAuditLog, &auditLog)
	if err != nil {
		return nil, fmt.Errorf("could not map from mongotype to apitype: %v", err)
	}

	return &auditLog, nil
}

func GetMetadata(ctx context.Context) (map[string][]string, error) {
	metadata, err := auditlogrepo.GetMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get metadata from auditlogrepo: %v", err)
	}
	return metadata, nil
}
