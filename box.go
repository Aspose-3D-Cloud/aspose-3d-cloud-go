/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

// The Box Entity
type Box struct {
	// Gets or sets the length of the box             
	Length float64 `json:"Length"`
	// Gets or sets the width of the box
	Width float64 `json:"Width"`
	// Gets or sets the height of the box
	Height float64 `json:"Height"`
	// Gets or sets the name of the box             
	Name string `json:"Name,omitempty"`
	// Gets or sets the lengthSegments of the box
	LengthSegments int32 `json:"LengthSegments"`
	// Gets or sets the widthSegments of the box
	WidthSegments int32 `json:"WidthSegments"`
	// Gets or sets the heightSegments of the box
	HeightSegments int32 `json:"HeightSegments"`
}
