package webull

// import "io"

// // Pager for paginating data
// type Pager struct {
// 	Next, Previous string
// }

// // HasMore for determining when end of pages are reached
// func (p Pager) HasMore() bool {
// 	return p.Next != ""
// }

// // GetNext will decode the next set of pages if they exist
// func (p *Pager) GetNext(c *Client, out interface{}) error {
// 	if p.Next == "" {
// 		return io.EOF
// 	}

// 	return c.GetAndDecode(p.Next, out)
// }
