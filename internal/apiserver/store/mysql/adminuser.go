/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"context"

	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/model"
	"blog-api/internal/pkg/util/gormutil"
	"blog-api/pkg/errors"
	"blog-api/pkg/fields"
	metav1 "blog-api/pkg/meta/v1"
	"gorm.io/gorm"
)

type adminUser struct {
	db *gorm.DB
}

func newAdminUser(db *gorm.DB) *adminUser {
	return &adminUser{db: db}
}

func (u *adminUser) Create(ctx context.Context, adminUserModel *model.AdminUser, opts metav1.CreateOptions) error {

	return u.db.Create(adminUserModel).Error
}

func (u *adminUser) List(cxt context.Context, opts metav1.ListOptions) (*model.AdminUserList, error) {
	userList := &model.AdminUserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, _ := selector.RequiresExactMatch("name")
	d := u.db.Where("name like ? and status = 1", "%"+username+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&userList.Items).
		Offset(-1).
		Limit(-1).
		Count(&userList.Total)

	return userList, d.Error
}

func (u *adminUser) Delete(ctx context.Context, account string, opts metav1.DeleteOptions) error {
	return nil
}

func (u *adminUser) Get(ctx context.Context, username string, opts metav1.GetOptions) (*model.AdminUser, error) {
	user := &model.AdminUser{}

	err := u.db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}
	}

	return user, nil
}
