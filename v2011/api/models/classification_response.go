/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="classification_response.go">
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

// The REST response with data on multi-class text classification.
type ClassificationResponse struct {
    // The REST response with data on multi-class text classification.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with data on multi-class text classification.
    BestClassName *string `json:"BestClassName,omitempty"`

    // The REST response with data on multi-class text classification.
    BestClassProbability *float64 `json:"BestClassProbability,omitempty"`

    // The REST response with data on multi-class text classification.
    BestResults []ClassificationResult `json:"BestResults,omitempty"`
}

type IClassificationResponse interface {
    IsClassificationResponse() bool
}
func (ClassificationResponse) IsClassificationResponse() bool {
    return true
}

func (ClassificationResponse) IsWordsResponse() bool {
    return true
}