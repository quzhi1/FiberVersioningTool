package lib

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type VersionChangeList []VersionDef

func (vcl VersionChangeList) RequestVersionAdaptor(c *fiber.Ctx) error {
	version := c.GetReqHeaders()["Nylas-Api-Version"]
	fmt.Println("Version: " + version)

	// loop through version change
	startVersion := false
	for _, versionChange := range vcl {
		// Skip if version is newer than version change
		if !startVersion {
			if version == versionChange.Version {
				fmt.Println("Version match, start applying version change: " + version)
				startVersion = true
			}
			continue
		}

		fmt.Println("Applying request version change for version " + versionChange.Version)
		// Change request body
		newRequestBody, err := versionChange.RequestBodyChange(c.Body())
		if err != nil {
			return err
		}
		fmt.Printf("New request body: %s\n", string(newRequestBody))
		c.Request().SetBodyRaw(newRequestBody)

		// Parse query string
		queryString := string(c.Request().URI().QueryString())
		oldValues, err := url.ParseQuery(queryString)
		if err != nil {
			return err
		}

		// Generate old query map
		oldQueryMap := map[string]string{}
		for queryKey, queryVals := range oldValues {
			oldQueryMap[queryKey] = strings.Join(queryVals, ",")
		}

		// Transform query parameters
		newQueryMap, err := versionChange.QueryParamChange(oldQueryMap)
		if err != nil {
			return err
		}

		// Change query params
		values := url.Values{}
		for queryKey, queryVal := range newQueryMap {
			values.Add(queryKey, queryVal)
		}
		query := values.Encode()
		c.Request().URI().SetQueryString(query)

		// Change request headers
		newRequestHeaders, err := versionChange.RequestHeaderChange(c.GetReqHeaders())
		if err != nil {
			return err
		}
		for headerKey, headerVal := range newRequestHeaders {
			c.Request().Header.Set(headerKey, headerVal)
		}
	}

	return c.Next()
}

func (vcl VersionChangeList) ResponseVersionAdaptor(c *fiber.Ctx) error {
	version := c.GetReqHeaders()["Nylas-Api-Version"]
	fmt.Println("Version: " + version)

	// loop through version change
	startVersion := false
	for _, versionChange := range vcl {
		// Skip if version is newer than version change
		if !startVersion {
			if version == versionChange.Version {
				fmt.Println("Version match, start applying version change: " + version)
				startVersion = true
			}
			continue
		}

		fmt.Println("Applying response version change for version " + versionChange.Version)

		// Change response body
		newResponseBody, err := versionChange.ResponseBodyChange(c.Response().Body())
		if err != nil {
			return err
		}
		fmt.Println("Setting response body: " + string(newResponseBody))
		c.Response().SetBodyRaw(newResponseBody)

		// Change response headers
		newResponseHeaders, err := versionChange.ResponseHeaderChange(c.GetRespHeaders())
		if err != nil {
			return err
		}
		for headerKey, headerVal := range newResponseHeaders {
			c.Set(headerKey, headerVal)
		}
	}

	return nil
}
