/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppPreOrderCreateRequestData struct for AppPreOrderCreateRequestData
type AppPreOrderCreateRequestData struct {
	Type          string                                    `json:"type"`
	Attributes    AppPreOrderCreateRequestDataAttributes    `json:"attributes,omitempty"`
	Relationships AppPreOrderCreateRequestDataRelationships `json:"relationships"`
}