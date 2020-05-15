/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

// File upload result
type FilesUploadResult struct {
	// List of uploaded file names
	Uploaded []string `json:"Uploaded,omitempty"`
	// List of errors.
	Errors []ModelError `json:"Errors,omitempty"`
}
