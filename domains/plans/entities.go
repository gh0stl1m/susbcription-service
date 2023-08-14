package plans

import (
  "time"
)

type Plan struct {
  ID string `gorm:"primaryKey"`
  PlanName string
  PlanAmount int
  CreatedAt time.Time
  UpdatedAt time.Time
}

type PlanRepository interface {
  FindOneBy(condition Plan) (*Plan, error)
  Find() ([]*Plan, error)
  SubscribeUserToPlan(userId, planId string) error
}
