/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

type PdfSaveOption struct {
	// Gets or sets  of the SaveFormat.
	SaveFormat SaveFormat `json:"SaveFormat,omitempty"`
	// Some files like OBJ depends on external file, the lookup paths will allows Aspose.3D to look for external file to load
	LookupPaths []string `json:"LookupPaths,omitempty"`
	// The file name of the exporting/importing scene. This is optional, but useful when serialize external assets like OBJ's material.
	FileName string `json:"FileName,omitempty"`
	// The file format like FBX,U3D,PDF ....
	FileFormat string `json:"FileFormat,omitempty"`
	// Gets or sets to flip the coordinate system of the scene during exporting.
	FlipCoordinateSystem bool `json:"FlipCoordinateSystem,omitempty"`
	// Render mode specifies the style in which the 3D artwork is rendered.
	RenderMode *PdfRenderMode `json:"RenderMode,omitempty"`
	// LightingScheme specifies the lighting to apply to 3D artwork.
	LightingScheme *PdfLightingScheme `json:"LightingScheme,omitempty"`
}
