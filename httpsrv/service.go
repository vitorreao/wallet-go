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
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
  Register(r gin.IRouter)
}

type service struct {
  prefix string
  handlers []handler
}

func (s *service) Register(r gin.IRouter) {
  g := r.Group(fmt.Sprintf("/%s", s.prefix))
  for _, h := range s.handlers {
    g.Handle(h.Method, h.Path, wrapH(h.Func))
  }
}


