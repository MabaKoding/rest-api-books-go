package controllers

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	mathRand "math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/go-playground/validator/v10"
	"github.com/leekchan/accounting"
)

var (
// ZapLogger *zap.Logger
)

func init() {
	// ZapLogger = logger.ZapLogger
}

type (
	BaseController struct {
		beego.Controller
		Settings map[string]interface{}
	}
)

func (c *BaseController) PublicContent(view string) {
	c.Layout = "basic-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.TplName = view + ".tpl"
}

func (c *BaseController) ConvertUnixtimeToDate(intDate interface{}) string {
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")
	var output string
	switch intDate.(type) {
	case float64:
		if intDate.(float64) > 0 {
			tm := time.Unix(int64(intDate.(float64)), 0).In(timeLocation)
			output = tm.Format("02-01-2006")
		}
	case int64:
		if intDate.(int64) > 0 {
			i := intDate.(int64)
			tm := time.Unix(int64(i), 0).In(timeLocation)
			output = tm.Format("02-01-2006")
		}
	case string:
		if intDate.(string) != "" {
			i, err := strconv.ParseInt(intDate.(string), 10, 64)
			if err != nil {
				// ZapLogger.Error(err.Error())
			} else {
				tm := time.Unix(int64(i), 0).In(timeLocation)
				output = tm.Format("02-01-2006")
			}
		}
	}

	return output
}

func (c *BaseController) ConvertUnixtimeToDateTime(intDate interface{}) string {
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")
	var output string
	switch intDate.(type) {
	case float64:
		if intDate.(float64) > 0 {
			tm := time.Unix(int64(intDate.(float64)), 0).In(timeLocation)
			output = tm.Format("02-01-2006 15:04:05 MST")
		}
	case int64:
		if intDate.(int64) > 0 {
			i := intDate.(int64)
			tm := time.Unix(int64(i), 0).In(timeLocation)
			output = tm.Format("02-01-2006 15:04:05 MST")
		}
	case string:
		if intDate.(string) != "" {
			i, err := strconv.ParseInt(intDate.(string), 10, 64)
			if err != nil {
				// ZapLogger.Error(err.Error())
			} else {
				tm := time.Unix(int64(i), 0).In(timeLocation)
				output = tm.Format("02-01-2006 15:04:05 MST")
			}
		}
	}

	return output
}

func (c *BaseController) TimeUnix() int64 {
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	createdOnUnix := time.Now().In(jakarta).Unix()

	return createdOnUnix
}

func (c *BaseController) GenerateOTP(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

func (c *BaseController) TrimPhoneNumber(phoneNumber string) string {
	phoneNumber = strings.TrimSpace(phoneNumber)

	if strings.Contains(phoneNumber, "-") {
		phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	}
	if strings.Contains(phoneNumber, " ") {
		phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	}
	phoneNumber = strings.ReplaceAll(phoneNumber, "'", "")
	if strings.HasPrefix(phoneNumber, "+628") {
		phoneNumber = strings.TrimPrefix(phoneNumber, "+628")
	} else if strings.HasPrefix(phoneNumber, "08628") {
		phoneNumber = strings.TrimPrefix(phoneNumber, "08628")
	} else if strings.HasPrefix(phoneNumber, "08") {
		phoneNumber = strings.TrimPrefix(phoneNumber, "08")
		// } else if strings.HasPrefix(phoneNumber, "8") {
		// 	phoneNumber = strings.TrimPrefix(phoneNumber, "8")
	}

	return phoneNumber
}

func (c *BaseController) FormatMoney(sum interface{}) string {
	ac := accounting.Accounting{Symbol: "", Precision: 2, Thousand: ".", Decimal: ","}

	return ac.FormatMoney(sum)
}

// https://siongui.github.io/2017/03/28/go-create-directory-if-not-exist/
func (c *BaseController) CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			// ZapLogger.Error(err.Error())
			return
		}
	}
}

func (c *BaseController) EncodeFilename(filename string) string {
	targetFileExt := filepath.Ext(filename)
	encodedName := c.RandomString(16)

	return encodedName + targetFileExt
}

func (c *BaseController) RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[mathRand.Intn(len(letter))]
	}

	return string(b)
}

/*
*

	Convert jwt.MapClaims to interface
	(usermerchant_session / publicuser_session)

*
*/
func (c *BaseController) DescriptToken(param string) (interface{}, error) {
	var output interface{}

	if param == "" {
		return nil, errors.New("Missing parameter token")
	}

	claims := c.Ctx.Input.GetData(param)
	tmp, _ := json.Marshal(claims)
	_ = json.Unmarshal(tmp, &output)

	return output, nil
}

func (c *BaseController) GetMyIP() string {
	myIP := c.Ctx.Input.IP()

	return myIP
}

func (c *BaseController) NewValidator() *validator.Validate {
	return validator.New()
}
