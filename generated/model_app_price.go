/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppPrice struct for AppPrice
type AppPrice struct {
	Type          string                `json:"type"`
	Id            string                `json:"id"`
	Relationships AppPriceRelationships `json:"relationships,omitempty"`
	Links         ResourceLinks         `json:"links"`
}
