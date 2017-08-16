// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	taskARN       = "t1"
	attachmentARN = "att1"
	mac           = "mac1"
	attachSent    = false
)

func TestMarshalUnmarshal(t *testing.T) {
	expiresAt := time.Now()
	attachment := &ENIAttachment{
		TaskARN:          taskARN,
		AttachmentARN:    attachmentARN,
		AttachStatusSent: attachSent,
		MACAddress:       mac,
		Status:           ENIAttachmentNone,
		ExpiresAt:        expiresAt,
	}
	bytes, err := json.Marshal(attachment)
	assert.NoError(t, err)
	var unmarshalledAttachment ENIAttachment
	err = json.Unmarshal(bytes, &unmarshalledAttachment)
	assert.NoError(t, err)
	assert.Equal(t, attachment.String(), unmarshalledAttachment.String())
}

func TestStartTimerErrorWhenExpiresAtIsInThePast(t *testing.T) {
	expiresAt := time.Now().Unix() - 1
	attachment := &ENIAttachment{
		TaskARN:          taskARN,
		AttachmentARN:    attachmentARN,
		AttachStatusSent: attachSent,
		MACAddress:       mac,
		Status:           ENIAttachmentNone,
		ExpiresAt:        time.Unix(expiresAt, 0),
	}
	assert.Error(t, attachment.StartTimer(func() {}))
}
