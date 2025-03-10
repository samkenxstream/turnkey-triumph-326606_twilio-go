/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c *ApiService) FetchAnnotation(CallSid string) (*InsightsV1Annotation, error) {
	path := "/v1/Voice/{CallSid}/Annotation"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Annotation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateAnnotation'
type UpdateAnnotationParams struct {
	//
	AnsweredBy *string `json:"AnsweredBy,omitempty"`
	//
	CallScore *int `json:"CallScore,omitempty"`
	//
	Comment *string `json:"Comment,omitempty"`
	//
	ConnectivityIssue *string `json:"ConnectivityIssue,omitempty"`
	//
	Incident *string `json:"Incident,omitempty"`
	//
	QualityIssues *string `json:"QualityIssues,omitempty"`
	//
	Spam *bool `json:"Spam,omitempty"`
}

func (params *UpdateAnnotationParams) SetAnsweredBy(AnsweredBy string) *UpdateAnnotationParams {
	params.AnsweredBy = &AnsweredBy
	return params
}
func (params *UpdateAnnotationParams) SetCallScore(CallScore int) *UpdateAnnotationParams {
	params.CallScore = &CallScore
	return params
}
func (params *UpdateAnnotationParams) SetComment(Comment string) *UpdateAnnotationParams {
	params.Comment = &Comment
	return params
}
func (params *UpdateAnnotationParams) SetConnectivityIssue(ConnectivityIssue string) *UpdateAnnotationParams {
	params.ConnectivityIssue = &ConnectivityIssue
	return params
}
func (params *UpdateAnnotationParams) SetIncident(Incident string) *UpdateAnnotationParams {
	params.Incident = &Incident
	return params
}
func (params *UpdateAnnotationParams) SetQualityIssues(QualityIssues string) *UpdateAnnotationParams {
	params.QualityIssues = &QualityIssues
	return params
}
func (params *UpdateAnnotationParams) SetSpam(Spam bool) *UpdateAnnotationParams {
	params.Spam = &Spam
	return params
}

func (c *ApiService) UpdateAnnotation(CallSid string, params *UpdateAnnotationParams) (*InsightsV1Annotation, error) {
	path := "/v1/Voice/{CallSid}/Annotation"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AnsweredBy != nil {
		data.Set("AnsweredBy", *params.AnsweredBy)
	}
	if params != nil && params.CallScore != nil {
		data.Set("CallScore", fmt.Sprint(*params.CallScore))
	}
	if params != nil && params.Comment != nil {
		data.Set("Comment", *params.Comment)
	}
	if params != nil && params.ConnectivityIssue != nil {
		data.Set("ConnectivityIssue", *params.ConnectivityIssue)
	}
	if params != nil && params.Incident != nil {
		data.Set("Incident", *params.Incident)
	}
	if params != nil && params.QualityIssues != nil {
		data.Set("QualityIssues", *params.QualityIssues)
	}
	if params != nil && params.Spam != nil {
		data.Set("Spam", fmt.Sprint(*params.Spam))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Annotation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
