/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppInfoLocalizationsResponse struct for AppInfoLocalizationsResponse
type AppInfoLocalizationsResponse struct {
	Data  []AppInfoLocalization `json:"data"`
	Links PagedDocumentLinks    `json:"links"`
	Meta  PagingInformation     `json:"meta,omitempty"`
}