/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="execute_mail_merge_test.go">
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

// Example of how to perform mail merge.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

// Test for executing mail merge online.
func Test_ExecuteMailMerge_ExecuteMailMergeOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "SampleExecuteTemplate.docx"
    localDataFile := "SampleExecuteTemplateData.txt"


    options := map[string]interface{}{
    }
    _, err := client.WordsApi.ExecuteMailMergeOnline(ctx, OpenFile(t, mailMergeFolder + "/" + localDocumentFile), OpenFile(t, mailMergeFolder + "/" + localDataFile), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for executing mail merge.
func Test_ExecuteMailMerge_ExecuteMailMerge(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/MailMerge"
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "SampleExecuteTemplate.docx"
    remoteFileName := "TestExecuteMailMerge.docx"
    localDataFile := ReadFile(t, mailMergeFolder + "/SampleMailMergeTemplateData.txt")

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(mailMergeFolder + "/" + localDocumentFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "data": localDataFile,
        "folder": remoteDataFolder,
        "withRegions": false,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.ExecuteMailMerge(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Document, "Validate ExecuteMailMerge response.");
    assert.True(t, strings.HasPrefix(actual.Document.FileName, "TestExecuteMailMerge.docx"), "Validate ExecuteMailMerge response.");
}
