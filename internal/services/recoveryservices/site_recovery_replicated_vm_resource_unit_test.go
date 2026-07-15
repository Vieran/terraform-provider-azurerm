// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package recoveryservices

import (
	"strings"
	"testing"
)

func TestSiteRecoveryReplicatedVMValidateManagedDiskIDs(t *testing.T) {
	diskId := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/source/providers/Microsoft.Compute/disks/example"

	tests := map[string]struct {
		disks   []interface{}
		wantErr bool
	}{
		"unique": {
			disks: []interface{}{
				map[string]interface{}{"disk_id": diskId},
				map[string]interface{}{"disk_id": diskId + "-2"},
			},
		},
		"duplicate": {
			disks: []interface{}{
				map[string]interface{}{"disk_id": diskId},
				map[string]interface{}{"disk_id": diskId},
			},
			wantErr: true,
		},
		"duplicate with different casing": {
			disks: []interface{}{
				map[string]interface{}{"disk_id": diskId},
				map[string]interface{}{"disk_id": strings.ToUpper(diskId)},
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := siteRecoveryReplicatedVMValidateManagedDiskIDs(test.disks)
			if (err != nil) != test.wantErr {
				t.Fatalf("expected error %t, got %v", test.wantErr, err)
			}
		})
	}
}
