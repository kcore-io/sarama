package sarama

type DeleteGroupsRequest struct {
	Version int16
	Groups  []string
}

func (r *DeleteGroupsRequest) Encode(pe packetEncoder) error {
	return pe.putStringArray(r.Groups)
}

func (r *DeleteGroupsRequest) Decode(pd packetDecoder, version int16) (err error) {
	r.Groups, err = pd.getStringArray()
	return
}

func (r *DeleteGroupsRequest) APIKey() int16 {
	return 42
}

func (r *DeleteGroupsRequest) APIVersion() int16 {
	return r.Version
}

func (r *DeleteGroupsRequest) HeaderVersion() int16 {
	return 1
}

func (r *DeleteGroupsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *DeleteGroupsRequest) RequiredVersion() KafkaVersion {
	switch r.Version {
	case 1:
		return V2_0_0_0
	case 0:
		return V1_1_0_0
	default:
		return V2_0_0_0
	}
}

func (r *DeleteGroupsRequest) AddGroup(group string) {
	r.Groups = append(r.Groups, group)
}
