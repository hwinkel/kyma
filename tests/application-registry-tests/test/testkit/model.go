/*
 *  © 2018 SAP SE or an SAP affiliate company.
 *  All rights reserved.
 *  Please see http://www.sap.com/corporate-en/legal/copyright/index.epx for additional trademark information and
 *  notices.
 */
package testkit

import (
	"bytes"
	"encoding/json"
)

var (
	ApiRawSpec    = Compact([]byte("{\"name\":\"api\"}"))
	EventsRawSpec = Compact([]byte("{\"name\":\"events\"}"))

	SwaggerApiSpec = Compact([]byte("{\"swagger\":\"2.0\"}"))
)

type Service struct {
	ID               string            `json:"id"`
	Provider         string            `json:"provider"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription,omitempty"`
	Identifier       string            `json:"identifier,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
}

type ServiceDetails struct {
	Provider         string            `json:"provider"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription,omitempty"`
	Identifier       string            `json:"identifier,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	Api              *API              `json:"api,omitempty"`
	Events           *Events           `json:"events,omitempty"`
	Documentation    *Documentation    `json:"documentation,omitempty"`
}

type PostServiceResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type API struct {
	TargetUrl        string          `json:"targetUrl"`
	Credentials      *Credentials    `json:"credentials,omitempty"`
	Spec             json.RawMessage `json:"spec,omitempty"`
	SpecificationUrl string          `json:"specificationUrl,omitempty"`
	ApiType          string          `json:"apiType"`
}

type Credentials struct {
	Oauth          *Oauth          `json:"oauth,omitempty"`
	Basic          *Basic          `json:"basic,omitempty"`
	CertificateGen *CertificateGen `json:"certificateGen, omitempty"`
}

type Oauth struct {
	URL          string `json:"url"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type Basic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CertificateGen struct {
	CommonName  string `json:"commonName"`
	Certificate string `json:"certificate"`
}

type Events struct {
	Spec json.RawMessage `json:"spec,omitempty"`
}

type Documentation struct {
	DisplayName string       `json:"displayName"`
	Description string       `json:"description"`
	Type        string       `json:"type"`
	Tags        []string     `json:"tags,omitempty"`
	Docs        []DocsObject `json:"docs,omitempty"`
}

type DocsObject struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Source string `json:"source"`
}

func Compact(src []byte) []byte {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, src)
	if err != nil {
		return src
	}
	return buffer.Bytes()
}

func (sd ServiceDetails) WithAPI(api *API) ServiceDetails {
	sd.Api = api
	return sd
}
