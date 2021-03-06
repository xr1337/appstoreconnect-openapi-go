/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// DiagnosticSignature struct for DiagnosticSignature
type DiagnosticSignature struct {
	Type       string                        `json:"type"`
	Id         string                        `json:"id"`
	Attributes DiagnosticSignatureAttributes `json:"attributes,omitempty"`
	Links      ResourceLinks                 `json:"links"`
}
