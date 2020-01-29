package model

import (
	"errors"
	//"fmt"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
)

//结构体
type accountReqeustParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

type User struct{
	User_id string	`json:"user_id"`
	User_name	string	`json:"user_name"`
	Password	string	`json:"password"`
	Signture	string	`json:"signture"`
	Image_url	string	`json:"image_url"`
	Background_url	string	`json:"background_url"`
	Fans_num	int	`json:"fans_num"`
	Following_num	int	`json:"following_num"`
}

type Following_fan struct{
	Following_id
}

//确认模拟登陆是否成功
func ConfirmUser(sid string, pwd string) bool {
	params,err := makeAccountPreflightRequest()
	if err != nil {
		log.Println(err)
		return false
	}

	jar,err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Println(err)
		return false
	}
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}
	//fmt.Println(params)
	result := makeAccountRequest( sid, pwd, params, &client)

	return result
}

// 预处理，打开 account.ccnu.edu.cn 获取模拟登陆需要的表单字段
func makeAccountPreflightRequest() (*accountReqeustParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &accountReqeustParams{}

	// 初始化 http client
	client := http.Client{
		//Timeout: TIMEOUT,
	}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		log.Println(err)
		return params, err
	}
	//request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	// 发起请求
	resp, err := client.Do(request)
	if err != nil {

		log.Println(err)
		return params, err
	}

	// 读取 Body
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return params, err
	}

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range resp.Cookies() {
		//fmt.Println(cookie.Value)
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}

	if JSESSIONID == "" {
		log.Println("Can not get JSESSIONID")
		return params, errors.New("Can not get JSESSIONID")
	}

	// 正则匹配 HTML 返回的表单字段
	ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	bodyStr := string(body)

	ltArr := ltReg.FindStringSubmatch(bodyStr)
	if len(ltArr) != 2 {
		log.Println("Can not get form paramater: lt")
		return params, errors.New("Can not get form paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		log.Println("Can not get form paramater: execution")
		return params, errors.New("Can not get form paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		log.Println("Can not get form paramater: _eventId")
		return params, errors.New("Can not get form paramater: _eventId")
	}
	_eventId = _eventIdArr[1]

	//log.Println("Get params successfully", lt, execution, _eventId)

	params.lt = lt
	params.execution = execution
	params._eventId = _eventId
	params.submit = "LOGIN"
	params.JSESSIONID = JSESSIONID

	return params, nil
}

func makeAccountRequest(sid, password string, params *accountReqeustParams, client *http.Client) bool {
	v := url.Values{}
	v.Set("username", sid)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)
	//fmt.Println(strings.NewReader(v.Encode()))
	request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")


	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	// check
	reg := regexp.MustCompile("class=\"success\"")
	matched := reg.MatchString(string(body))
	if !matched {
		//log.Println("Wrong sid or pwd")
		return false
	}

	//log.Println("Login successfully")
	return true
}

//注册用户
func CreateUser(user_id string,user_name string,password string) {
	DB.Self.Model(&User{}).Create(&User{User_id: user_id,User_name: user_name,Password: password})
}

func CheckUserByUser_id(user_id string) bool {
    var l User
    res := DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

func CreateToken(user_id string)string  {
	keyInfo := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"
	dataByte,_:= json.Marshal(user_id)
	var dataStr = string(dataByte)

	//使用Claim保存json
	//这里是个例子，并包含了一个故意签发一个已过期的token
	data := jwt.StandardClaims{Subject:dataStr,ExpiresAt:time.Now().Unix()+100000000}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256,data)
	//生成token字符串
	token,_ := tokenInfo.SignedString([]byte(keyInfo))
	return token
}

func Viewing(user_id string) User {
	var l User
    DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).First(&l)
	
	return l
}

func Token_info(Token string) (string,bool) {
	keyInfo := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"
	//将token字符串转换为token对象（结构体更确切点吧，go很孤单，没有对象。。。）
	tokenInfo , _ := jwt.Parse(Token, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo,nil
	})

	
	//校验错误（基本）
	err := tokenInfo.Claims.Valid()
	if err!=nil{
		println(err.Error())
		return err.Error(),false
	}
	
	finToken := tokenInfo.Claims.(jwt.MapClaims)
	//校验下token是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(),true)
	
	var a string

	if succ {
		return  a,false
	}else{
		return finToken["sub"].(string),true
		//return true
	}

	//fmt.Println("succ",succ)
    //获取token中保存的用户信息
	//fmt.Println(finToken["sub"])
	//return finToken["sub"]
}

func Background_modify(user_id string,background_url string)  {
	DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).Update(User{Background_url: background_url})
}

func Image_modify(user_id string,image_url string)  {
	DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).Update(User{Image_url: image_url})
}

func Signture_modify(user_id string,signture string)  {
	DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).Update(User{Signture: signture})
}