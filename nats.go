/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package akira

import (
	"time"

	"github.com/nats-io/nats"
)

// Connector provides an interface to NATS. This allows the service to be
// mocked so an actual NATS instance isn't required for unit testing.
type Connector interface {
	Close()
	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)
	Publish(subj string, data []byte) error
	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)
}
