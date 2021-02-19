/*
 *    Copyright (C) 2021 dione bigode
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      dione bigode <jamarson@gmail.com>
 */

package imapbee

import (
	"github.com/muesli/beehive/bees"
)

// ImapBeeFactory is a factory for ImapBees.
type ImapBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *ImapBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := ImapBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *ImapBeeFactory) ID() string {
	return "ImapBee"
}

// Name returns the name of this Bee.
func (factory *ImapBeeFactory) Name() string {
	return "Email - IMAP"
}

// Description returns the description of this Bee.
func (factory *ImapBeeFactory) Description() string {
	return "Bee for new emails event via IMAP"
}

// Options returns the options available to configure this Bee.
func (factory *ImapBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "server",
			Description: "IP or URL for the IMAP server with port (example: imap.gmail.com:993)",
			Type:        "string",
			Mandatory:   true,
		},
		{
			Name:        "account",
			Description: "Email account to be used with @",
			Type:        "string",
			Mandatory:   true,
		},
		{
			Name:        "password",
			Description: "The password used for the email account",
			Type:        "string",
			Mandatory:   true,
		},
	}
	return opts
}

// Events describes the available events provided by this Bee.
func (factory *ImapBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "email",
			Description: "An email receieved after the bee started working",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "sender",
					Description: "Who sent the email",
					Type:        "string",
				}, {
					Name:        "content",
					Description: "Email content",
					Type:        "string",
				},
				{
					Name:        "timestamp",
					Description: "When the new email was recieved",
					Type:        "timestamp",
				},
			},
		},
	}

	return events
}

func init() {
	f := ImapBeeFactory{}
	bees.RegisterFactory(&f)
}
