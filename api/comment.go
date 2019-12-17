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

// Comment.
type Comment struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets returns or sets the author name for a comment.
	Author string `json:"Author,omitempty"`

	// Gets or sets content of comment.
	Content *StoryChildNodes `json:"Content,omitempty"`

	// Gets or sets the date and time that the comment was made.
	DateTime time.Time `json:"DateTime,omitempty"`

	// Gets or sets returns or sets the initials of the user associated with a specific comment.
	Initial string `json:"Initial,omitempty"`

	// Gets or sets link to comment range end node.
	RangeEnd *DocumentPosition `json:"RangeEnd,omitempty"`

	// Gets or sets link to comment range start node.
	RangeStart *DocumentPosition `json:"RangeStart,omitempty"`

	// Gets or sets this is a convenience property that allows to easily get or set text of the comment.
	Text string `json:"Text,omitempty"`
}
