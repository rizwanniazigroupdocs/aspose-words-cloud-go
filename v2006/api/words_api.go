/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

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
*/

package api

import (
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"os"
	"io/ioutil"
	"encoding/json"
    "fmt"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2006/api/models"
	"errors"
)

// Linger please
var (
	_ context.Context
)

type WordsApiService service

/* WordsApiService Accepts all revisions in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RevisionsModificationResponse*/
func (a *WordsApiService) AcceptAllRevisions(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RevisionsModificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.RevisionsModificationResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/revisions/acceptAll"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Appends documents to original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) AppendDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/appendDocument"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Apply a style to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return WordsResponse*/
func (a *WordsApiService) ApplyStyleToDocumentElement(ctx context.Context, localVarOptionals map[string]interface{}) ( models.WordsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.WordsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{styledNodePath}/style"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes document \&quot;build report\&quot; operation.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) BuildReport(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/buildReport"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes document \&quot;build report\&quot; online operation.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) BuildReportOnline(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/buildReport"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Classifies raw text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ClassificationResponse*/
func (a *WordsApiService) Classify(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ClassificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ClassificationResponse
	)

	// create path and map variables
	localVarPath := "/words/classify"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Classifies document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ClassificationResponse*/
func (a *WordsApiService) ClassifyDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ClassificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ClassificationResponse
	)

	// create path and map variables
	localVarPath := "/words/{documentName}/classify"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Compares document with original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) CompareDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/compareDocument"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts document from the request&#39;s content to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) ConvertDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/convert"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Copy file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) CopyFile(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/file/copy/{srcPath}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Copy folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) CopyFolder(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/folder/copy/{srcPath}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Copy and insert a new style to the document, returns a copied style.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StyleResponse*/
func (a *WordsApiService) CopyStyle(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StyleResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.StyleResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/styles/copy"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Creates new document. Document is created with format which is recognized from file extensions. Supported extensions: \&quot;.doc\&quot;, \&quot;.docx\&quot;, \&quot;.docm\&quot;, \&quot;.dot\&quot;, \&quot;.dotm\&quot;, \&quot;.dotx\&quot;, \&quot;.flatopc\&quot;, \&quot;.fopc\&quot;, \&quot;.flatopc_macro\&quot;, \&quot;.fopc_macro\&quot;, \&quot;.flatopc_template\&quot;, \&quot;.fopc_template\&quot;, \&quot;.flatopc_template_macro\&quot;, \&quot;.fopc_template_macro\&quot;, \&quot;.wordml\&quot;, \&quot;.wml\&quot;, \&quot;.rtf\&quot;.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) CreateDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/create"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Create the folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) CreateFolder(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/folder/{path}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Adds new or update existing document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentPropertyResponse*/
func (a *WordsApiService) CreateOrUpdateDocumentProperty(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentPropertyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/documentProperties/{propertyName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Remove all tab stops.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) DeleteAllParagraphTabStops(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstops"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Remove all tab stops.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) DeleteAllParagraphTabStopsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/tabstops"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Resets border properties to default values.             
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BorderResponse*/
func (a *WordsApiService) DeleteBorder(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BorderResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.BorderResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/borders/{borderType}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Resets borders properties to default values.             
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BordersResponse*/
func (a *WordsApiService) DeleteBorders(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BordersResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.BordersResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/borders"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes comment from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteComment(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/comments/{commentIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteDocumentProperty(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/documentProperties/{propertyName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes drawing object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteDrawingObject(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes drawing object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteDrawingObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteField(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/fields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/fields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes fields from section paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFields(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/fields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes fields from section paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFieldsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/fields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Delete file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFile(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/file/{path}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Delete folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFolder(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/folder/{path}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes footnote from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFootnote(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/footnotes/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes footnote from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFootnoteWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/footnotes/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes form field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFormField(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/formfields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes form field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteFormFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/formfields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes header/footer from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteHeaderFooter(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{sectionPath}/headersfooters/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes document headers and footers.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteHeadersFooters(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{sectionPath}/headersfooters"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes macros from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteMacros(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/macros"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes OfficeMath object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteOfficeMathObject(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes OfficeMath object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteOfficeMathObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/OfficeMathObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes paragraph from section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteParagraph(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Delete paragraph list format, returns updated list format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphListFormatResponse*/
func (a *WordsApiService) DeleteParagraphListFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphListFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.ParagraphListFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/listFormat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Delete paragraph list format, returns updated list format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphListFormatResponse*/
func (a *WordsApiService) DeleteParagraphListFormatWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphListFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.ParagraphListFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/listFormat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Remove the i-th tab stop.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) DeleteParagraphTabStop(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstop"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Remove the i-th tab stop.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) DeleteParagraphTabStopWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/tabstop"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes paragraph from section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteParagraphWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes run from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteRun(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes section from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteSection(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/sections/{sectionIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteTable(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteTableCell(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tableRowPath}/cells/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteTableRow(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tablePath}/rows/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) DeleteTableWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes watermark (for deleting last watermark from the document).
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) DeleteWatermark(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/watermarks/deleteLast"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Download file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) DownloadFile(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/storage/file/{path}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes document mail merge operation.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) ExecuteMailMerge(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/MailMerge"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes document mail merge online.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) ExecuteMailMergeOnline(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/MailMerge"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets the list of fonts, available for document processing.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return AvailableFontsResponse*/
func (a *WordsApiService) GetAvailableFonts(ctx context.Context, localVarOptionals map[string]interface{}) ( models.AvailableFontsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.AvailableFontsResponse
	)

	// create path and map variables
	localVarPath := "/words/fonts/available"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document bookmark data by its name.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BookmarkResponse*/
func (a *WordsApiService) GetBookmarkByName(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BookmarkResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.BookmarkResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/bookmarks/{bookmarkName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document bookmarks common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BookmarksResponse*/
func (a *WordsApiService) GetBookmarks(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BookmarksResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.BookmarksResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/bookmarks"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a border.
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BorderResponse*/
func (a *WordsApiService) GetBorder(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BorderResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.BorderResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/borders/{borderType}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a collection of borders.
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BordersResponse*/
func (a *WordsApiService) GetBorders(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BordersResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.BordersResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/borders"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets comment from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return CommentResponse*/
func (a *WordsApiService) GetComment(ctx context.Context, localVarOptionals map[string]interface{}) ( models.CommentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.CommentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/comments/{commentIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets comments from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return CommentsResponse*/
func (a *WordsApiService) GetComments(ctx context.Context, localVarOptionals map[string]interface{}) ( models.CommentsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.CommentsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/comments"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) GetDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{documentName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing object common info by its index or convert to format specified.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndex(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing object common info by its index or convert to format specified.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndexWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads drawing object image data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) GetDocumentDrawingObjectImageData(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}/imageData"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads drawing object image data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) GetDocumentDrawingObjectImageDataWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects/{index}/imageData"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets drawing object OLE data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) GetDocumentDrawingObjectOleData(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}/oleData"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets drawing object OLE data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) GetDocumentDrawingObjectOleDataWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects/{index}/oleData"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing objects common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjects(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DrawingObjectsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing objects common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DrawingObjectsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document field names.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNames(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldNamesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FieldNamesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/mailMerge/FieldNames"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document field names.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNamesOnline(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldNamesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FieldNamesResponse
	)

	// create path and map variables
	localVarPath := "/words/mailMerge/FieldNames"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document hyperlink by its index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return HyperlinkResponse*/
func (a *WordsApiService) GetDocumentHyperlinkByIndex(ctx context.Context, localVarOptionals map[string]interface{}) ( models.HyperlinkResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.HyperlinkResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/hyperlinks/{hyperlinkIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document hyperlinks common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return HyperlinksResponse*/
func (a *WordsApiService) GetDocumentHyperlinks(ctx context.Context, localVarOptionals map[string]interface{}) ( models.HyperlinksResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.HyperlinksResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/hyperlinks"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document properties info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentPropertiesResponse*/
func (a *WordsApiService) GetDocumentProperties(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentPropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/documentProperties"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document property info by the property name.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentPropertyResponse*/
func (a *WordsApiService) GetDocumentProperty(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentPropertyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/documentProperties/{propertyName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document protection common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ProtectionDataResponse*/
func (a *WordsApiService) GetDocumentProtection(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ProtectionDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ProtectionDataResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/protection"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document statistics.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StatDataResponse*/
func (a *WordsApiService) GetDocumentStatistics(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StatDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.StatDataResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/statistics"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Exports the document into the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) GetDocumentWithFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldResponse*/
func (a *WordsApiService) GetField(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/fields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldResponse*/
func (a *WordsApiService) GetFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/fields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldsResponse*/
func (a *WordsApiService) GetFields(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FieldsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/fields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldsResponse*/
func (a *WordsApiService) GetFieldsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FieldsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/fields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get all files and folders within a folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FilesList*/
func (a *WordsApiService) GetFilesList(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FilesList,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FilesList
	)

	// create path and map variables
	localVarPath := "/words/storage/folder/{path}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads footnote by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnoteResponse*/
func (a *WordsApiService) GetFootnote(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/footnotes/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads footnote by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnoteResponse*/
func (a *WordsApiService) GetFootnoteWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/footnotes/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets footnotes from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnotesResponse*/
func (a *WordsApiService) GetFootnotes(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnotesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FootnotesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/footnotes"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets footnotes from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnotesResponse*/
func (a *WordsApiService) GetFootnotesWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnotesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FootnotesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/footnotes"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns representation of an one of the form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldResponse*/
func (a *WordsApiService) GetFormField(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/formfields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns representation of an one of the form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldResponse*/
func (a *WordsApiService) GetFormFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/formfields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets form fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldsResponse*/
func (a *WordsApiService) GetFormFields(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FormFieldsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/formfields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets form fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldsResponse*/
func (a *WordsApiService) GetFormFieldsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FormFieldsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/formfields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a header/footer from the document by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooter(ctx context.Context, localVarOptionals map[string]interface{}) ( models.HeaderFooterResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.HeaderFooterResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/headersfooters/{headerFooterIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a header/footer from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooterOfSection(ctx context.Context, localVarOptionals map[string]interface{}) ( models.HeaderFooterResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.HeaderFooterResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/sections/{sectionIndex}/headersfooters/{headerFooterIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of header/footers from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return HeaderFootersResponse*/
func (a *WordsApiService) GetHeaderFooters(ctx context.Context, localVarOptionals map[string]interface{}) ( models.HeaderFootersResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.HeaderFootersResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{sectionPath}/headersfooters"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents one of the lists contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ListResponse*/
func (a *WordsApiService) GetList(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ListResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/lists/{listId}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of lists that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ListsResponse*/
func (a *WordsApiService) GetLists(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ListsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ListsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/lists"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads OfficeMath object by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObject(ctx context.Context, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.OfficeMathObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads OfficeMath object by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.OfficeMathObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/OfficeMathObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets OfficeMath objects from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjects(ctx context.Context, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.OfficeMathObjectsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets OfficeMath objects from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjectsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.OfficeMathObjectsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/OfficeMathObjects"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents one of the paragraphs contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphResponse*/
func (a *WordsApiService) GetParagraph(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Represents all the formatting for a paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/format"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Represents all the formatting for a paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormatWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/format"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Represents list format for a paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphListFormatResponse*/
func (a *WordsApiService) GetParagraphListFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphListFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphListFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/listFormat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Represents list format for a paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphListFormatResponse*/
func (a *WordsApiService) GetParagraphListFormatWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphListFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphListFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/listFormat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get all tab stops for the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) GetParagraphTabStops(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstops"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get all tab stops for the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) GetParagraphTabStopsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/tabstops"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents one of the paragraphs contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphResponse*/
func (a *WordsApiService) GetParagraphWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of paragraphs that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphs(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of paragraphs that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphsWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.ParagraphLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets the text from the range.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RangeTextResponse*/
func (a *WordsApiService) GetRangeText(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RangeTextResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.RangeTextResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents run of text contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RunResponse*/
func (a *WordsApiService) GetRun(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RunResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.RunResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents font of run.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FontResponse*/
func (a *WordsApiService) GetRunFont(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FontResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.FontResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs/{index}/font"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents collection of runs in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RunsResponse*/
func (a *WordsApiService) GetRuns(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RunsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.RunsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets document section by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SectionResponse*/
func (a *WordsApiService) GetSection(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.SectionResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/sections/{sectionIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets page setup of section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SectionPageSetupResponse*/
func (a *WordsApiService) GetSectionPageSetup(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SectionPageSetupResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.SectionPageSetupResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/sections/{sectionIndex}/pageSetup"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of sections that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SectionLinkCollectionResponse*/
func (a *WordsApiService) GetSections(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SectionLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.SectionLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/sections"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents one of the styles contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StyleResponse*/
func (a *WordsApiService) GetStyle(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StyleResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.StyleResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/styles/{styleName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets a style from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StyleResponse*/
func (a *WordsApiService) GetStyleFromDocumentElement(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StyleResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.StyleResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{styledNodePath}/style"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of styles contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StylesResponse*/
func (a *WordsApiService) GetStyles(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StylesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.StylesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/styles"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableResponse*/
func (a *WordsApiService) GetTable(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableCellResponse*/
func (a *WordsApiService) GetTableCell(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableCellResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableCellResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tableRowPath}/cells/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table cell format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableCellFormatResponse*/
func (a *WordsApiService) GetTableCellFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableCellFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableCellFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tableRowPath}/cells/{index}/cellformat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TablePropertiesResponse*/
func (a *WordsApiService) GetTableProperties(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables/{index}/properties"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TablePropertiesResponse*/
func (a *WordsApiService) GetTablePropertiesWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables/{index}/properties"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableRowResponse*/
func (a *WordsApiService) GetTableRow(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableRowResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableRowResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tablePath}/rows/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table row format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableRowFormatResponse*/
func (a *WordsApiService) GetTableRowFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableRowFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableRowFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tablePath}/rows/{index}/rowformat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableResponse*/
func (a *WordsApiService) GetTableWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of tables that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableLinkCollectionResponse*/
func (a *WordsApiService) GetTables(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of tables that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableLinkCollectionResponse*/
func (a *WordsApiService) GetTablesWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.TableLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds comment to document, returns inserted comment data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return CommentResponse*/
func (a *WordsApiService) InsertComment(ctx context.Context, localVarOptionals map[string]interface{}) ( models.CommentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.CommentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/comments"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds drawing object to document, returns added  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectResponse*/
func (a *WordsApiService) InsertDrawingObject(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds drawing object to document, returns added  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectResponse*/
func (a *WordsApiService) InsertDrawingObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds field to document, returns inserted field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldResponse*/
func (a *WordsApiService) InsertField(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/fields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds field to document, returns inserted field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldResponse*/
func (a *WordsApiService) InsertFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/fields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds footnote to document, returns added footnote&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnoteResponse*/
func (a *WordsApiService) InsertFootnote(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/footnotes"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds footnote to document, returns added footnote&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnoteResponse*/
func (a *WordsApiService) InsertFootnoteWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/footnotes"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds form field to paragraph, returns added form field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldResponse*/
func (a *WordsApiService) InsertFormField(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/formfields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds form field to paragraph, returns added form field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldResponse*/
func (a *WordsApiService) InsertFormFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/formfields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts to document header or footer.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return HeaderFooterResponse*/
func (a *WordsApiService) InsertHeaderFooter(ctx context.Context, localVarOptionals map[string]interface{}) ( models.HeaderFooterResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.HeaderFooterResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{sectionPath}/headersfooters"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds list to document, returns added list&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ListResponse*/
func (a *WordsApiService) InsertList(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.ListResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/lists"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Insert or resplace tab stop if a tab stop with the position exists.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) InsertOrUpdateParagraphTabStop(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstops"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Insert or resplace tab stop if a tab stop with the position exists.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TabStopsResponse*/
func (a *WordsApiService) InsertOrUpdateParagraphTabStopWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TabStopsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.TabStopsResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/tabstops"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts document page numbers.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) InsertPageNumbers(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/PageNumbers"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds paragraph to document, returns added paragraph&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphResponse*/
func (a *WordsApiService) InsertParagraph(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds paragraph to document, returns added paragraph&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphResponse*/
func (a *WordsApiService) InsertParagraphWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds run to document, returns added paragraph&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RunResponse*/
func (a *WordsApiService) InsertRun(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RunResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.RunResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds a style to the document, returns an added style.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StyleResponse*/
func (a *WordsApiService) InsertStyle(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StyleResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.StyleResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/styles/insert"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table to document, returns added table&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableResponse*/
func (a *WordsApiService) InsertTable(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table cell to table, returns added cell&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableCellResponse*/
func (a *WordsApiService) InsertTableCell(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableCellResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.TableCellResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tableRowPath}/cells"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table row to table, returns added row&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableRowResponse*/
func (a *WordsApiService) InsertTableRow(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableRowResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.TableRowResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tablePath}/rows"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table to document, returns added table&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableResponse*/
func (a *WordsApiService) InsertTableWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts document watermark image.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) InsertWatermarkImage(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/watermarks/images"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts document watermark text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) InsertWatermarkText(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/watermarks/texts"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Loads new document from web into the file with any supported format of data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SaveResponse*/
func (a *WordsApiService) LoadWebDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SaveResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.SaveResponse
	)

	// create path and map variables
	localVarPath := "/words/loadWebDocument"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Move file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) MoveFile(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/file/move/{srcPath}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Move folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return */
func (a *WordsApiService) MoveFolder(ctx context.Context, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/storage/folder/move/{srcPath}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Protects document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ProtectionDataResponse*/
func (a *WordsApiService) ProtectDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ProtectionDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ProtectionDataResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/protection"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Rejects all revisions in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RevisionsModificationResponse*/
func (a *WordsApiService) RejectAllRevisions(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RevisionsModificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.RevisionsModificationResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/revisions/rejectAll"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes the range from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) RemoveRange(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders drawing object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderDrawingObject(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders drawing object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderDrawingObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders math object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderMathObject(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders math object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderMathObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/OfficeMathObjects/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders page to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderPage(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/pages/{pageIndex}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders paragraph to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderParagraph(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders paragraph to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderParagraphWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders table to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderTable(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders table to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return string*/
func (a *WordsApiService) RenderTableWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.string
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables/{index}/render"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Replaces document text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ReplaceTextResponse*/
func (a *WordsApiService) ReplaceText(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ReplaceTextResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ReplaceTextResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/replaceText"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Replaces the content in the range.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) ReplaceWithText(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Resets font&#39;s cache.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (a *WordsApiService) ResetCache(ctx context.Context) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := "/words/fonts/cache"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Converts document to destination format with detailed settings and saves result to storage.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SaveResponse*/
func (a *WordsApiService) SaveAs(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SaveResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.SaveResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/saveAs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Saves the selected range as a new document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) SaveAsRange(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}/SaveAs"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts document to tiff with detailed settings and saves result to storage.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SaveResponse*/
func (a *WordsApiService) SaveAsTiff(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SaveResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.SaveResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/saveAs/tiff"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Searches text in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SearchResponse*/
func (a *WordsApiService) Search(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SearchResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		successPayload models.SearchResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/search"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Splits document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SplitDocumentResponse*/
func (a *WordsApiService) SplitDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SplitDocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.SplitDocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/split"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Unprotects document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ProtectionDataResponse*/
func (a *WordsApiService) UnprotectDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ProtectionDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		successPayload models.ProtectionDataResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/protection"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates document bookmark.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BookmarkResponse*/
func (a *WordsApiService) UpdateBookmark(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BookmarkResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.BookmarkResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/bookmarks/{bookmarkName}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates border properties.             
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return BorderResponse*/
func (a *WordsApiService) UpdateBorder(ctx context.Context, localVarOptionals map[string]interface{}) ( models.BorderResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.BorderResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/borders/{borderType}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the comment, returns updated comment data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return CommentResponse*/
func (a *WordsApiService) UpdateComment(ctx context.Context, localVarOptionals map[string]interface{}) ( models.CommentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.CommentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/comments/{commentIndex}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates drawing object, returns updated  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectResponse*/
func (a *WordsApiService) UpdateDrawingObject(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates drawing object, returns updated  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DrawingObjectResponse*/
func (a *WordsApiService) UpdateDrawingObjectWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/drawingObjects/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates field&#39;s properties, returns updated field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FieldResponse*/
func (a *WordsApiService) UpdateField(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/fields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates (reevaluate) fields in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return DocumentResponse*/
func (a *WordsApiService) UpdateFields(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/updateFields"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates footnote&#39;s properties, returns updated run&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnoteResponse*/
func (a *WordsApiService) UpdateFootnote(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/footnotes/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates footnote&#39;s properties, returns updated run&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FootnoteResponse*/
func (a *WordsApiService) UpdateFootnoteWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/footnotes/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates properties of form field, returns updated form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldResponse*/
func (a *WordsApiService) UpdateFormField(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/formfields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates properties of form field, returns updated form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FormFieldResponse*/
func (a *WordsApiService) UpdateFormFieldWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/formfields/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates list properties, returns updated list.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ListResponse*/
func (a *WordsApiService) UpdateList(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ListResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/lists/{listId}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates list level in document list, returns updated list.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ListResponse*/
func (a *WordsApiService) UpdateListLevel(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ListResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/lists/{listId}/listLevels/{listLevel}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates paragraph format properties, returns updated format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphFormatResponse*/
func (a *WordsApiService) UpdateParagraphFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/format"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates paragraph format properties, returns updated format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphFormatResponse*/
func (a *WordsApiService) UpdateParagraphFormatWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/format"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates paragraph list format properties, returns updated list format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphListFormatResponse*/
func (a *WordsApiService) UpdateParagraphListFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphListFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ParagraphListFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/listFormat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates paragraph list format properties, returns updated list format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return ParagraphListFormatResponse*/
func (a *WordsApiService) UpdateParagraphListFormatWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.ParagraphListFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.ParagraphListFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/paragraphs/{index}/listFormat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates run&#39;s properties, returns updated run&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return RunResponse*/
func (a *WordsApiService) UpdateRun(ctx context.Context, localVarOptionals map[string]interface{}) ( models.RunResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.RunResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs/{index}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates font properties, returns updated font data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FontResponse*/
func (a *WordsApiService) UpdateRunFont(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FontResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FontResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{paragraphPath}/runs/{index}/font"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates page setup of section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return SectionPageSetupResponse*/
func (a *WordsApiService) UpdateSectionPageSetup(ctx context.Context, localVarOptionals map[string]interface{}) ( models.SectionPageSetupResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.SectionPageSetupResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/sections/{sectionIndex}/pageSetup"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates style properties, returns an updated style.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return StyleResponse*/
func (a *WordsApiService) UpdateStyle(ctx context.Context, localVarOptionals map[string]interface{}) ( models.StyleResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.StyleResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/styles/{styleName}/update"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table cell format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableCellFormatResponse*/
func (a *WordsApiService) UpdateTableCellFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableCellFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.TableCellFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tableRowPath}/cells/{index}/cellformat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TablePropertiesResponse*/
func (a *WordsApiService) UpdateTableProperties(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{nodePath}/tables/{index}/properties"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TablePropertiesResponse*/
func (a *WordsApiService) UpdateTablePropertiesWithoutNodePath(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/tables/{index}/properties"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table row format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return TableRowFormatResponse*/
func (a *WordsApiService) UpdateTableRowFormat(ctx context.Context, localVarOptionals map[string]interface{}) ( models.TableRowFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.TableRowFormatResponse
	)

	// create path and map variables
	localVarPath := "/words/{name}/{tablePath}/rows/{index}/rowformat"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Upload file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @return FilesUploadResult*/
func (a *WordsApiService) UploadFile(ctx context.Context, localVarOptionals map[string]interface{}) ( models.FilesUploadResult,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		successPayload models.FilesUploadResult
	)

	// create path and map variables
	localVarPath := "/words/storage/file/{path}"
	
	localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
	localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return successPayload, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
