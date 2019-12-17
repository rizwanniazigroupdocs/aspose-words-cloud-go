/*
 * Aspose.Words Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 19.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// container class for doc/dot save options.
type DocSaveOptionsData struct {

	// Gets or sets a value determining how colors are rendered. { Normal | Grayscale}.
	ColorMode string `json:"ColorMode,omitempty"`

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

	// Gets or sets When false, small metafiles are not compressed for performance reason. Default value is true, all metafiles are compressed regardless of its size.
	AlwaysCompressMetafiles bool `json:"AlwaysCompressMetafiles,omitempty"`

	// Gets or sets password.
	Password string `json:"Password,omitempty"`

	// Gets or sets When false, PictureBullet data is not saved to output document. Default value is true.
	SavePictureBullet bool `json:"SavePictureBullet,omitempty"`

	// Gets or sets determine whether or not save RoutingSlip data saved to output document.
	SaveRoutingSlip bool `json:"SaveRoutingSlip,omitempty"`
}
