package main

import (
	"net/http"
	"strconv"
)

// GetAccountSettings is a wrapper over account/settings.
// https://developer.twitter.com/en/docs/accounts-and-users/manage-account-settings/api-reference/get-account-settings
func GetAccountSettings() (string, error) {
	response, err := client.Get("https://api.twitter.com/1.1/account/settings.json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bits, err := readResponse(response.Body)
	if err != nil {
		return "", err
	}
	return string(bits), nil
}

// GetAccountVerifyCredentials is a wrapper over account/verify_credentials
// https://developer.twitter.com/en/docs/accounts-and-users/manage-account-settings/api-reference/get-account-verify_credentials
func GetAccountVerifyCredentials(includeEntities, skipStatus, includeEmail bool) (string, error) {
	request, err := http.NewRequest("GET", "https://api.twitter.com/1.1/account/verify_credentials.json", nil)
	q := request.URL.Query()
	q.Add("include_entities", strconv.FormatBool(includeEntities))
	q.Add("skip_status", strconv.FormatBool(skipStatus))
	q.Add("include_email", strconv.FormatBool(includeEmail))
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bits, err := readResponse(response.Body)
	if err != nil {
		return "", err
	}

	return string(bits), nil
}
