 /**
 *
 *   Copyright (c) 2020 Aspose.3D Cloud
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
 *
 */
package aspose3dcloud

import (
	//"os"
	"fmt"
	"context"
	"github.com/antihax/optional"		
	"testing"
	//"io/ioutil"
	//"unsafe"
)

/*
func TestCopyFile(t *testing.T) {
	srcPath := "3DTest/Aspose.pdf"
	destPath := "3DTest/Aspose.Go.pdf"

    ctx := context.Background()
    localVarOptionals := new(CopyFileOpts)
    localVarOptionals.SrcStorageName = optional.NewString("") //like static fun
    localVarOptionals.DestStorageName = optional.NewString("")
    localVarOptionals.VersionId = optional.NewString("")
    
    httpResponse,err := GetBaseTest().ThreeDCloudAPI.CopyFile(ctx, srcPath, destPath, localVarOptionals)
    	
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestCopyFile --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
	
}

func TestUploadFile(t *testing.T) {
	name := "Aspose.pdf"
	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name) 
	if err != nil {
		t.Error(err)
	}

    ctx := context.Background()
    localVarOptionals := new(UploadFileOpts)
    localVarOptionals.StorageName = optional.NewString("")  
    
    _,httpResponse,err := GetBaseTest().ThreeDCloudAPI.UploadFile(ctx, GetBaseTest().remoteFolder + "/" + name, file, localVarOptionals)
    	
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestUploadFile --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
	
}


func TestDownloadFile(t *testing.T) {
	name := "Aspose.pdf"
    ctx := context.Background()
    localVarOptionals := new(DownloadFileOpts)
    localVarOptionals.StorageName = optional.NewString("")
    localVarOptionals.VersionId  = optional.NewString("")  	
	dataRet,_, httpResponse, err := GetBaseTest().ThreeDCloudAPI.DownloadFile(ctx, GetBaseTest().remoteFolder + "/" + name, localVarOptionals)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestDownloadFile --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
		fmt.Printf("httpResponse  %d \n",httpResponse.ContentLength)	
		downName := "AsposeDown.pdf"					
        downPath := GetBaseTest().localTestDataFolder + "/" + downName
        downFile, err := os.OpenFile(downPath, os.O_WRONLY | os.O_CREATE, 0666)
        if err != nil {
			t.Error(err)
		}
        downFile.Write(dataRet)  
        downFile.Close()
	}
	
}

func TestDeleteNodes(t *testing.T) {
	name := "oaptest.pdf"
	objectaddressingpath :="//*[(@Type = 'Camera') or (@Name = 'light')]" 
    ctx := context.Background()
    localVarOptionals := new(DeleteNodesOpts)
    localVarOptionals.Folder = optional.NewString("3DTest") 
    localVarOptionals.Storage = optional.NewString("")
    
    _,httpResponse,err := GetBaseTest().ThreeDCloudAPI.DeleteNodes(ctx, name, objectaddressingpath, localVarOptionals)
    	
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestDeleteNodes --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestPostConvertByFormat(t *testing.T) {
	name := "oaptest.fbx"
	newformat := "pdf";
	newfilename := "oaptestGo.pdf";

  ctx := context.Background()
  localVarOptionals := new(PostConvertByFormatOpts)
  localVarOptionals.Folder = optional.NewString("3DTest") 
  localVarOptionals.IsOverwrite = optional.NewBool(true)      
  localVarOptionals.Storage = optional.NewString("")
    
    _,httpResponse,err := GetBaseTest().ThreeDCloudAPI.PostConvertByFormat(ctx, name, newformat, newfilename, localVarOptionals)
    	
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestPostConvertByFormat --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}
*/

 
func TestPostConvertByOpt(t *testing.T) {
	name := "Aspose.pdf"	
    newfilename := "saveasOptGo.pdf" 
       
    // -OK-
    var saveopt = new(PlySaveOption) 
    saveopt.SaveFormat = SaveFormat_PLY   
      
    ctx := context.Background()
    localVarOptionals := new(PostConvertByOptOpts)
    localVarOptionals.Folder = optional.NewString("3DTest") 
    localVarOptionals.IsOverwrite = optional.NewBool(true)      
    localVarOptionals.Storage = optional.NewString("")
    
    //var ptr unsafe.Pointer = unsafe.Pointer(saveopt)   	
    _,httpResponse,err := GetBaseTest().ThreeDCloudAPI.PostConvertByPlySaveOption(ctx, name, saveopt, newfilename, localVarOptionals)   	
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestPostConvertByOpt --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}


 
func TestPostModel(t *testing.T) {
    name := "3DTest/Aspose.pdf"
    
    var modeldata ModelData
    modeldata.Transform = new(Transform)
    modeldata.Transform.Translation = new(Vector3)
    modeldata.Transform.Translation.X = 100
    modeldata.Transform.Translation.Y = 100   
    modeldata.Transform.Translation.Z = 100     
    modeldata.Entity = new(Entity)  
    modeldata.Entity.Torus = new(Torus) 
    modeldata.Entity.Torus.Arc = 0 
    modeldata.Entity.Torus.Radius = 10
    modeldata.Entity.Torus.Tube = 10     
    modeldata.Entity.Torus.RadialSegments = 10   
    modeldata.Entity.Torus.TubularSegments = 10 
          
    ctx := context.Background()	    	    
    localVarOptionals := new(PostModelOpts)
    localVarOptionals.Newformat = optional.NewString("pdf") 
    localVarOptionals.Folder = optional.NewString("")      
    localVarOptionals.Storage = optional.NewString("")   	
	_,httpResponse,err := GetBaseTest().ThreeDCloudAPI.PostModel(ctx, name, modeldata, localVarOptionals)	
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d: TestPostConvertByOpt --> %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}	
}	
 
// to continue for more test cases