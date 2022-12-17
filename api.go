package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"
	"time"

	. "github.com/dirtman/sitepkg"
)

// Supported API authentication methods:
const AuthMethodBasic = "Basic"
const AuthMethodBearer = "Bearer"

// For convenience, hold the basic API parameter:
type APIConfig struct {
	BaseURL      string
	AuthMethod   string
	AuthToken    string
	AuthUser     string
	AuthPassword string
}

var API APIConfig

/*****************************************************************************\
  API GET, POST, PUT and DELETE Requests
\*****************************************************************************/

func APIGet(url_request string) (body []byte, err error) {
	return APIRequest("GET", url_request, nil)
}
func APIPost(url_request string, data interface{}) (body []byte, err error) {
	return APIRequest("POST", url_request, data)
}
func APIPut(url_request string, data interface{}) (body []byte, err error) {
	return APIRequest("PUT", url_request, data)
}
func APIDelete(url_request string, data interface{}) (body []byte, err error) {
	return APIRequest("DELETE", url_request, data)
}

func APIRequest(update_type string, url_request string, data interface{}) ([]byte, error) {

	var body, data_json []byte
	var payload io.Reader
	var err error

	// Set the timeout for the request:
	seconds, err := GetIntOpt("HTTPTimeout")
	if err != nil {
		if !strings.Contains(err.Error(), ConfErrNoSuchOption) {
			return body, Error("failure getting current setting of \"HTTPTimeout\": %s", err)
		} else {
			seconds = 10
		}
	}
	timeout := time.Duration(time.Duration(seconds) * time.Second)
	url := API.BaseURL + url_request
	if Debug {
		Show("Method: \"%s\".", update_type)
		Show("URL: \"%s\".", url)
	}

	if data != nil {
		switch t := data.(type) {
		default:
			return body, Error("unknown body type provided (%T)", t)
		case string:
			payload = bytes.NewBuffer([]byte(data.(string)))
			if Debug {
				Show("Data: \"%v\".", data)
				Show("Payload: \"%s\".", data.(string))
			}
		case interface{}:
			if data != nil {
				if data_json, err = json.Marshal(data); err != nil {
					return body, err
				}
				payload = bytes.NewBuffer(data_json)
			}
			if Debug {
				Show("Data: \"%v\".", data)
				Show("Payload: \"%s\".", data_json)
			}
		}
	}

	// Create a new http.Request
	req, err := http.NewRequest(update_type, url, payload)
	if err != nil {
		return body, Error("error getting url \"%s\": %v", url, err)
	}

	// Create a new http.Request
	err = SetHTTPAuth(req)
	if err != nil {
		return body, Error("failure setting up API %s authentication: %s", API.AuthMethod, err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Make the specified request
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return body, Error("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, Error("HTTP failure reading response: %v", err)
	}

	if Debug {
		Show("HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		if body != nil {
			Show("%s:\n%s", "Response body", body)
		}
	}

	// For convenience to the caller, convert an empty body to nil:
	if len(body) == 0 {
		body = nil
	}

	// We'll treat a 2xx as success and return a nil error; for all other codes,
	// keep the original integer code intact, as first part of error string, to
	// allow the calling program to easily decipher it as an error or not.
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return body, nil
	}
	return body, Error("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

/*****************************************************************************\
  Handle the authentication for an http.Request.
  Globals: API - holds the configuration parmaters of the API.
\*****************************************************************************/

func SetHTTPAuth(req *http.Request) (err error) {

	if Debug {
		Show("AuthMethod: \"%v\".", API.AuthMethod)
	}

	// Basic Auth requires a user and a password (curl -u option)
	if API.AuthMethod == AuthMethodBasic {
		req.SetBasicAuth(API.AuthUser, API.AuthPassword)
		if Debug {
			Show("APIAuthUser: \"%v\".", API.AuthUser)
		}
		if Debug {
			//	Show("APIAuthPassword: \"%v\".", API.AuthPassword)
			Show("APIAuthPassword: \"%v\".", "**************")
		}

		// Bearer auth requires a "Bearer: AuthToken" header.
	} else if API.AuthMethod == AuthMethodBearer {
		req.Header.Set("Authorization", "Bearer "+API.AuthToken)
		if Debug {
			Show("Setting Auth Header: \"%s\".", "Bearer "+API.AuthToken)
		}

		// Unsupported AuthMethod
	} else {
		return Error("unknown APIAuthMethod \"%s\"", API.AuthMethod)
	}
	return nil
}

/*****************************************************************************\
  Setup the basic API paramters.
\*****************************************************************************/

func InitAPI() error {

	var url, token, id, password, auth_type string
	var err error

	// Get up the base URL.
	if url, err = GetStringOpt("APIBaseURL"); err != nil {
		return Error("failure getting APIBaseURL: %v", err)
	} else if url == "" {
		return Error("APIBaseURL in not configured")
	}
	API.BaseURL = url

	// Get the API authentication method.
	if auth_type, err = GetStringOpt("APIAuthMethod"); err != nil {
		return Error("failure getting APIAuthMethod: %v", err)
	} else if auth_type == "" {
		return Error("APIAuthMethod in not configured")
	} else if auth_type != AuthMethodBearer && auth_type != AuthMethodBasic {
		return Error("Unsupported API authentication method \"%s\".", auth_type)
	}
	API.AuthMethod = auth_type

	// Check if an APIAuthToken has been specified; if so were are done here.
	// Note that a token is allowed for both Bearer and Basic authentication. For
	// for Basic auth, APIAuthToken should have the form "APIAuthUser:APIAuthPassword"
	if token, err = GetStringOpt("APIAuthToken"); token != "" && err == nil {
		API.AuthToken = token
		return nil
	}
	if password, err = GetStringOpt("Password"); err != nil {
		return Error("failure getting Password: %v", err)
	}
	if id, err = GetStringOpt("APIAuthTokenID"); err != nil {
		return Error("failure getting APIAuthTokenID: %v", err)
	}

	// For Basic auth mode, if a password is specified, assume APIAuthTokenID,
	// if specified, is the username.  If not specified, assume the current user.
	if auth_type == AuthMethodBasic && password != "" {
		if id == "" {
			currentUser, err := user.Current()
			if err != nil {
				return Error("failure getting current user: %v", err)
			}
			id = currentUser.Username
		}
		API.AuthUser = id
		API.AuthPassword = password
		return nil
	}

	// We haven't found our credentials yet, so assume APIAuthTokenID is a file
	// containing our credentials.

	if id == "" {
		return Error("ran out of authentication methods to try.")
	} else if token, err = GetSecret(id); err != nil {
		return Error("failure getting APIAuthToken from secrets file: %v", err)
	} else if token == "" {
		return Error("empty APIAuthToken retrieved from secrets file")
	}
	API.AuthToken = token

	// For Basic auth, API.AuthToken should have the form "APIAuthUser:APIAuthPassword"
	if auth_type == AuthMethodBasic {
		auth_slice := strings.Split(API.AuthToken, ":")
		API.AuthUser = auth_slice[0]
		API.AuthPassword = auth_slice[1]

		if API.AuthUser == "" {
			return Error("failure getting user from APIAuthToken \"%s\"", API.AuthToken)
		}
		if API.AuthPassword == "" {
			return Error("failure getting password from APIAuthToken \"%s\"", API.AuthToken)
		}
	}

	return nil
}
