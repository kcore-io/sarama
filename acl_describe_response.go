package sarama

import "time"

// DescribeAclsResponse is a describe acl response type
type DescribeAclsResponse struct {
	Version      int16
	ThrottleTime time.Duration
	Err          KError
	ErrMsg       *string
	ResourceAcls []*ResourceAcls
}

func (d *DescribeAclsResponse) Encode(pe packetEncoder) error {
	pe.putInt32(int32(d.ThrottleTime / time.Millisecond))
	pe.putInt16(int16(d.Err))

	if err := pe.putNullableString(d.ErrMsg); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(d.ResourceAcls)); err != nil {
		return err
	}

	for _, resourceAcl := range d.ResourceAcls {
		if err := resourceAcl.encode(pe, d.Version); err != nil {
			return err
		}
	}

	return nil
}

func (d *DescribeAclsResponse) Decode(pd packetDecoder, version int16) (err error) {
	throttleTime, err := pd.getInt32()
	if err != nil {
		return err
	}
	d.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

	kerr, err := pd.getInt16()
	if err != nil {
		return err
	}
	d.Err = KError(kerr)

	errmsg, err := pd.getString()
	if err != nil {
		return err
	}
	if errmsg != "" {
		d.ErrMsg = &errmsg
	}

	n, err := pd.getArrayLength()
	if err != nil {
		return err
	}
	d.ResourceAcls = make([]*ResourceAcls, n)

	for i := 0; i < n; i++ {
		d.ResourceAcls[i] = new(ResourceAcls)
		if err := d.ResourceAcls[i].Decode(pd, version); err != nil {
			return err
		}
	}

	return nil
}

func (d *DescribeAclsResponse) APIKey() int16 {
	return 29
}

func (d *DescribeAclsResponse) APIVersion() int16 {
	return d.Version
}

func (d *DescribeAclsResponse) HeaderVersion() int16 {
	return 0
}

func (d *DescribeAclsResponse) IsValidVersion() bool {
	return d.Version >= 0 && d.Version <= 1
}

func (d *DescribeAclsResponse) RequiredVersion() KafkaVersion {
	switch d.Version {
	case 1:
		return V2_0_0_0
	default:
		return V0_11_0_0
	}
}

func (r *DescribeAclsResponse) throttleTime() time.Duration {
	return r.ThrottleTime
}
