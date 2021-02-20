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

	"github.com/emersion/go-imap"
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

func (mod *ImapBee) checkForEmails() {
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
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")

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

	for {
		select {
		case <-mod.SigChan:
			return
		case <-time.After(time.Duration(15) * time.Second):
			mod.LogDebugf("Retrieving data from IMAP:")
			log.Println("imap pls d")
			mod.checkForEmails()
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
