/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="base_test.go">
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

package api_test

import (
    "bytes"
    "context"
    "crypto/rand"
    "time"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "regexp"
    "runtime"
    "strings"
    "testing"

    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2009/api"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2009/api/models"
)

var remoteBaseTestDataFolder string = "Temp/SdkTests/Go/TestData"

var baseTestOutPath string = "TestOut/Go"

var commonTestFile string = GetLocalPath("Common", "test_multi_pages.docx")

func CreateRandomGuid() string {
    b := make([]byte, 16)
    rand.Read(b)
    uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return uuid
}

func GetLocalPath(folderPath string, filename string) string {
    _, currentFilePath, _, _ := runtime.Caller(0)
    currentFolder := filepath.Dir(currentFilePath)
    return filepath.Join(currentFolder, "testdata", folderPath, filename)
}

func GetLocalFile(filePath string) string {
    _, currentFilePath, _, _ := runtime.Caller(0)
    currentFolder := filepath.Dir(currentFilePath)
    return filepath.Join(currentFolder, "testdata", filePath)
}

func ReadFile(t *testing.T, path string) (result string) {
    data, err := ioutil.ReadFile(GetLocalFile(path))
    if err != nil {
        t.Error(err)
    }

    return string(data)
}

func OpenFile(t *testing.T, path string) (result *os.File) {
    data, err := os.Open(GetLocalFile(path))
    if err != nil {
        t.Error(err)
    }

    return data
}

func CreateTime(year int, month int, day int, hour int, min int, sec int) (result time.Time) {
    return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
}

func ReadConfiguration(t *testing.T) (cfg *models.Configuration) {
    configuration, err := models.NewConfiguration(GetConfigFilePath())

    if err != nil {
        t.Error(err)
    }

    return configuration
}

func GetConfigFilePath() (path string) {
    _, filename, _, _ := runtime.Caller(0)
    configFilePath := filepath.Join(filepath.Dir(filename), "../config.json")
    _, fileErr := os.Stat(configFilePath)
    if os.IsNotExist(fileErr) {
        configFilePath = filepath.Join(filepath.Dir(filename), "../../config.json")
    }
    return configFilePath
}

func PrepareTest(t *testing.T, config *models.Configuration) (apiClient *api.APIClient, ctx context.Context) {
    client, err := api.NewAPIClient(config)

    if err != nil {
        t.Error(err)
    }

    context, err := client.NewContextWithToken(nil)

    if err != nil {
        t.Error(err)
    }

    return client, context
}

func UploadFileToStorage(t *testing.T, fileName string, path string) (*api.APIClient, context.Context) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)

    UploadNextFileToStorage(t, ctx, client, fileName, path)

    return client, ctx
}

func UploadNextFileToStorage(t *testing.T, ctx context.Context, client *api.APIClient, fileName string, path string) {
    file, fileErr := os.Open(fileName)
    if fileErr != nil {
        t.Error(fileErr)
    }

    _, _, err := client.WordsApi.UploadFile(ctx, file, path, nil)
    if err != nil {
        t.Error(err)
    }
}

func TestDebugMode(t *testing.T) {
    config := ReadConfiguration(t)
    config.DebugMode = true
    client, ctx := PrepareTest(t, config)

    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    _, _, err := client.WordsApi.CreateDocument(ctx, map[string]interface{}{
        "fileName": "test.docx",
    })

    if err != nil {
        t.Error(err)
    }

    logData := buf.String()

    url := "PUT /v4.0/words/create?FileName=test.docx"
    success := " 200 OK"

    if !strings.Contains(logData, url) {
        t.Errorf("%s doesn't contain the string '%s'", logData, url)
    }

    if !strings.Contains(logData, success) {
        t.Errorf("%s doesn't contain the string '%s'", logData, success)
    }
}

func TestCoverage(t *testing.T) {
    _, currentFileName, _, _ := runtime.Caller(0)
    currentFolder := filepath.Dir(currentFileName)
    apiFileName := filepath.Join(currentFolder, "../api/words_api.go")
    api, err := ioutil.ReadFile(apiFileName)

    if err != nil {
        t.Error(err)
    }

    declareFuncRegex := regexp.MustCompile(`func \(a \*WordsApiService\) (\w+)`)
    allApiFuncNames := declareFuncRegex.FindAllSubmatch(api, -1)

    var exists = struct{}{}

    notTestedApiFuncNames := make(map[string]struct{})

    for _, name := range allApiFuncNames {
        notTestedApiFuncNames[string(name[1])] = exists
    }

    testFileNames, _ := filepath.Glob(filepath.Join(currentFolder, "*.go"))

    for _, testFileName := range testFileNames {
        tests, err := ioutil.ReadFile(testFileName)
        if err != nil {
            t.Error(err)
        }

        callFuncRegex := regexp.MustCompile(`\.WordsApi\.(\w+)`)
        testedlApiFuncNames := callFuncRegex.FindAllSubmatch(tests, -1)

        for _, name := range testedlApiFuncNames {
            delete(notTestedApiFuncNames, string(name[1]))
        }
    }

    if len(notTestedApiFuncNames) > 0 {
        report := "Not covered methods:\n"

        for name := range notTestedApiFuncNames {
            report = report + fmt.Sprintf("\t%s\n", name)
        }

        t.Error(report)
    }
}

func TestGetAvailableFonts(t *testing.T) {

    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)

    _, _, err := client.WordsApi.GetAvailableFonts(ctx, nil)

    if err != nil {
        t.Error(err)
    }
}

func TestResetCache(t *testing.T) {

    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)

    _, err := client.WordsApi.ResetCache(ctx)

    if err != nil {
        t.Error(err)
    }
}

func TestCreateWordsApi(t *testing.T) {
    config := ReadConfiguration(t)
    wordsApi, ctx, _ := api.CreateWordsApi(config)

    _, err := wordsApi.ResetCache(ctx)

    if err != nil {
        t.Error(err)
    }
}
