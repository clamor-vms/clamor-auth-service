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
    "os"
    "encoding/json"
    "net/http"

    clamor "github.com/clamor-vms/clamor-go-core"

    "clamor/services"
)

//Login Controller
type LoginController struct {
    authService services.IAuthService
}
func NewLoginController(authService services.IAuthService) *LoginController {
    return &LoginController{authService: authService}
}

func (l *LoginController) Get(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (l *LoginController) Post(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    //Parse request into struct
    decoder := json.NewDecoder(r.Body)
    var data PostLoginRequest
    err := decoder.Decode(&data)

    if err != nil {
        //if json doesn't map to struct return error
        return clamor.ControllerResponse{Status: http.StatusInternalServerError, Body: clamor.EmptyResponse{}}
    } else {
        //check if the username / password match
        user, err := l.authService.GetAuth(data.Email)

        if(err != nil || user.Password != data.Password) {
            //If not return an error
            return clamor.ControllerResponse{Status: http.StatusInternalServerError, Body: clamor.EmptyResponse{}}
        } else {
            //generate a jwt
            tokenStr, err := clamor.GenerateJWTStr(clamor.JWTData{ UserId: user.ID, Email: user.Email }, []byte(os.Getenv("JWT_SECRET")))

            if err != nil {
                //if we error jwt generating we should throw an error.
                return clamor.ControllerResponse{Status: http.StatusInternalServerError, Body: clamor.EmptyResponse{}}
            } else {
                //and there was much rejoicing
                return clamor.ControllerResponse{Status: http.StatusOK, Body: PostLoginResponse{JWT: tokenStr}}
            }
        }
    }
}
func (l *LoginController) Put(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (l *LoginController) Delete(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
