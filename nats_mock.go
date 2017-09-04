/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package akira

import (
	"time"

	"github.com/nats-io/nats"
)

type FakeConnector struct {
	Events   map[string][]*nats.Msg
	Handlers map[string]nats.MsgHandler
}

func NewFakeConnector() Connector {
	return &FakeConnector{
		Events:   make(map[string][]*nats.Msg),
		Handlers: make(map[string]nats.MsgHandler),
	}
}

func (f *FakeConnector) Reset() {
	f.ResetEvents()
	f.ResetHandlers()
}

func (f *FakeConnector) ResetEvents() {
	f.Events = make(map[string][]*nats.Msg)
}

func (f *FakeConnector) ResetHandlers() {
	f.Handlers = make(map[string]nats.MsgHandler)
}

func (f *FakeConnector) Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	msg := &nats.Msg{Subject: subj, Data: data}
	f.Events[subj] = append(f.Events[subj], msg)

	f.Handlers[subj](msg)

	return msg, nil
}

func (f *FakeConnector) Publish(subj string, data []byte) error {
	msg := &nats.Msg{Subject: subj, Data: data}
	f.Events[subj] = append(f.Events[subj], msg)
	return nil
}

func (f *FakeConnector) Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	f.Handlers[subj] = cb
	return nil, nil
}
