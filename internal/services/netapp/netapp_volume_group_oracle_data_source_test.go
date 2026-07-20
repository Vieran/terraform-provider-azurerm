// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package netapp_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type NetAppVolumeGroupOracleDataSource struct{}

func TestAccNetAppVolumeGroupOracleDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_netapp_volume_group_oracle", "test")
	d := NetAppVolumeGroupOracleDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: d.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("resource_group_name").Exists(),
				check.That(data.ResourceName).Key("volume.1.volume_spec_name").HasValue("ora-log"),
				check.That(data.ResourceName).Key("volume.0.mount_target.0.ip_address").Exists(),
				check.That(data.ResourceName).Key("volume.0.mount_target.0.smb_server_fqdn").IsEmpty(),
			),
		},
	})
}

func (d NetAppVolumeGroupOracleDataSource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_netapp_volume_group_oracle" "test" {
  name                = azurerm_netapp_volume_group_oracle.test.name
  resource_group_name = azurerm_netapp_volume_group_oracle.test.resource_group_name
  account_name        = azurerm_netapp_volume_group_oracle.test.account_name
}
`, NetAppVolumeGroupOracleResource{}.basicAvailabilityZone(data))
}
