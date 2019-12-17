/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// Represents all formatting for a table row.
type TableRowFormat struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets true if the text in a table row is allowed to split across a page break.
	AllowBreakAcrossPages bool `json:"AllowBreakAcrossPages,omitempty"`

	// Gets or sets true if the row is repeated as a table heading on every page when the table spans more than one page.
	HeadingFormat bool `json:"HeadingFormat,omitempty"`

	// Gets or sets the height of the table row in points.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the rule for determining the height of the table row.
	HeightRule string `json:"HeightRule,omitempty"`
}
