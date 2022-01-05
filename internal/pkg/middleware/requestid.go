/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import "github.com/gin-gonic/gin"
import uuid "github.com/satori/go.uuid"

const (
	XRequestIDKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(XRequestIDKey)
		if rid == "" {
			rid = uuid.NewV4().String()
			c.Request.Header.Set(XRequestIDKey, rid)
			c.Set(XRequestIDKey, rid)
		}

		c.Writer.Header().Set(XRequestIDKey, rid)
		c.Next()
	}
}
