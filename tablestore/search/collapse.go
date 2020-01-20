package search

import "github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore/otsprotocol"

type Collapse struct {
	FieldName string
}

func (c *Collapse) ProtoBuffer() (*otsprotocol.Collapse, error) {
	pb := &otsprotocol.Collapse{
		FieldName: &c.FieldName,
	}
	return pb, nil
}
