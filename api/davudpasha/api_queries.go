package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

// QueriesApi service type.
type QueriesApi common.Service

// SearchQueriesOptionalParameters holds optional parameters for SearchQueries.
type SearchQueriesOptionalParameters struct {
	Body *QueriesSearchRequest
}
