// Code generated by tutone: DO NOT EDIT
package dashboards

import (
	"github.com/newrelic/newrelic-client-go/pkg/entities"
)

// Create a `DashboardEntity`
func (a *Dashboards) DashboardCreate(
	accountID int,
	dashboard DashboardInput,
) (*DashboardCreateResult, error) {

	resp := DashboardCreateQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"dashboard": dashboard,
	}

	if err := a.client.NerdGraphQuery(DashboardCreateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.DashboardCreateResult, nil
}

type DashboardCreateQueryResponse struct {
	DashboardCreateResult DashboardCreateResult `json:"DashboardCreate"`
}

const DashboardCreateMutation = `mutation(
	$accountId: Int!,
	$dashboard: DashboardInput!,
) { dashboardCreate(
	accountId: $accountId,
	dashboard: $dashboard,
) {
	entityResult {
		accountId
		createdAt
		description
		guid
		name
		owner {
			email
			userId
		}
		pages {
			createdAt
			description
			guid
			name
			owner {
				email
				userId
			}
			updatedAt
			widgets {
				configuration {
					area {
						nrqlQueries {
							accountId
							query
						}
					}
					bar {
						nrqlQueries {
							accountId
							query
						}
					}
					billboard {
						nrqlQueries {
							accountId
							query
						}
						thresholds {
							alertSeverity
							value
						}
					}
					line {
						nrqlQueries {
							accountId
							query
						}
					}
					markdown {
						text
					}
					pie {
						nrqlQueries {
							accountId
							query
						}
					}
					table {
						nrqlQueries {
							accountId
							query
						}
					}
				}
				id
				layout {
					column
					height
					row
					width
				}
				linkedEntities {
					__typename
					account {
						id
						name
						reportingEventTypes
					}
					accountId
					domain
					entityType
					goldenMetrics {
						context {
							account
							guid
						}
						metrics {
							definition {
								eventId
								eventObjectId
								facet
								from
								select
								where
							}
							name
							query
							title
						}
					}
					goldenTags {
						context {
							account
							guid
						}
						tags {
							key
						}
					}
					guid
					indexedAt
					name
					permalink
					reporting
					tags {
						key
						values
					}
					type
					... on ApmApplicationEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						apmBrowserSummary {
							ajaxRequestThroughput
							ajaxResponseTimeAverage
							jsErrorRate
							pageLoadThroughput
							pageLoadTimeAverage
						}
						apmSummary {
							apdexScore
							errorRate
							hostCount
							instanceCount
							nonWebResponseTimeAverage
							nonWebThroughput
							responseTimeAverage
							throughput
							webResponseTimeAverage
							webThroughput
						}
						applicationId
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						language
						runningAgentVersions {
							maxVersion
							minVersion
						}
						settings {
							apdexTarget
							serverSideConfig
						}
						tags {
							key
							values
						}
					}
					... on ApmDatabaseInstanceEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						host
						portOrPath
						tags {
							key
							values
						}
						vendor
					}
					... on ApmExternalServiceEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						externalSummary {
							responseTimeAverage
							throughput
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						host
						tags {
							key
							values
						}
					}
					... on BrowserApplicationEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						agentInstallType
						alertSeverity
						applicationId
						browserSummary {
							ajaxRequestThroughput
							ajaxResponseTimeAverage
							jsErrorRate
							pageLoadThroughput
							pageLoadTimeAverage
							pageLoadTimeMedian
							spaResponseTimeAverage
							spaResponseTimeMedian
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						runningAgentVersions {
							maxVersion
							minVersion
						}
						servingApmApplicationId
						settings {
							apdexTarget
						}
						tags {
							key
							values
						}
					}
					... on DashboardEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						dashboardParentGuid
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						owner {
							email
							userId
						}
						permissions
						tags {
							key
							values
						}
					}
					... on ExternalEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on GenericEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on GenericInfrastructureEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						integrationTypeCode
						tags {
							key
							values
						}
					}
					... on InfrastructureAwsLambdaFunctionEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						integrationTypeCode
						runtime
						tags {
							key
							values
						}
					}
					... on InfrastructureHostEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						hostSummary {
							cpuUtilizationPercent
							diskUsedPercent
							memoryUsedPercent
							networkReceiveRate
							networkTransmitRate
							servicesCount
						}
						tags {
							key
							values
						}
					}
					... on MobileApplicationEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						applicationId
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						mobileSummary {
							appLaunchCount
							crashCount
							crashRate
							httpErrorRate
							httpRequestCount
							httpRequestRate
							httpResponseTimeAverage
							mobileSessionCount
							networkFailureRate
							usersAffectedCount
						}
						tags {
							key
							values
						}
					}
					... on SecureCredentialEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						description
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						secureCredentialId
						secureCredentialSummary {
							failingMonitorCount
							monitorCount
						}
						tags {
							key
							values
						}
						updatedAt
					}
					... on SyntheticMonitorEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						monitorId
						monitorSummary {
							locationsFailing
							locationsRunning
							status
							successRate
						}
						monitorType
						monitoredUrl
						period
						tags {
							key
							values
						}
					}
					... on ThirdPartyServiceEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on UnavailableEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on WorkloadEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						createdAt
						createdByUser {
							email
							gravatar
							id
							name
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
						updatedAt
						workloadStatus {
							description
							statusSource
							statusValue
							summary
						}
					}
				}
				rawConfiguration
				title
				visualization {
					id
				}
			}
		}
		permissions
		updatedAt
	}
	errors {
		description
		type
	}
} }`

