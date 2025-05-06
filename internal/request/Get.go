package request

import (
	"github.com/ElderLab/riot-wrapper/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"strings"
)

// Get is a method that sends a GET request to the specified resource
func (c *Client) Get(resource string, paramsURL []ParamURL) ([]byte, error) {
	if strings.Contains(resource, ":") {
		idToDelete := make([]int, 0)
		for i, param := range paramsURL {
			ok := strings.Contains(resource, ":"+param.Key)
			resource = strings.Replace(resource, ":"+param.Key, url.PathEscape(param.Value), -1)
			if ok {
				idToDelete = append(idToDelete, i)
			}
		}
		comp := 0
		for _, i := range idToDelete {
			paramsURL = append(paramsURL[:i-comp], paramsURL[i+1-comp:]...)
			comp++
		}
	}
	request := fiber.Get(c.baseURL + resource)
	//request.Debug()
	request.Set("X-Riot-Token", c.xRiotToken)

	// Query parameters
	AddQuery(request, paramsURL)

	statusCode, data, errs := request.Bytes()
	//save the last status code
	c.lastStatus = statusCode

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	//check if the status code is not 200
	if statusCode/100 != 2 {
		// if the client was rate limited
		if statusCode == 429 {
			//return the error message
			return data, errors.NewRateLimitedError()
		}
		//return the error message
		return data, errors.NewStatusError(statusCode)
	}

	return data, nil
}
