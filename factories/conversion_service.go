package factories

import (
	"time"
	"math/big"
	"strconv"
)

func ConvertDateStrToObj(dateStr *string) (time.Time, error) {
	parsedDate, err := time.Parse(DATE_LAYOUT, *dateStr)
	return parsedDate, err
}

func ConvertTimeStrToObj(timeStr *string) time.Time {
	parsedTime,_ := time.Parse(TIME_LAYOUT, *timeStr)
	return parsedTime
}

func ConvertTimeObjToStr(timeObj *time.Time) string {
	return timeObj.Format(TIME_LAYOUT)
}

func ConvertDateObjToStr(dateObj *time.Time) string {
	return dateObj.Format(DATE_LAYOUT)
}

func ConvertStrNumToRatNum(amountStr *string) *big.Rat {
	rat := new(big.Rat)
	rat.SetString(*amountStr)
	return rat
}

func ConvertBigRatToStrDollarAmt(amount *big.Rat) string {
	floatVal,_ := amount.Float64()
	return strconv.FormatFloat(floatVal, 'f', 2, 64)
}