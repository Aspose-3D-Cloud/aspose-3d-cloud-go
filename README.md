![](https://img.shields.io/badge/api-v3.0-lightgrey) [![GitHub license](https://img.shields.io/github/license/aspose-3d-cloud/aspose-3d-cloud-go)](https://github.com/aspose-3d-cloud/aspose-3d-cloud-go/blob/master/LICENSE) ![GitHub last commit](https://img.shields.io/github/last-commit/Aspose-3D-Cloud/aspose-3d-cloud-go) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Aspose-3D-Cloud/aspose-3d-cloud-go)

# Go SDK for 3D File Formats

[Aspose.3D Cloud SDK for Go](https://products.aspose.cloud/3d/go) empowers your Go Lang applications to connect with 3D document formats. The APIs let engineers read, convert, build, alter and control the substance of the [3D document formats](https://docs.aspose.cloud/3d/supported-document-formats/) in the Cloud. All without any 3D modeling and rendering software installed on the machine.

## Process 3D Objects & Files in the Cloud

- [Create 3D entities](https://docs.aspose.cloud/3d/create-a-new-entity-with-size/) - Box, Sphere, Plane, Torus, Cylinder.
- Perform processing of 3D models and attributes.
- Perform transformation, translation, rotation & scaling of 3D objects.
- [Convert 3D files](https://docs.aspose.cloud/3d/converting-between-formats-using-aspose-3d-cloud/) to other format.
- Convert whole of the 3D file or convert a specific part only.
- [Extract 3D scenes](https://docs.aspose.cloud/3d/extract-and-save-a-scene-in-a-different-format/) and export to various formats.
- Perform parametric modeling, 3D modeling & data processing.
- Create cloud-based folder structure & perform cloud-based conversion of 3D files.
- [Extract 3D raw data from PDF](https://docs.aspose.cloud/3d/extract-raw-data-from-a-pdf-file/). Password-protected 3D PDF file supported as well.
- Supports working with triangulate meshes, triangulate whole or part of the 3D scene.
- [Remove 3D objects from a scene](https://docs.aspose.cloud/3d/delete-nodes-from-a-scene/).


## Read & Write 3D Formats

**Additive Manufacturing Format** AMF, **3D Studio** 3DS, **AVEVA** RVM, **Google Draco** DRC, **Collada** DAE, **AutoCAD** DXF, **Autodesk** FBX, **Wavefront** OBJ, **3D** PDF, **Polygon File Format** PLY, **Standard Triangle Language** STL, **Universal3D** U3D, **GL Transmission Format** GLTF

## Read-Only Formats

**Siemens** JT, **DirectX** X, **3D Manufacturing Format** 3MF, **3D Studio Max** ASE

## Write 3D Formats As

HTML

## Enhancements Version 20.5
- Added support to read/write more file formats.
- Added support for HTML file format conversion.
- New feature added to Triangulate part of the scene (Specified by OAP) and save the scene to a different file.

## Get Started with Aspose.3D Cloud SDK for Go

Register an account at [Aspose Cloud Dashboard](https://dashboard.aspose.cloud/#/apps) to get you application information. Next, download the repository and refer to `three_d_cloud_test.go`or download the [Go Module](https://pkg.go.dev/github.com/Aspose-3D-Cloud/aspose-3d-cloud-go/v20?tab=overview) using `import "github.com/Aspose-3D-Cloud/aspose-3d-cloud-go/v20"1` from your application.

## Convert 3D to PDF Conversion in GO

```
	// Get your ClientId and ClientSecret from https://dashboard.aspose.cloud (free registration required).
	threeDCloudApi , ctx, _ := api.NewThreeDCloudApiService("MY_CLIENT_ID", "MY_CLIENT_SECRET", "ap.aspose.cloud")
	
	response, httpResponse, err := threeDCloudApi.PostConvertByFormat("sample.3d", "pdf", "output.pdf")
```


## Aspose.3D Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Swift | Perl | GO |
|---|---|---|---|---|---|---|---|---|
| [GitHub](https://github.com/aspose-3d-cloud/aspose-3d-cloud-dotnet) | [GitHub](https://github.com/aspose-3d-cloud/aspose-3d-cloud-java) | [GitHub](https://github.com/aspose-3d-cloud/aspose-3d-cloud-php) |[GitHub](https://github.com/Aspose-3D-Cloud/aspose-3d-cloud-python) | [GitHub](https://github.com/Aspose-3D-Cloud/aspose-3d-cloud-ruby)  | [GitHub](https://github.com/Aspose-3D-Cloud/aspose-3d-cloud-node) | [GitHub](https://github.com/aspose-3d-cloud/aspose-3d-cloud-swift) | [GitHub](https://github.com/Aspose-3D-Cloud/aspose-3d-cloud-perl) | [GitHub](https://github.com/Aspose-3D-Cloud/aspose-3d-cloud-go) |
| [NuGet](https://www.nuget.org/packages/Aspose.3D-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-3d-cloud) | [Composer](https://packagist.org/packages/aspose/3d-sdk-php) | [PIP](https://pypi.org/project/aspose3dcloud/) | [GEM](https://rubygems.org/gems/aspose_3d_cloud)  | [NPM](https://www.npmjs.com/package/aspose3dcloud) | [POD](https://cocoapods.org/pods/Aspose3DCloud) |  [CPAN](https://metacpan.org/release/AsposeThreeDCloud-ThreeDCloudApi) | [GO](https://pkg.go.dev/github.com/Aspose-3D-Cloud/aspose-3d-cloud-go/v20?tab=overview) |

[Home](https://www.aspose.cloud) | [Product Page](https://products.aspose.cloud/3d/go) | [Documentation](https://docs.aspose.cloud/3d/) | [Live Demo](https://products.aspose.app/3d/family) | [API Reference](https://apireference.aspose.cloud/3d/) | [Code Samples](https://github.com/Aspose-3D-Cloud/aspose-3d-cloud-go) | [Blog](https://blog.aspose.cloud/category/3d/) | [Free Support](https://forum.aspose.cloud/c/3d) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
