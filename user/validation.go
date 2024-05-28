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
	"fmt"
	"regexp"

	"github.com/vitorreao/wallet-go/httperr"
)

const (
  usernameMaxLength = 16
  passwordMaxLength = 32
  passwordMinLength = 8
)

var userNameRegex = regexp.MustCompile("\\A[A-Za-z]+[0-9A-Za-z]*\\z")
var passwordRegex = regexp.MustCompile("\\A([0-9A-Za-z]|\\(|\\)|\\{|\\}|\\[|\\]|\\?|\\!|\\@|\\#|\\$|\\%|\\^|\\&|\\*)+\\z")
var phoneRegex = regexp.MustCompile(`\A\+(9[976][0-9]|8[987530][0-9]|6[987][0-9]|5[90][0-9]|42[0-9]|3[875][0-9]|2[98654321][0-9]|9[8543210]|8[6421]|6[6543210]|5[87654321]|
4[987654310]|3[9643210]|2[70]|7|1)[0-9]{1,14}\z`)

func validateCreateUserReq(req *CreateUserRequest) error {
  if req == nil {
    return httperr.NewBadRequest("Request body not found")
  }
  if err := validateUsername(req.Username); err != nil {
    return err
  }
  if err := validatePassword(req.Password); err != nil {
    return err
  }
  if err := validatePhone(req.PhoneNum); err != nil {
    return err
  }
  return nil
}

func validateUsername(username string) error {
  if username == "" {
    return httperr.NewBadRequest("Username cannot be empty.")
  }
  if len(username) > usernameMaxLength {
    return httperr.NewBadRequest(fmt.Sprintf(
      "Username is too long. Keep it below %d characters.",
      usernameMaxLength,
    ))
  }
  if !userNameRegex.MatchString(username) {
    return httperr.NewBadRequest(
      "Username has characters that are not allowed. " +
      "Use only alphanumeric characters.")
  }
  return nil
}

func validatePassword(password string) error {
  if password == "" {
    return httperr.NewBadRequest("Password cannot be empty.")
  }
  if len(password) > passwordMaxLength {
    return httperr.NewBadRequest("Your password is too long!")
  }
  if len(password) < passwordMinLength {
    return httperr.NewBadRequest("Your password is too short.")
  }
  if !passwordRegex.MatchString(password) {
    return httperr.NewBadRequest("Your password contains illegal characteres.")
  }
  return nil
}

func validatePhone(phone string) error {
  if !phoneRegex.MatchString(phone) {
    return httperr.NewBadRequest("Invalid phone number.")
  }
  return nil
}

