/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package provisioning

import (
	"agent/log"
	"agent/provisioning/system"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
				log.Informationln("Recieved new node configuration")
				persistConfiguration(node)
				p.Current = node
				p.notifyListeners(node)
			}
			time.Sleep(wait)
		} else {
			node, changed, wait, err := checkForNodeConfigurationUpdates(p.Current)
			if err == nil && changed {
				log.Informationln("Recieved update for node configuration")
				persistConfiguration(node)
				p.Current = node
				p.notifyListeners(node)
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

func encodeConfiguration(node Node) (data []byte, err error) {
	data, err = json.Marshal(node)
	if err != nil {
		log.Errorln("Could not encode Node configuration:", err)
	}
	return
}

func persistConfiguration(node Node) error {
	if node.IsValid() {
		data, err := encodeConfiguration(node)
		err = ioutil.WriteFile("node.json", data, 0700)
		if err != nil {
			log.Errorln("Could not write node.json file:", err)
			return err
		}
	} else {
		err := os.Remove("node.json")
		if err != nil {
			log.Errorln("Could not delete node.json file:", err)
			return err
		}
	}
	return nil
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
			return Node{}, waitWhileInError, err
		}
		return node, waitWhileConfigured, nil

	default:
		log.Errorln("Unexpected response code from provisioning endpoint:", response.StatusCode)
		return Node{}, waitWhileInError, errors.New("Unexpected status code")
	}
}

func checkForNodeConfigurationUpdates(current Node) (Node, bool, time.Duration, error) {
	data, err := encodeConfiguration(current)
	if err != nil {
		return Node{}, false, waitWhileInError, err
	}

	hash := sha256.Sum256(data)
	information := url.Values{}
	information.Add("NodeId", current.NodeId)
	information.Add("Hash", base64.StdEncoding.EncodeToString(hash[0:32]))

	response, err := http.PostForm(provisioningEndpoint+"/Check", information)
	if err != nil {
		log.Errorln("Error while checking configuration updates for node:", err)
		return Node{}, false, waitWhileInError, nil
	}

	switch response.StatusCode {
	case http.StatusNotFound:
		log.Warningln("Node was configured - but could not find it while checking for updates")
		return Node{}, true, waitWhileNotConfigured, nil

	case http.StatusUnauthorized:
		log.Informationln("Node configuration was revoked")
		return Node{}, true, waitWhileNotAuthorized, nil

	case http.StatusNotModified:
		log.Informationln("Configuration for node not changed")
		return Node{}, false, waitWhileConfigured, nil

	case http.StatusOK:
		node, err := readRecievedConfiguration(response)
		if err != nil {
			log.Errorln("Recieved node configuration update not valid")
			return Node{}, false, waitWhileInError, err
		}
		return node, true, waitWhileConfigured, nil

	default:
		log.Errorln("Unexpected response code from provisioning endpoint:", response.StatusCode)
		return Node{}, false, waitWhileConfigured, errors.New("Unexpected status code")
	}
}
