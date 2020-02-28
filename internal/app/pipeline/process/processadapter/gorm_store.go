// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package processadapter

import (
	"context"
	"time"

	"emperror.dev/errors"
	"github.com/jinzhu/gorm"

	"github.com/banzaicloud/pipeline/internal/app/pipeline/process"
)

// TableName constants
const (
	processTableName      = "processes"
	processEventTableName = "process_events"
)

type processModel struct {
	ID           string `gorm:"primary_key"`
	ParentID     string
	OrgID        uint                `gorm:"not null"`
	Name         string              `gorm:"not null"`
	ResourceType string              `gorm:"not null"`
	ResourceID   string              `gorm:"not null"`
	Status       string              `gorm:"not null"`
	StartedAt    time.Time           `gorm:"index:idx_start_time_end_time;default:current_timestamp;not null"`
	FinishedAt   *time.Time          `gorm:"index:idx_start_time_end_time;default:'1970-01-01 00:00:01';not null"`
	Events       []processEventModel `gorm:"foreignkey:ProcessID"`
}

// TableName changes the default table name.
func (processModel) TableName() string {
	return processTableName
}

type processEventModel struct {
	ProcessID string
	Log       string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Timestamp time.Time `gorm:"index:idx_timestamp;default:current_timestamp;not null"`
}

// TableName changes the default table name.
func (processEventModel) TableName() string {
	return processEventTableName
}

// GormStore is a notification store using Gorm for data persistence.
type GormStore struct {
	db *gorm.DB
}

// NewGormStore returns a new GormStore.
func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{
		db: db,
	}
}

// List returns the list of active processes.
func (s *GormStore) List(ctx context.Context, query process.Process) ([]process.Process, error) {
	var processes []processModel

	err := s.db.Find(&processes, query).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to find processes")
	}

	result := []process.Process{}

	for _, pm := range processes {

		p := process.Process{
			ID:           pm.ID,
			ParentID:     pm.ParentID,
			OrgID:        pm.OrgID,
			Name:         pm.Name,
			StartedAt:    pm.StartedAt,
			FinishedAt:   pm.FinishedAt,
			ResourceID:   pm.ResourceID,
			ResourceType: pm.ResourceType,
			Status:       pm.Status,
		}

		var processEvents []processEventModel

		err := s.db.Model(&pm).Related(&processEvents, "Events").Error
		if err != nil {
			return nil, errors.Wrap(err, "failed to find process events")
		}

		for _, em := range processEvents {
			p.Events = append(p.Events, process.ProcessEvent{
				ProcessID: em.ProcessID,
				Name:      em.Name,
				Log:       em.Log,
				Timestamp: em.Timestamp,
			})
		}

		result = append(result, p)
	}

	return result, nil
}

// LogProcess logs a process entry
func (s *GormStore) LogProcess(ctx context.Context, p process.Process) error {

	existing := processModel{ID: p.ID, ParentID: p.ParentID}

	err := s.db.Find(&existing).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			pm := processModel{
				ID:           p.ID,
				ParentID:     p.ParentID,
				OrgID:        p.OrgID,
				Name:         p.Name,
				ResourceType: p.ResourceType,
				ResourceID:   p.ResourceID,
				Status:       p.Status,
				StartedAt:    p.StartedAt,
			}

			err := s.db.Create(&pm).Error
			return errors.Wrap(err, "failed to create process")
		}

		return err
	}

	existing.Status = p.Status
	existing.FinishedAt = p.FinishedAt

	err = s.db.Save(&existing).Error
	if err != nil {
		return errors.Wrap(err, "failed to update process")
	}

	return nil
}

// LogEvent logs a process event
func (s *GormStore) LogEvent(ctx context.Context, p process.ProcessEvent) error {

	pem := processEventModel{
		ProcessID: p.ProcessID,
		Name:      p.Name,
		Log:       p.Log,
		Timestamp: p.Timestamp,
	}

	err := s.db.Create(&pem).Error
	return errors.Wrap(err, "failed to create process event")
}
