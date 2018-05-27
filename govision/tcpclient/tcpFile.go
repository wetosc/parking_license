package tcpClient

import "io"

func (c *Client) ReadFileAsync(callback func(*Client, []byte)) {
	var data []byte
	go func() {
		for {
			buffer := make([]byte, 1024)
			n, err := c.conn.Read(buffer)
			buffer = buffer[:n]
			data = append(data, buffer...)

			if err == io.EOF {
				callback(c, data)
				checkError(err, "[TCPClient] The connection closed")
				break
			} else {
				checkError(err, "[TCPClient] Error reading data")
			}
			c.checkClosed(err)
		}
	}()
}
