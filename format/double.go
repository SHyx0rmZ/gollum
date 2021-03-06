// Copyright 2015 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package format

import (
	"github.com/trivago/gollum/core"
)

// Double formatter plugin
//
// Double is a formatter that doubles the message and glues both parts
// together by using a separator. Both parts of the new message may be
// formatted differently
//
// Parameters
//
// - Separator: This value sets the separator string placed between both parts.
// This parameter is set to ":" by default.
//
// - UseLeftStreamID: Use the stream id of the left side as the final stream id
// for the message if this value is "true".
// This parameter is set to "false" by default.
//
// - Left: A optional formatter list which are used for the left side.
// This parameter is set to "empty list" by default.
//
// - Right: A optional formatter list which are used for the right side.
// This parameter is set to "empty list" by default.
//
// Examples
//
// This example create a message where you find a "input|base64" pair of the original console input:
//
//  exampleConsumer:
//    Type: consumer.Console
//    Streams: "*"
//    Modulators:
//      - format.Double:
//	      Separator: "|"
//    	  Right:
//          - format.Base64Encode
//
type Double struct {
	core.SimpleFormatter `gollumdoc:"embed_type"`
	separator            []byte              `config:"Separator" default:":"`
	leftStreamID         bool                `config:"UseLeftStreamID" default:"false"`
	left                 core.FormatterArray `config:"Left"`
	right                core.FormatterArray `config:"Right"`
	applyTo              string
}

func init() {
	core.TypeRegistry.Register(Double{})
}

// Configure initializes this formatter with values from a plugin config.
func (format *Double) Configure(conf core.PluginConfigReader) {
	format.applyTo = conf.GetString("ApplyTo", "")
}

// ApplyFormatter update message payload
func (format *Double) ApplyFormatter(msg *core.Message) error {
	leftMsg := msg.Clone()
	rightMsg := msg.Clone()

	// pre-process
	if format.applyTo != "" {
		leftMsg.StorePayload(format.GetAppliedContent(msg))
		rightMsg.StorePayload(format.GetAppliedContent(msg))
	}

	// apply sub-formatter
	if err := format.left.ApplyFormatter(leftMsg); err != nil {
		return err
	}

	if err := format.right.ApplyFormatter(rightMsg); err != nil {
		return err
	}

	// update content
	format.SetAppliedContent(msg, format.getCombinedContent(leftMsg.GetPayload(), rightMsg.GetPayload()))

	// handle streamID
	if format.leftStreamID {
		msg.SetStreamID(leftMsg.GetStreamID())
	} else {
		msg.SetStreamID(rightMsg.GetStreamID())
	}

	// fin
	return nil
}

func (format *Double) getCombinedContent(leftContent []byte, rightContent []byte) []byte {
	size := len(leftContent) + len(format.separator) + len(rightContent)
	content := make([]byte, 0, size)

	content = append(content, leftContent...)
	content = append(content, format.separator...)
	content = append(content, rightContent...)

	return content

}
