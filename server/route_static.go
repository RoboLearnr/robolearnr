package main

//go:generate go-bindata -o bindata.go ../build/web/ ../build/web/static/js ../build/web/static/css ../build/web/sounds

import (
	"github.com/Code-Hex/echo-static"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
)

func HandleStatic() echo.MiddlewareFunc {
	return static.ServeRoot("/", NewAssets("../build/web/"))
}

func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}
