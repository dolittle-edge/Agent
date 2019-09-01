/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package provisioning

import (
	"agent/log"
	"agent/provisioning/system"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	//provisioningEndpoint   = "https://edge.dolittle.studio/api/Provisioning"
	provisioningEndpoint   = "http://localhost:5000/api/Provisioning"
	waitWhileNotConfigured = 1 * time.Second
	waitWhileConfigured    = 10 * time.Second
	waitWhileNotAuthorized = 10 * time.Second
	waitWhileInError       = 10 * time.Second
)

// Provider provides node configuration
type Provider struct {
	Current   Node
	listeners []chan<- Node
}

// NewProvider instanciates a new Provider
func NewProvider() *Provider {
	p := new(Provider)

	p.Current.isValid = false
	go p.runProvider()

	return p
}

// Listen notifies a listener about updates to the node configuration
func (p *Provider) Listen(listener chan<- Node) {
	listener <- p.Current
	p.listeners = append(p.listeners, listener)
}

func (p *Provider) notifyListeners(node Node) {
	for _, l := range p.listeners {
		select {
		case l <- node:
		}
	}
}

func (p *Provider) runProvider() {
	persisted, err := readPersistedConfiguration()
	if err == nil {
		p.Current = persisted
		p.notifyListeners(persisted)
	}

	for {
		if !p.Current.IsValid() {
			node, wait, err := getNodeConfiguration()
			if err == nil {
				p.Current = node
			}
			time.Sleep(wait)
		}
	}
}

func decodeConfiguration(data []byte) (node Node, err error) {
	err = json.Unmarshal(data, &node)
	if err != nil {
		log.Errorln("Could not decode Node configuration:", err)
		return
	}
	node.isValid = true
	return
}

func readPersistedConfiguration() (Node, error) {
	data, err := ioutil.ReadFile("node.json")
	if err != nil {
		return Node{}, err
	}
	return decodeConfiguration(data)
}

func readRecievedConfiguration(response *http.Response) (Node, error) {
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Node{}, err
	}
	return decodeConfiguration(data)
}

func getNodeConfiguration() (Node, time.Duration, error) {
	information, err := system.ReadSystemInformation()
	if err != nil {
		log.Errorln("Could not get system information in BIOS:", err)
		return Node{}, waitWhileInError, err
	}

	informationJSON, err := json.Marshal(information)
	if err != nil {
		log.Errorln("Error while marshalling system information:", err)
		return Node{}, waitWhileInError, err
	}
	informationReader := bytes.NewReader(informationJSON)

	// TODO: Close buffers
	response, err := http.Post(provisioningEndpoint+"/Get", "application/json", informationReader)
	if err != nil {
		log.Errorln("Error while getting configuration for node:", err)
		return Node{}, waitWhileNotConfigured, err
	}

	switch response.StatusCode {
	case http.StatusNotFound:
		log.Informationln("Node configuration not found - not provisioned.")
		return Node{}, waitWhileNotConfigured, errors.New("Node not provisioned")

	case http.StatusUnauthorized:
		log.Warningln("Node not authorized to get configuration.")
		return Node{}, waitWhileNotAuthorized, errors.New("Node not authorized")

	case http.StatusOK:
		node, err := readRecievedConfiguration(response)
		if err != nil {
			log.Errorln("Recieved node configuration not valid")
			return Node{}, waitWhileNotConfigured, err
		}
		return node, waitWhileConfigured, nil

	default:
		log.Errorln("Unexpected response code from provisioning endpoint:", response.StatusCode)
		return Node{}, waitWhileNotConfigured, errors.New("Unexpected status code")
	}
}
