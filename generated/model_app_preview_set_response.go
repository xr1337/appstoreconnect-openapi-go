/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppPreviewSetResponse struct for AppPreviewSetResponse
type AppPreviewSetResponse struct {
	Data     AppPreviewSet `json:"data"`
	Included []AppPreview  `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}