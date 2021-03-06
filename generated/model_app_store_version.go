/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppStoreVersion struct for AppStoreVersion
type AppStoreVersion struct {
	Type          string                       `json:"type"`
	Id            string                       `json:"id"`
	Attributes    AppStoreVersionAttributes    `json:"attributes,omitempty"`
	Relationships AppStoreVersionRelationships `json:"relationships,omitempty"`
	Links         ResourceLinks                `json:"links"`
}
