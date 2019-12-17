/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// Container for the page's statistical data.
type PageStatData struct {

	// Gets or sets page number.
	PageNumber int32 `json:"PageNumber"`

	// Gets or sets total count of words in the page.
	WordCount int32 `json:"WordCount"`

	// Gets or sets total count of paragraphs in the page.
	ParagraphCount int32 `json:"ParagraphCount"`

	// Gets or sets detailed statistics of footnotes.
	FootnotesStatData *FootnotesStatData `json:"FootnotesStatData,omitempty"`
}
