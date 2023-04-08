package newx

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/noahlsl/noah/help"
	"github.com/noahlsl/noah/tools/osx"
	"github.com/pkg/errors"
)

var project string

func New(args ...string) {

	if len(args) == 0 {
		help.Help()
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建项目
	projectName := fmt.Sprintf("%s/%s", wd, args[0])
	project = args[0]
	err = exec.Command("mkdir", projectName).Err
	if err != nil {
		log.Fatalln(err)
	}

	// 创建docs
	err = os.MkdirAll(fmt.Sprintf("%s/docs", projectName), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeEtc(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeMain(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeGitignore(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeMakeFile(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeBuildSh(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeREADME(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeMod(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	// 创建internal
	internalPath := fmt.Sprintf("%s/internal", projectName)
	err = exec.Command("mkdir", internalPath).Err
	if err != nil {
		log.Fatalln(err)
	}

	// config
	err = makeConfig(internalPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeSvc(internalPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeModel(internalPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeRouter(internalPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeHandler(internalPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = makeLogic(internalPath)
	if err != nil {
		log.Fatalln(err)
	}

	color.Green("The project created successfully")
	color.Green("Welcome to noah")
}

func makeConfig(path string) error {

	cfg := fmt.Sprintf("%s/config", path)
	err := exec.Command("mkdir", cfg).Err
	if err != nil {
		return errors.WithStack(err)
	}

	content :=
		`package config

import (
	"github.com/noahlsl/noah/tools/logx"
	"github.com/spf13/viper"
)

type Conf struct {
	Name string // 服务名称
	Port int    // 监听端口
	LogConf logx.Conf // 日志配置文件
}

var defaultCfgFile = "etc/config.yaml"

// MustLoad 配置文件读取
// addr ETCD地址
// name 服务名称
// env  服务环境
func MustLoad(addr, name, env string) *Conf {

	var c = &Conf{
		Name: name,
	}
	// 本地配置文件
	if addr == "" {
		// 解析配置文件
		v := viper.New()
		v.AddConfigPath("etc/") // 查找配置文件所在的路径
		v.AddConfigPath("etc")
		v.SetConfigType("yaml")
		v.SetConfigName("config")
		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}

		err = v.Unmarshal(&c)
		if err != nil {
			panic(err)
		}

		return c
	}

	// 远程配置文件

	return c
}

`
	fileName := fmt.Sprintf("%s/config.go", cfg)
	return osx.Write(fileName, content)
}
func makeMain(path string) error {

	fileName := fmt.Sprintf("%s/main.go", path)
	if osx.Exists(fileName) {
		log.Printf("the %s is exist\n", fileName)
		return nil
	}

	content := fmt.Sprintf(
		`package main

import (
	"flag"

	"%s/internal/config"
	"%s/internal/router"
	"%s/internal/svc"

	"github.com/noahlsl/noah/models/versionx"
)

var (
	serverName string // 服务名称
	buildTime  string // 编译时间
	commitId   string // 最新CommitId
	branch     string // 代码分支名称

	a = flag.String("a", "", "the etcd address")
	e = flag.String("e", "dev", "the server env")
	n = flag.String("n", "%s", "the server name")
)

// @title Basic Server
// @version 1.0 版本
// @description Basic 基础模版
// @BasePath /%s/v1  基础路径
func main() {
	flag.Parse()

	// 配置文件
	c := config.MustLoad(*a, *e, *n)
	// 编译信息
	ver := versionx.NewVersion(serverName, buildTime, commitId, branch,c.Port)
	// 依赖连接
	svcCtx := svc.NewServiceContext(c, ver)
	// 路由注册
	r := router.NewRouter(svcCtx)
	// 服务启动
	svcCtx.Start(r)
}

`, project, project, project, project, project)

	return osx.Write(fileName, content)
}
func makeGitignore(path string) error {

	content :=
		`
*.exe
*.exe~
*.dll
*.so
*.dylib


*.test

*.out


go.work

.idea
`
	fileName := fmt.Sprintf("%s/.gitignore", path)
	return osx.Write(fileName, content)
}
func makeMakeFile(path string) error {
	content := fmt.Sprintf(
		`# 定义服务名称
server=%s

# 初始化及更新依赖
init:
	go mod tidy

# 根据API文件生成代码
gen:
	sh gen.sh "$(server)"

# 工具API文件生成对接文档
swag:
	swag init

# 打包编译
build:
	sh build.sh "$(server)"
`, project)
	fileName := fmt.Sprintf("%s/Makefile", path)
	return osx.Write(fileName, content)
}
func makeBuildSh(path string) error {
	content := fmt.Sprintf(
		`#!/bin/sh

commit=$(git show -s --format=%%H)
buildTime=$(date "+%%Y-%%m-%%d %%H:%%M:%%S")
serverName=%s
branch=$(git symbolic-ref --short HEAD)
ldflags="-X 'main.commitId=$commit' -X 'main.buildTime=$buildTime' -X 'main.serverName=$serverName' -X 'main.branch=$branch'"

# shellcheck disable=SC2154
go build -ldflags "$ldflags" -o "$serverName"
`, project)
	fileName := fmt.Sprintf("%s/build.sh", path)
	return osx.Write(fileName, content)
}
func makeREADME(path string) error {
	content := fmt.Sprintf(
		`# 1.Server Info %s
`, project)

	fileName := fmt.Sprintf("%s/README.md", path)
	return osx.Write(fileName, content)
}
func makeSvc(path string) error {

	content := fmt.Sprintf(
		`package svc

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"%s/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/noahlsl/noah/models/versionx"
	"github.com/noahlsl/noah/tools/logx"

)

type ServiceContext struct {
	c        *config.Conf
	Ver      *versionx.Version
}

// NewServiceContext 构建依赖
func NewServiceContext(c *config.Conf, ver *versionx.Version) *ServiceContext {

	// 初始化日志依赖
	logx.MustUpSet(&c.LogConf)
	return &ServiceContext{
		c:        c,
		Ver:      ver,
	}
}

// Start 开始服务
func (s *ServiceContext) Start(r *gin.Engine) {

	logx.Infof("The server starting at %%d...", s.c.Port)
	logx.Infof("The swagger url is http://127.0.0.1:%%d/%%s/swagger/index.html", s.c.Port, s.c.Name)

	server := http.Server{
		Addr:         fmt.Sprintf(":%%d", s.c.Port),
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 在此阻塞
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	logx.Infof("The server stopping at %%d...", s.c.Port)
	// 资源回收
	s.Close()
	err := server.Shutdown(ctx)
	if err != nil {
		logx.Error(err)
	}
}

// Close 资源回收
func (s *ServiceContext) Close() {
// TODO Do the Close logic

}

`, project)

	fileName := fmt.Sprintf("%s/svc/svc.go", path)
	return osx.Write(fileName, content)
}
func makeModel(path string) error {
	content := fmt.Sprintf(`
package model
type Res struct {
	Code  int         
	Msg   string     
	Data  interface{} 
	Trace string      
	Ts    int64     
}
`)
	fileName := fmt.Sprintf("%s/model/common.go", path)

	return osx.Write(fileName, content)
}
func makeRouter(path string) error {
	content := fmt.Sprintf(
		`package router

import (
	_ "%s/docs"
	"%s/internal/handler"
	"%s/internal/svc"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	sgf "github.com/swaggo/files"
	sg "github.com/swaggo/gin-swagger"
)

func NewRouter(svcCtx *svc.ServiceContext) *gin.Engine {

	// 不打印日志
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 性能监控中间件
	r.Use(ginprom.PromMiddleware(nil))
	// 跨域
	// 获取编译版本信息
	r.GET("/base/version", handler.GetVersion(svcCtx))
	// 获取性能监控
	r.GET("/base/metrics", ginprom.PromHandler(promhttp.Handler()))
	// 获取swagger对接文档
	r.GET("/base/swagger/*any", sg.WrapHandler(sgf.Handler))

	//v1 := r.Group("/v1")
	// TODO 写其他逻辑路由

	return r
}

`, project, project, project)

	fileName := fmt.Sprintf("%s/router/router.go", path)
	return osx.Write(fileName, content)
}
func makeHandler(path string) error {
	content := fmt.Sprintf(
		`package handler

import (
	"%s/internal/logic"
	_ "%s/internal/model"
	"%s/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/noahlsl/noah/tools/resultx"
	"github.com/noahlsl/noah/consts"
)

// GetVersion
// @Summary 获取当前服务编译版本信息
// @Description  获取当前服务编译版本信息
// @Description 编译版本信息
// @Schemes
// ShowAccount godoc
// @Tags 基础接口
// @Accept       json
// @Produce      json
// @Router /base/version [GET]
// @Success      200  {object}  model.Res
// @Failure      400  {object}  model.Res
func GetVersion(svcCtl *svc.ServiceContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		l := logic.NewVersionCtl(svcCtl)
		version, err := l.GetVersion()
		if err != nil {
			resultx.ResultErr(c, consts.ErrCodeSysBadRequest, err)
			return
		}

		resultx.Result(c, version)
	}
}

`, project, project, project)

	fileName := fmt.Sprintf("%s/handler/version.go", path)
	return osx.Write(fileName, content)
}

func makeMod(path string) error {
	content := fmt.Sprintf(`module %s`, project)
	fileName := fmt.Sprintf("%s/go.mod", path)
	return osx.Write(fileName, content)
}

func makeEtc(path string) error {
	content := fmt.Sprintf(
		`Name: %s
Host: 127.0.0.1
Port: 8080
`,
		project)
	fileName := fmt.Sprintf("%s/etc/config.yaml", path)
	return osx.Write(fileName, content)
}

func makeLogic(path string) error {
	content := fmt.Sprintf(
		`package logic

import (
	"%s/internal/svc"

	"github.com/noahlsl/noah/models/versionx"
)

type VersionCtl struct {
	svcCtx *svc.ServiceContext
}

func NewVersionCtl(ctx *svc.ServiceContext) *VersionCtl {
	return &VersionCtl{
		svcCtx: ctx,
	}
}

func (s *VersionCtl) GetVersion() (*versionx.Version, error) {

	return s.svcCtx.Ver, nil
}
`,
		project)
	fileName := fmt.Sprintf("%s/logic/version.go", path)
	return osx.Write(fileName, content)
}
