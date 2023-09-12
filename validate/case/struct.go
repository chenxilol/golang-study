package _case

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type User struct {
	Name  string `json:"name" v:"max=15,min=2"`
	Age   int    `json:"age" v:"gte=10,lte=30"`
	Phone string ` json:"phone"`
	Email string `json:"email"`
}

func StructValidate() {
	v := validate
	v.SetTagName("v")
	r := gin.Default()
	r.POST("/c", func(context *gin.Context) {

		a, _ := strconv.Atoi(context.PostForm("age"))
		u := &User{
			Name:  context.PostForm("name"),
			Age:   a,
			Phone: "121",
			Email: "212",
		}
		err := v.Struct(u)
		outRes("s", &err)
		context.JSON(http.StatusOK, u)
	})
	r.Run(":9999")
}

type User1 struct {
	Name            string         `json:"name" v:"required,alphaunicode"`
	Age             uint8          `json:"age" v:"gte=10,lte=30"`
	Phone           string         ` json:"phone" v:"required,e164"` // e164电话格式
	Email           string         `json:"email" V:"required,email"` // email是电话格式
	FavouriteColor1 string         `v:"required"`
	FavouriteColor2 string         `v:"required"`
	Address         *Address       `json:"address" v:"required"`
	ContactUser     []*ContactUser `json:"contactUser" v:"required,gte=1,dive"`                       // dive 深入一层验证，如果要多个验证，那么就要多个dive
	Hobby           []string       `json:"hobby" v:"required,gte=2,dive,required,gte=2,alphaunicode"` // 如果填入空字符也会通过校验，所以用dive进行深度验证到每一个字符
}
type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}
type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=20,lte=130"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"` // 当email、address都没填iphone是必填的，但是iphone不是必填的，他可以是空字符串，会导致校验不通过，因为有e164，所以用omitempty如果是空的就跳过以后得校验规则
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `v:"required_without_all=Email Phone"`
}

func StructValidate1() {
	v := validate
	//	v.SetTagName("v")
	address := &Address{
		Province: "河南",
		City:     "郑州",
	}
	contactUser1 := &ContactUser{
		Name:    "chenxi",
		Age:     19,
		Phone:   "+8617797776520",
		Email:   "",
		Address: nil,
	}
	contactUser2 := &ContactUser{
		Name:    "chenxi",
		Age:     19,
		Phone:   "+8617797776520",
		Email:   "",
		Address: nil,
	}
	user := &User1{
		Name:            "chen",
		Age:             19,
		Phone:           "+8617719996251",
		Email:           "chenxilola@gmail.com",
		FavouriteColor1: "#ffff",
		FavouriteColor2: "rgb(255,255,255)",
		Address:         address,
		ContactUser:     []*ContactUser{contactUser1, contactUser2},
		Hobby:           []string{"乒乓球", "跑步"},
	}
	err := v.Struct(user)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}
}
