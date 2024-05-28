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

package httperr

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
  error
  Code() int
}

type httpError struct {
  code int
  msg string
}

func (e *httpError) Error() string {
  return fmt.Sprintf("(HTTP %d) %s", e.code, e.msg)
}

func (e *httpError) Code() int {
  return e.code
}

func FromError(err error) HTTPError {
  if httpErr, ok := err.(*httpError); ok {
    return httpErr
  }
  return &httpError{
    code: http.StatusInternalServerError,
    msg: err.Error(),
  }
}

func NewBadRequest(msg string) error {
  return &httpError{
    code: http.StatusBadRequest,
    msg: msg,
  }
}

func NewNotImplemented(msg string) error {
  return &httpError{
    code: http.StatusNotImplemented,
    msg: msg,
  }
}

