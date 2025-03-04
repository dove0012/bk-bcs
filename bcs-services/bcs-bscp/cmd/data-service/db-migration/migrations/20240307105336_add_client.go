/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package migrations

import (
	"time"

	"gorm.io/gorm"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/cmd/data-service/db-migration/migrator"
)

func init() {
	// add current migration to migrator
	migrator.GetMigrator().AddMigration(&migrator.Migration{
		Version: "20240307105336",
		Name:    "20240307105336_add_client",
		Mode:    migrator.GormMode,
		Up:      mig20240307105336Up,
		Down:    mig20240307105336Down,
	})
}

// mig20240307105336Up for up migration
func mig20240307105336Up(tx *gorm.DB) error {
	// Clients : client
	type Clients struct {
		ID uint `gorm:"column:id;type:bigint(1) unsigned;primary_key;autoIncrement:false"`

		BizID uint   `gorm:"type:bigint(1) unsigned not null;uniqueIndex:idx_bizID_appID_uid,priority:1"`
		APPID uint   `gorm:"type:bigint(1) unsigned not null;uniqueIndex:idx_bizID_appID_uid,priority:2"`
		UID   string `gorm:"column:uid;type:varchar(64);NOT NULL;uniqueIndex:idx_bizID_appID_uid,priority:3"`

		ClientVersion             string    `gorm:"column:client_version;type:varchar(150);default:'';NOT NULL"`
		ClientType                string    `gorm:"column:client_type;type:varchar(20);default:'';NOT NULL"`
		Ip                        string    `gorm:"column:ip;type:varchar(16);default:'';NOT NULL"`
		Labels                    string    `gorm:"column:labels;type:json;default:NULL"`
		Annotations               string    `gorm:"column:annotations;type:json;default:NULL"`
		FirstConnectTime          time.Time `gorm:"column:first_connect_time;type:datetime(6);default:NULL"`
		LastHeartbeatTime         time.Time `gorm:"column:last_heartbeat_time;type:datetime(6);default:NULL;index:idx_lastHeartbeatTime_onlineStatus,priority:1"`  // nolint
		OnlineStatus              string    `gorm:"column:online_status;type:varchar(20);default:'';NOT NULL;index:idx_lastHeartbeatTime_onlineStatus,priority:2"` // nolint
		CpuUsage                  float64   `gorm:"column:cpu_usage;type:double unsigned;default:0;NOT NULL"`
		CpuMaxUsage               float64   `gorm:"column:cpu_max_usage;type:double unsigned;default:0;NOT NULL"`
		MemoryUsage               uint64    `gorm:"column:memory_usage;type:bigint(1) unsigned;default:0;NOT NULL"`
		MemoryMaxUsage            uint64    `gorm:"column:memory_max_usage;type:bigint(1) unsigned;default:0;NOT NULL"`
		CurrentReleaseID          uint      `gorm:"column:current_release_id;type:bigint(1) unsigned;default:0;NOT NULL"`
		TargetReleaseID           uint      `gorm:"column:target_release_id;type:bigint(1) unsigned;default:0;NOT NULL"`
		ReleaseChangeStatus       string    `gorm:"column:release_change_status;type:varchar(20);default:'';NOT NULL"`
		ReleaseChangeFailedReason string    `gorm:"column:release_change_failed_reason;type:varchar(20);default:'';NOT NULL"`
		FailedDetailReason        string    `gorm:"column:failed_detail_reason;type:varchar(255);default:'';NOT NULL"`
	}

	// ClientEvents : ClientEvents
	type ClientEvents struct {
		ID       uint   `gorm:"column:id;type:bigint(1) unsigned;primary_key"`
		ClientID uint   `gorm:"column:client_id;type:bigint(1) unsigned;NOT NULL;index:idx_clientID,priority:1"`
		CursorID string `gorm:"column:cursor_id;type:varchar(128);NOT NULL;uniqueIndex:idx_bizID_appID_uid_cursorID,priority:1"` // nolint

		BizID uint `gorm:"type:bigint(1) unsigned not null;uniqueIndex:idx_bizID_appID_uid_cursorID,priority:2"`
		AppID uint `gorm:"type:bigint(1) unsigned not null;uniqueIndex:idx_bizID_appID_uid_cursorID,priority:3"`

		UID                       string    `gorm:"column:uid;type:varchar(64);NOT NULL;uniqueIndex:idx_bizID_appID_uid_cursorID,priority:4"` // nolint
		ClientMode                string    `gorm:"column:client_mode;type:varchar(20);default:'';NOT NULL;"`
		OriginalReleaseID         uint      `gorm:"column:original_release_id;type:bigint(1) unsigned;default:0;NOT NULL"`
		TargetReleaseID           uint      `gorm:"column:target_release_id;type:bigint(1) unsigned;default:0;NOT NULL"`
		StartTime                 time.Time `gorm:"column:start_time;type:datetime(6);default:NULL"`
		EndTime                   time.Time `gorm:"column:end_time;type:datetime(6);default:NULL"`
		TotalSeconds              float64   `gorm:"column:total_seconds;type:double unsigned;default:0;NOT NULL"`
		TotalFileSize             float64   `gorm:"column:total_file_size;type:double unsigned;default:0;NOT NULL"`
		DownloadFileSize          float64   `gorm:"column:download_file_size;type:double unsigned;default:0;NOT NULL"`
		TotalFileNum              uint      `gorm:"column:total_file_num;type:int(3) unsigned;default:0;NOT NULL"`
		DownloadFileNum           uint      `gorm:"column:download_file_num;type:int(3) unsigned;default:0;NOT NULL"`
		ReleaseChangeStatus       string    `gorm:"column:release_change_status;type:varchar(20);default:;NOT NULL"`
		ReleaseChangeFailedReason string    `gorm:"column:release_change_failed_reason;type:varchar(20);default:'';NOT NULL"`
		FailedDetailReason        string    `gorm:"column:failed_detail_reason;type:varchar(255);default:'';NOT NULL"`
	}

	// IDGenerators : ID生成器
	type IDGenerators struct {
		ID        uint      `gorm:"type:bigint(1) unsigned not null;primaryKey"`
		Resource  string    `gorm:"type:varchar(50) not null;uniqueIndex:idx_resource"`
		MaxID     uint      `gorm:"type:bigint(1) unsigned not null"`
		UpdatedAt time.Time `gorm:"type:datetime(6) not null"`
	}

	if err := tx.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").
		AutoMigrate(&Clients{}, &ClientEvents{}); err != nil {
		return err
	}

	now := time.Now()

	if result := tx.Create([]IDGenerators{
		{Resource: "clients", MaxID: 0, UpdatedAt: now},
		{Resource: "client_events", MaxID: 0, UpdatedAt: now},
	}); result.Error != nil {
		return result.Error
	}

	return nil

}

// mig20240307105336Down for down migration
func mig20240307105336Down(tx *gorm.DB) error {

	// IDGenerators : ID生成器
	type IDGenerators struct {
		ID        uint      `gorm:"type:bigint(1) unsigned not null;primaryKey"`
		Resource  string    `gorm:"type:varchar(50) not null;uniqueIndex:idx_resource"`
		MaxID     uint      `gorm:"type:bigint(1) unsigned not null"`
		UpdatedAt time.Time `gorm:"type:datetime(6) not null"`
	}

	if err := tx.Migrator().
		DropTable("clients", "client_events"); err != nil {
		return err
	}

	var resources = []string{
		"clients",
		"client_events",
	}
	if result := tx.Where("resource IN ?", resources).Delete(&IDGenerators{}); result.Error != nil {
		return result.Error
	}

	return nil
}
