/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppInfoUpdateRequestData struct for AppInfoUpdateRequestData
type AppInfoUpdateRequestData struct {
	Type          string                                `json:"type"`
	Id            string                                `json:"id"`
	Relationships AppInfoUpdateRequestDataRelationships `json:"relationships,omitempty"`
}