/*
 *    Copyright (C) 2014 Christian Muehlhaeuser
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
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

// beehive's Serial module.
package serialbee

import (
	"github.com/huin/goserial"
	"github.com/muesli/beehive/modules"
	"io"
	"log"
	"bytes"
	"encoding/binary"
	"time"
)

type SerialBee struct {
	modules.Module

	conn io.ReadWriteCloser

	device   string
	baudrate     int
}

func (mod *SerialBee) Action(action modules.Action) []modules.Placeholder {
	outs := []modules.Placeholder{}
	text := ""

	switch action.Name {
	case "send":
		for _, opt := range action.Options {
			if opt.Name == "text" {
				text = opt.Value.(string)
			}
		}

		bufOut := new(bytes.Buffer)
		err := binary.Write(bufOut, binary.LittleEndian, text)
		if err != nil {
			panic(err)
		}

		for _, v := range [][]byte{bufOut.Bytes()} {
			_, err = mod.conn.Write(v)
			if err != nil {
				panic(err)
			}
		}

	default:
		// unknown action
		return outs
	}

	return outs
}

func (mod *SerialBee) Run(eventChan chan modules.Event) {
	if mod.baudrate == 0 || mod.device == "" {
		return
	}

	var err error
	c := &goserial.Config{Name: mod.device, Baud: mod.baudrate}
	mod.conn, err = goserial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	for {
/*			if len(v.Text) > 0 {
				text := strings.TrimSpace(v.Text)

				ev := modules.Event{
					Bee:  mod.Name(),
					Name: "message",
					Options: []modules.Placeholder{
						modules.Placeholder{
							Name:  "user",
							Type:  "string",
							Value: v.Remote,
						},
						modules.Placeholder{
							Name:  "text",
							Type:  "string",
							Value: text,
						},
					},
				}
				eventChan <- ev
			}*/
			time.Sleep(1 * time.Second)
	}
}