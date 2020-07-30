/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

import (
	"time"
)

// CertificateAttributes struct for CertificateAttributes
type CertificateAttributes struct {
	Name               string           `json:"name,omitempty"`
	CertificateType    CertificateType  `json:"certificateType,omitempty"`
	DisplayName        string           `json:"displayName,omitempty"`
	SerialNumber       string           `json:"serialNumber,omitempty"`
	Platform           BundleIdPlatform `json:"platform,omitempty"`
	ExpirationDate     time.Time        `json:"expirationDate,omitempty"`
	CertificateContent string           `json:"certificateContent,omitempty"`
}