/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package akira

import (
	"strings"
	"testing"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/stretchr/testify/assert"
)

func TestFakeRequest(t *testing.T) {
	cases := []struct {
		Name, Subject     string
		Message, Expected []byte
		Error             error
	}{
		{"valid_subject", "test.message", []byte(`{"message": "ping"}`), []byte(`{"message": "pong"}`), nil},
		{"invalid_subject", "test.broken", []byte(`{"message": "ping"}`), nil, nats.ErrTimeout},
	}

	fc := NewFakeConnector()
	fc.Subscribe("test.message", func(msg *nats.Msg) {
		if strings.Contains(string(msg.Data), "ping") {
			fc.Publish(msg.Reply, []byte(`{"message": "pong"}`))
		}
	})

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			resp, err := fc.Request(tc.Subject, tc.Message, time.Second)
			assert.Equal(t, tc.Error, err)
			if resp != nil {
				assert.Equal(t, tc.Expected, resp.Data)
			}
		})
	}
}

func TestFakeSubcribe(t *testing.T) {
	cases := []struct {
		Name, Subject string
		Function      nats.MsgHandler
		Expected      []byte
	}{
		{"valid_event", "test.message", func(msg *nats.Msg) { msg.Data = []byte(`{"message": "pong"}`) }, []byte(`{"message": "pong"}`)},
		{"invalid_event", "test.broken", func(msg *nats.Msg) { msg.Data = []byte(`{"message": "pong"}`) }, nil},
	}

	fc := NewFakeConnector()
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := fc.Subscribe("test.message", tc.Function)
			assert.Nil(t, err)
			resp, _ := fc.Request(tc.Subject, nil, time.Second)
			if resp != nil {
				assert.Equal(t, tc.Expected, resp.Data)
			}
		})
	}
}

func TestFakePublish(t *testing.T) {
	cases := []struct {
		Name, Subject string
		Event         []byte
	}{
		{"valid_event", "test.message", []byte(`{"message": "ping"}`)},
	}

	fc := NewFakeConnector()
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			err := fc.Publish(tc.Subject, tc.Event)
			assert.Nil(t, err)
			x := fc.(*FakeConnector)
			assert.Equal(t, len(x.Events[tc.Subject]), 1)
			assert.Equal(t, x.Events[tc.Subject][0].Data, tc.Event)
		})
	}
}

func TestFakeReset(t *testing.T) {
	fc := NewFakeConnector()
	x := fc.(*FakeConnector)

	fc.Subscribe("test.message", func(msg *nats.Msg) {})

	x.Reset()

	_, err := fc.Request("test.message", nil, time.Second)
	assert.Equal(t, nats.ErrTimeout, err)
}
