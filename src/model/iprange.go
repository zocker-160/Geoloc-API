package model

type IPRange struct {
	StartIP, EndIP uint32
}

func (c *IPRange) IsInRange(ip uint32) bool {
	return c.EndIP >= ip && ip >= c.StartIP
}