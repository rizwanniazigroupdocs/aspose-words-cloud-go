/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="watermark_test.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

// Example of how to work with watermarks.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for adding watermark image.
func Test_Watermark_InsertWatermarkImage(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Watermark"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsertWatermarkImage.docx"
    remoteImagePath := remoteDataFolder + "/TestInsertWatermarkImage.png"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)
    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/aspose-cloud.png"), remoteImagePath)


    options := map[string]interface{}{
        "imageFile": nil,
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
        "image": remoteImagePath,
    }
    actual, _, err := client.WordsApi.InsertWatermarkImage(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Document, "Validate InsertWatermarkImage response.");
    assert.True(t, strings.HasPrefix(actual.Document.FileName, "TestInsertWatermarkImage.docx"), "Validate InsertWatermarkImage response.");
}

// Test for adding watermark text.
func Test_Watermark_InsertWatermarkText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Watermark"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsertWatermarkText.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestWatermarkText := models.WatermarkText{
        Text: "This is the text",
        RotationAngle: 90,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.InsertWatermarkText(ctx, remoteFileName, requestWatermarkText, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Document, "Validate InsertWatermarkText response.");
    assert.True(t, strings.HasPrefix(actual.Document.FileName, "TestInsertWatermarkText.docx"), "Validate InsertWatermarkText response.");
}

// Test for deleting watermark.
func Test_Watermark_DeleteWatermark(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Watermark"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteWatermark.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.DeleteWatermark(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Document, "Validate DeleteWatermark response.");
    assert.True(t, strings.HasPrefix(actual.Document.FileName, "TestDeleteWatermark.docx"), "Validate DeleteWatermark response.");
}
