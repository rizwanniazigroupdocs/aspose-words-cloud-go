/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package api

// Container class for epub save options.
type EpubSaveOptionsData struct {

	// Gets or sets format of save.
	SaveFormat string `json:"SaveFormat,omitempty"`

	// Gets or sets name of destination file.
	FileName string `json:"FileName,omitempty"`

	// Gets or sets a value determining how DrawingML shapes are rendered. { Fallback | DrawingML }.
	DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

	// Gets or sets a value determining how DrawingML effects are rendered. { Simplified | None | Fine }.
	DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

	// Gets or sets controls zip output or not. Default value is false.
	ZipOutput bool `json:"ZipOutput,omitempty"`

	// Gets or sets a value determining whether the Aspose.Words.Properties.BuiltInDocumentProperties.LastSavedTime property is updated before saving.
	UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

	// Gets or sets value determining whether content of  is updated before saving.
	UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

	// Gets or sets a value determining if fields should be updated before saving the document to a fixed page format. Default value for this property is. true
	UpdateFields bool `json:"UpdateFields,omitempty"`

	// Gets or sets specifies whether negative left and right indents of paragraphs are allowed (not normalized).
	AllowNegativeIndent bool `json:"AllowNegativeIndent,omitempty"`

	// Gets or sets specifies a prefix which is added to all CSS class names. Default value is an empty string and generated CSS class names have no common prefix.  If this value is not empty, all CSS classes generated by Aspose.Words will start with the specified prefix.This might be useful, for example, if you add custom CSS to generated documents and want to prevent class name conflicts. If the value is not null or empty, it must be a valid CSS identifier.
	CssClassNamePrefix string `json:"CssClassNamePrefix,omitempty"`

	// Gets or sets specifies the name of the CSS file written when the document is exported to HTML.
	CssStyleSheetFileName string `json:"CssStyleSheetFileName,omitempty"`

	// Gets or sets specifies how CSS styles are exported.
	CssStyleSheetType string `json:"CssStyleSheetType,omitempty"`

	// Gets or sets specifies how the document should be split when saving.
	DocumentSplitCriteria string `json:"DocumentSplitCriteria,omitempty"`

	// Gets or sets specifies the maximum level of headings at which to split the document.
	DocumentSplitHeadingLevel int32 `json:"DocumentSplitHeadingLevel,omitempty"`

	// Gets or sets specifies the encoding to use when exporting.
	Encoding string `json:"Encoding,omitempty"`

	// Gets or sets specifies whether to export built-in and custom document properties.
	ExportDocumentProperties bool `json:"ExportDocumentProperties,omitempty"`

	// Gets or sets controls how drop-down form fields are saved to HTML. Default value is false.
	ExportDropDownFormFieldAsText bool `json:"ExportDropDownFormFieldAsText,omitempty"`

	// Gets or sets specifies whether font resources should be exported.
	ExportFontResources bool `json:"ExportFontResources,omitempty"`

	// Gets or sets specifies whether fonts resources should be embedded to HTML in Base64 encoding.  Default is false.
	ExportFontsAsBase64 bool `json:"ExportFontsAsBase64,omitempty"`

	// Gets or sets specifies how headers and footers are output.
	ExportHeadersFootersMode string `json:"ExportHeadersFootersMode,omitempty"`

	// Gets or sets specifies whether images are saved in Base64 format.
	ExportImagesAsBase64 bool `json:"ExportImagesAsBase64,omitempty"`

	// Gets or sets specifies whether language information is exported.
	ExportLanguageInformation bool `json:"ExportLanguageInformation,omitempty"`

	// Gets or sets controls how list labels are output.
	ExportListLabels string `json:"ExportListLabels,omitempty"`

	// Gets or sets specifies whether original URL should be used as the URL of the linked images. Default value is false.
	ExportOriginalUrlForLinkedImages bool `json:"ExportOriginalUrlForLinkedImages,omitempty"`

	// Gets or sets specifies whether page margins is exported to HTML, MHTML or EPUB. Default is false.
	ExportPageMargins bool `json:"ExportPageMargins,omitempty"`

	// Gets or sets specifies whether page setup is exported.
	ExportPageSetup bool `json:"ExportPageSetup,omitempty"`

	// Gets or sets specifies whether font sizes should be output in relative units when saving.
	ExportRelativeFontSize bool `json:"ExportRelativeFontSize,omitempty"`

	// Gets or sets specifies whether to write the roundtrip information when saving to HTML Default value is true.
	ExportRoundtripInformation bool `json:"ExportRoundtripInformation,omitempty"`

	// Gets or sets controls how textboxes represented by Aspose.Words.Drawing.Shape are saved to HTML, MHTML or EPUB. Default value is false.    When set to true, exports textboxes as inline \"svg\" elements. When false, exports as \"image\" elements.
	ExportTextBoxAsSvg bool `json:"ExportTextBoxAsSvg,omitempty"`

	// Gets or sets controls how text input form fields are saved.
	ExportTextInputFormFieldAsText bool `json:"ExportTextInputFormFieldAsText,omitempty"`

	// Gets or sets specifies whether to write page numbers to table of contents when saving.
	ExportTocPageNumbers bool `json:"ExportTocPageNumbers,omitempty"`

	// Gets or sets specifies whether to write the DOCTYPE declaration when saving.
	ExportXhtmlTransitional bool `json:"ExportXhtmlTransitional,omitempty"`

	// Gets or sets controls which font resources need subsetting when saving.
	FontResourcesSubsettingSizeThreshold int32 `json:"FontResourcesSubsettingSizeThreshold,omitempty"`

	// Gets or sets specifies the physical folder where fonts are saved when exporting a document.
	FontsFolder string `json:"FontsFolder,omitempty"`

	// Gets or sets specifies the name of the folder used to construct font URIs.
	FontsFolderAlias string `json:"FontsFolderAlias,omitempty"`

	// Gets or sets specifies version of HTML standard that should be used when saving the document to HTML or MHTML. Default value is Aspose.Words.Saving.HtmlVersion.Xhtml.
	HtmlVersion string `json:"HtmlVersion,omitempty"`

	// Gets or sets specifies the output resolution for images when exporting.
	ImageResolution int32 `json:"ImageResolution,omitempty"`

	// Gets or sets specifies the physical folder where images are saved when exporting a document.
	ImagesFolder string `json:"ImagesFolder,omitempty"`

	// Gets or sets specifies the name of the folder used to construct image URIs.
	ImagesFolderAlias string `json:"ImagesFolderAlias,omitempty"`

	// Gets or sets specifies in what format metafiles are saved when exporting to HTML, MHTML, or EPUB. Default value is Aspose.Words.Saving.HtmlMetafileFormat.Png, meaning that metafiles are rendered to raster PNG images.  Metafiles are not natively displayed by HTML browsers. By default, Aspose.Words converts WMF and EMF images into PNG files when exporting to HTML.Other options are to convert metafiles to SVG images or to export them as is without conversion. Some image transforms, in particular image cropping, will not be applied to metafile images if they are exported to HTML without conversion.
	MetafileFormat string `json:"MetafileFormat,omitempty"`

	// Gets or sets controls how OfficeMath objects are exported to HTML, MHTML or EPUB.  Default value is HtmlOfficeMathOutputMode.Image.
	OfficeMathOutputMode string `json:"OfficeMathOutputMode,omitempty"`

	// Gets or sets specifies whether or not use pretty formats output.
	PrettyFormat bool `json:"PrettyFormat,omitempty"`

	// Gets or sets specifies whether font family names used in the document are resolved and substituted according to FontSettings when being written into HTML-based formats. default value is false.
	ResolveFontNames bool `json:"ResolveFontNames,omitempty"`

	// Gets or sets specifies a physical folder where all resources like images, fonts, and external CSS are saved when a document is exported to HTML. Default is an empty string.
	ResourceFolder string `json:"ResourceFolder,omitempty"`

	// Gets or sets specifies the name of the folder used to construct URIs of all resources written into an HTML document.  Default is an empty string.
	ResourceFolderAlias string `json:"ResourceFolderAlias,omitempty"`

	// Gets or sets specifies whether images are scaled by Aspose.Words to the bounding shape size when exporting.
	ScaleImageToShapeSize bool `json:"ScaleImageToShapeSize,omitempty"`

	// Gets or sets controls how table, row and cell widths are exported.
	TableWidthOutputMode string `json:"TableWidthOutputMode,omitempty"`

	// Gets or sets specifies the maximum level of headings populated to the navigation map when exporting.
	EpubNavigationMapLevel int32 `json:"EpubNavigationMapLevel,omitempty"`
}
type IEpubSaveOptionsData interface {
	IsEpubSaveOptionsData() bool
}
func (EpubSaveOptionsData) IsEpubSaveOptionsData() bool {
	return true;
}
func (EpubSaveOptionsData) IsHtmlSaveOptionsData() bool {
	return true;
}

