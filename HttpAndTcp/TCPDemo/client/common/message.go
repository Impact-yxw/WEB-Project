package common

const (
	LoginMesType = "LoginMes"
	LoginResType = "LoginRes"
)

type Message struct {
	Type string `jdon:"type"` //类型
	Data string `json:"data"` //内容
}

//客户端发送消息
type LoginMes struct {
	Uid   string `json:"uid"`
	Pwd   string `json:"pwd"`
	Uname string `json:"uname"`
}

//服务器返回消息
type LoginRes struct {
	Code int `json:"code"` //返回状态码
	//500 未注册
	//200 登陆成功

	Error string `jdon:"error"` //返回错误消息
}
