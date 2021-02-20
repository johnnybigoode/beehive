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
	"log"
	"time"

	"github.com/emersion/go-imap/client"

	"github.com/muesli/beehive/bees"
)

type ImapBee struct {
	bees.Bee

	username string
	password string
	server   string
}

// Action triggers the action passed to it.
func (mod *ImapBee) Action(action bees.Action) []bees.Placeholder {
	outs := []bees.Placeholder{}
	//send goes here
	panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	return outs
}

func (mod *ImapBee) getTotalMessages() int {
	log.Println("Connecting to server...")
	server := mod.server
	username := mod.username
	password := mod.password

	// Connect to server
	c, err := client.DialTLS(server, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select("INBOX", true) //read-only = true
	if err != nil {
		log.Fatal(err)
	}

	totalMessages := mbox.Messages
	log.Println("totalMessages: ")
	log.Println(totalMessages)
	log.Println("Done!")
	return int(totalMessages)
}

func (mod *ImapBee) Run(eventChan chan bees.Event) {
	log.Println("Server get test")
	log.Println("mod.server")
	log.Println(mod.server)
	log.Println("mod")
	log.Println(mod)
	log.Println("mod.username")
	log.Println(mod.username)
	log.Println("mod.passwotd")
	log.Println(mod.password)

	// if mod.interval < 1 {
	// 	mod.interval = defaultUpdateInterval
	// }

	//oldIP := mod.getIP("", eventChan)

	totalMessages := mod.getTotalMessages()
	previousTotalMessages := 0

	for {
		select {
		case <-mod.SigChan:
			return
		case <-time.After(time.Duration(15) * time.Second):
			mod.LogDebugf("Retrieving data from IMAP:")
			previousTotalMessages = totalMessages
			totalMessages := mod.getTotalMessages()
			if totalMessages > previousTotalMessages {
				log.Println("You got mail!")
			}
		}
	}
	/*	ev := bees.Event{
			Bee: mod.Name(),
			Name:      "hello",
			Options:   []bees.Placeholder{},
		}

		eventChan <- ev*/
}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *ImapBee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)

	options.Bind("account", &mod.username)
	options.Bind("password", &mod.password)
	options.Bind("server", &mod.server)
}
