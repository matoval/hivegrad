package main

import (
	config "github.com/matoval/hivegrad/pkg/Config"
	conn "github.com/matoval/hivegrad/pkg/Conn"
	grad "github.com/matoval/hivegrad/pkg/Grad"
)

func main() {
  config.NewConfig().SetConfigPath("./").SetConfigName("config").SetConfigType("yml").LoadConfig()
  conn.NewConn()
  a := grad.New(-4.0, nil, "")
  b := grad.New(2.0, nil, "")
  c := a.Add(b)
  d := a.Mul(b).Add(b.Pow(3))
  c = c.Add(c.Add(grad.New(1.0, nil, "")))
  c = c.Add(grad.New(1.0, nil, "").Add(c))
  d = d.Add(d.Mul(2.0).Add(b.Add(a).Relu()))

  d.Backward()

}
