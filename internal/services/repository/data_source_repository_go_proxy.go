package repository

import (
	"github.com/datadrivers/terraform-provider-nexus/internal/schema/common"
	repositorySchema "github.com/datadrivers/terraform-provider-nexus/internal/schema/repository"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRepositoryGoProxy() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get an existing go proxy repository.",

		Read: dataSourceRepositoryGoProxyRead,

		Schema: map[string]*schema.Schema{
			// Common schemas
			"id":     common.DataSourceID,
			"name":   repositorySchema.DataSourceName,
			"online": repositorySchema.DataSourceOnline,
			// Proxy schemas
			"cleanup":        repositorySchema.DataSourceCleanup,
			"http_client":    repositorySchema.DataSourceHTTPClientWithPreemptiveAuth,
			"negative_cache": repositorySchema.DataSourceNegativeCache,
			"proxy":          repositorySchema.DataSourceProxy,
			"routing_rule":   repositorySchema.DataSourceRoutingRule,
			"storage":        repositorySchema.DataSourceStorage,
		},
	}
}

func dataSourceRepositoryGoProxyRead(resourceData *schema.ResourceData, m interface{}) error {
	resourceData.SetId(resourceData.Get("name").(string))

	return resourceGoProxyRepositoryRead(resourceData, m)
}
