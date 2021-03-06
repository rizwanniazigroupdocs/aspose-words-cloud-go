/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="lists_test.go">
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

// Example of how to work with lists.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2010/api/models"
)

// Test for getting lists from document.
func Test_Lists_GetLists(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestGetLists.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetLists(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting list from document.
func Test_Lists_GetList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestGetList.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetList(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating list from document.
func Test_Lists_UpdateList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestUpdateList.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestListUpdate := models.ListUpdate{
        IsRestartAtEachSection: true,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateList(ctx, remoteFileName, requestListUpdate, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating list level from document.
func Test_Lists_UpdateListLevel(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestUpdateListLevel.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestListUpdate := models.ListLevelUpdate{
        Alignment: "Right",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateListLevel(ctx, remoteFileName, requestListUpdate, int32(1), int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for inserting list from document.
func Test_Lists_InsertList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestInsertList.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestListInsert := models.ListInsert{
        Template: "OutlineLegal",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertList(ctx, remoteFileName, requestListInsert, options)

    if err != nil {
        t.Error(err)
    }

}
