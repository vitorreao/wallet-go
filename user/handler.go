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

package user

import (
	"context"

	"github.com/vitorreao/wallet-go/httperr"
	"github.com/vitorreao/wallet-go/httpsrv"
)

type Handler interface {
  CreateUser(
    ctx context.Context,
    req *httpsrv.Request[CreateUserRequest],
  ) (*httpsrv.Response[CreateUserResponse], error)
}

type handler struct {}

func NewHandler() Handler {
  return &handler{}
}

func (h *handler) CreateUser(
  ctx context.Context,
  req *httpsrv.Request[CreateUserRequest],
) (*httpsrv.Response[CreateUserResponse], error) {
  return nil, httperr.NewNotImplemented("Create user is not available yet")
}

