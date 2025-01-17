package request

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
)

// AddQuery is a function that adds query parameters to a request.
func AddQuery(request *fiber.Agent, paramsURL []ParamURL) {
	if len(paramsURL) != 0 {
		query := ""
		for i, param := range paramsURL {
			query += param.Key + "=" + url.QueryEscape(param.Value)
			if i != len(paramsURL)-1 {
				query += "&"
			}
		}
		request.QueryString(query)
	}
}
