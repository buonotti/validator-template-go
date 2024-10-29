package main

// Variable describes a variable that should be interpolated in the base url and the query parameters
type Variable struct {
	Name       string   `yaml:"name" json:"name" validate:"required"`         // Name is the name of the variable
	IsConstant bool     `yaml:"constant" json:"constant" validate:"required"` // IsConstant is true if the value of the variable is constant or else false
	Values     []string `yaml:"values" json:"values" validate:"required"`     // Values are all the possible values of the variable (only 1 in case of a constant)
}

type JwtLoginOptions struct {
	Url          string         `yaml:"url" json:"url" validate:"required"`   // Url is the url to the login endpoint
	LoginPayload map[string]any `yaml:"login_payload" json:"login_payload"`   // LoginPayload is the json or yml payload to send
	TokenKeyName string         `yaml:"token_key_name" json:"token_key_name"` // TokenKeyName is the name of the key in the response which contains the token
}

// QueryDefinition is a query parameter that should be added to the call
type QueryDefinition struct {
	Name  string `yaml:"name" json:"name" validate:"required"`   // Name is the name of the query parameter
	Value string `yaml:"value" json:"value" validate:"required"` // Value is the value of the query parameter
}

// Endpoint is the definition of an endpoint to test with all its query
// parameters, variables and its result schema
type Endpoint struct {
	Name               string            `yaml:"name,omitempty" json:"name,omitempty" validate:"required"`                      // Name is the name of the endpoint
	IsEnabled          bool              `yaml:"enabled,omitempty" json:"enabled,omitempty"`                                    // IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)
	BaseUrl            string            `yaml:"base_url,omitempty" json:"baseUrl,omitempty" validate:"required"`               // BaseUrl is the base path of the endpoint
	Method             string            `yaml:"method,omitempty" json:"method,omitempty"`                                      // Method is the name of the http-method to use for the request
	Payload            map[string]any    `yaml:"payload,omitempty" json:"payload,omitempty"`                                    // Payload is the payload to use in case of a POST or PUT request
	Authorization      string            `yaml:"authorization,omitempty" json:"authorization,omitempty"`                        // Authorization is the value to set for the authorization header
	JwtLogin           JwtLoginOptions   `yaml:"jwt_login,omitempty" json:"jwt_login,omitempty"`                                // JwtLogin are options to auto-get a login token for a request.
	Headers            map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`                                    // Headers are additional headers to set for the request
	ExcludedValidators []string          `yaml:"excluded_validators,omitempty" json:"excludedValidators,omitempty"`             // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []QueryDefinition `yaml:"query_parameters,omitempty" json:"queryParameters,omitempty"`                   // QueryParameters are all the query parameters that should be added to the call
	Format             string            `yaml:"format,omitempty" json:"format,omitempty" validate:"required"`                  // Format is the response format of the
	Variables          []Variable        `yaml:"variables,omitempty" json:"variables,omitempty"`                                // Variables are all the variables that should be interpolated in the base url and the query parameters
	OkCode             int               `yaml:"ok_code,omitempty" json:"ok_code,omitempty"`                                    // The expected status code
	ResponseSchema     map[string]any    `yaml:"response_schema,omitempty" json:"responseSchema,omitempty" validate:"required"` // ResponseSchema describes how the response should look like
}

type EndpointResponse struct {
	StatusCode int    // StatusCode is the status code of the response
	RawData    any    // RawData is the raw data of the response
	Url        string // Url is the full url of the request
}

type ValidationItem struct {
	Response   EndpointResponse
	Definition Endpoint
}
