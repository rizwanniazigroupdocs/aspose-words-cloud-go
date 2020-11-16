/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="available_fonts_response.go">
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

// The REST response with data on system, additional and custom fonts, available for document processing.
type AvailableFontsResponse struct {
    // The REST response with data on system, additional and custom fonts, available for document processing.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with data on system, additional and custom fonts, available for document processing.
    AdditionalFonts []FontInfo `json:"AdditionalFonts,omitempty"`

    // The REST response with data on system, additional and custom fonts, available for document processing.
    CustomFonts []FontInfo `json:"CustomFonts,omitempty"`

    // The REST response with data on system, additional and custom fonts, available for document processing.
    SystemFonts []FontInfo `json:"SystemFonts,omitempty"`
}

type IAvailableFontsResponse interface {
    IsAvailableFontsResponse() bool
}
func (AvailableFontsResponse) IsAvailableFontsResponse() bool {
    return true
}

func (AvailableFontsResponse) IsWordsResponse() bool {
    return true
}
