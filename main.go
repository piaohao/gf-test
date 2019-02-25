package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

func main() {
	v := g.View()
	c := g.Config()
	s := g.Server()

	v.AddPath("template")
	c.AddPath("config")
	c.SetFileName("config.yml")

	glog.SetPath("/data/logs/gf-test")

	s.BindHandler("/test", func(r *ghttp.Request) {
		r.Response.WriteTpl("user.html", g.Map{"name": "piaohao", "age": 18, "score": 100})
	})
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(8888)
	s.Run()
}
