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

package services

import (
    "github.com/jinzhu/gorm"

    "skaioskit/models"
)

type IAuthService interface {
    CreateAuth(models.Auth) models.Auth
    UpdateAuth(models.Auth) models.Auth
    GetAuth(string) (models.Auth, error)
    EnsureAuthTable()
}

type AuthService struct {
    db *gorm.DB
}
func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{db: db}
}
func (p *AuthService) CreateAuth(auth models.Auth) models.Auth {
    p.db.Create(&auth)
    return auth
}
func (p *AuthService) UpdateAuth(auth models.Auth) models.Auth {
    p.db.Save(&auth)
    return auth
}
func (p *AuthService) GetAuth(name string) (models.Auth, error) {
    var auth models.Auth
    err := p.db.Where(&models.Auth{Name: name}).First(&auth).Error
    return auth, err
}
func (p *AuthService) EnsureAuthTable() {
    p.db.AutoMigrate(&models.Auth{})
    p.db.Model(&models.Auth{}).AddUniqueIndex("idx_auth_code", "code")
}
func (p *AuthService) EnsureAuth(auth models.Auth) {
    existing, err := p.GetAuth(auth.Name)
    if err != nil {
        p.CreateAuth(auth)
    } else {
        existing.Name = auth.Name
        p.UpdateAuth(existing)
    }
}
