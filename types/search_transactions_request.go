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

package types

// SearchTransactionsRequest SearchTransactionsRequest is used to search for transactions matching a
// set of provided conditions in canonical blocks.
type SearchTransactionsRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	Operator          *Operator          `json:"operator,omitempty"`
	// max_block is the largest block index to consider when searching for transactions. If this
	// field is not populated, the current block is considered the max_block. If you do not specify
	// a max_block, it is possible a newly synced block will interfere with paginated transaction
	// queries (as the offset could become invalid with newly added rows).
	MaxBlock *int64 `json:"max_block,omitempty"`
	// offset is the offset into the query result to start returning transactions. If any search
	// conditions are changed, the query offset will change and you must restart your search
	// iteration.
	Offset *int64 `json:"offset,omitempty"`
	// limit is the maximum number of transactions to return in one call. The implementation may
	// return <= limit transactions.
	Limit                 *int64                 `json:"limit,omitempty"`
	TransactionIdentifier *TransactionIdentifier `json:"transaction_identifier,omitempty"`
	AccountIdentifier     *AccountIdentifier     `json:"account_identifier,omitempty"`
	CoinIdentifier        *CoinIdentifier        `json:"coin_identifier,omitempty"`
	Currency              *Currency              `json:"currency,omitempty"`
	// status is the network-specific operation type.
	Status *string `json:"status,omitempty"`
	// type is the network-specific operation type.
	Type *string `json:"type,omitempty"`
	// address is AccountIdentifier.Address. This is used to get all transactions related to an
	// AccountIdentifier.Address, regardless of SubAccountIdentifier.
	Address *string `json:"address,omitempty"`
	// success is a synthetic condition populated by parsing network-specific operation statuses
	// (using the mapping provided in `/network/options`).
	Success *bool `json:"success,omitempty"`
}
