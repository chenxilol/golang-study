package _case

import "time"

func SingleFieldValidate() {
	v := validate
	var err error
	var b bool
	err = v.Var(b, "boolean")
	outRes("boolean", &err)
	// 切片格式
	var a = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	err = v.Var(a, "max=15,min=2")
	outRes("slice", &err)

	// map集合的验证

	mp1 := map[string]string{
		"A": "123456",
		"B": "9876665",
		"1": "sdadsfdf",
	}
	err = v.Var(mp1, "gte=3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number")
	outRes("mp1", &err)
	// 时间格式
	var timeStr = time.Now().Format("2016-01-02 15:04:05")
	err = v.Var(timeStr, "datetime=2006-01-02 15:02:05")
	outRes("datetime", &err)

	// 字段是否相等
	s, s1 := "abc", "abc"
	err = v.VarWithValue(s, s1, "eqfield")
	outRes("eqfield", &err)

	// 验证字段是否小于第一个, 即检查第一个字段是否小于第二个字段

	i1, i2 := 4, 10
	err = v.VarWithValue(i1, i2, "ltfield")
	outRes("ltfield", &err)
}
