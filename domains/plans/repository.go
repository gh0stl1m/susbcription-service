package plans

import (
  "gorm.io/gorm"
)

type PlanRepository struct {
  db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) IPlanRepository {

  return &PlanRepository{ db }
}

func (pr *PlanRepository) FindOneBy(condition Plan) (*Plan, error) {

  plan := Plan{}

  result := pr.db.Model(condition).First(&plan)

  if result.Error != nil {

    return nil, result.Error
  }

  return &plan, nil
}

func (pr *PlanRepository) Find() ([]*Plan, error) {

  plans := []*Plan{}

  result := pr.db.Find(&plans)

  if result.Error != nil {

    return nil, result.Error
  }

  return plans, nil
}

func (pr *PlanRepository) SubscribeUserToPlan(userId, planId string) error {


  result := pr.db.Exec("INSERT INTO user_plans(user_id, plan_id) VALUES ($1, $2)", userId, planId)

  if result.Error != nil {

    return result.Error
  }

  return nil
}
