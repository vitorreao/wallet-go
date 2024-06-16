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

package httpsrv

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorreao/wallet-go/httperr"
)

type Service interface {
  Register(r gin.IRouter)
}

type Option func (g *gin.RouterGroup)

type service struct {
  prefix string
  opts []Option
}

func (s *service) Register(r gin.IRouter) {
  g := r.Group(fmt.Sprintf("/%s", s.prefix))
  for _, opt := range s.opts {
    opt(g)
  }
}

func NewService(prefix string, opts ...Option) Service {
  return &service{
    prefix: prefix,
    opts: opts,
  }
}

func WithPost[TReq any, TRes any](
  path string,
  hf HandlerFunc[TReq, TRes],
) Option {
  return withHandler(http.MethodPost, path, hf)
}

func withHandler[TReq any, TRes any](
  method string,
  path string,
  hf HandlerFunc[TReq, TRes],
) Option {
  return func (g *gin.RouterGroup) {
    g.Handle(method, path, func(c *gin.Context) {
      // TODO: get context from gin context
      ctx := context.Background()
      req := Request[TReq]{}
      if err := c.ShouldBindJSON(&req.Body); err != nil {
        c.JSON(http.StatusBadRequest,
          fmt.Sprintf("Error deserializing request body: %s", err.Error()))
        return
      }
      res, err := hf(ctx, &req)
      herr := httperr.FromError(err)
      if herr != nil {
        c.JSON(herr.Code(), herr.Error())
        return
      }
      if res == nil {
        c.Status(http.StatusOK)
        return
      }
      c.JSON(res.Code, res.Data)
    })
  }
}

