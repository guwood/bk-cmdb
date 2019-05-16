/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package x19_05_16_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addDefaultCategory(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	categoryName := "Default"

	// insert first category
	cond := metadata.BizLabelNotExist.Clone()
	cond.Set("name", categoryName)
	cond.Set("parent_id", 0)
	count, err := db.Table(common.BKTableNameServiceCategory).Find(cond).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	firstID, err := db.NextSequence(ctx, common.BKTableNameServiceCategory)
	if err != nil {
		return err
	}

	firstCategory := metadata.ServiceCategory{
		ID:              int64(firstID),
		Name:            categoryName,
		RootID:          int64(firstID),
		ParentID:        0,
		SupplierAccount: "0",
		IsBuiltIn:       true,
	}
	err = db.Table(common.BKTableNameServiceCategory).Insert(ctx, firstCategory)
	if err != nil {
		return err
	}

	// insert second category
	cond = metadata.BizLabelNotExist.Clone()
	cond.Set("name", categoryName)
	cond.Set("parent_id", firstCategory)
	count, err = db.Table(common.BKTableNameServiceCategory).Find(cond).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	secondID, err := db.NextSequence(ctx, common.BKTableNameServiceCategory)
	if err != nil {
		return err
	}

	secondCategory := metadata.ServiceCategory{
		ID:              int64(secondID),
		Name:            categoryName,
		RootID:          int64(firstID),
		ParentID:        int64(firstID),
		SupplierAccount: "0",
		IsBuiltIn:       true,
	}
	err = db.Table(common.BKTableNameServiceCategory).Insert(ctx, secondCategory)
	if err != nil {
		return err
	}

	return nil
}
