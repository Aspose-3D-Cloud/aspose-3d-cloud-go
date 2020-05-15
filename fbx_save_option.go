/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

type FbxSaveOption struct {
	// Gets or sets  of the SaveFormat.
	SaveFormat SaveFormat `json:"SaveFormat,omitempty"`
	// Some files like OBJ depends on external file, the lookup paths will allows Aspose.3D to look for external file to load
	LookupPaths []string `json:"LookupPaths,omitempty"`
	// The file name of the exporting/importing scene. This is optional, but useful when serialize external assets like OBJ's material.
	FileName string `json:"FileName,omitempty"`
	// The file format like FBX,U3D,PDF ....
	FileFormat string `json:"FileFormat,omitempty"`
	//  Compression large binary data in the FBX file, default value is true
	EnableCompression bool `json:"EnableCompression,omitempty"`
	// Gets or sets whether reuse repeated curve data by increasing last data's ref count
	FoldRepeatedCurveData bool `json:"FoldRepeatedCurveData,omitempty"`
	// Gets or sets whether export legacy material properties, used for back compatibility. This option is turned on by default
	ExportLegacyMaterialProperties bool `json:"ExportLegacyMaterialProperties,omitempty"`
	// Gets or sets whether generate a Video instance for Aspose.ThreeD.Shading.Texture when exporting as FBX.
	VideoForTexture bool `json:"VideoForTexture,omitempty"`
	// Gets or sets whether always generate a Aspose.ThreeD.Entities.VertexElementMaterial for geometries if the attached node contains materials. This is turned off by default.
	GenerateVertexElementMaterial bool `json:"GenerateVertexElementMaterial,omitempty"`
}
