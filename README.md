![](https://img.shields.io/badge/api-v4.0-lightgrey) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/aspose-words-cloud/aspose-words-cloud-go) [![GitHub license](https://img.shields.io/github/license/aspose-words-cloud/aspose-words-cloud-go)](https://github.com/aspose-words-cloud/aspose-words-cloud-go)

# Go SDK for Word Documents Processing in the Cloud
[Aspose.Wors Cloud SDK for Go](https://products.aspose.cloud/words/go) wraps Aspose.Words REST API so you could seamlessly integrate Microsoft Word® document generation, manipulation, conversion & inspection features into your own Go applications.
Feel free to explore the [Developer's Guide](https://docs.aspose.cloud/display/wordscloud/Developer+Guide) & [API Reference](https://apireference.aspose.cloud/words/) to know all about Aspose.Words Cloud API. 

## Document Processing Services in REST

- Convert between various document-related formats, including Word to PDF & vice versa.
- Mail merge and report generation in the Cloud.
- Split & merge Word documents.
- Access Word document metadata.
- Find and replace text.
- Add & remove watermarks and protection.
- Read & write access to Document Object Model.
- Get document information in JSON or XML representation.
- Convert an encrypted PDF document into Word document format.
- [Fetch statistical data](https://docs.aspose.cloud/display/wordscloud/Get+Document+Statistics) of a document.
- [Remove all macros](https://docs.aspose.cloud/display/wordscloud/Remove+all+Macros+from+Document) contained in a specific document.
- [Convert a document to desired file format](https://docs.aspose.cloud/display/wordscloud/Convert+Document+to+Destination+Format+with+Detailed+Settings+and+Save+Result+to+Storage) along with detailed settings.

Check out the [Developer's Guide](https://docs.aspose.cloud/display/wordscloud/Developer+Guide) to know more about Aspose Words Cloud REST API.


## Read & Write Document Formats

**Microsoft Word:** DOC, DOCX, RTF, DOT, DOTX, DOTM, FlatOPC (XML)
**OpenOffice:** ODT, OTT
**WordprocessingML:** XML
**Web:** HTML, MHTML, HtmlFixed
**Text:** TXT
**Fixed Layout:** PDF

## Save Document As

**Fixed Layout:** PDF/A, XPS, OpenXPS, PS
**Images:** JPEG, PNG, BMP, SVG, TIFF, EMF
**Others:** PCL

## Enhancements in Version 20.11

- In configuration json file appSid / appKey has been replaced to clientId / clientSecret.
- In Words API initialization methods clientId parameter precedes clientSecret parameter.
- Fixed the problem with object updating methods (they ignored parameters with default type values).

## Enhancements in Version 20.10

- Internal API changes.


## Enhancements in Version 20.9

- Added Batch API feature


## Enhancements in Version 20.8

- Added new api method (PUT '/words/{name}/compatibility/optimize') which is allows to optimize the document contents as well as default Aspose.Words behavior to a particular versions of MS Word
- Added 'ApplyBaseDocumentHeadersAndFootersToAppendingDocuments' option to 'DocumentEntryList' for AppendDocument API
- WithoutNodePath methods have been removed, pass null values instead


## Enhancements in Version 20.7

- Added 'Markdown' save format
- Added endpoint to update paragraph format without node path (PUT '/words/{name}/paragraphs/{index}/format')


## Getting Started with Aspose.Words Cloud SDK for Go

Firstly, create an account at [Aspose for Cloud](https://dashboard.aspose.cloud/#/apps) to get your application information and free quota to use the API. 

#### Install Aspose.Words-Cloud

From Visual Stuio Code:

	Add "github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api" and "github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models" in the import section of your code

From the command line:

	go get -v github.com/aspose-words-cloud/aspose-words-cloud-go/2007/api

The complete source code is available at [GitHub Repository](https://github.com/aspose-words-cloud/aspose-words-cloud-go).

### SDK Dependencies

The libraray doesn't uses any non-Google Golang packages.

## Sample usage via the SDK

The examples below show how your application have to initiate and get a text of the first paragraph using Aspose.Words-Cloud library:

Config.json file:
```csharp
{
	"ClientId": "your client id",
	"ClientSecret": "your client secret",
	"BaseUrl": "https://api.aspose.cloud"
} 
```
Go code:

```
	// Start README example

	// init words cloud api
	config, _ := models.NewConfiguration(configFilePath)
	wordsApi, ctx, _ := api.CreateWordsApi(config)

	// upload test.docx to a cloud
	// remote.docx is a name in the cloud
	file, _ := os.Open(localFilePath)
	wordsApi.UploadFile(ctx, file, remotePath, nil)

	// get a text for the first paragraph of the first section
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, _, _ := wordsApi.GetParagraphs(ctx, remoteName, options)

	fmt.Println(result.Paragraphs.ParagraphLinkList[0].Text)

	// End README example
```
| .NET | Java | PHP | Python | Ruby | Node.js | Android | Dart|Go|
|---|---|---|---|---|---|---|--|--|
| [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-dotnet) | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-java) | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-php) | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-python) | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-ruby)  | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-node) | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-android) | [GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-dart) |[GitHub](https://github.com/aspose-words-cloud/aspose-words-cloud-go) |
| [NuGet](https://www.nuget.org/packages/Aspose.Words-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-words-cloud) | [Composer](https://packagist.org/packages/aspose-cloud/aspose-words-cloud) | [PIP](https://pypi.org/project/aspose.words-cloud/) | [GEM](https://rubygems.org/gems/aspose_words_cloud)  | [NPM](https://www.npmjs.com/package/asposewordscloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-words-cloud) | [Pub.Dev](https://pub.dev/packages/aspose_words_cloud) | [Go.Dev](https://pkg.go.dev/github.com/aspose-words-cloud/aspose-words-cloud-go/) | 

[Product Page](https://products.aspose.cloud/words/go) | [Documentation](https://docs.aspose.cloud/display/wordscloud/Home) | [API Reference](https://apireference.aspose.cloud/words/) | [Code Samples](https://github.com/aspose-words-cloud/aspose-words-cloud-go) | [Blog](https://blog.aspose.cloud/category/words/) | [Free Support](https://forum.aspose.cloud/c/words) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
