/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// ProfileRelationshipsBundleId struct for ProfileRelationshipsBundleId
type ProfileRelationshipsBundleId struct {
	Links AppCategoryRelationshipsSubcategoriesLinks                   `json:"links,omitempty"`
	Data  BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData `json:"data,omitempty"`
}
