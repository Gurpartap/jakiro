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
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/Gurpartap/jakiro"
)

func IndexUserHandler(c jakiro.Context) {
	users := make(Users, 0, len(DBUsersTable))

	var keys []int
	for k := range DBUsersTable {
		keys = append(keys, int(k))
	}

	sort.Ints(keys)

	// Select users sorted by ID.
	for _, k := range keys {
		users = append(users, DBUsersTable[int64(k)])
	}

	c.JSON(200, users)
	return
}

func CreateUserHandler(c jakiro.Context) {
	user := new(User)
	err := user.FromJSON(c.Body())

	if err != nil {
		// 400 Bad Request
		c.Error(400, err)
		return
	}

	// Insert the user.
	user.Id = int64(len(DBUsersTable) + 1)
	user.CreatedAt = time.Now()
	DBUsersTable[user.Id] = user

	// 201 Created
	c.JSON(201, user)
}

func ReadUserHandler(c jakiro.Context) {
	idStr := c.Params()["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		// 500 Internal Server Error
		c.Error(500, err)
		return
	}

	// Select the user.
	user, exists := DBUsersTable[id]

	if exists == false {
		c.Error(404, errors.New("User not found"))
		return
	}

	// 200 OK
	c.JSON(200, user)
	return
}

func DestroyUserHandler(c jakiro.Context) {
	idStr := c.Params()["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		// 500 Internal Server Error
		c.Error(500, err)
		return
	}

	_, exists := DBUsersTable[id]

	if exists == false {
		c.Error(404, errors.New("User not found"))
		return
	}

	// Delete the user.
	delete(DBUsersTable, id)

	// 202 Accepted
	c.Write(202, nil)
	return
}
