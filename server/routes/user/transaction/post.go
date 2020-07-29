package transaction

import (
	"assignment/entities/models"
	commonRepo "assignment/entities/repos/common"
	"assignment/utils"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type PostResponse struct {
	User *models.User
}

type PostRequest struct {
	State         string
	Amount        utils.Decimal2
	TransactionId string
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

		userId := ctx.UserValue("userId")
		logger = logger.WithField("userId", userId)
		uId, e := strconv.Atoi(userId.(string))
		if e != nil {
			err = errors.New("cannot convert userId to int")
			return
		}

		body := ctx.PostBody()
		req := new(PostRequest)
		if e := json.Unmarshal(body, req); e != nil {
			err = errors.New("cannot parse request")
			logger = logger.WithField("request", body)
			return
		}
		logger = logger.WithField("request", req)
		if req.TransactionId == "" {
			err = errors.New("no transaction id")
			return
		}

		transactionType := ctx.Request.Header.Peek("Source-Type")

		transaction := models.Transaction{
			State:         models.ConvertTransactionStateStoI(req.State),
			Type:          models.ConvertTransactionTypeStoI(string(transactionType)),
			UserId:        uId,
			Amount:        req.Amount,
			TransactionId: req.TransactionId,
		}
		commonRepository := commonRepo.GetRepo()
		user, err := commonRepository.CreateUserTransaction(&transaction)
		if err != nil {
			return
		}
		resp.User = user
	}
}
