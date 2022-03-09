// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// IPSECPolicyAuthentication authentication for IPSec policy
// swagger:model IPSECPolicyAuthentication
type IPSECPolicyAuthentication string

const (

	// IPSECPolicyAuthenticationHmacMd596 captures enum value "hmac-md5-96"
	IPSECPolicyAuthenticationHmacMd596 IPSECPolicyAuthentication = "hmac-md5-96"

	// IPSECPolicyAuthenticationHmacSha256128 captures enum value "hmac-sha-256-128"
	IPSECPolicyAuthenticationHmacSha256128 IPSECPolicyAuthentication = "hmac-sha-256-128"

	// IPSECPolicyAuthenticationHmacSha196 captures enum value "hmac-sha1-96"
	IPSECPolicyAuthenticationHmacSha196 IPSECPolicyAuthentication = "hmac-sha1-96"

	// IPSECPolicyAuthenticationNone captures enum value "none"
	IPSECPolicyAuthenticationNone IPSECPolicyAuthentication = "none"
)

// for schema
var ipSECPolicyAuthenticationEnum []interface{}

func init() {
	var res []IPSECPolicyAuthentication
	if err := json.Unmarshal([]byte(`["hmac-md5-96","hmac-sha-256-128","hmac-sha1-96","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipSECPolicyAuthenticationEnum = append(ipSECPolicyAuthenticationEnum, v)
	}
}

func (m IPSECPolicyAuthentication) validateIPSECPolicyAuthenticationEnum(path, location string, value IPSECPolicyAuthentication) error {
	if err := validate.Enum(path, location, value, ipSECPolicyAuthenticationEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this IP s e c policy authentication
func (m IPSECPolicyAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIPSECPolicyAuthenticationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}