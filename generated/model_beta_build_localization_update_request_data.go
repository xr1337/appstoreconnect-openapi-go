/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaBuildLocalizationUpdateRequestData struct for BetaBuildLocalizationUpdateRequestData
type BetaBuildLocalizationUpdateRequestData struct {
	Type       string                                           `json:"type"`
	Id         string                                           `json:"id"`
	Attributes BetaBuildLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}