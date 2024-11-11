package controllers

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

// ErrorController defines custom error output
type ErrorController struct {
	beego.Controller
}

// Error401 shows 401 Unauthorized
func (c *ErrorController) Error401() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "401",
		"title": "Unauthorized",
	}

	c.Data["json"] = output
}

// Error402 shows 402 Payment Required
func (c *ErrorController) Error402() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "402",
		"title": "Payment Required",
	}

	c.Data["json"] = output
}

// Error403 shows 403 Forbidden
func (c *ErrorController) Error403() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "403",
		"title": "Forbidden",
	}

	c.Data["json"] = output
}

// Error404 shows 404 Not Found
func (c *ErrorController) Error404() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "404",
		"title": "Not Found",
	}
	http.StatusText(404)

	c.Data["json"] = output
}

// Error405 shows 405 Method Not Allowed
func (c *ErrorController) Error405() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "405",
		"title": "Method Not Allowed",
	}

	c.Data["json"] = output
}

// Error413 shows 413 Payload Too Large
func (c *ErrorController) Error413() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "413",
		"title": "Payload Too Large",
	}

	c.Data["json"] = output
}

// Error417 shows 417 Invalid xsrf token
func (c *ErrorController) Error417() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "417",
		"title": "Invalid xsrf token",
	}

	c.Data["json"] = output
}

// Error422 shows 422 '_xsrf' argument missing from POST
func (c *ErrorController) Error422() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "422",
		"title": "'_xsrf' argument missing from POST",
	}

	c.Data["json"] = output
}

// Error500 shows 500 Internal Server Error
func (c *ErrorController) Error500() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "500",
		"title": "Internal Server Error",
	}

	c.Data["json"] = output
}

// Error501 shows 501 Not Implemented
func (c *ErrorController) Error501() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "501",
		"title": "Not Implemented",
	}

	c.Data["json"] = output
}

// Error502 shows 502 Bad Gateway
func (c *ErrorController) Error502() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "502",
		"title": "Bad Gateway",
	}

	c.Data["json"] = output
}

// Error503 shows 503 Service Unavailable
func (c *ErrorController) Error503() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "503",
		"title": "Service Unavailable",
	}

	c.Data["json"] = output
}

// Error504 shows 504 Gateway Timeout
func (c *ErrorController) Error504() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "504",
		"title": "Gateway Timeout",
	}

	c.Data["json"] = output
}

// ErrorDb shows 503 database is now down
func (c *ErrorController) ErrorDb() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false
	output["error"] = map[string]string{
		"code":  "503",
		"title": "Service Unavailable",
	}

	c.Data["json"] = output
}
