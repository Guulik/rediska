package cli

import (
	"fmt"
	"github.com/tidwall/resp"
)

func (c *CliClient) sendCommand(command string, args ...string) error {
	values := []resp.Value{resp.StringValue(command)}
	for _, arg := range args {
		values = append(values, resp.StringValue(arg))
	}

	var stringRequest string
	for _, val := range values {
		stringRequest += val.String()
	}

	req, err := resp.ArrayValue(values).MarshalRESP()
	//req, err := resp.StringValue(stringRequest).MarshalRESP()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to Marshal RESP request: "), err)
		return err
	}

	_, err = c.conn.Write(req)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to send data to server: "), err)
		return err
	}
	//TODO: Delete me
	fmt.Println("sent: ", resp.ArrayValue(values))
	return nil
}
