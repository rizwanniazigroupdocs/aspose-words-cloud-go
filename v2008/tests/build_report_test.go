/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="build_report_test.go">
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
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2008/api/models"
)

// Test for build report online.
func Test_BuildReport_BuildReportOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    reportingFolder := "DocumentActions/Reporting"
    localDocumentFile := "ReportTemplate.docx"
    localDataFile := ReadFile(t, reportingFolder + "/ReportData.json")

    requestReportEngineSettings := models.ReportEngineSettings{
        DataSourceType: "Json",
        DataSourceName: "persons",
    }

    options := map[string]interface{}{
    }
    _, err := client.WordsApi.BuildReportOnline(ctx, OpenFile(t, reportingFolder + "/" + localDocumentFile), localDataFile, requestReportEngineSettings, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for build report.
func Test_BuildReport_BuildReport(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Reporting"
    reportingFolder := "DocumentActions/Reporting"
    localDocumentFile := "ReportTemplate.docx"
    remoteFileName := "TestBuildReport.docx"
    localDataFile := ReadFile(t, reportingFolder + "/ReportData.json")

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(reportingFolder + "/" + localDocumentFile), remoteDataFolder + "/" + remoteFileName)

    requestReportEngineSettingsReportBuildOptions := []string{
        "AllowMissingMembers",
        "RemoveEmptyParagraphs",
    }
    requestReportEngineSettings := models.ReportEngineSettings{
        DataSourceType: "Json",
        ReportBuildOptions: requestReportEngineSettingsReportBuildOptions,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.BuildReport(ctx, remoteFileName, localDataFile, requestReportEngineSettings, options)

    if err != nil {
        t.Error(err)
    }
}
