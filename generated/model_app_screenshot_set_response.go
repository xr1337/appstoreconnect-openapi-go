/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppScreenshotSetResponse struct for AppScreenshotSetResponse
type AppScreenshotSetResponse struct {
	Data     AppScreenshotSet `json:"data"`
	Included []AppScreenshot  `json:"included,omitempty"`
	Links    DocumentLinks    `json:"links"`
}