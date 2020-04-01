# Go API client for openapi

## Agent REST API
### Event SDK REST Web Service
Using the Event SDK REST Web Service, it is possible to integrate custom health checks and other event sources into Instana. Each one running the Instana Agent can be used to feed in manual events. The agent has an endpoint which listens on `http://localhost:42699/com.instana.plugin.generic.event` and accepts the following JSON via a POST request:

```json
{
    \"title\": <string>,
    \"text\": <string>,
    \"severity\": <integer> , -1, 5 or 10
    \"timestamp\": <integer>, timestamp in milliseconds from epoch
    \"duration\": <integer>, duration in milliseconds
}
```

*Title* and *text* are used for display purposes.

*Severity* is an optional integer of -1, 5 and 10. A value of -1 or EMPTY will generate a Change. A value of 5 will generate a *warning Issue*, and a value of 10 will generate a *critical Issue*.

When absent, the event is treated as a change without severity. *Timestamp* is the timestamp of the event, but it is optional, in which case the current time is used. *Duration* can be used to mark a timespan for the event. It also is optional, in which case the event will be marked as \"instant\" rather than \"from-to.\"

The endpoint also accepts a batch of events, which then need to be given as an array:

```json
[
    {
    // event as above
    },
    {
    // event as above
    }
]
```

#### Ruby Example

```ruby
duration = (Time.now.to_f * 1000).floor - deploy_start_time_in_ms
payload = {}
payload[:title] = 'Deployed MyApp'
payload[:text] = 'pglombardo deployed MyApp@revision'
payload[:duration] = duration

uri = URI('http://localhost:42699/com.instana.plugin.generic.event')
req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json')
req.body = payload.to_json
Net::HTTP.start(uri.hostname, uri.port) do |http|
    http.request(req)
end
```

#### Curl Example

```bash
curl -XPOST http://localhost:42699/com.instana.plugin.generic.event -H \"Content-Type: application/json\" -d '{\"title\":\"Custom API Events \", \"text\": \"Failure Redeploying Service Duration\", \"duration\": 5000, \"severity\": -1}'
```

#### PowerShell Example

For Powershell you can either use the standard Cmdlets `Invoke-WebRequest` or `Invoke-RestMethod`. The parameters to be provided are basically the same.

```bash
Invoke-RestMethod
    -Uri http://localhost:42699/com.instana.plugin.generic.event
    -Method POST
    -Body '{\"title\":\"PowerShell Event \", \"text\": \"You used PowerShell to create this event!\", \"duration\": 5000, \"severity\": -1}'
```

```bash
Invoke-WebRequest
    -Uri http://localhost:42699/com.instana.plugin.generic.event
    -Method Post
    -Body '{\"title\":\"PowerShell Event \", \"text\": \"You used PowerShell to create this event!\", \"duration\": 5000, \"severity\": -1}'
```
## Backend REST API
The Instana API allows retrieval and configuration of key data points. Among others, this API enables automatic reaction and further analysis of identified incidents as well as reporting capabilities.

The API documentation referes to two crucial parameters that you need to know about before reading further:
base: This is the base URL of a tenant unit, e.g. `https://test-example.instana.io`. This is the same URL that is used to access the Instana user interface.
apiToken: Requests against the Instana API require valid API tokens. An initial API token can be generated via the Instana user interface. Any additional API tokens can be generated via the API itself.

### Example
Here is an Example to use the REST API with Curl. First lets get all the available metrics with possible aggregations with a GET call.

```bash
curl --request GET \\
  --url https://test-instana.instana.io/api/application-monitoring/catalog/metrics \\
  --header 'authorization: apiToken xxxxxxxxxxxxxxxx'
```

Next we can get every call grouped by the endpoint name that has an error count greater then zero. As a metric we could get the mean error rate for example.

```bash
curl --request POST \\
  --url https://test-instana.instana.io/api/application-monitoring/analyze/call-groups \\
  --header 'authorization: apiToken xxxxxxxxxxxxxxxx' \\
  --header 'content-type: application/json' \\
  --data '{
  \"group\":{
      \"groupbyTag\":\"endpoint.name\"
  },
  \"tagFilters\":[
   {
    \"name\":\"call.error.count\",
    \"value\":\"0\",
    \"operator\":\"GREATER_THAN\"
   }
  ],
  \"metrics\":[
   {
    \"metric\":\"errors\",
    \"aggregation\":\"MEAN\"
   }
  ]
  }'
```


### Rate Limiting
A rate limit is applied to API usage. Up to 5,000 calls per hour can be made. How many remaining calls can be made and when this call limit resets, can inspected via three headers that are part of the responses of the API server.

**X-RateLimit-Limit:** Shows the maximum number of calls that may be executed per hour.

**X-RateLimit-Remaining:** How many calls may still be executed within the current hour.

**X-RateLimit-Reset:** Time when the remaining calls will be reset to the limit. For compatibility reasons with other rate limited APIs, this date is not the date in milliseconds, but instead in seconds since 1970-01-01T00:00:00+00:00.

## Generating REST API clients

