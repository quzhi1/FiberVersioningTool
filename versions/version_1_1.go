package versions

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/quzhi1/fiber-versioning-tool/lib"
	"github.com/quzhi1/fiber-versioning-tool/schema"
)

var Version1_1 = lib.VersionDef{
	Version:              "1.1",
	RequestBodyChange:    RequestBodyChange1_1,
	ResponseBodyChange:   ResponseBodyChange1_1,
	QueryParamChange:     QueryParamChange1_1,
	RequestHeaderChange:  RequestHeaderChange1_1,
	ResponseHeaderChange: ResponseHeaderChange1_1,
}

func RequestBodyChange1_1(requestBodyByte []byte) ([]byte, error) {
	// Unmarshal via old format
	var requestBody1_0 schema.RequestBodyVersion1_0
	err := json.Unmarshal(requestBodyByte, &requestBody1_0)
	if err != nil {
		return nil, err
	}

	// Split name
	splitted := strings.Split(requestBody1_0.Name, " ")
	if len(splitted) != 2 {
		return nil, errors.New("wrong name format. It should be in format: Firstname Lastname")
	}

	// Marshal with new format
	requestBody1_1 := schema.RequestBodyVersion1_1{
		FirstName: splitted[0],
		LastName:  splitted[1],
	}
	return json.Marshal(requestBody1_1)
}

func ResponseBodyChange1_1(responseBodyByte []byte) ([]byte, error) {
	// Unmarshal via new format
	var responseBody1_1 schema.ResponseBodyVersion1_1
	err := json.Unmarshal(responseBodyByte, &responseBody1_1)
	if err != nil {
		return nil, err
	}

	// Marshal with old format
	responseBody1_0 := schema.ResponseBodyVersion1_0{
		Id:          responseBody1_1.Id,
		Name:        responseBody1_1.Name,
		CreatedTime: time.Now().Unix(),
	}
	return json.Marshal(responseBody1_0)
}

func QueryParamChange1_1(oldQParams map[string]string) (map[string]string, error) {
	newQParams := map[string]string{}
	fmt.Printf("Query params: %v\n", oldQParams)
	langCode, exist := oldQParams["language-code"]
	if exist {
		fmt.Printf("Got language code: %s\n", langCode)
		splitted := strings.Split(langCode, "-")
		if len(splitted) != 2 {
			return nil, errors.New("invalid language-code format")
		} else {
			newQParams["lang"] = splitted[0]
			newQParams["region"] = splitted[1]
		}
	}
	fmt.Printf("Setting new query params: %v\n", newQParams)
	return newQParams, nil
}

func RequestHeaderChange1_1(oldHeaders map[string]string) (map[string]string, error) {
	newHeaders := map[string]string{}
	newHeaders["Client-Metadata"] = oldHeaders["Metadata"]
	return newHeaders, nil
}

func ResponseHeaderChange1_1(oldHeaders map[string]string) (map[string]string, error) {
	newHeaders := map[string]string{}
	newHeaders["Metadata"] = oldHeaders["Client-Metadata"]
	return newHeaders, nil
}
