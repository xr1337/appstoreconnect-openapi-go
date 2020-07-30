/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaGroup struct for BetaGroup
type BetaGroup struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Attributes    BetaGroupAttributes    `json:"attributes,omitempty"`
	Relationships BetaGroupRelationships `json:"relationships,omitempty"`
	Links         ResourceLinks          `json:"links"`
}