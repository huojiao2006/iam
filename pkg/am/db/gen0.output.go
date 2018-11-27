// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Code generated. DO NOT EDIT.

package db

import (
	"github.com/golang/protobuf/proto"

	"openpitrix.io/iam/pkg/pb/am"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ pbam.DbSchema
var _ proto.Message

var DbSchemaTableList = []proto.Message{
	new(pbam.DbTableSchema_Role),
	new(pbam.DbTableSchema_RoleBinding),
}
