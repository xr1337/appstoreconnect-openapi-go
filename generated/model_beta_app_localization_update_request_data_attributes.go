/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaAppLocalizationUpdateRequestDataAttributes struct for BetaAppLocalizationUpdateRequestDataAttributes
type BetaAppLocalizationUpdateRequestDataAttributes struct {
	FeedbackEmail     string `json:"feedbackEmail,omitempty"`
	MarketingUrl      string `json:"marketingUrl,omitempty"`
	PrivacyPolicyUrl  string `json:"privacyPolicyUrl,omitempty"`
	TvOsPrivacyPolicy string `json:"tvOsPrivacyPolicy,omitempty"`
	Description       string `json:"description,omitempty"`
}
