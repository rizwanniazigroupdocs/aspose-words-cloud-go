/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"time"
)

// container class for details of digital signature.
type PdfDigitalSignatureDetailsData struct {

	// Gets or sets certificate's filename using for signing.
	CertificateFilename string `json:"CertificateFilename,omitempty"`

	// Gets or sets hash algorithm.
	HashAlgorithm string `json:"HashAlgorithm,omitempty"`

	// Gets or sets location of the signing.
	Location string `json:"Location,omitempty"`

	// Gets or sets reason for the signing.
	Reason string `json:"Reason,omitempty"`

	// Gets or sets date of the signing.
	SignatureDate time.Time `json:"SignatureDate,omitempty"`
}