// Delete an existing `DashboardEntity`
func (a *Dashboards) DashboardDelete(
	gUID entities.EntityGUID,
) (*DashboardDeleteResult, error) {

	resp := DashboardDeleteQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQuery(DashboardDeleteMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.DashboardDeleteResult, nil
}

type DashboardDeleteQueryResponse struct {
	DashboardDeleteResult DashboardDeleteResult `json:"DashboardDelete"`
}

const DashboardDeleteMutation = `mutation(
	$guid: EntityGuid!,
) { dashboardDelete(
	guid: $guid,
) {
	errors {
		description
		type
	}
	status
} }`

// Update an existing `DashboardEntity`
func (a *Dashboards) DashboardUpdate(
	dashboard DashboardInput,
	gUID entities.EntityGUID,
) (*DashboardUpdateResult, error) {

	resp := DashboardUpdateQueryResponse{}
	vars := map[string]interface{}{
		"dashboard": dashboard,
		"guid":      gUID,
	}

	if err := a.client.NerdGraphQuery(DashboardUpdateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.DashboardUpdateResult, nil
}

type DashboardUpdateQueryResponse struct {
	DashboardUpdateResult DashboardUpdateResult `json:"DashboardUpdate"`
}

const DashboardUpdateMutation = `mutation(
	$dashboard: DashboardInput!,
	$guid: EntityGuid!,
) { dashboardUpdate(
	dashboard: $dashboard,
	guid: $guid,
) {
	entityResult {
		accountId
		createdAt
		description
		guid
		name
		owner {
			email
			userId
		}
		pages {
			createdAt
			description
			guid
			name
			owner {
				email
				userId
			}
			updatedAt
			widgets {
				configuration {
					area {
						nrqlQueries {
							accountId
							query
						}
					}
					bar {
						nrqlQueries {
							accountId
							query
						}
					}
					billboard {
						nrqlQueries {
							accountId
							query
						}
						thresholds {
							alertSeverity
							value
						}
					}
					line {
						nrqlQueries {
							accountId
							query
						}
					}
					markdown {
						text
					}
					pie {
						nrqlQueries {
							accountId
							query
						}
					}
					table {
						nrqlQueries {
							accountId
							query
						}
					}
				}
				id
				layout {
					column
					height
					row
					width
				}
				linkedEntities {
					__typename
					account {
						id
						name
						reportingEventTypes
					}
					accountId
					domain
					entityType
					goldenMetrics {
						context {
							account
							guid
						}
						metrics {
							definition {
								eventId
								eventObjectId
								facet
								from
								select
								where
							}
							name
							query
							title
						}
					}
					goldenTags {
						context {
							account
							guid
						}
						tags {
							key
						}
					}
					guid
					indexedAt
					name
					permalink
					reporting
					tags {
						key
						values
					}
					type
					... on ApmApplicationEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						apmBrowserSummary {
							ajaxRequestThroughput
							ajaxResponseTimeAverage
							jsErrorRate
							pageLoadThroughput
							pageLoadTimeAverage
						}
						apmSummary {
							apdexScore
							errorRate
							hostCount
							instanceCount
							nonWebResponseTimeAverage
							nonWebThroughput
							responseTimeAverage
							throughput
							webResponseTimeAverage
							webThroughput
						}
						applicationId
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						language
						runningAgentVersions {
							maxVersion
							minVersion
						}
						settings {
							apdexTarget
							serverSideConfig
						}
						tags {
							key
							values
						}
					}
					... on ApmDatabaseInstanceEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						host
						portOrPath
						tags {
							key
							values
						}
						vendor
					}
					... on ApmExternalServiceEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						externalSummary {
							responseTimeAverage
							throughput
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						host
						tags {
							key
							values
						}
					}
					... on BrowserApplicationEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						agentInstallType
						alertSeverity
						applicationId
						browserSummary {
							ajaxRequestThroughput
							ajaxResponseTimeAverage
							jsErrorRate
							pageLoadThroughput
							pageLoadTimeAverage
							pageLoadTimeMedian
							spaResponseTimeAverage
							spaResponseTimeMedian
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						runningAgentVersions {
							maxVersion
							minVersion
						}
						servingApmApplicationId
						settings {
							apdexTarget
						}
						tags {
							key
							values
						}
					}
					... on DashboardEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						dashboardParentGuid
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						owner {
							email
							userId
						}
						permissions
						tags {
							key
							values
						}
					}
					... on ExternalEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on GenericEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on GenericInfrastructureEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						integrationTypeCode
						tags {
							key
							values
						}
					}
					... on InfrastructureAwsLambdaFunctionEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						integrationTypeCode
						runtime
						tags {
							key
							values
						}
					}
					... on InfrastructureHostEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						hostSummary {
							cpuUtilizationPercent
							diskUsedPercent
							memoryUsedPercent
							networkReceiveRate
							networkTransmitRate
							servicesCount
						}
						tags {
							key
							values
						}
					}
					... on MobileApplicationEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						applicationId
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						mobileSummary {
							appLaunchCount
							crashCount
							crashRate
							httpErrorRate
							httpRequestCount
							httpRequestRate
							httpResponseTimeAverage
							mobileSessionCount
							networkFailureRate
							usersAffectedCount
						}
						tags {
							key
							values
						}
					}
					... on SecureCredentialEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						description
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						secureCredentialId
						secureCredentialSummary {
							failingMonitorCount
							monitorCount
						}
						tags {
							key
							values
						}
						updatedAt
					}
					... on SyntheticMonitorEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						monitorId
						monitorSummary {
							locationsFailing
							locationsRunning
							status
							successRate
						}
						monitorType
						monitoredUrl
						period
						tags {
							key
							values
						}
					}
					... on ThirdPartyServiceEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on UnavailableEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
					}
					... on WorkloadEntityOutline {
						__typename
						account {
							id
							name
							reportingEventTypes
						}
						alertSeverity
						createdAt
						createdByUser {
							email
							gravatar
							id
							name
						}
						goldenMetrics {
							context {
								account
								guid
							}
							metrics {
								definition {
									eventId
									eventObjectId
									facet
									from
									select
									where
								}
								name
								query
								title
							}
						}
						goldenTags {
							context {
								account
								guid
							}
							tags {
								key
							}
						}
						tags {
							key
							values
						}
						updatedAt
						workloadStatus {
							description
							statusSource
							statusValue
							summary
						}
					}
				}
				rawConfiguration
				title
				visualization {
					id
				}
			}
		}
		permissions
		updatedAt
	}
	errors {
		description
		type
	}
} }`
