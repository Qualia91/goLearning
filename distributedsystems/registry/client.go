package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesUrl, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code %v", res.StatusCode)
	}
	return nil
}

func ShutdownService(serviceURL string) error {
	// create http delete request
	req, err := http.NewRequest(
		http.MethodDelete,
		ServicesUrl,
		bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return fmt.Errorf("unable to create request")
	}
	// add header
	req.Header.Add("Content-Type", "test/plain")
	// send request
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK || err != nil {
		return fmt.Errorf("failed to de-reg service. Responded with code %v", res.StatusCode)
	}
	return nil
}
