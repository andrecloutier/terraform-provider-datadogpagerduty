Terraform Provider
==================

Custom provider which allows per service configuration of the DataDog PagerDuty integration.

This is a stop gap till [something like this gets merged into the main client](https://github.com/terraform-providers/terraform-provider-datadog/pull/164).

Sample Usage
-----

```
provider "pagerduty" {
  token = "aaaaaa"
}

provider "datadog" {
  api_key = "bbbbbb"
  app_key = "cccccc"
}

provider "datadogpagerduty" {
  api_key = "bbbbbb"
  app_key = "cccccc"
}

data "pagerduty_vendor" "datadog" {
  name = "Datadog"
}

resource "pagerduty_service" "alert_query_service" {
  name                    = "Another Alert Query Service"
  auto_resolve_timeout    = 14400
  acknowledgement_timeout = 600
  escalation_policy       = "P2EUG11"
  alert_creation          = "create_alerts_and_incidents"
}

resource "pagerduty_service_integration" "alert_query_service_datadog" {
  name    = "${data.pagerduty_vendor.datadog.name}"
  service = "${pagerduty_service.alert_query_service.id}"
  vendor  = "${data.pagerduty_vendor.datadog.id}"
}

resource "datadogpagerduty_service_integration" "pd_aqs" {
  service_name = "AnotherAlertQueryServices"
  service_key  = "${pagerduty_service_integration.alert_query_service_datadog.integration_key}"
}

resource "datadog_monitor" "foo" {
  name                = "Monitor for AnotherAQS"
  type                = "metric alert"
  message             = "Monitor triggered.  Notify: ${datadog_pagerduty_service_integration.pd_aqs.notify_handle}"
  query               = "avg(last_1h):avg:aws.ec2.cpuutilization{*} > 50"
  notify_no_data      = false
  require_full_window = false
  tags                = ["terraform:true"]
}
```
