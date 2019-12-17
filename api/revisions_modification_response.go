/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// response of the modification operations for the revisions collection (now these are acceptAll and rejectAll).
type RevisionsModificationResponse struct {

	// Gets or sets request Id.
	RequestId string `json:"RequestId,omitempty"`

	// Gets or sets result of the modification operations for the revisions collection.
	Result *ModificationOperationResult `json:"Result,omitempty"`
}
