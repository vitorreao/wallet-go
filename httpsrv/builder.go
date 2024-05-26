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
	"net/http"
)

type Builder struct {
  prefix string
  handlers map[string]Handler
}

func (b *Builder) Build() Service {
  handlers := []Handler{}
  for _, handler := range b.handlers {
    handlers = append(handlers, handler)
  }
  return &service{
    prefix: b.prefix,
    handlers: handlers,
  }
}

func (b *Builder) WithPrefix(prefix string) *Builder {
  if b == nil {
    return nil
  }
  b.prefix = prefix
  return b
}

func (b *Builder) WithGet(path string, f HandlerFunc) *Builder {
  return b.WithHandle(http.MethodGet, path, f)
}

func (b *Builder) WithPost(path string, f HandlerFunc) *Builder {
  return b.WithHandle(http.MethodPost, path, f)
}

func (b *Builder) WithHandle(
  httpMethod string,
  path string,
  f HandlerFunc,
) *Builder {
  if b == nil {
    return b
  }
  if b.handlers == nil {
    b.handlers = map[string]Handler{}
  }
  b.handlers[handlerKey(httpMethod, path)] = Handler{
    Method: httpMethod,
    Path: path,
    Func: f,
  }
  return b
}

func handlerKey(httpMethod, path string) string {
  return fmt.Sprintf("%s:%s", httpMethod, path)
}

