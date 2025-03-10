# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.28.2
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://numbers.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*RegulatoryComplianceBundlesApi* | [**CreateBundle**](docs/RegulatoryComplianceBundlesApi.md#createbundle) | **Post** /v2/RegulatoryCompliance/Bundles | 
*RegulatoryComplianceBundlesApi* | [**DeleteBundle**](docs/RegulatoryComplianceBundlesApi.md#deletebundle) | **Delete** /v2/RegulatoryCompliance/Bundles/{Sid} | 
*RegulatoryComplianceBundlesApi* | [**FetchBundle**](docs/RegulatoryComplianceBundlesApi.md#fetchbundle) | **Get** /v2/RegulatoryCompliance/Bundles/{Sid} | 
*RegulatoryComplianceBundlesApi* | [**ListBundle**](docs/RegulatoryComplianceBundlesApi.md#listbundle) | **Get** /v2/RegulatoryCompliance/Bundles | 
*RegulatoryComplianceBundlesApi* | [**UpdateBundle**](docs/RegulatoryComplianceBundlesApi.md#updatebundle) | **Post** /v2/RegulatoryCompliance/Bundles/{Sid} | 
*RegulatoryComplianceBundlesCopiesApi* | [**CreateBundleCopy**](docs/RegulatoryComplianceBundlesCopiesApi.md#createbundlecopy) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Copies | 
*RegulatoryComplianceBundlesCopiesApi* | [**ListBundleCopy**](docs/RegulatoryComplianceBundlesCopiesApi.md#listbundlecopy) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Copies | 
*RegulatoryComplianceBundlesEvaluationsApi* | [**CreateEvaluation**](docs/RegulatoryComplianceBundlesEvaluationsApi.md#createevaluation) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
*RegulatoryComplianceBundlesEvaluationsApi* | [**FetchEvaluation**](docs/RegulatoryComplianceBundlesEvaluationsApi.md#fetchevaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid} | 
*RegulatoryComplianceBundlesEvaluationsApi* | [**ListEvaluation**](docs/RegulatoryComplianceBundlesEvaluationsApi.md#listevaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
*RegulatoryComplianceBundlesItemAssignmentsApi* | [**CreateItemAssignment**](docs/RegulatoryComplianceBundlesItemAssignmentsApi.md#createitemassignment) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
*RegulatoryComplianceBundlesItemAssignmentsApi* | [**DeleteItemAssignment**](docs/RegulatoryComplianceBundlesItemAssignmentsApi.md#deleteitemassignment) | **Delete** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
*RegulatoryComplianceBundlesItemAssignmentsApi* | [**FetchItemAssignment**](docs/RegulatoryComplianceBundlesItemAssignmentsApi.md#fetchitemassignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
*RegulatoryComplianceBundlesItemAssignmentsApi* | [**ListItemAssignment**](docs/RegulatoryComplianceBundlesItemAssignmentsApi.md#listitemassignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
*RegulatoryComplianceBundlesReplaceItemsApi* | [**CreateReplaceItems**](docs/RegulatoryComplianceBundlesReplaceItemsApi.md#createreplaceitems) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ReplaceItems | 
*RegulatoryComplianceEndUserTypesApi* | [**FetchEndUserType**](docs/RegulatoryComplianceEndUserTypesApi.md#fetchendusertype) | **Get** /v2/RegulatoryCompliance/EndUserTypes/{Sid} | 
*RegulatoryComplianceEndUserTypesApi* | [**ListEndUserType**](docs/RegulatoryComplianceEndUserTypesApi.md#listendusertype) | **Get** /v2/RegulatoryCompliance/EndUserTypes | 
*RegulatoryComplianceEndUsersApi* | [**CreateEndUser**](docs/RegulatoryComplianceEndUsersApi.md#createenduser) | **Post** /v2/RegulatoryCompliance/EndUsers | 
*RegulatoryComplianceEndUsersApi* | [**DeleteEndUser**](docs/RegulatoryComplianceEndUsersApi.md#deleteenduser) | **Delete** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
*RegulatoryComplianceEndUsersApi* | [**FetchEndUser**](docs/RegulatoryComplianceEndUsersApi.md#fetchenduser) | **Get** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
*RegulatoryComplianceEndUsersApi* | [**ListEndUser**](docs/RegulatoryComplianceEndUsersApi.md#listenduser) | **Get** /v2/RegulatoryCompliance/EndUsers | 
*RegulatoryComplianceEndUsersApi* | [**UpdateEndUser**](docs/RegulatoryComplianceEndUsersApi.md#updateenduser) | **Post** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
*RegulatoryComplianceRegulationsApi* | [**FetchRegulation**](docs/RegulatoryComplianceRegulationsApi.md#fetchregulation) | **Get** /v2/RegulatoryCompliance/Regulations/{Sid} | 
*RegulatoryComplianceRegulationsApi* | [**ListRegulation**](docs/RegulatoryComplianceRegulationsApi.md#listregulation) | **Get** /v2/RegulatoryCompliance/Regulations | 
*RegulatoryComplianceSupportingDocumentTypesApi* | [**FetchSupportingDocumentType**](docs/RegulatoryComplianceSupportingDocumentTypesApi.md#fetchsupportingdocumenttype) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid} | 
*RegulatoryComplianceSupportingDocumentTypesApi* | [**ListSupportingDocumentType**](docs/RegulatoryComplianceSupportingDocumentTypesApi.md#listsupportingdocumenttype) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes | 
*RegulatoryComplianceSupportingDocumentsApi* | [**CreateSupportingDocument**](docs/RegulatoryComplianceSupportingDocumentsApi.md#createsupportingdocument) | **Post** /v2/RegulatoryCompliance/SupportingDocuments | 
*RegulatoryComplianceSupportingDocumentsApi* | [**DeleteSupportingDocument**](docs/RegulatoryComplianceSupportingDocumentsApi.md#deletesupportingdocument) | **Delete** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 
*RegulatoryComplianceSupportingDocumentsApi* | [**FetchSupportingDocument**](docs/RegulatoryComplianceSupportingDocumentsApi.md#fetchsupportingdocument) | **Get** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 
*RegulatoryComplianceSupportingDocumentsApi* | [**ListSupportingDocument**](docs/RegulatoryComplianceSupportingDocumentsApi.md#listsupportingdocument) | **Get** /v2/RegulatoryCompliance/SupportingDocuments | 
*RegulatoryComplianceSupportingDocumentsApi* | [**UpdateSupportingDocument**](docs/RegulatoryComplianceSupportingDocumentsApi.md#updatesupportingdocument) | **Post** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 


## Documentation For Models

 - [ListBundleCopyResponse](docs/ListBundleCopyResponse.md)
 - [ListBundleResponse](docs/ListBundleResponse.md)
 - [ListBundleResponseMeta](docs/ListBundleResponseMeta.md)
 - [ListEndUserResponse](docs/ListEndUserResponse.md)
 - [ListEndUserTypeResponse](docs/ListEndUserTypeResponse.md)
 - [ListEvaluationResponse](docs/ListEvaluationResponse.md)
 - [ListItemAssignmentResponse](docs/ListItemAssignmentResponse.md)
 - [ListRegulationResponse](docs/ListRegulationResponse.md)
 - [ListSupportingDocumentResponse](docs/ListSupportingDocumentResponse.md)
 - [ListSupportingDocumentTypeResponse](docs/ListSupportingDocumentTypeResponse.md)
 - [NumbersV2Bundle](docs/NumbersV2Bundle.md)
 - [NumbersV2BundleCopy](docs/NumbersV2BundleCopy.md)
 - [NumbersV2EndUser](docs/NumbersV2EndUser.md)
 - [NumbersV2EndUserType](docs/NumbersV2EndUserType.md)
 - [NumbersV2Evaluation](docs/NumbersV2Evaluation.md)
 - [NumbersV2ItemAssignment](docs/NumbersV2ItemAssignment.md)
 - [NumbersV2Regulation](docs/NumbersV2Regulation.md)
 - [NumbersV2ReplaceItems](docs/NumbersV2ReplaceItems.md)
 - [NumbersV2SupportingDocument](docs/NumbersV2SupportingDocument.md)
 - [NumbersV2SupportingDocumentType](docs/NumbersV2SupportingDocumentType.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

