package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"     //加密
)
var choice int   //操作
var tokenMap = make(map[string]string)
var curToken string
var	userInfo = make(map[string][]byte)   //map存用户信息；key存用户名；value存加密密码

func main(){
	//操作事件
    fmt.Println("--------通行证--------")
	fmt.Println("1，注册")
    fmt.Println("2，登录")
	fmt.Println("3，查看登录")
    fmt.Println("4，退出登录")
    fmt.Println("5，删除账号")
    fmt.Println("0, 退出")

	for{
	//读取
        fmt.Print("请选择操作：")
        fmt.Scan(&choice)
   
    	//不同操作情况
	    switch choice{
	    	case 1:
	    	    register()
            //注册
	    	case 2:
	    	    login()
	        //登录 成功后生成token
			case 3:
                viewLogin()	
			//用token查看登录
			case 4:
                logout()
			//删除token用户数据保留
			case 5:
                deleteUser()
			//删掉用户数据

	    	case 0:
		        fmt.Println("再见")
	        //body
	    	    return
	    	default:
		        fmt.Println("不符合要求，请输入数字0/1/2")
	    }   
	}
}
//注册
func register() {
	var name string  //用户名
	var password string  //密码
	fmt.Print("请输入用户名：")
	fmt.Scan(&name)

	_, exist :=userInfo[name]    //看是否存在

	if exist{
		fmt.Println("该账号已存在")
		return
	}

	fmt.Print("请输入密码：")
	fmt.Scan(&password)

	hashpwd, err :=bcrypt.GenerateFromPassword([]byte(password),10)
	if err != nil{
		fmt.Println("密码加密失败")   //加密错误
		return
	}

	userInfo[name]=hashpwd
	fmt.Println("注册成功")
}
//登录
func login() {
	var name string  //用户名
	var password string  //密码
	fmt.Print("请输入用户名：")
	fmt.Scan(&name)

	hashPwd,exist:=userInfo[name]   //看是否存在

	if !exist{
		fmt.Println("账号不存在")
		return
	}
	
	//if exist{
		fmt.Print("请输入密码：")
		fmt.Scan(&password)
		err:=bcrypt.CompareHashAndPassword(hashPwd,[]byte(password)) //验证密码是否正确
		
  
    if err!=nil{
		fmt.Println("密码错误")
	}else{
		fmt.Println("登录成功")

		token:="tok_"+name
		tokenMap[token]=name
		curToken= token
		fmt.Println("你的token：",token)
	}
	//}
}
//查看登录
func viewLogin(){
	var inputToken string
	fmt.Print("请输入token:")
	fmt.Scan(&inputToken)

	userName,yes:=tokenMap[inputToken]
	if yes{
		fmt.Println("该用户为：",userName)
	}else{
		fmt.Println("token无效")
	
	}
}
//退出登录
func logout(){
	if curToken==""{
		fmt.Println("您还未登录，不用退出")
		return
	}else{
		delete (tokenMap,curToken)
		curToken=""
		fmt.Println("已退出登录，下次需要重新输密码")
	}

}
//删除账号
func deleteUser(){
	var name2 string
	fmt.Print("请输入要删除的账号：")
	fmt.Scan(&name2)

	_,ok:=userInfo[name2]
	if ok{
		delete(userInfo,name2)

		for Tok,delname:=range tokenMap{
			if delname==name2{
				delete(tokenMap,Tok)
			}
		}
		//如果要删除现在的账号

		if tokenMap[curToken]==name2{
			curToken=""
		}

	    fmt.Println("删除成功")
		return
	}else{
		fmt.Println("该账号不存在，请先注册")
		return
	}

}