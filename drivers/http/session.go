package http

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gh0stl1m/subscription-service/drivers/redis"
)

func CreateSession() *scs.SessionManager {

  redisConn := redis.NewConnection()
  session := scs.New()
  session.Store = redisstore.New(redisConn)
  session.Lifetime = 24 * time.Hour
  session.Cookie.Persist = true
  session.Cookie.SameSite = http.SameSiteLaxMode
  session.Cookie.Secure = true

  return session
}

