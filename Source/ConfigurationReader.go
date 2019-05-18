/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

import (
	"encoding/json"
	"io/ioutil"
)

// ReadConfiguration reads and deserializes the configuration
func ReadConfiguration() Node {
	data, err := ioutil.ReadFile("node.json")
	if err != nil {
		panic(err)
	}

	node := Node{}

	json.Unmarshal([]byte(data), &node)

	return node
}
