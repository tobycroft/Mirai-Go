package Calc

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}

func Mt_rand(min, max int64) int64 {
	rand2.Seed(Seed())
	if min == max {
		return min
	} else {
		r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		return r.Int63n(max-min+1) + min
	}
}

func Seed() int64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(999999999))
	return num.Int64() + time.Now().UnixNano()
}

func Rand(min, max int) int {
	rand2.Seed(Seed())
	if min == max {
		return min
	} else {
		var randNum int
		if max-min < 0 {
			randNum = rand2.Intn(min-max) + min
		} else {
			randNum = rand2.Intn(max-min) + min
		}
		return randNum
	}
}

func Any2Int64(any interface{}) int64 {
	ret, err := String2Int64(Any2String(any))
	if err != nil {
		return -99999998
	}
	return ret
}

func Any2Int64_2(any interface{}) (int64, error) {
	return String2Int64(Any2String(any))
}

func Any2Float64(any interface{}) float64 {
	ret, err := String2Float64(Any2String(any))
	if err != nil {
		return 0
	}
	return ret
}
func Any2Float64_2(any interface{}) (float64, error) {
	return String2Float64(Any2String(any))
}

func Any2Int(any interface{}) int {
	ret, err := String2Int(Any2String(any))
	if err != nil {
		return -99999998
	}
	return ret
}

func Hex2Dec(val string) int64 {
	val = strings.TrimLeft(val, "0x")
	if val == "" {
		return 0
	}
	n := new(big.Int)
	n, _ = n.SetString(val, 16)

	return n.Int64()
}

func Dec2Hex(val int64) string {
	n := strconv.FormatInt(val, 16)
	return n
}

func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

func Transfer2Eth(value float64, decimal int) float64 {
	return value / math.Pow10(Any2Int(decimal))
}

func DateThisWeek() string {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Format("2006-01-02")
	return weekStart
}

func DateThisMonth() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation).Format("2006-01-02")
	//lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	//fmt.Println(firstOfMonth)
	//fmt.Println(lastOfMonth)
	return firstOfMonth
}

func DateThisMonthEnd() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1).Format("2006-01-02")

	//fmt.Println(firstOfMonth)
	//fmt.Println(lastOfMonth)
	return lastOfMonth
}

func DateThisday() string {
	timeStr := time.Now().Format("2006-01-02")

	return timeStr
}
