/*  Copyright (C) 2024 Vitor de Albuquerque Torreao
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorreao/wallet-go/httpsrv"
	"github.com/vitorreao/wallet-go/user"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
  en := gin.Default()
  en.Use(gin.Recovery())
  return en
}

func NewApiGroup(e *gin.Engine) gin.IRouter {
  return e.Group("/api")
}

type RouteParams struct {
  fx.In
  Router gin.IRouter
  Services []httpsrv.Service `group:"services"`
}

func main() {
  fx.New(
    fx.Provide(NewGinEngine),
    fx.Provide(NewApiGroup),
    user.Module,
    fx.Invoke(func (params RouteParams) {
      for _, srv := range params.Services {
        srv.Register(params.Router)
      }
    }),
    fx.Invoke(func(en *gin.Engine) {
      en.Run()
    }),
  ).Run()
}

