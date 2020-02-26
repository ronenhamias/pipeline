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
	processTableName = "processes"
)

type processModel struct {
	ID           string `gorm:"primary_key"`
	ParentID     string
	OrgID        uint       `gorm:"not null"`
	Name         string     `gorm:"not null"`
	ResourceType string     `gorm:"not null"`
	ResourceID   string     `gorm:"not null"`
	Status       string     `gorm:"not null"`
	StartedAt    time.Time  `gorm:"index:idx_start_time_end_time;default:current_timestamp;not null"`
	FinishedAt   *time.Time `gorm:"index:idx_start_time_end_time;default:'1970-01-01 00:00:01';not null"`
}

// TableName changes the default table name.
func (processModel) TableName() string {
	return processTableName
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

	for _, n := range processes {
		result = append(result, process.Process{
			ID:           n.ID,
			ParentID:     n.ParentID,
			OrgID:        n.OrgID,
			Name:         n.Name,
			StartedAt:    n.StartedAt,
			FinishedAt:   n.FinishedAt,
			ResourceID:   n.ResourceID,
			ResourceType: n.ResourceType,
			Status:       n.Status,
		})
	}

	return result, nil
}

// Log logs a process entry
func (s *GormStore) Log(ctx context.Context, p process.Process) error {

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
