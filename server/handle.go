package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
	"github.com/thinkeridea/go-extend/exnet"

	"gifl/rds"
)

type Info struct {
	Ip        string    `json:"ip"`
	PublicIp  string    `json:"public ip"`
	VisitTime time.Time `json:"time"`
}

func Write(ctx *gear.Context) error {
	key := ctx.Req.URL.Query().Get("q")
	if key == "" {
		return ctx.HTML(http.StatusNotFound, "请添加输入参数[q]")
	}

	inf := &Info{
		Ip:        exnet.ClientIP(ctx.Req),
		PublicIp:  exnet.ClientPublicIP(ctx.Req),
		VisitTime: time.Now().Local(),
	}

	bf, err := rds.RDB.Get(ctx, key).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return ctx.HTML(http.StatusInternalServerError, "数据库数据获取失败")
		}
	}

	var infs []*Info
	switch {
	case len(bf) > 0:
		if err = json.Unmarshal(bf, &infs); err != nil {
			return ctx.HTML(http.StatusInternalServerError, "数据库数据反序列化")
		}
	default:
		infs = make([]*Info, 0)
	}

	infs = append([]*Info{inf}, infs...)
	buf, _ := json.Marshal(infs)

	status := rds.RDB.Set(ctx, key, buf, 0)
	logging.Info(status)

	return ctx.HTML(http.StatusOK, "hello")
}

func Read(ctx *gear.Context) error {
	key := ctx.Req.URL.Query().Get("q")
	if key == "" {
		return ctx.HTML(http.StatusNotFound, "请添加输入参数[q]")
	}

	bf, err := rds.RDB.Get(ctx, key).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return ctx.HTML(http.StatusInternalServerError, "数据库数据获取失败")
		}
	}

	var infs []*Info
	switch {
	case len(bf) > 0:
		if err = json.Unmarshal(bf, &infs); err != nil {
			return ctx.HTML(http.StatusInternalServerError, "数据库数据反序列化")
		}
	default:
		return ctx.OkHTML("没有查到对应的访问数据")
	}

	var content string
	for _, data := range infs {
		content += fmt.Sprintf("%s<br>%s<br>%s<br><hr>", data.Ip, data.PublicIp, data.VisitTime.String())
	}

	return ctx.OkHTML(content)
}

func Other(ctx *gear.Context) error {
	return ctx.HTML(http.StatusNotFound, "server not found")
}
