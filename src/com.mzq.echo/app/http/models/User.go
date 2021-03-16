package models

import (
	"com.mzq.echo/app/core"
	"com.mzq.echo/app/core/config"
	"com.mzq.echo/app/http/exception"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type User struct {
	Id int	`json:"uid"`
	Password string	`json:"-"`
	UserName string `json:"username"`
}

type UserQuery struct {
	Id int	`json:"uid"`
	UserName string `json:"username"`
}

var currentUser *User;


const (
	APP_UID = "uid"
)

func Verify(userName, password string) (*User, error) {
	var u = new(User)

	find := core.Db().Find(u, "user_name=?", userName)
	if find.Error != nil {
		return nil, exception.NewExceptionMzq(exception.ERROR_NO_USER, "用户不存在")
	}
	if u.Password ==  password {
		return u, nil
	}

	return nil, exception.NewExceptionMzq(exception.ERROR_PASSWORD, "密码错误")
}

func (user *User)CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.Id,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(AppKey()))
	if err != nil {
		return "", exception.NewExceptionMzq(exception.ERROR_CREATE_TOKEN, err.Error())
	}

	return tokenString, nil
}

func FindUid(uid int) (*User, error)  {
	var u = new(User)
	find := core.Db().Find(u, "id=?", uid)
	if find.Error != nil {
		return nil, exception.NewExceptionMzq(exception.ERROR_NO_USER, "用户不存在")
	}
	return u, nil
}

func FindUidString(uid string) (*User, error)  {
	id, _ := strconv.Atoi(uid)
	return FindUid(id)
}


func ListUser(userQuery UserQuery) (*[]User, int, error) {
	var u = new([]User)
	db := core.Db()
	if userQuery.UserName != "" {
		db = db.Where("user_name=?", userQuery.UserName)
	}
	if userQuery.Id > 0 {
		db = db.Where("id=?", userQuery.Id)
	}
	var count int

	find := db.Find(u)
	if find.Error != nil {
		return nil, 0, exception.NewExceptionMzq(exception.ERROR_NO_USER, "用户不存在")
	}
	find.Count(&count)
	return u, count, nil
}

func CreateUser(userName, password string) (*User,  error) {
	var u = new(User)
	u.UserName = userName
	u.Password = password
	db := core.Db()

	find := db.Create(u)
	if find.Error != nil {
		return nil, exception.NewExceptionMzq(exception.ERROR_NO_USER, "創建失敗")
	}
	return u,  nil
}

func (user *User)UpdateUser(userName, password string) (*User,  error) {
	var u = new(User)
	u.UserName = userName
	u.Password = password
	db := core.Db()
	find := db.Model(user).Update(u)
	if find.Error != nil {
		return nil, exception.NewExceptionMzq(exception.ERROR_NO_USER, "用户不存在")
	}
	return user,  nil
}



func CheckLogin(context echo.Context) (error) {
	token := context.Get(APP_UID)
	claims := token.(*jwt.Token)
	tokenClaims :=claims.Claims.(jwt.MapClaims)
	uid := tokenClaims["uid"].(float64)
	uInfo, err := FindUid(int(uid))
	if err != nil {
		return exception.NewExceptionMzq(exception.ERROR_CREATE_TOKEN, "请登录")
	}
	currentUser = uInfo
	return  nil
}


func LoginUser() *User {
	return  currentUser
}

func AppKey() string {
	return config.App().AppKey
}