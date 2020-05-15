/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

type DracoSaveOption struct {
	// Gets or sets  of the SaveFormat.
	SaveFormat SaveFormat `json:"SaveFormat,omitempty"`
	// Some files like OBJ depends on external file, the lookup paths will allows Aspose.3D to look for external file to load
	LookupPaths []string `json:"LookupPaths,omitempty"`
	// The file name of the exporting/importing scene. This is optional, but useful when serialize external assets like OBJ's material.
	FileName string `json:"FileName,omitempty"`
	// The file format like FBX,U3D,PDF ....
	FileFormat string `json:"FileFormat,omitempty"`
	// Quantization bits for position, default value is 14
	PositionBits int32 `json:"PositionBits,omitempty"`
	// Quantization bits for texture coordinate, default value is 12
	TextureCoordinateBits int32 `json:"TextureCoordinateBits,omitempty"`
	// Quantization bits for vertex color, default value is 10
	ColorBits int32 `json:"ColorBits,omitempty"`
	// Quantization bits for normal vectors, default value is 10
	NormalBits int32 `json:"NormalBits,omitempty"`
	// Compression level, default value is Aspose.ThreeD.Formats.DracoCompressionLevel.Standard.
	CompressionLevel *DracoCompressionLevel `json:"CompressionLevel,omitempty"`
}
