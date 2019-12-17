/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// Paragraph format element.             
type ParagraphFormat struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets a flag indicating whether inter-character spacing is automatically adjusted between regions of Latin text and regions of East Asian text in the current paragraph.             
	AddSpaceBetweenFarEastAndAlpha bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

	// Gets or sets a flag indicating whether inter-character spacing is automatically adjusted between regions of numbers and regions of East Asian text in the current paragraph.             
	AddSpaceBetweenFarEastAndDigit bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

	// Gets or sets text alignment for the paragraph.             
	Alignment string `json:"Alignment,omitempty"`

	// Gets or sets whether this is a right-to-left paragraph.             
	Bidi bool `json:"Bidi,omitempty"`

	// Gets or sets the position for a drop cap text.             
	DropCapPosition string `json:"DropCapPosition,omitempty"`

	// Gets or sets the value (in points) for a first line or hanging indent. Use a positive value to set a first-line indent, and use a negative value to set a hanging indent.             
	FirstLineIndent float64 `json:"FirstLineIndent,omitempty"`

	// Gets or sets True when the paragraph is an item in a bulleted or numbered list.
	IsListItem bool `json:"IsListItem,omitempty"`

	// Gets or sets true if all lines in the paragraph are to remain on the same page.             
	KeepTogether bool `json:"KeepTogether,omitempty"`

	// Gets or sets true if the paragraph is to remains on the same page as the paragraph that follows it.             
	KeepWithNext bool `json:"KeepWithNext,omitempty"`

	// Gets or sets the value (in points) that represents the left indent for paragraph.             
	LeftIndent float64 `json:"LeftIndent,omitempty"`

	// Gets or sets the line spacing (in points) for the paragraph.             
	LineSpacing float64 `json:"LineSpacing,omitempty"`

	// Gets or sets the line spacing for the paragraph.             
	LineSpacingRule string `json:"LineSpacingRule,omitempty"`

	// Gets or sets the number of lines of the paragraph text used to calculate the drop cap height.             
	LinesToDrop int32 `json:"LinesToDrop,omitempty"`

	// Gets or sets when true,  and  will be ignored between the paragraphs of the same style.             
	NoSpaceBetweenParagraphsOfSameStyle bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

	// Gets or sets specifies the outline level of the paragraph in the document.             
	OutlineLevel string `json:"OutlineLevel,omitempty"`

	// Gets or sets true if a page break is forced before the paragraph.             
	PageBreakBefore bool `json:"PageBreakBefore,omitempty"`

	// Gets or sets the value (in points) that represents the right indent for paragraph.             
	RightIndent float64 `json:"RightIndent,omitempty"`

	// Gets or sets the amount of spacing (in points) after the paragraph.             
	SpaceAfter float64 `json:"SpaceAfter,omitempty"`

	// Gets or sets true if the amount of spacing after the paragraph is set automatically.             
	SpaceAfterAuto bool `json:"SpaceAfterAuto,omitempty"`

	// Gets or sets the amount of spacing (in points) before the paragraph.             
	SpaceBefore float64 `json:"SpaceBefore,omitempty"`

	// Gets or sets true if the amount of spacing before the paragraph is set automatically.             
	SpaceBeforeAuto bool `json:"SpaceBeforeAuto,omitempty"`

	// Gets or sets the locale independent style identifier of the paragraph style applied to this formatting.             
	StyleIdentifier string `json:"StyleIdentifier,omitempty"`

	// Gets or sets the name of the paragraph style applied to this formatting.             
	StyleName string `json:"StyleName,omitempty"`

	// Gets or sets specifies whether the current paragraph should be exempted from any hyphenation which is applied in the document settings.             
	SuppressAutoHyphens bool `json:"SuppressAutoHyphens,omitempty"`

	// Gets or sets specifies whether the current paragraph's lines should be exempted from line numbering which is applied in the parent section.             
	SuppressLineNumbers bool `json:"SuppressLineNumbers,omitempty"`

	// Gets or sets true if the first and last lines in the paragraph are to remain on the same page as the rest of the paragraph.             
	WidowControl bool `json:"WidowControl,omitempty"`
}
