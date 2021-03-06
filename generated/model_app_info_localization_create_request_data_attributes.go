/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppInfoLocalizationCreateRequestDataAttributes struct for AppInfoLocalizationCreateRequestDataAttributes
type AppInfoLocalizationCreateRequestDataAttributes struct {
	Locale            string `json:"locale"`
	Name              string `json:"name,omitempty"`
	Subtitle          string `json:"subtitle,omitempty"`
	PrivacyPolicyUrl  string `json:"privacyPolicyUrl,omitempty"`
	PrivacyPolicyText string `json:"privacyPolicyText,omitempty"`
}
