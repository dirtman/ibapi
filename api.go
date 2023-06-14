package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/term"
	"io"
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"
	"syscall"
	"time"

	. "github.com/dirtman/sitepkg"
)

// Supported API authentication methods:
const AuthMethodBasic = "basic"
const AuthMethodBearer = "bearer"

// For convenience, hold the basic API parameter:
type APIConfig struct {
	BaseURL    string
	AuthMethod string
	AuthToken  string
}

var API APIConfig

// Fetch exit status; distinguish "not found" from other errors.
type FetchStatus int

const FetchStatusOK FetchStatus = 0
const FetchStatusError FetchStatus = 1
const FetchStatusNotFound FetchStatus = 2

/*****************************************************************************\
  API GET, POST, PUT and DELETE Requests
\*****************************************************************************/

func APIGet(url string, headers ...http.Header) (body []byte, err error) {
	return APIRequest("GET", url, nil, headers...)
}
func APIPost(url string, data interface{}, headers ...http.Header) (body []byte, err error) {
	return APIRequest("POST", url, data, headers...)
}
func APIPut(url string, data interface{}, headers ...http.Header) (body []byte, err error) {
	return APIRequest("PUT", url, data, headers...)
}
func APIDelete(url string, data interface{}, headers ...http.Header) (body []byte, err error) {
	return APIRequest("DELETE", url, data, headers...)
}
func APIPatch(url string, data interface{}, headers ...http.Header) (body []byte, err error) {
	return APIRequest("PATCH", url, data, headers...)
}

