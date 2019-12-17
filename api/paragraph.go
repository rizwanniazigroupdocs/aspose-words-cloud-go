/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// Paragraph element.
type Paragraph struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets node id.
	NodeId string `json:"NodeId,omitempty"`

	// Gets or sets child nodes.
	ChildNodes []NodeLink `json:"ChildNodes,omitempty"`
}
