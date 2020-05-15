/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

type ColladaSaveOption struct {
	// Gets or sets  of the SaveFormat.
	SaveFormat SaveFormat `json:"SaveFormat,omitempty"`
	// Some files like OBJ depends on external file, the lookup paths will allows Aspose.3D to look for external file to load
	LookupPaths []string `json:"LookupPaths,omitempty"`
	// The file name of the exporting/importing scene. This is optional, but useful when serialize external assets like OBJ's material.
	FileName string `json:"FileName,omitempty"`
	// The file format like FBX,U3D,PDF ....
	FileFormat string `json:"FileFormat,omitempty"`
	// Gets or sets whether the exported XML document is indented.
	Indented bool `json:"Indented,omitempty"`
	// Gets or sets  of the Transform Style.
	TransformStyle *ColladaTransformStyle `json:"TransformStyle,omitempty"`
}
