/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// Collection of links to paragraphs.
type ParagraphLinkCollection struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets collection of paragraph's links.
	ParagraphLinkList []ParagraphLink `json:"ParagraphLinkList,omitempty"`
}
