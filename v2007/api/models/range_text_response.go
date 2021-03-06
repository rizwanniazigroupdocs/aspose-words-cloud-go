/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="range_text_response.go">
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

package models

// This response should be returned by the service when handling:
// GET https://api.aspose.cloud/v4.0/words/Test.doc/range/{0}/{1}/.
type RangeTextResponse struct {
    // This response should be returned by the service when handling:
    // GET https://api.aspose.cloud/v4.0/words/Test.doc/range/{0}/{1}/.
    RequestId string `json:"RequestId,omitempty"`

    // This response should be returned by the service when handling:
    // GET https://api.aspose.cloud/v4.0/words/Test.doc/range/{0}/{1}/.
    Text string `json:"Text,omitempty"`
}

type IRangeTextResponse interface {
    IsRangeTextResponse() bool
}
func (RangeTextResponse) IsRangeTextResponse() bool {
    return true
}

func (RangeTextResponse) IsWordsResponse() bool {
    return true
}
