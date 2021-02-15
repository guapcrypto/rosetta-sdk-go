// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/guapcrypto/rosetta-sdk-go/asserter"
	"github.com/guapcrypto/rosetta-sdk-go/types"
)

// A SearchAPIController binds http requests to an api service and writes the service results to the
// http response
type SearchAPIController struct {
	service  SearchAPIServicer
	asserter *asserter.Asserter
}

// NewSearchAPIController creates a default api controller
func NewSearchAPIController(
	s SearchAPIServicer,
	asserter *asserter.Asserter,
) Router {
	return &SearchAPIController{
		service:  s,
		asserter: asserter,
	}
}

// Routes returns all of the api route for the SearchAPIController
func (c *SearchAPIController) Routes() Routes {
	return Routes{
		{
			"SearchTransactions",
			strings.ToUpper("Post"),
			"/search/transactions",
			c.SearchTransactions,
		},
	}
}

// SearchTransactions - [INDEXER] Search for Transactions
func (c *SearchAPIController) SearchTransactions(w http.ResponseWriter, r *http.Request) {
	searchTransactionsRequest := &types.SearchTransactionsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&searchTransactionsRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that SearchTransactionsRequest is correct
	if err := c.asserter.SearchTransactionsRequest(searchTransactionsRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.SearchTransactions(r.Context(), searchTransactionsRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}
