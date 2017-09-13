/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package akira

import (
	"time"

	"github.com/nats-io/nats"
	"github.com/nats-io/nuid"
)

// FakeConnector : A fake nats connector for testing nats handlers
type FakeConnector struct {
	Events   map[string][]*nats.Msg
	Handlers map[string]nats.MsgHandler
	gen      *nuid.NUID
}

// NewFakeConnector : Returns a new fake connector
func NewFakeConnector() Connector {
	return &FakeConnector{
		Events:   make(map[string][]*nats.Msg),
		Handlers: make(map[string]nats.MsgHandler),
		gen:      nuid.New(),
	}
}

// Reset : resets all handlers and events
func (f *FakeConnector) Reset() {
	f.ResetEvents()
	f.ResetHandlers()
}

// ResetEvents : Resets cache of collected events
func (f *FakeConnector) ResetEvents() {
	f.Events = make(map[string][]*nats.Msg)
}

// ResetHandlers : Resets all handlers
func (f *FakeConnector) ResetHandlers() {
	f.Handlers = make(map[string]nats.MsgHandler)
}

// Close : Resets all handlers and events
func (f *FakeConnector) Close() {
	f.Reset()
}

// Request : Make a request
func (f *FakeConnector) Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	msg := &nats.Msg{
		Subject: subj,
		Reply:   "_INBOX." + f.gen.Next(),
		Data:    data,
	}

	f.Events[subj] = append(f.Events[subj], msg)

	if f.Handlers[subj] == nil {
		return nil, nats.ErrTimeout
	}

	f.Handlers[subj](msg)

	if len(f.Events[msg.Reply]) < 1 {
		return nil, nats.ErrTimeout
	}

	return f.Events[msg.Reply][0], nil
}

// Publish : Publish an event
func (f *FakeConnector) Publish(subj string, data []byte) error {
	msg := &nats.Msg{Subject: subj, Data: data}
	f.Events[subj] = append(f.Events[subj], msg)
	return nil
}

// Subscribe : Subscribe to an event stream
func (f *FakeConnector) Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	f.Handlers[subj] = cb
	return nil, nil
}

// QueueSubscribe : Subscribe to an event stream
func (f *FakeConnector) QueueSubscribe(subj string, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	f.Handlers[subj] = cb
	return nil, nil
}