func APIRequest(method string, url string, data interface{}, headers ...http.Header) ([]byte, error) {

	var body, dataJson []byte
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
	urlPath := API.BaseURL + url
	if Debug {
		Show("Method: \"%s\".", method)
		Show("URL: \"%s\".", urlPath)
	}

	if data != nil {
		switch t := data.(type) {
		default:
			return body, Error("unknown body type provided (%T)", t)
		case string:
			payload = bytes.NewBuffer([]byte(data.(string)))
			if Debug {
				Show("Data (string): \"%v\".", data)
				Show("Payload: \"%s\".", data.(string))
			}
		case interface{}:
			if data != nil {
				if dataJson, err = json.Marshal(data); err != nil {
					return body, err
				}
				payload = bytes.NewBuffer(dataJson)
			}
			if Debug {
				Show("Data (any): \"%v\".", data)
				Show("Payload: \"%s\".", dataJson)
			}
		}
	}

	// Create a new http.Request.
	req, err := http.NewRequest(method, urlPath, payload)
	if err != nil {
		return body, Error("error getting url \"%s\": %v", urlPath, err)
	}

	// Configure auth for the new request.
	err = SetHTTPAuth(req)
	if err != nil {
		return body, Error("failure setting up API %s authentication: %s", API.AuthMethod, err)
	}

	// Set the headers for the request.  If no headers are provided, we'll assume a
	// JSON Content-Type is desired.
	if len(headers) == 0 {
		req.Header.Set("Content-Type", "application/json")
	} else {
		for _, httpHeader := range headers {
			for header, values := range httpHeader {
				for _, value := range values {
					req.Header.Set(header, value)
				}
			}
		}
	}

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
			// Show("%s:\n%s", "Response body", body)
			Show("Response body:")
			ShowBody(body)
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

	// Basic Auth requires a username and a password (curl -u option).  At this
	// point, API.AuthToken should contain the username and password, separated
	// by a ":".
	if API.AuthMethod == AuthMethodBasic {
		if API.AuthToken == "" {
			return Error("basic auth requires a username and password")
		}
		authSlice := strings.Split(API.AuthToken, ":")
		username := authSlice[0]
		password := authSlice[1]

		if username == "" {
			return Error("failure getting user from APIAuthToken \"%s\"", API.AuthToken)
		} else if password == "" {
			return Error("failure getting password from APIAuthToken \"%s\"", API.AuthToken)
		}
		req.SetBasicAuth(username, password)
		if Debug {
			Show("Username: \"%v\".", username)
			Show("Password: \"%v\".", "**************")
		}

		// Bearer auth requires a "Bearer: AuthToken" header.  At this point,
		// API.AuthToken should contain the bearer token.
	} else if API.AuthMethod == AuthMethodBearer {
		req.Header.Set("Authorization", "Bearer "+API.AuthToken)
		if Debug {
			// Don't show the whole token, but showing a bit may be helpful.
			show := len(API.AuthToken)
			if show > 20 {
				show = 20
			}
			Show("Setting Auth Header: \"%s\".", "Bearer "+API.AuthToken[0:show]+"...")
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

func InitAPI(args ...string) error {

	var url, authType, token, tokenID string
	var err error

	// The caller can set an authType directly; otherwise, the caller must set
	// the "APIAuthMethod" option, and we will look up the authType ourselves.
	if len(args) > 1 {
		return Error("InitAPI: bad call (too many args)")
	} else if len(args) == 1 {
		authType = args[0]
	} else if authType, err = GetStringOpt("APIAuthMethod"); err != nil {
		return Error("failure getting APIAuthMethod option: %v", err)
	}
	if authType = strings.ToLower(authType); authType == "" {
		return Error("empty APIAuthMethod specified")
	} else if authType != AuthMethodBearer && authType != AuthMethodBasic {
		return Error("Unsupported API authentication method \"%s\".", authType)
	}

	if url, err = GetStringOpt("APIBaseURL"); err != nil {
		return Error("failure getting APIBaseURL: %v", err)
	} else if url == "" {
		return Error("APIBaseURL in not configured")
	} else if token, err = GetStringOpt("APIAuthToken"); err != nil {
		return Error("failure getting APIAuthToken option: %v", err)
	} else if tokenID, err = GetStringOpt("APIAuthTokenID"); err != nil {
		return Error("failure getting APIAuthTokenID option: %v", err)
	}
	API.BaseURL = url
	API.AuthMethod = authType

	// Specifying a Username/Password takes precedence over both APIAuthToken
	// and APIAuthTokenID.  Note this is only applicable with Basic auth method.
	if authType == AuthMethodBasic {
		var username, password string
		var tokenSpecified = (token != "" || tokenID != "")
		if username, password, err = GetBasicAuthCreds(tokenSpecified); err != nil {
			return Error("failure getting username and password: %v", err)
		} else if username != "" && password != "" {
			// We have a username and password, so we are done.  Store the
			// username and password in API.AuthToken, separated by ":".
			API.AuthToken = strings.Join([]string{username, password}, ":")
			return nil
		}
	}

	// Check if an APIAuthToken has been specified; if so we are done here.
	// Note that a token is allowed for both Bearer and Basic authentication.
	// For Basic auth, APIAuthToken should have the form "username:password".
	// Also note that APIAuthToken takes precedence over APIAuthTokenID.
	if token != "" {
		API.AuthToken = token
		return nil
	}

	// We haven't found our credentials yet, so try APIAuthTokenID.
	if tokenID == "" {
		return Error("ran out of authentication methods to try.")
	} else if token, err = GetSecret(tokenID); err != nil {
		return Error("failure getting APIAuthToken from secrets file: %v", err)
	} else if token == "" {
		return Error("empty APIAuthToken retrieved from secrets file")
	}
	API.AuthToken = token
	return nil
}

// GetBasicAuthCreds processes the optional Username, Password and
// PromptForPassword options.

func GetBasicAuthCreds(altMethodsAvailable bool) (string, string, error) {

	var username, password string
	var prompt bool
	var err error

	if username, err = GetStringOpt("Username"); err != nil {
		return "", "", Error("failure getting username option: %v", err)
	} else if password, err = GetStringOpt("Password"); err != nil {
		return "", "", Error("failure getting Password option: %v", err)
	} else if prompt, err = GetBoolOpt("PromptForPassword"); err != nil {
		return "", "", Error("failure getting PromptForPassword option: %v", err)
	}

	// If a --Username option was not provided, and either: 1) no other
	// auth methods have been specified, or 2) a password has been specified,
	// or 3) the PromptForPassword option has been specified, then set
	// username to the current user.:
	if username == "" && (!altMethodsAvailable || password != "" || prompt) {
		if currentUser, err := user.Current(); err != nil {
			return "", "", Error("failure getting current user: %v", err)
		} else if username = currentUser.Username; username == "" {
			return "", "", Error("failure getting current user: got empty string")
		}
	}
	if username != "" && password == "" {
		if password, err = PromptForPassword(username); err != nil {
			return "", "", Error("%v", err)
		}
	}
	return username, password, nil
}

// PromptForPassword prompts the user for the password.

func PromptForPassword(username string) (string, error) {

	var bytepw []byte
	var err error

	fmt.Printf("Password for %s: ", username)
	if bytepw, err = term.ReadPassword(int(syscall.Stdin)); err != nil {
		return "", Error("failure prompting for password: %v", err)
	}
	fmt.Printf("\n")
	if len(bytepw) == 0 {
		return "", Error("failure prompting for password: zero length password not supported")
	}
	return string(bytepw), nil
}

// SetAPIOptions sets the required API related options, and if used, it must be
// called before ConfigureOptions (else setting these options will have no
// effect).  The caller can optionally set the required options himself instead of
// calling SetAPIOptions.  Note if the caller does not provide an authMethod, the
// "APIAuthMethod" option will be set, allowing the user to choose the authMethod.

func SetAPIOptions(args ...string) error {

	var authType string
	authTypes := strings.Join([]string{AuthMethodBasic, AuthMethodBearer}, ", ")
	authMethodHelp := fmt.Sprintf("API authentication type (%s)", authTypes)
	authTokenHelp := "API bearer access token"
	authIDHelp := "Name of a file that contains the API access token"

	// The caller can set an authType directly; otherwise, we will set
	// the "APIAuthMethod" option and allow the user to choose..
	if len(args) > 1 {
		return Error("SetAPIOptions: bad call (too many args)")
	} else if len(args) == 1 {
		authType = args[0]
	} else {
		SetStringOpt("APIAuthMethod", "", true, AuthMethodBasic, authMethodHelp)
	}
	if authType == "" || authType == AuthMethodBasic {
		authTokenHelp = "API access token (username:password)"
		SetStringOpt("Username", "u", true, "", "Username for API access")
		SetStringOpt("Password", "p", true, "", "Password for API access")
		SetBoolOpt("PromptForPassword", "P", false, false, "Prompt for password for API access")
	}

	// Set common API options:
	SetStringOpt("APIBaseURL", "", true, "", "API base URL")
	SetStringOpt("APIAuthToken", "", true, "", authTokenHelp)
	SetStringOpt("APIAuthTokenID", "", true, "", authIDHelp)
	SetStringOpt("SecretsDir", "", true, "", "Location of \"APIAuthTokenID\"")
	SetIntOpt("HTTPTimeout", "", true, 60, "Timeout in seconds of the HTTP connection")

	return nil
}

// ShowBody attempts to "pretty print" the returned body of an API request.
// It is generally used for debugging purposes.
// The old Mulesoft-based API returned a body as indented JSON,
// but the new java-based API returns normal (non-indented) JSON.
// The mailhome API endpoint returns a string.

func ShowBody(body []byte, indentOpts ...string) error {

	prefix := ""
	indent := "  "

	if len(body) == 0 {
		return Error("body has no bytes")
	} else if len(indentOpts) > 2 {
		return Error("bug: bad call to ShowBody()")
	}

	if len(indentOpts) > 0 {
		prefix = indentOpts[0]
	}
	if len(indentOpts) > 1 {
		indent = indentOpts[1]
	}

	// If the body starts with a '[' or '{', assume it's a JSON sequence.
	if body[0] == 91 || body[0] == 123 {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, body, prefix, indent)
		if err != nil {
			Warn("Failure indenting body for printing: %v", err)
			Print("%s\n", body)
		}
		Fprint(DefaultShow, "%s\n", prettyJSON.Bytes())
		return err
	}
	Print("%s\n", body)
	return nil
}
