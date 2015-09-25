// Copyright (c) 2015 Gurpartap Singh
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"encoding/json"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int64
	CreatedAt    time.Time
	FullName     string
	Email        string
	PasswordHash string
}

type Users []*User

type UserFromJSON struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserToJSON struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
}

// Implements JSONDecodable for User
func (user *User) FromJSON(from []byte) error {
	a := new(UserFromJSON)

	err := json.Unmarshal(from, a)

	if err != nil {
		return err
	}

	// Validations
	if len(a.Email) == 0 || len(a.Password) == 0 {
		return errors.New("Missing email or password")
	}

	// Encrypt password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	(*user) = User{
		FullName:     a.FirstName + " " + a.LastName,
		Email:        a.Email,
		PasswordHash: string(passwordHash),
	}

	return nil
}

// Implements JSONEncodable for User
func (user *User) ToJSON() ([]byte, error) {
	return json.MarshalIndent(UserToJSON{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		FullName:  user.FullName,
		Email:     user.Email,
	}, "", "	")
}

// Implements JSONEncodable for Users
func (users *Users) ToJSON() ([]byte, error) {
	JSONs := make([]UserToJSON, 0)
	for _, user := range *users {
		j := UserToJSON{
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			FullName:  user.FullName,
			Email:     user.Email,
		}
		JSONs = append(JSONs, j)
	}
	return json.MarshalIndent(JSONs, "", "	")
}
