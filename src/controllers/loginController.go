/*
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package controllers

import (
    "encoding/json"
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/services"
)

//Login Controller
type LoginController struct {
    authService services.IAuthService
}
func NewLoginController(authService services.IAuthService) *LoginController {
    return &LoginController{authService: authService}
}

func (l *LoginController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (l *LoginController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    //Parse request into struct
    decoder := json.NewDecoder(r.Body)
    var data PostLoginRequest
    err := decoder.Decode(&data)

    if err != nil {
        //if json doesn't map to struct return error
        return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
    } else {
        //check if the username / password match
        user, err := l.authService.GetAuth(data.Email)

        if(err != nil || user.Password != data.Password) {
            //If not return an error
            return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
        } else {
            //generate a jwt
            tokenStr, err := skaioskit.GenerateJWTStr(skaioskit.JWTData{ UserId: user.ID, Email: user.Email }, []byte("foobar"))

            if err != nil {
                //if we error jwt generating we should throw an error.
                return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
            } else {
                //and there was much rejoicing
                return skaioskit.ControllerResponse{Status: http.StatusOK, Body: PostLoginResponse{JWT: tokenStr}}
            }
        }
    }
}
func (l *LoginController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (l *LoginController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
