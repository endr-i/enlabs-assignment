package transactions

import (
	transactionModel "assignment/models/transaction"
	userModel "assignment/models/user"
	"assignment/utils"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type TransactionResponse struct {
	User *userModel.User
}

type TransactionRequest struct {
	State         string
	Amount        utils.Decimal2
	TransactionId string
}

func Handle() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		resp := new(TransactionResponse)
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

		body := ctx.PostBody()
		req := new(TransactionRequest)
		if e := json.Unmarshal(body, req); e != nil {
			err = errors.New("cannot parse request")
			logger = logger.WithField("request", body)
			return
		}
		logger = logger.WithField("request", req)
		userRepo := userModel.GetRepo()
		user, e := userRepo.User(uId)
		if e != nil {
			err = errors.New("user is absent")
			logger.WithError(e).Error()
			return
		}

		transactionRepo := transactionModel.GetRepo()

		resp.User = user
	}
}
