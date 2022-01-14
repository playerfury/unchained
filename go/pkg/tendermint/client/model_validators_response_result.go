/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ValidatorsResponseResult struct for ValidatorsResponseResult
type ValidatorsResponseResult struct {
	BlockHeight string `json:"block_height"`
	Validators []ValidatorPriority `json:"validators"`
	Count *string `json:"count,omitempty"`
	Total *string `json:"total,omitempty"`
}

// NewValidatorsResponseResult instantiates a new ValidatorsResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatorsResponseResult(blockHeight string, validators []ValidatorPriority) *ValidatorsResponseResult {
	this := ValidatorsResponseResult{}
	this.BlockHeight = blockHeight
	this.Validators = validators
	return &this
}

// NewValidatorsResponseResultWithDefaults instantiates a new ValidatorsResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatorsResponseResultWithDefaults() *ValidatorsResponseResult {
	this := ValidatorsResponseResult{}
	return &this
}

// GetBlockHeight returns the BlockHeight field value
func (o *ValidatorsResponseResult) GetBlockHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *ValidatorsResponseResult) GetBlockHeightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *ValidatorsResponseResult) SetBlockHeight(v string) {
	o.BlockHeight = v
}

// GetValidators returns the Validators field value
func (o *ValidatorsResponseResult) GetValidators() []ValidatorPriority {
	if o == nil {
		var ret []ValidatorPriority
		return ret
	}

	return o.Validators
}

// GetValidatorsOk returns a tuple with the Validators field value
// and a boolean to check if the value has been set.
func (o *ValidatorsResponseResult) GetValidatorsOk() (*[]ValidatorPriority, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Validators, true
}

// SetValidators sets field value
func (o *ValidatorsResponseResult) SetValidators(v []ValidatorPriority) {
	o.Validators = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ValidatorsResponseResult) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorsResponseResult) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ValidatorsResponseResult) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *ValidatorsResponseResult) SetCount(v string) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ValidatorsResponseResult) GetTotal() string {
	if o == nil || o.Total == nil {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorsResponseResult) GetTotalOk() (*string, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ValidatorsResponseResult) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *ValidatorsResponseResult) SetTotal(v string) {
	o.Total = &v
}

func (o ValidatorsResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["block_height"] = o.BlockHeight
	}
	if true {
		toSerialize["validators"] = o.Validators
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableValidatorsResponseResult struct {
	value *ValidatorsResponseResult
	isSet bool
}

func (v NullableValidatorsResponseResult) Get() *ValidatorsResponseResult {
	return v.value
}

func (v *NullableValidatorsResponseResult) Set(val *ValidatorsResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatorsResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatorsResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatorsResponseResult(val *ValidatorsResponseResult) *NullableValidatorsResponseResult {
	return &NullableValidatorsResponseResult{value: val, isSet: true}
}

func (v NullableValidatorsResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatorsResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

