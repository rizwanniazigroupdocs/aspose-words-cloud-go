/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// This response should be returned by the service when handling:  POST /{name}/split .
type SplitDocumentResponse struct {

	// Gets or sets request Id.
	RequestId string `json:"RequestId,omitempty"`

	// Gets or sets resylt of splitting document.
	SplitResult *SplitDocumentResult `json:"SplitResult,omitempty"`
}
