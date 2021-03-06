/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_test.go">
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

// Example of how to get comments from document.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2010/api/models"
)

// Test for getting comment by specified comment's index.
func Test_Comment_GetComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetComment(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting all comments from document.
func Test_Comment_GetComments(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetComments.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetComments(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for adding comment.
func Test_Comment_InsertComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsertComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCommentRangeStartNode := models.NodeLink{
        NodeId: "0.3.0.3",
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: &requestCommentRangeStartNode,
        Offset: int32(0),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: "0.3.0.3",
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: &requestCommentRangeEndNode,
        Offset: int32(0),
    }
    requestComment := models.CommentInsert{
        RangeStart: &requestCommentRangeStart,
        RangeEnd: &requestCommentRangeEnd,
        Initial: "IA",
        Author: "Imran Anwar",
        Text: "A new Comment",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertComment(ctx, remoteFileName, requestComment, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating comment.
func Test_Comment_UpdateComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCommentRangeStartNode := models.NodeLink{
        NodeId: "0.3.0",
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: &requestCommentRangeStartNode,
        Offset: int32(0),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: "0.3.0",
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: &requestCommentRangeEndNode,
        Offset: int32(0),
    }
    requestComment := models.CommentUpdate{
        RangeStart: &requestCommentRangeStart,
        RangeEnd: &requestCommentRangeEnd,
        Initial: "IA",
        Author: "Imran Anwar",
        Text: "A new Comment",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateComment(ctx, remoteFileName, int32(0), requestComment, options)

    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteComment.
func Test_Comment_DeleteComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, err := client.WordsApi.DeleteComment(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}
