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

All URIs are relative to *https://studio.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FlowsApi* | [**DeleteFlow**](docs/FlowsApi.md#deleteflow) | **Delete** /v1/Flows/{Sid} | 
*FlowsApi* | [**FetchFlow**](docs/FlowsApi.md#fetchflow) | **Get** /v1/Flows/{Sid} | 
*FlowsApi* | [**ListFlow**](docs/FlowsApi.md#listflow) | **Get** /v1/Flows | 
*FlowsEngagementsApi* | [**CreateEngagement**](docs/FlowsEngagementsApi.md#createengagement) | **Post** /v1/Flows/{FlowSid}/Engagements | 
*FlowsEngagementsApi* | [**DeleteEngagement**](docs/FlowsEngagementsApi.md#deleteengagement) | **Delete** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
*FlowsEngagementsApi* | [**FetchEngagement**](docs/FlowsEngagementsApi.md#fetchengagement) | **Get** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
*FlowsEngagementsApi* | [**ListEngagement**](docs/FlowsEngagementsApi.md#listengagement) | **Get** /v1/Flows/{FlowSid}/Engagements | 
*FlowsEngagementsContextApi* | [**FetchEngagementContext**](docs/FlowsEngagementsContextApi.md#fetchengagementcontext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Context | 
*FlowsEngagementsStepsApi* | [**FetchStep**](docs/FlowsEngagementsStepsApi.md#fetchstep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{Sid} | 
*FlowsEngagementsStepsApi* | [**ListStep**](docs/FlowsEngagementsStepsApi.md#liststep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps | 
*FlowsEngagementsStepsContextApi* | [**FetchStepContext**](docs/FlowsEngagementsStepsContextApi.md#fetchstepcontext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{StepSid}/Context | 
*FlowsExecutionsApi* | [**CreateExecution**](docs/FlowsExecutionsApi.md#createexecution) | **Post** /v1/Flows/{FlowSid}/Executions | 
*FlowsExecutionsApi* | [**DeleteExecution**](docs/FlowsExecutionsApi.md#deleteexecution) | **Delete** /v1/Flows/{FlowSid}/Executions/{Sid} | 
*FlowsExecutionsApi* | [**FetchExecution**](docs/FlowsExecutionsApi.md#fetchexecution) | **Get** /v1/Flows/{FlowSid}/Executions/{Sid} | 
*FlowsExecutionsApi* | [**ListExecution**](docs/FlowsExecutionsApi.md#listexecution) | **Get** /v1/Flows/{FlowSid}/Executions | 
*FlowsExecutionsApi* | [**UpdateExecution**](docs/FlowsExecutionsApi.md#updateexecution) | **Post** /v1/Flows/{FlowSid}/Executions/{Sid} | 
*FlowsExecutionsContextApi* | [**FetchExecutionContext**](docs/FlowsExecutionsContextApi.md#fetchexecutioncontext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Context | 
*FlowsExecutionsStepsApi* | [**FetchExecutionStep**](docs/FlowsExecutionsStepsApi.md#fetchexecutionstep) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid} | 
*FlowsExecutionsStepsApi* | [**ListExecutionStep**](docs/FlowsExecutionsStepsApi.md#listexecutionstep) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps | 
*FlowsExecutionsStepsContextApi* | [**FetchExecutionStepContext**](docs/FlowsExecutionsStepsContextApi.md#fetchexecutionstepcontext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context | 


## Documentation For Models

 - [ListEngagementResponse](docs/ListEngagementResponse.md)
 - [ListExecutionResponse](docs/ListExecutionResponse.md)
 - [ListExecutionStepResponse](docs/ListExecutionStepResponse.md)
 - [ListFlowResponse](docs/ListFlowResponse.md)
 - [ListFlowResponseMeta](docs/ListFlowResponseMeta.md)
 - [ListStepResponse](docs/ListStepResponse.md)
 - [StudioV1Engagement](docs/StudioV1Engagement.md)
 - [StudioV1EngagementContext](docs/StudioV1EngagementContext.md)
 - [StudioV1Execution](docs/StudioV1Execution.md)
 - [StudioV1ExecutionContext](docs/StudioV1ExecutionContext.md)
 - [StudioV1ExecutionStep](docs/StudioV1ExecutionStep.md)
 - [StudioV1ExecutionStepContext](docs/StudioV1ExecutionStepContext.md)
 - [StudioV1Flow](docs/StudioV1Flow.md)
 - [StudioV1Step](docs/StudioV1Step.md)
 - [StudioV1StepContext](docs/StudioV1StepContext.md)


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

