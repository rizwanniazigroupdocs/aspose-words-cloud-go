/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// Section element.
type HeaderFooter struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets paragraph's text.
	Type_ string `json:"Type,omitempty"`

	// Gets or sets child nodes.
	ChildNodes []NodeLink `json:"ChildNodes,omitempty"`

	// Gets or sets link to DrawingObjects resource.
	DrawingObjects *LinkElement `json:"DrawingObjects,omitempty"`

	// Gets or sets link to Paragraphs resource.
	Paragraphs *LinkElement `json:"Paragraphs,omitempty"`
}
