package common

import (
	"assignment/entities/models"
	"assignment/utils"
	"strconv"
)

const (
	cancelOddQuery = `with
    trs as (
        select *
        from (
                 select
                     t.id,
                     row_number() over (order by t.id) as number,
                     t.amount,
                     t.status,
                     t.state,
                     t.date_create,
                     t.user_id
                 from transactions t
                 where t.status = :successStatus
                 order by date_create desc
                 limit :limit
             ) as raw
        where
            raw.number % 2 = 1
    ),
    usrs as (
        select
            trs.user_id,
            sum(
                    case
                        when trs.state = :loseState
                            then trs.amount
                        else (
                            case
                                when trs.state = :winState
                                    then -1 * trs.amount
                                end
                            )
                        end
                ) as diff
        from trs
        group by trs.user_id
    ),
    updUsrs as (
        update users
        set balance = balance + (
            case when balance + usrs.diff < 0
            then -1 * balance
            else usrs.diff
            end
            )
        from usrs
        where usrs.user_id = users.id
    )
update transactions
set
    status = :cancelStatus
from trs
where trs.id=transactions.id`
)

func getCancelOddQuery(n int) string {
	return utils.StringReplace(cancelOddQuery, map[string]string{
		":successStatus": strconv.Itoa(models.TransactionStatusSuccess),
		":cancelStatus":  strconv.Itoa(models.TransactionStatusCancelled),
		":winState":      strconv.Itoa(models.TransactionStateWin),
		":loseState":     strconv.Itoa(models.TransactionStateLose),
		":limit":         strconv.Itoa(n),
	})
}
