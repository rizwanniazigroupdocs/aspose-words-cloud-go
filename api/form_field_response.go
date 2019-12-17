/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// This response should be returned by the service when handling: GET http://api.aspose.com/v4.0/words/Test.doc/paragraphs/{0}/formfields/{1}.
type FormFieldResponse struct {

	// Gets or sets request Id.
	RequestId string `json:"RequestId,omitempty"`

	// Gets or sets field information.
	FormField *FormField `json:"FormField,omitempty"`
}
