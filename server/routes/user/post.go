package user

import (
	"assignment/entities/models"
	userRepo "assignment/entities/repos/user"
	"assignment/utils"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"time"
)

type PostRequest struct {
	Name string
}

type PostResponse struct {
	User *models.User
}

func HandlePost() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		resp := new(PostResponse)
		var err error
		logger := log.WithTime(time.Now())
		defer func() {
			r := utils.NewResponse(resp, err, logger)
			r.Print(ctx)
		}()

		body := ctx.PostBody()
		req := new(PostRequest)
		if e := json.Unmarshal(body, req); e != nil {
			err = errors.New("cannot parse request")
			logger = logger.WithField("request", body)
			return
		}
		logger = logger.WithField("request", req)

		userRepository := userRepo.GetRepo()
		user, err := userRepository.Create(req.Name)
		if err != nil {
			return
		}
		resp.User = user
	}
}
