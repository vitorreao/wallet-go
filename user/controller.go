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
)

type Controller interface {
  // CreateUserRegistration adds a new registration to the database. This
  // registration represents an intent to register a new user. A User account
  // is only actually created when the user confirms the registration afterwards.
  CreateUserRegistration(
    ctx context.Context,
    req CreateUserRequest,
  ) (*CreateUserResponse, error)
}

type controller struct {}

func NewController() Controller {
  return &controller{}
}

func (c *controller) CreateUserRegistration(
  ctx context.Context,
  req CreateUserRequest,
) (*CreateUserResponse, error) {
  return nil, httperr.NewNotImplemented("CreateUserRegistration is not implemented yet")
}

