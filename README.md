geetestsdk
==========

[Geetest](http://www.geetest.com) SDK for Golang

使用
====

`go get github.com/Bluek404/geetestsdk`

然后在项目中`import`

```go
import "github.com/Bluek404/geetestsdk"
```

使用key创建一个`GeetestSDK`实例

```go
sdk := geetestsdk.New("3c3f7cfcdcf21e216bf4ed7930c191e2")
```

请把上面的`3c3f7cfcdcf21e216bf4ed7930c191e2`替换成你的key

然后进行验证

```go
ok, err := sdk.Validate(challenge, validate, seccode)
if err != nil {
	log.Println(err)
}
```

如果`ok`为`true`则验证成功

`challenge, validate, seccode`是从表单获取的数据

```go
challenge := r.FormValue("geetest_challenge")
validate := r.FormValue("geetest_validate")
seccode := r.FormValue("geetest_seccode")
```

表单获取数据的方法根据所使用的web框架而定

更多请查看*example*文件夹内示例
