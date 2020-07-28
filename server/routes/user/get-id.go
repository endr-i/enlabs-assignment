package user

import (
	"assignment/entities/models"
	userRepo "assignment/entities/repos/user"
	"assignment/utils"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type GetResponse struct {
	User *models.User
}

func HandleGet() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		resp := new(PostResponse)
		var err error
		logger := log.WithTime(time.Now())
		defer func() {
			r := utils.NewResponse(resp, err, logger)
			r.Print(ctx)
		}()

		userId := ctx.UserValue("userId")
		logger = logger.WithField("userId", userId)
		uId, e := strconv.Atoi(userId.(string))
		if e != nil {
			err = errors.New("cannot convert userId to int")
			return
		}
		userRepository := userRepo.GetRepo()
		user, e := userRepository.User(uId)
		if e != nil {
			err = errors.New("user is absent")
			logger.WithError(e).Error()
			return
		}
		resp.User = user
	}
}
