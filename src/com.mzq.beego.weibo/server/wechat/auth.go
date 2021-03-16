package wechat

import (
	"crypto/tls"
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	url2 "net/url"
)

type BaseError struct {
	errCode int
	errMsg  string
}

func (err BaseError) Error() string {
	return "errCode：" + string(err.errCode) + " errMsg:" + err.errMsg
}

type Auth struct {
	BaseError
	host   string
	appId  string
	secret string
}

func NewAuth(appId string, secret string) *Auth {
	baseError := new(BaseError)
	return &Auth{
		BaseError: *baseError,
		host:      "https://api.weixin.qq.com",
		secret:    secret,
		appId:     appId,
	}
}

func (auth *Auth) ErrCode() int {
	return auth.BaseError.errCode
}

func (auth *Auth) ErrMsg() string {
	return auth.BaseError.errMsg
}

func (auth *Auth) GetAccessToken(jsCode string) (string, int, error) {
	//	https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	var url string = auth.host + "/cgi-bin/token?"
	values := url2.Values{}
	values.Add("appid", auth.appId)
	values.Add("secret", auth.secret)
	values.Add("grant_type", "client_credential")
	url = url + values.Encode()
	req := httplib.Get(url).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	bytes, err := req.Bytes()
	if err != nil {
		return "", 0, err
	}
	var i interface{}
	err = json.Unmarshal(bytes, &i)
	if err != nil {
		return "", 0, err
	}
	m := i.(map[string]interface{})
	auth.errCode = m["errcode"].(int)
	auth.errMsg = m["errmsg"].(string)
	if auth.errCode != 0 {
		return "", 0, auth
	}
	accessToken := m["access_token"].(string)
	expiresIn := m["expires_in"].(int)
	return accessToken, expiresIn, nil
}

func (auth *Auth) Code2Session(jsCode string) (string, string, string, error) {
	//	https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	var url string = auth.host + "/sns/jscode2session?"
	values := url2.Values{}
	values.Add("appid", auth.appId)
	values.Add("secret", auth.secret)
	values.Add("js_code", jsCode)
	values.Add("grant_type", "authorization_code")
	url = url + values.Encode()
	req := httplib.Get(url).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	bytes, err := req.Bytes()
	if err != nil {
		return "", "", "", err
	}
	var i interface{}
	err = json.Unmarshal(bytes, &i)
	if err != nil {
		return "", "", "", err
	}
	m := i.(map[string]interface{})
	auth.errCode = m["errcode"].(int)
	auth.errMsg = m["errmsg"].(string)
	if auth.errCode != 0 {
		return "", "", "", auth
	}
	openid := m["openid"].(string)
	sessionKey := m["session_key"].(string)
	unionId := m["unionid"].(string)
	return openid, sessionKey, unionId, nil
}

// GetPaidUnionId(accessToken, openid , transactionId )
// GetPaidUnionId(accessToken, openid , mchId, outTradeNo)
func (auth *Auth) GetPaidUnionId(accessToken string, openid string, transactionId string, arg ...string) (string, error) {
	//	 https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID
	var url string = auth.host + "/wxa/getpaidunionid?"
	values := url2.Values{}
	values.Add("access_token", accessToken)
	values.Add("openid", openid)
	if len(arg) > 0 {
		values.Add("mch_id", transactionId)
		values.Add("out_trade_no", arg[0])
	} else {
		values.Add("transaction_id", transactionId)
	}
	url = url + values.Encode()
	req := httplib.Get(url).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	bytes, err := req.Bytes()
	if err != nil {
		return "", err
	}
	var i interface{}
	err = json.Unmarshal(bytes, &i)
	if err != nil {
		return "", err
	}
	m := i.(map[string]interface{})
	auth.errCode = m["errcode"].(int)
	auth.errMsg = m["errmsg"].(string)
	if auth.errCode != 0 {
		return "", auth
	}
	unionId := m["unionid"].(string)
	return unionId, nil
}
