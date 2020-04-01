/*
 * Introduction to Instana public APIs
 *
 * ## Agent REST API ### Event SDK REST Web Service Using the Event SDK REST Web Service, it is possible to integrate custom health checks and other event sources into Instana. Each one running the Instana Agent can be used to feed in manual events. The agent has an endpoint which listens on `http://localhost:42699/com.instana.plugin.generic.event` and accepts the following JSON via a POST request:  ```json {     \"title\": <string>,     \"text\": <string>,     \"severity\": <integer> , -1, 5 or 10     \"timestamp\": <integer>, timestamp in milliseconds from epoch     \"duration\": <integer>, duration in milliseconds } ```  *Title* and *text* are used for display purposes.  *Severity* is an optional integer of -1, 5 and 10. A value of -1 or EMPTY will generate a Change. A value of 5 will generate a *warning Issue*, and a value of 10 will generate a *critical Issue*.  When absent, the event is treated as a change without severity. *Timestamp* is the timestamp of the event, but it is optional, in which case the current time is used. *Duration* can be used to mark a timespan for the event. It also is optional, in which case the event will be marked as \"instant\" rather than \"from-to.\"  The endpoint also accepts a batch of events, which then need to be given as an array:  ```json [     {     // event as above     },     {     // event as above     } ] ```  #### Ruby Example  ```ruby duration = (Time.now.to_f * 1000).floor - deploy_start_time_in_ms payload = {} payload[:title] = 'Deployed MyApp' payload[:text] = 'pglombardo deployed MyApp@revision' payload[:duration] = duration  uri = URI('http://localhost:42699/com.instana.plugin.generic.event') req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json') req.body = payload.to_json Net::HTTP.start(uri.hostname, uri.port) do |http|     http.request(req) end ```  #### Curl Example  ```bash curl -XPOST http://localhost:42699/com.instana.plugin.generic.event -H \"Content-Type: application/json\" -d '{\"title\":\"Custom API Events \", \"text\": \"Failure Redeploying Service Duration\", \"duration\": 5000, \"severity\": -1}' ```  #### PowerShell Example  For Powershell you can either use the standard Cmdlets `Invoke-WebRequest` or `Invoke-RestMethod`. The parameters to be provided are basically the same.  ```bash Invoke-RestMethod     -Uri http://localhost:42699/com.instana.plugin.generic.event     -Method POST     -Body '{\"title\":\"PowerShell Event \", \"text\": \"You used PowerShell to create this event!\", \"duration\": 5000, \"severity\": -1}' ```  ```bash Invoke-WebRequest     -Uri http://localhost:42699/com.instana.plugin.generic.event     -Method Post     -Body '{\"title\":\"PowerShell Event \", \"text\": \"You used PowerShell to create this event!\", \"duration\": 5000, \"severity\": -1}' ``` ## Backend REST API The Instana API allows retrieval and configuration of key data points. Among others, this API enables automatic reaction and further analysis of identified incidents as well as reporting capabilities.  The API documentation referes to two crucial parameters that you need to know about before reading further: base: This is the base URL of a tenant unit, e.g. `https://test-example.instana.io`. This is the same URL that is used to access the Instana user interface. apiToken: Requests against the Instana API require valid API tokens. An initial API token can be generated via the Instana user interface. Any additional API tokens can be generated via the API itself.  ### Example Here is an Example to use the REST API with Curl. First lets get all the available metrics with possible aggregations with a GET call.  ```bash curl --request GET \\   --url https://test-instana.instana.io/api/application-monitoring/catalog/metrics \\   --header 'authorization: apiToken xxxxxxxxxxxxxxxx' ```  Next we can get every call grouped by the endpoint name that has an error count greater then zero. As a metric we could get the mean error rate for example.  ```bash curl --request POST \\   --url https://test-instana.instana.io/api/application-monitoring/analyze/call-groups \\   --header 'authorization: apiToken xxxxxxxxxxxxxxxx' \\   --header 'content-type: application/json' \\   --data '{   \"group\":{       \"groupbyTag\":\"endpoint.name\"   },   \"tagFilters\":[    {     \"name\":\"call.error.count\",     \"value\":\"0\",     \"operator\":\"GREATER_THAN\"    }   ],   \"metrics\":[    {     \"metric\":\"errors\",     \"aggregation\":\"MEAN\"    }   ]   }' ```   ### Rate Limiting A rate limit is applied to API usage. Up to 5,000 calls per hour can be made. How many remaining calls can be made and when this call limit resets, can inspected via three headers that are part of the responses of the API server.  **X-RateLimit-Limit:** Shows the maximum number of calls that may be executed per hour.  **X-RateLimit-Remaining:** How many calls may still be executed within the current hour.  **X-RateLimit-Reset:** Time when the remaining calls will be reset to the limit. For compatibility reasons with other rate limited APIs, this date is not the date in milliseconds, but instead in seconds since 1970-01-01T00:00:00+00:00.  ## Generating REST API clients  The API is specified using the [OpenAPI v3](https://github.com/OAI/OpenAPI-Specification) (previously known as Swagger) format. You can download the current specification at our [GitHub API documentation](https://instana.github.io/openapi/openapi.yaml).  OpenAPI tries to solve the issue of ever-evolving APIs and clients lagging behind. To generate a client library for your language, you can use the [OpenAPI client generators](https://github.com/OpenAPITools/openapi-generator).  To generate a client library for Go to interact with our backend, you can use the following script (you need a JDK and `wget`):  ```bash //Download the generator to your current working directory: wget http://central.maven.org/maven2/org/openapitools/openapi-generator-cli/3.2.3/openapi-generator-cli-3.2.3.jar -O openapi-generator-cli.jar  //generate a client library that you can vendor into your repository java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g go \\     -o pkg/instana/openapi \\     --skip-validate-spec  //(optional) format the Go code according to the Go code standard gofmt -s -w pkg/instana/openapi ```  The generated clients contain comprehensive READMEs. To use the client from the example above, you can start right away:  ```go import instana \"./pkg/instana/openapi\"  // readTags will read all available application monitoring tags along with their type and category func readTags() {  configuration := instana.NewConfiguration()  configuration.Host = \"tenant-unit.instana.io\"  configuration.BasePath = \"https://tenant-unit.instana.io\"   client := instana.NewAPIClient(configuration)  auth := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{   Key:    apiKey,   Prefix: \"apiToken\",  })   tags, _, err := client.ApplicationCatalogApi.GetTagsForApplication(auth)  if err != nil {   fmt.Fatalf(\"Error calling the API, aborting.\")  }   for _, tag := range tags {   fmt.Printf(\"%s (%s): %s\\n\", tag.Category, tag.Type, tag.Name)  } } ```
 *
 * API version: 1.174.285
 * Contact: support@instana.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WebsiteMonitoringBeacon struct for WebsiteMonitoringBeacon
type WebsiteMonitoringBeacon struct {
	WebsiteId                    string            `json:"websiteId"`
	WebsiteLabel                 string            `json:"websiteLabel"`
	Page                         string            `json:"page,omitempty"`
	Phase                        string            `json:"phase,omitempty"`
	Timestamp                    int64             `json:"timestamp,omitempty"`
	ClockSkew                    int64             `json:"clockSkew,omitempty"`
	Duration                     int64             `json:"duration,omitempty"`
	BatchSize                    int64             `json:"batchSize,omitempty"`
	AccurateTimingsAvailable     bool              `json:"accurateTimingsAvailable,omitempty"`
	Deprecations                 []string          `json:"deprecations,omitempty"`
	PageLoadId                   string            `json:"pageLoadId"`
	SessionId                    string            `json:"sessionId,omitempty"`
	BeaconId                     string            `json:"beaconId"`
	BackendTraceId               string            `json:"backendTraceId,omitempty"`
	Type                         string            `json:"type"`
	CustomEventName              string            `json:"customEventName,omitempty"`
	Meta                         map[string]string `json:"meta,omitempty"`
	LocationUrl                  string            `json:"locationUrl"`
	LocationOrigin               string            `json:"locationOrigin"`
	LocationPath                 string            `json:"locationPath,omitempty"`
	ErrorCount                   int64             `json:"errorCount,omitempty"`
	ErrorMessage                 string            `json:"errorMessage,omitempty"`
	ErrorId                      string            `json:"errorId,omitempty"`
	ErrorType                    string            `json:"errorType,omitempty"`
	StackTrace                   string            `json:"stackTrace,omitempty"`
	StackTraceParsingStatus      int32             `json:"stackTraceParsingStatus,omitempty"`
	ParsedStackTrace             []StackTraceLine  `json:"parsedStackTrace,omitempty"`
	StackTraceReadability        int32             `json:"stackTraceReadability,omitempty"`
	ComponentStack               string            `json:"componentStack,omitempty"`
	UserIp                       string            `json:"userIp,omitempty"`
	UserId                       string            `json:"userId,omitempty"`
	UserName                     string            `json:"userName,omitempty"`
	UserEmail                    string            `json:"userEmail,omitempty"`
	UserLanguages                []string          `json:"userLanguages,omitempty"`
	DeviceType                   string            `json:"deviceType,omitempty"`
	ConnectionType               string            `json:"connectionType,omitempty"`
	BrowserName                  string            `json:"browserName,omitempty"`
	BrowserVersion               string            `json:"browserVersion,omitempty"`
	OsName                       string            `json:"osName,omitempty"`
	OsVersion                    string            `json:"osVersion,omitempty"`
	WindowHidden                 bool              `json:"windowHidden,omitempty"`
	WindowWidth                  int32             `json:"windowWidth,omitempty"`
	WindowHeight                 int32             `json:"windowHeight,omitempty"`
	Latitude                     float64           `json:"latitude,omitempty"`
	Longitude                    float64           `json:"longitude,omitempty"`
	AccuracyRadius               int64             `json:"accuracyRadius,omitempty"`
	City                         string            `json:"city,omitempty"`
	Subdivision                  string            `json:"subdivision,omitempty"`
	SubdivisionCode              string            `json:"subdivisionCode,omitempty"`
	Country                      string            `json:"country,omitempty"`
	CountryCode                  string            `json:"countryCode,omitempty"`
	Continent                    string            `json:"continent,omitempty"`
	ContinentCode                string            `json:"continentCode,omitempty"`
	HttpCallUrl                  string            `json:"httpCallUrl,omitempty"`
	HttpCallOrigin               string            `json:"httpCallOrigin,omitempty"`
	HttpCallPath                 string            `json:"httpCallPath,omitempty"`
	HttpCallMethod               string            `json:"httpCallMethod,omitempty"`
	HttpCallStatus               int32             `json:"httpCallStatus,omitempty"`
	HttpCallCorrelationAttempted bool              `json:"httpCallCorrelationAttempted,omitempty"`
	HttpCallAsynchronous         bool              `json:"httpCallAsynchronous,omitempty"`
	Initiator                    string            `json:"initiator,omitempty"`
	ResourceType                 string            `json:"resourceType,omitempty"`
	CacheInteraction             string            `json:"cacheInteraction,omitempty"`
	EncodedBodySize              int64             `json:"encodedBodySize,omitempty"`
	DecodedBodySize              int64             `json:"decodedBodySize,omitempty"`
	TransferSize                 int64             `json:"transferSize,omitempty"`
	UnloadTime                   int64             `json:"unloadTime,omitempty"`
	RedirectTime                 int64             `json:"redirectTime,omitempty"`
	AppCacheTime                 int64             `json:"appCacheTime,omitempty"`
	DnsTime                      int64             `json:"dnsTime,omitempty"`
	TcpTime                      int64             `json:"tcpTime,omitempty"`
	SslTime                      int64             `json:"sslTime,omitempty"`
	RequestTime                  int64             `json:"requestTime,omitempty"`
	ResponseTime                 int64             `json:"responseTime,omitempty"`
	ProcessingTime               int64             `json:"processingTime,omitempty"`
	OnLoadTime                   int64             `json:"onLoadTime,omitempty"`
	BackendTime                  int64             `json:"backendTime,omitempty"`
	FrontendTime                 int64             `json:"frontendTime,omitempty"`
	DomTime                      int64             `json:"domTime,omitempty"`
	ChildrenTime                 int64             `json:"childrenTime,omitempty"`
	FirstPaintTime               int64             `json:"firstPaintTime,omitempty"`
	FirstContentfulPaintTime     int64             `json:"firstContentfulPaintTime,omitempty"`
	CspBlockedUri                string            `json:"cspBlockedUri,omitempty"`
	CspEffectiveDirective        string            `json:"cspEffectiveDirective,omitempty"`
	CspOriginalPolicy            string            `json:"cspOriginalPolicy,omitempty"`
	CspDisposition               string            `json:"cspDisposition,omitempty"`
	CspSample                    string            `json:"cspSample,omitempty"`
	CspSourceFile                string            `json:"cspSourceFile,omitempty"`
	CspLineNumber                int64             `json:"cspLineNumber,omitempty"`
	CspColumnNumber              int64             `json:"cspColumnNumber,omitempty"`
}
