package plans

import (
  "time"
  "github.com/gh0stl1m/subscription-service/domains/users"
)

type Plan struct {
  ID string `gorm:"primaryKey"`
  PlanName string
  PlanAmount int
  CreatedAt time.Time
  UpdatedAt time.Time
}

type PlanRepository interface {
  FindOne(id string) (*Plan, error)
  Find() ([]*Plan, error)
  SubscribeUserToPlan(user users.User, plan Plan) error
}