The API is specified using the [OpenAPI v3](https://github.com/OAI/OpenAPI-Specification) (previously known as Swagger) format. You can download the current specification at our [GitHub API documentation](https://instana.github.io/openapi/openapi.yaml).

OpenAPI tries to solve the issue of ever-evolving APIs and clients lagging behind. To generate a client library for your language, you can use the [OpenAPI client generators](https://github.com/OpenAPITools/openapi-generator).

To generate a client library for Go to interact with our backend, you can use the following script (you need a JDK and `wget`):

```bash
//Download the generator to your current working directory:
wget http://central.maven.org/maven2/org/openapitools/openapi-generator-cli/3.2.3/openapi-generator-cli-3.2.3.jar -O openapi-generator-cli.jar

//generate a client library that you can vendor into your repository
java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g go \\
    -o pkg/instana/openapi \\
    --skip-validate-spec

//(optional) format the Go code according to the Go code standard
gofmt -s -w pkg/instana/openapi
```

The generated clients contain comprehensive READMEs. To use the client from the example above, you can start right away:

```go
import instana \"./pkg/instana/openapi\"

// readTags will read all available application monitoring tags along with their type and category
func readTags() {
 configuration := instana.NewConfiguration()
 configuration.Host = \"tenant-unit.instana.io\"
 configuration.BasePath = \"https://tenant-unit.instana.io\"

 client := instana.NewAPIClient(configuration)
 auth := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{
  Key:    apiKey,
  Prefix: \"apiToken\",
 })

 tags, _, err := client.ApplicationCatalogApi.GetTagsForApplication(auth)
 if err != nil {
  fmt.Fatalf(\"Error calling the API, aborting.\")
 }

 for _, tag := range tags {
  fmt.Printf(\"%s (%s): %s\\n\", tag.Category, tag.Type, tag.Name)
 }
}
```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.174.285
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://instana.com](http://instana.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://unit-tenant.instana.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APITokenApi* | [**DeleteApiToken**](docs/APITokenApi.md#deleteapitoken) | **Delete** /api/settings/api-tokens/{apiTokenId} | Delete API token
*APITokenApi* | [**GetApiToken**](docs/APITokenApi.md#getapitoken) | **Get** /api/settings/api-tokens/{apiTokenId} | API token
*APITokenApi* | [**GetApiTokens**](docs/APITokenApi.md#getapitokens) | **Get** /api/settings/api-tokens | All API tokes
*APITokenApi* | [**PutApiToken**](docs/APITokenApi.md#putapitoken) | **Put** /api/settings/api-tokens/{apiTokenId} | Create or update an API token
*ApplicationAnalyzeApi* | [**GetCallGroup**](docs/ApplicationAnalyzeApi.md#getcallgroup) | **Post** /api/application-monitoring/analyze/call-groups | Get grouped call metrics
*ApplicationAnalyzeApi* | [**GetTrace**](docs/ApplicationAnalyzeApi.md#gettrace) | **Get** /api/application-monitoring/analyze/traces/{id} | Get trace detail
*ApplicationAnalyzeApi* | [**GetTraceGroups**](docs/ApplicationAnalyzeApi.md#gettracegroups) | **Post** /api/application-monitoring/analyze/trace-groups | Get grouped trace metrics
*ApplicationAnalyzeApi* | [**GetTraces**](docs/ApplicationAnalyzeApi.md#gettraces) | **Post** /api/application-monitoring/analyze/traces | Get all traces
*ApplicationCatalogApi* | [**GetApplicationCatalogMetrics**](docs/ApplicationCatalogApi.md#getapplicationcatalogmetrics) | **Get** /api/application-monitoring/catalog/metrics | Get Metric catalog
*ApplicationCatalogApi* | [**GetApplicationCatalogTags**](docs/ApplicationCatalogApi.md#getapplicationcatalogtags) | **Get** /api/application-monitoring/catalog/tags | Get filter tag catalog
*ApplicationMetricsApi* | [**GetApplicationMetrics**](docs/ApplicationMetricsApi.md#getapplicationmetrics) | **Post** /api/application-monitoring/metrics/applications | Get Application Metrics
*ApplicationMetricsApi* | [**GetEndpointsMetrics**](docs/ApplicationMetricsApi.md#getendpointsmetrics) | **Post** /api/application-monitoring/metrics/endpoints | Get Endpoint metrics
*ApplicationMetricsApi* | [**GetServicesMetrics**](docs/ApplicationMetricsApi.md#getservicesmetrics) | **Post** /api/application-monitoring/metrics/services | Get Service metrics
*ApplicationResourcesApi* | [**ApplicationResourcesEndpoints**](docs/ApplicationResourcesApi.md#applicationresourcesendpoints) | **Get** /api/application-monitoring/applications/services/endpoints | Get endpoints
*ApplicationResourcesApi* | [**GetApplicationServices**](docs/ApplicationResourcesApi.md#getapplicationservices) | **Get** /api/application-monitoring/applications/services | Get applications/services
*ApplicationResourcesApi* | [**GetApplications**](docs/ApplicationResourcesApi.md#getapplications) | **Get** /api/application-monitoring/applications | Get applications
*ApplicationResourcesApi* | [**GetServices**](docs/ApplicationResourcesApi.md#getservices) | **Get** /api/application-monitoring/services | Get services
*ApplicationSettingsApi* | [**AddApplicationConfig**](docs/ApplicationSettingsApi.md#addapplicationconfig) | **Post** /api/application-monitoring/settings/application | Add application configuration
*ApplicationSettingsApi* | [**AddServiceConfig**](docs/ApplicationSettingsApi.md#addserviceconfig) | **Post** /api/application-monitoring/settings/service | Add service configuration
*ApplicationSettingsApi* | [**CreateEndpointConfig**](docs/ApplicationSettingsApi.md#createendpointconfig) | **Post** /api/application-monitoring/settings/http-endpoint | Create endpoint configuration
*ApplicationSettingsApi* | [**DeleteApplicationConfig**](docs/ApplicationSettingsApi.md#deleteapplicationconfig) | **Delete** /api/application-monitoring/settings/application/{id} | Delete application configuration
*ApplicationSettingsApi* | [**DeleteEndpointConfig**](docs/ApplicationSettingsApi.md#deleteendpointconfig) | **Delete** /api/application-monitoring/settings/http-endpoint/{id} | Delete endpoint configuration
*ApplicationSettingsApi* | [**DeleteServiceConfig**](docs/ApplicationSettingsApi.md#deleteserviceconfig) | **Delete** /api/application-monitoring/settings/service/{id} | Delete service configuration
*ApplicationSettingsApi* | [**GetApplicationConfig**](docs/ApplicationSettingsApi.md#getapplicationconfig) | **Get** /api/application-monitoring/settings/application/{id} | Application configuration
*ApplicationSettingsApi* | [**GetApplicationConfigs**](docs/ApplicationSettingsApi.md#getapplicationconfigs) | **Get** /api/application-monitoring/settings/application | All Application configurations
*ApplicationSettingsApi* | [**GetEndpointConfig**](docs/ApplicationSettingsApi.md#getendpointconfig) | **Get** /api/application-monitoring/settings/http-endpoint/{id} | Endpoint configuration
*ApplicationSettingsApi* | [**GetEndpointConfigs**](docs/ApplicationSettingsApi.md#getendpointconfigs) | **Get** /api/application-monitoring/settings/http-endpoint | All Endpoint configurations
*ApplicationSettingsApi* | [**GetServiceConfig**](docs/ApplicationSettingsApi.md#getserviceconfig) | **Get** /api/application-monitoring/settings/service/{id} | Service configuration
*ApplicationSettingsApi* | [**GetServiceConfigs**](docs/ApplicationSettingsApi.md#getserviceconfigs) | **Get** /api/application-monitoring/settings/service | All service configurations
*ApplicationSettingsApi* | [**OrderServiceConfig**](docs/ApplicationSettingsApi.md#orderserviceconfig) | **Put** /api/application-monitoring/settings/service/order | Order of service configuration
*ApplicationSettingsApi* | [**PutApplicationConfig**](docs/ApplicationSettingsApi.md#putapplicationconfig) | **Put** /api/application-monitoring/settings/application/{id} | Update application configuration
*ApplicationSettingsApi* | [**PutServiceConfig**](docs/ApplicationSettingsApi.md#putserviceconfig) | **Put** /api/application-monitoring/settings/service/{id} | Update service configuration
*ApplicationSettingsApi* | [**UpdateEndpointConfig**](docs/ApplicationSettingsApi.md#updateendpointconfig) | **Put** /api/application-monitoring/settings/http-endpoint/{id} | Update endpoint configuration
*AuditLogApi* | [**GetAuditLogs**](docs/AuditLogApi.md#getauditlogs) | **Get** /api/settings/auditlog | Audit log
*DefaultApi* | [**CreateSourceMapConfig**](docs/DefaultApi.md#createsourcemapconfig) | **Post** /api/website-monitoring/config/{websiteId}/sourceMap | 
*DefaultApi* | [**Delete**](docs/DefaultApi.md#delete) | **Delete** /api/events/settings/website-alert-configs/{id} | 
*DefaultApi* | [**DeleteSourceMapConfig**](docs/DefaultApi.md#deletesourcemapconfig) | **Delete** /api/website-monitoring/config/{websiteId}/sourceMap/{sourceMapConfigId} | 
*DefaultApi* | [**Disable**](docs/DefaultApi.md#disable) | **Put** /api/events/settings/website-alert-configs/{id}/disable | 
*DefaultApi* | [**Enable**](docs/DefaultApi.md#enable) | **Put** /api/events/settings/website-alert-configs/{id}/enable | 
*DefaultApi* | [**GetSourceMapConfig**](docs/DefaultApi.md#getsourcemapconfig) | **Get** /api/website-monitoring/config/{websiteId}/sourceMap/{sourceMapConfigId} | 
*DefaultApi* | [**GetSourceMapConfigs**](docs/DefaultApi.md#getsourcemapconfigs) | **Get** /api/website-monitoring/config/{websiteId}/sourceMap | 
*DefaultApi* | [**UpdateSourceMapConfig**](docs/DefaultApi.md#updatesourcemapconfig) | **Put** /api/website-monitoring/config/{websiteId}/sourceMap/{sourceMapConfigId} | 
*EventSettingsApi* | [**Create**](docs/EventSettingsApi.md#create) | **Post** /api/events/settings/website-alert-configs | Create Website Alert Config
*EventSettingsApi* | [**DeleteAlert**](docs/EventSettingsApi.md#deletealert) | **Delete** /api/events/settings/alerts/{id} | Delete alerting
*EventSettingsApi* | [**DeleteAlertingChannel**](docs/EventSettingsApi.md#deletealertingchannel) | **Delete** /api/events/settings/alertingChannels/{id} | Delete alerting channel
*EventSettingsApi* | [**DeleteBuiltInEventSpecification**](docs/EventSettingsApi.md#deletebuiltineventspecification) | **Delete** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Delete built-in event specification
*EventSettingsApi* | [**DeleteCustomEventSpecification**](docs/EventSettingsApi.md#deletecustomeventspecification) | **Delete** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Delete custom event specification
*EventSettingsApi* | [**DisableBuiltInEventSpecification**](docs/EventSettingsApi.md#disablebuiltineventspecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/disable | Disable built-in event specification
*EventSettingsApi* | [**DisableCustomEventSpecification**](docs/EventSettingsApi.md#disablecustomeventspecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/disable | Disable custom event specification
*EventSettingsApi* | [**EnableBuiltInEventSpecification**](docs/EventSettingsApi.md#enablebuiltineventspecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/enable | Enable built-in event specification
*EventSettingsApi* | [**EnableCustomEventSpecification**](docs/EventSettingsApi.md#enablecustomeventspecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/enable | Enable custom event specification
*EventSettingsApi* | [**Find**](docs/EventSettingsApi.md#find) | **Get** /api/events/settings/website-alert-configs/{id} | Get Website Alert Config
*EventSettingsApi* | [**FindAllActive**](docs/EventSettingsApi.md#findallactive) | **Get** /api/events/settings/website-alert-configs | All Website Alert Configs
*EventSettingsApi* | [**FindVersions**](docs/EventSettingsApi.md#findversions) | **Get** /api/events/settings/website-alert-configs/{id}/versions | Get versions of Website Alert Config
*EventSettingsApi* | [**GetAlert**](docs/EventSettingsApi.md#getalert) | **Get** /api/events/settings/alerts/{id} | Alerting
*EventSettingsApi* | [**GetAlertingChannel**](docs/EventSettingsApi.md#getalertingchannel) | **Get** /api/events/settings/alertingChannels/{id} | Alerting channel
*EventSettingsApi* | [**GetAlertingChannels**](docs/EventSettingsApi.md#getalertingchannels) | **Get** /api/events/settings/alertingChannels | All alerting channels
*EventSettingsApi* | [**GetAlertingConfigurationInfos**](docs/EventSettingsApi.md#getalertingconfigurationinfos) | **Get** /api/events/settings/alerts/infos | All alerting configuration info
*EventSettingsApi* | [**GetAlerts**](docs/EventSettingsApi.md#getalerts) | **Get** /api/events/settings/alerts | All Alerting
*EventSettingsApi* | [**GetBuiltInEventSpecification**](docs/EventSettingsApi.md#getbuiltineventspecification) | **Get** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Built-in event specifications
*EventSettingsApi* | [**GetBuiltInEventSpecifications**](docs/EventSettingsApi.md#getbuiltineventspecifications) | **Get** /api/events/settings/event-specifications/built-in | All built-in event specification
*EventSettingsApi* | [**GetCustomEventSpecification**](docs/EventSettingsApi.md#getcustomeventspecification) | **Get** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Custom event specification
*EventSettingsApi* | [**GetCustomEventSpecifications**](docs/EventSettingsApi.md#getcustomeventspecifications) | **Get** /api/events/settings/event-specifications/custom | All custom event specifications
*EventSettingsApi* | [**GetEventSpecificationInfos**](docs/EventSettingsApi.md#geteventspecificationinfos) | **Get** /api/events/settings/event-specifications/infos | Summary of all built-in and custom event specifications
*EventSettingsApi* | [**GetEventSpecificationInfosByIds**](docs/EventSettingsApi.md#geteventspecificationinfosbyids) | **Post** /api/events/settings/event-specifications/infos | All built-in and custom event specifications
*EventSettingsApi* | [**GetSystemRules**](docs/EventSettingsApi.md#getsystemrules) | **Get** /api/events/settings/event-specifications/custom/systemRules | All system rules for custom event specifications
*EventSettingsApi* | [**PutAlert**](docs/EventSettingsApi.md#putalert) | **Put** /api/events/settings/alerts/{id} | Update alerting
*EventSettingsApi* | [**PutAlertingChannel**](docs/EventSettingsApi.md#putalertingchannel) | **Put** /api/events/settings/alertingChannels/{id} | Update alerting channel
*EventSettingsApi* | [**PutCustomEventSpecification**](docs/EventSettingsApi.md#putcustomeventspecification) | **Put** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Update custom event specification
*EventSettingsApi* | [**SendTestAlerting**](docs/EventSettingsApi.md#sendtestalerting) | **Put** /api/events/settings/alertingChannels/test | Test alerting channel
*EventSettingsApi* | [**Update**](docs/EventSettingsApi.md#update) | **Post** /api/events/settings/website-alert-configs/{id} | Update Website Alert Config
*EventsApi* | [**GetEvent**](docs/EventsApi.md#getevent) | **Get** /api/events/{eventId} | Get Event
*EventsApi* | [**GetEvents**](docs/EventsApi.md#getevents) | **Get** /api/events | Get alerts
*HealthApi* | [**GetHealthState**](docs/HealthApi.md#gethealthstate) | **Get** /api/instana/health | Basic health traffic light
*HealthApi* | [**GetVersion**](docs/HealthApi.md#getversion) | **Get** /api/instana/version | API version information
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogMetrics**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogmetrics) | **Get** /api/infrastructure-monitoring/catalog/metrics/{plugin} | Get metric catalog
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogPlugins**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogplugins) | **Get** /api/infrastructure-monitoring/catalog/plugins | Get plugin catalog
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogSearchFields**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogsearchfields) | **Get** /api/infrastructure-monitoring/catalog/search | get search field catalog
*InfrastructureMetricsApi* | [**GetInfrastructureMetrics**](docs/InfrastructureMetricsApi.md#getinfrastructuremetrics) | **Post** /api/infrastructure-monitoring/metrics | Get infrastructure metrics
*InfrastructureMetricsApi* | [**GetSnapshot**](docs/InfrastructureMetricsApi.md#getsnapshot) | **Get** /api/infrastructure-monitoring/snapshots/{id} | Get snapshot details
*InfrastructureMetricsApi* | [**GetSnapshots**](docs/InfrastructureMetricsApi.md#getsnapshots) | **Get** /api/infrastructure-monitoring/snapshots | Search snapshots
*InfrastructureResourcesApi* | [**GetInfrastructureViewTree**](docs/InfrastructureResourcesApi.md#getinfrastructureviewtree) | **Get** /api/infrastructure-monitoring/graph/views | Get view tree
*InfrastructureResourcesApi* | [**GetMonitoringState**](docs/InfrastructureResourcesApi.md#getmonitoringstate) | **Get** /api/infrastructure-monitoring/monitoring-state | Monitored host count
*InfrastructureResourcesApi* | [**GetRelatedHosts**](docs/InfrastructureResourcesApi.md#getrelatedhosts) | **Get** /api/infrastructure-monitoring/graph/related-hosts/{snapshotId} | Related hosts
*InfrastructureResourcesApi* | [**SoftwareVersions**](docs/InfrastructureResourcesApi.md#softwareversions) | **Get** /api/infrastructure-monitoring/software/versions | Get installed software
*MaintenanceConfigurationApi* | [**DeleteMaintenanceConfig**](docs/MaintenanceConfigurationApi.md#deletemaintenanceconfig) | **Delete** /api/settings/maintenance/{id} | Delete maintenance configuration
*MaintenanceConfigurationApi* | [**GetMaintenanceConfig**](docs/MaintenanceConfigurationApi.md#getmaintenanceconfig) | **Get** /api/settings/maintenance/{id} | Maintenance configuration
*MaintenanceConfigurationApi* | [**GetMaintenanceConfigs**](docs/MaintenanceConfigurationApi.md#getmaintenanceconfigs) | **Get** /api/settings/maintenance | All maintenance configurations
*MaintenanceConfigurationApi* | [**PutMaintenanceConfig**](docs/MaintenanceConfigurationApi.md#putmaintenanceconfig) | **Put** /api/settings/maintenance/{id} | Create or update maintenance configuration
*ReleasesApi* | [**DeleteRelease**](docs/ReleasesApi.md#deleterelease) | **Delete** /api/releases/{releaseId} | Delete release
*ReleasesApi* | [**GetAllReleases**](docs/ReleasesApi.md#getallreleases) | **Get** /api/releases | Get all releases
*ReleasesApi* | [**GetRelease**](docs/ReleasesApi.md#getrelease) | **Get** /api/releases/{releaseId} | Get release
*ReleasesApi* | [**PostRelease**](docs/ReleasesApi.md#postrelease) | **Post** /api/releases | Create release
*ReleasesApi* | [**PutRelease**](docs/ReleasesApi.md#putrelease) | **Put** /api/releases/{releaseId} | Update release
*SyntheticCallsApi* | [**DeleteSyntheticCall**](docs/SyntheticCallsApi.md#deletesyntheticcall) | **Delete** /api/settings/synthetic-calls | Delete synthetic call configurations
*SyntheticCallsApi* | [**GetSyntheticCalls**](docs/SyntheticCallsApi.md#getsyntheticcalls) | **Get** /api/settings/synthetic-calls | Synthetic call configurations
*SyntheticCallsApi* | [**UpdateSyntheticCall**](docs/SyntheticCallsApi.md#updatesyntheticcall) | **Put** /api/settings/synthetic-calls | Update synthetic call configurations
*UsageApi* | [**GetAllUsage**](docs/UsageApi.md#getallusage) | **Get** /api/instana/usage/api | API usage by customer
*UsageApi* | [**GetHostsPerDay**](docs/UsageApi.md#gethostsperday) | **Get** /api/instana/usage/hosts/{day}/{month}/{year} | Host count day / month / year
*UsageApi* | [**GetHostsPerMonth**](docs/UsageApi.md#gethostspermonth) | **Get** /api/instana/usage/hosts/{month}/{year} | Host count month / year
*UsageApi* | [**GetUsagePerDay**](docs/UsageApi.md#getusageperday) | **Get** /api/instana/usage/api/{day}/{month}/{year} | API usage day / month / year
*UsageApi* | [**GetUsagePerMonth**](docs/UsageApi.md#getusagepermonth) | **Get** /api/instana/usage/api/{month}/{year} | API usage month / year
*UserApi* | [**DeleteRole**](docs/UserApi.md#deleterole) | **Delete** /api/settings/roles/{roleId} | Delete role
*UserApi* | [**GetInvitations**](docs/UserApi.md#getinvitations) | **Get** /api/settings/users/invitations | All pending invitations
*UserApi* | [**GetRole**](docs/UserApi.md#getrole) | **Get** /api/settings/roles/{roleId} | Role
*UserApi* | [**GetRoles**](docs/UserApi.md#getroles) | **Get** /api/settings/roles | All roles
*UserApi* | [**GetUsers**](docs/UserApi.md#getusers) | **Get** /api/settings/users | All users (without invitations)
*UserApi* | [**GetUsersIncludingInvitations**](docs/UserApi.md#getusersincludinginvitations) | **Get** /api/settings/users/overview | All users (incl. invitations)
*UserApi* | [**PutRole**](docs/UserApi.md#putrole) | **Put** /api/settings/roles/{roleId} | Create or update role
*UserApi* | [**RemoveUserFromTenant**](docs/UserApi.md#removeuserfromtenant) | **Delete** /api/settings/users/{userId} | Remove user from tenant
*UserApi* | [**RevokePendingInvitations**](docs/UserApi.md#revokependinginvitations) | **Delete** /api/settings/users/invitations | Revoke pending invitation
*UserApi* | [**SendInvitation**](docs/UserApi.md#sendinvitation) | **Post** /api/settings/users/invitations | Send user invitation
*UserApi* | [**SetRole**](docs/UserApi.md#setrole) | **Put** /api/settings/users/{userId}/role | Add user to role
*WebsiteAnalyzeApi* | [**GetBeaconGroups**](docs/WebsiteAnalyzeApi.md#getbeacongroups) | **Post** /api/website-monitoring/analyze/beacon-groups | Get grouped beacon metrics
*WebsiteAnalyzeApi* | [**GetBeacons**](docs/WebsiteAnalyzeApi.md#getbeacons) | **Post** /api/website-monitoring/analyze/beacons | Get all beacons
*WebsiteCatalogApi* | [**GetWebsiteCatalogMetrics**](docs/WebsiteCatalogApi.md#getwebsitecatalogmetrics) | **Get** /api/website-monitoring/catalog/metrics | Metric catalog
*WebsiteCatalogApi* | [**GetWebsiteCatalogTags**](docs/WebsiteCatalogApi.md#getwebsitecatalogtags) | **Get** /api/website-monitoring/catalog/tags | Filter tag catalog
*WebsiteConfigurationApi* | [**Delete1**](docs/WebsiteConfigurationApi.md#delete1) | **Delete** /api/website-monitoring/config/{websiteId} | Remove website
*WebsiteConfigurationApi* | [**Get**](docs/WebsiteConfigurationApi.md#get) | **Get** /api/website-monitoring/config | Get configured websites
*WebsiteConfigurationApi* | [**Post**](docs/WebsiteConfigurationApi.md#post) | **Post** /api/website-monitoring/config | Configure new website
*WebsiteConfigurationApi* | [**Rename**](docs/WebsiteConfigurationApi.md#rename) | **Put** /api/website-monitoring/config/{websiteId} | Rename website
*WebsiteMetricsApi* | [**GetBeaconMetrics**](docs/WebsiteMetricsApi.md#getbeaconmetrics) | **Post** /api/website-monitoring/metrics | Get beacon metrics
*WebsiteMetricsApi* | [**GetPageLoad**](docs/WebsiteMetricsApi.md#getpageload) | **Get** /api/website-monitoring/page-load | Get page load


## Documentation For Models

 - [AbstractIntegration](docs/AbstractIntegration.md)
 - [AbstractRule](docs/AbstractRule.md)
 - [AlertingConfiguration](docs/AlertingConfiguration.md)
 - [AlertingConfigurationWithLastUpdated](docs/AlertingConfigurationWithLastUpdated.md)
 - [ApiToken](docs/ApiToken.md)
 - [AppDataMetricConfiguration](docs/AppDataMetricConfiguration.md)
 - [Application](docs/Application.md)
 - [ApplicationConfig](docs/ApplicationConfig.md)
 - [ApplicationItem](docs/ApplicationItem.md)
 - [ApplicationMetricResult](docs/ApplicationMetricResult.md)
 - [ApplicationResult](docs/ApplicationResult.md)
 - [AuditLogEntry](docs/AuditLogEntry.md)
 - [AuditLogResponse](docs/AuditLogResponse.md)
 - [BeaconGroupsResult](docs/BeaconGroupsResult.md)
 - [BeaconResult](docs/BeaconResult.md)
 - [BinaryOperatorDto](docs/BinaryOperatorDto.md)
 - [BuiltInEventSpecification](docs/BuiltInEventSpecification.md)
 - [BuiltInEventSpecificationWithLastUpdated](docs/BuiltInEventSpecificationWithLastUpdated.md)
 - [CallGroupsItem](docs/CallGroupsItem.md)
 - [CallGroupsResult](docs/CallGroupsResult.md)
 - [CloudfoundryPhysicalContext](docs/CloudfoundryPhysicalContext.md)
 - [ConfigVersion](docs/ConfigVersion.md)
 - [CursorPagination](docs/CursorPagination.md)
 - [CustomEventSpecification](docs/CustomEventSpecification.md)
 - [CustomEventSpecificationWithLastUpdated](docs/CustomEventSpecificationWithLastUpdated.md)
 - [EmailIntegration](docs/EmailIntegration.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointItem](docs/EndpointItem.md)
 - [EndpointMetricResult](docs/EndpointMetricResult.md)
 - [EndpointResult](docs/EndpointResult.md)
 - [EntityVerificationRule](docs/EntityVerificationRule.md)
 - [EventFilteringConfiguration](docs/EventFilteringConfiguration.md)
 - [EventResult](docs/EventResult.md)
 - [EventSpecificationInfo](docs/EventSpecificationInfo.md)
 - [FixedHttpPathSegmentMatchingRule](docs/FixedHttpPathSegmentMatchingRule.md)
 - [FullTrace](docs/FullTrace.md)
 - [GetApplications](docs/GetApplications.md)
 - [GetCallGroups](docs/GetCallGroups.md)
 - [GetCombinedMetrics](docs/GetCombinedMetrics.md)
 - [GetEndpoints](docs/GetEndpoints.md)
 - [GetServices](docs/GetServices.md)
 - [GetTraceGroups](docs/GetTraceGroups.md)
 - [GetTraces](docs/GetTraces.md)
 - [GetWebsiteBeaconGroups](docs/GetWebsiteBeaconGroups.md)
 - [GetWebsiteBeacons](docs/GetWebsiteBeacons.md)
 - [GetWebsiteMetrics](docs/GetWebsiteMetrics.md)
 - [GoogleChatIntegration](docs/GoogleChatIntegration.md)
 - [Group](docs/Group.md)
 - [HealthState](docs/HealthState.md)
 - [HistoricBaseline](docs/HistoricBaseline.md)
 - [HttpEndpointConfig](docs/HttpEndpointConfig.md)
 - [HttpEndpointRule](docs/HttpEndpointRule.md)
 - [HttpPathSegmentMatchingRule](docs/HttpPathSegmentMatchingRule.md)
 - [HyperParam](docs/HyperParam.md)
 - [InfrastructureMetricResult](docs/InfrastructureMetricResult.md)
 - [IngestionOffsetCursor](docs/IngestionOffsetCursor.md)
 - [InstanaVersionInfo](docs/InstanaVersionInfo.md)
 - [KubernetesPhysicalContext](docs/KubernetesPhysicalContext.md)
 - [LogEntryActor](docs/LogEntryActor.md)
 - [MaintenanceConfig](docs/MaintenanceConfig.md)
 - [MaintenanceConfigWithLastUpdated](docs/MaintenanceConfigWithLastUpdated.md)
 - [MaintenanceWindow](docs/MaintenanceWindow.md)
 - [MatchAllHttpPathSegmentMatchingRule](docs/MatchAllHttpPathSegmentMatchingRule.md)
 - [MetricConfiguration](docs/MetricConfiguration.md)
 - [MetricDescription](docs/MetricDescription.md)
 - [MetricInstance](docs/MetricInstance.md)
 - [MetricItem](docs/MetricItem.md)
 - [MonitoringState](docs/MonitoringState.md)
 - [Office365Integration](docs/Office365Integration.md)
 - [OpsgenieIntegration](docs/OpsgenieIntegration.md)
 - [Order](docs/Order.md)
 - [PagerdutyIntegration](docs/PagerdutyIntegration.md)
 - [Pagination](docs/Pagination.md)
 - [PathParameterHttpPathSegmentMatchingRule](docs/PathParameterHttpPathSegmentMatchingRule.md)
 - [PhysicalContext](docs/PhysicalContext.md)
 - [PluginResult](docs/PluginResult.md)
 - [PrometheusWebhookIntegration](docs/PrometheusWebhookIntegration.md)
 - [Release](docs/Release.md)
 - [ReleaseWithMetadata](docs/ReleaseWithMetadata.md)
 - [Role](docs/Role.md)
 - [RuleInput](docs/RuleInput.md)
 - [SearchFieldResult](docs/SearchFieldResult.md)
 - [Service](docs/Service.md)
 - [ServiceConfig](docs/ServiceConfig.md)
 - [ServiceItem](docs/ServiceItem.md)
 - [ServiceMatchingRule](docs/ServiceMatchingRule.md)
 - [ServiceMetricResult](docs/ServiceMetricResult.md)
 - [ServiceResult](docs/ServiceResult.md)
 - [SlackIntegration](docs/SlackIntegration.md)
 - [SlownessWebsiteAlertRule](docs/SlownessWebsiteAlertRule.md)
 - [SnapshotItem](docs/SnapshotItem.md)
 - [SnapshotPreview](docs/SnapshotPreview.md)
 - [SnapshotResult](docs/SnapshotResult.md)
 - [SoftwareUser](docs/SoftwareUser.md)
 - [SoftwareVersion](docs/SoftwareVersion.md)
 - [Span](docs/Span.md)
 - [SpanRelation](docs/SpanRelation.md)
 - [SpecificJsErrorsWebsiteAlertRule](docs/SpecificJsErrorsWebsiteAlertRule.md)
 - [SplunkIntegration](docs/SplunkIntegration.md)
 - [StackTraceItem](docs/StackTraceItem.md)
 - [StackTraceLine](docs/StackTraceLine.md)
 - [StaticThreshold](docs/StaticThreshold.md)
 - [StatusCodeWebsiteAlertRule](docs/StatusCodeWebsiteAlertRule.md)
 - [SyntheticCallConfig](docs/SyntheticCallConfig.md)
 - [SyntheticCallRule](docs/SyntheticCallRule.md)
 - [SyntheticCallWithDefaultsConfig](docs/SyntheticCallWithDefaultsConfig.md)
 - [SystemRule](docs/SystemRule.md)
 - [SystemRuleLabel](docs/SystemRuleLabel.md)
 - [Tag](docs/Tag.md)
 - [TagFilter](docs/TagFilter.md)
 - [TagMatcherDto](docs/TagMatcherDto.md)
 - [Threshold](docs/Threshold.md)
 - [ThresholdRule](docs/ThresholdRule.md)
 - [TimeFrame](docs/TimeFrame.md)
 - [Trace](docs/Trace.md)
 - [TraceGroupsItem](docs/TraceGroupsItem.md)
 - [TraceGroupsResult](docs/TraceGroupsResult.md)
 - [TraceItem](docs/TraceItem.md)
 - [TraceResult](docs/TraceResult.md)
 - [TreeNode](docs/TreeNode.md)
 - [TreeNodeResult](docs/TreeNodeResult.md)
 - [UnsupportedHttpPathSegmentMatchingRule](docs/UnsupportedHttpPathSegmentMatchingRule.md)
 - [UsageResult](docs/UsageResult.md)
 - [UsageResultItems](docs/UsageResultItems.md)
 - [UserImpactWebsiteTimeThreshold](docs/UserImpactWebsiteTimeThreshold.md)
 - [UserResult](docs/UserResult.md)
 - [UsersResult](docs/UsersResult.md)
 - [ValidatedAlertingChannelInputInfo](docs/ValidatedAlertingChannelInputInfo.md)
 - [ValidatedAlertingConfiguration](docs/ValidatedAlertingConfiguration.md)
 - [ValidatedMaintenanceConfigWithStatus](docs/ValidatedMaintenanceConfigWithStatus.md)
 - [VictorOpsIntegration](docs/VictorOpsIntegration.md)
 - [ViolationsInPeriodWebsiteTimeThreshold](docs/ViolationsInPeriodWebsiteTimeThreshold.md)
 - [ViolationsInSequenceWebsiteTimeThreshold](docs/ViolationsInSequenceWebsiteTimeThreshold.md)
 - [WebhookIntegration](docs/WebhookIntegration.md)
 - [Website](docs/Website.md)
 - [WebsiteAlertConfig](docs/WebsiteAlertConfig.md)
 - [WebsiteAlertConfigWithMetadata](docs/WebsiteAlertConfigWithMetadata.md)
 - [WebsiteAlertRule](docs/WebsiteAlertRule.md)
 - [WebsiteBeaconGroupsItem](docs/WebsiteBeaconGroupsItem.md)
 - [WebsiteBeaconTagGroup](docs/WebsiteBeaconTagGroup.md)
 - [WebsiteBeaconsItem](docs/WebsiteBeaconsItem.md)
 - [WebsiteMetricResult](docs/WebsiteMetricResult.md)
 - [WebsiteMonitoringBeacon](docs/WebsiteMonitoringBeacon.md)
 - [WebsiteMonitoringMetricDescription](docs/WebsiteMonitoringMetricDescription.md)
 - [WebsiteMonitoringMetricsConfiguration](docs/WebsiteMonitoringMetricsConfiguration.md)
 - [WebsiteTimeThreshold](docs/WebsiteTimeThreshold.md)


## Documentation For Authorization



## ApiKeyAuth

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```



## Author

support@instana.com

