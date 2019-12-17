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
type TableCellFormat struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add below the contents of cell.
	BottomPadding float64 `json:"BottomPadding,omitempty"`

	// Gets or sets if true, fits text in the cell, compressing each paragraph to the width of the cell.
	FitText bool `json:"FitText,omitempty"`

	// Gets or sets specifies how the cell is merged horizontally with other cells in the row.
	HorizontalMerge string `json:"HorizontalMerge,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add to the left of the contents of cell.
	LeftPadding float64 `json:"LeftPadding,omitempty"`

	// Gets or sets returns or sets the orientation of text in a table cell.
	Orientation string `json:"Orientation,omitempty"`

	// Gets or sets returns or sets the preferred width of the cell.
	PreferredWidth *PreferredWidth `json:"PreferredWidth,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add to the right of the contents of cell.
	RightPadding float64 `json:"RightPadding,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add above the contents of cell.
	TopPadding float64 `json:"TopPadding,omitempty"`

	// Gets or sets returns or sets the vertical alignment of text in the cell.
	VerticalAlignment string `json:"VerticalAlignment,omitempty"`

	// Gets or sets specifies how the cell is merged with other cells vertically.
	VerticalMerge string `json:"VerticalMerge,omitempty"`

	// Gets or sets the width of the cell in points.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets if true, wrap text for the cell.
	WrapText bool `json:"WrapText,omitempty"`
}
