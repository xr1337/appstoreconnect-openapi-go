/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaTesterResponse struct for BetaTesterResponse
type BetaTesterResponse struct {
	Data     BetaTester    `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}
