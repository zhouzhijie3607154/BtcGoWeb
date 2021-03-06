package moudles

//import "github.com/astaxie/beego/orm"

//用户名结构体
type User struct { //form解析 orm映射的数据库表中列名解析 json 解析
	Id       int    `form:"id" orm:"column(id)" json:"id"`
	Name     string `form:"username" orm:"column(user_name)" json:"username"`
	Password string `form:"password" orm:"column(user_pwd)" json:"password"`
	Status   int    `form:"status" orm:"column(user_status)" json:"status"`
	TimeStamp  int64    `form:"timestamp" orm:"column(user_create_time)" json:"timestamp"`
}


