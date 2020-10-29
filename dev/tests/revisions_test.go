/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="revisions_test.go">
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

// Example of how to accept all revisions in document.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

// Test for accepting revisions in document.
func Test_Revisions_AcceptAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Revisions"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestAcceptAllRevisions.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.AcceptAllRevisions(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Result, "Validate AcceptAllRevisions response.");
    assert.NotNil(t, actual.Result.Dest, "Validate AcceptAllRevisions response.");
    assert.Equal(t, "TestOut/NET/TestAcceptAllRevisions.docx", actual.Result.Dest.Href, "Validate AcceptAllRevisions response.");
}

// Test for rejecting revisions in document.
func Test_Revisions_RejectAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Revisions"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestRejectAllRevisions.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.RejectAllRevisions(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Result, "Validate RejectAllRevisions response.");
    assert.NotNil(t, actual.Result.Dest, "Validate RejectAllRevisions response.");
    assert.Equal(t, "TestOut/NET/TestRejectAllRevisions.docx", actual.Result.Dest.Href, "Validate RejectAllRevisions response.");
}
