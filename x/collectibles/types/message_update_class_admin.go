package types

func NewMsgUpdateClassAdmin(creator string, classId string, newAdmin string) *MsgUpdateClassAdmin {
	return &MsgUpdateClassAdmin{
		Creator:  creator,
		ClassId:  classId,
		NewAdmin: newAdmin,
	}
}
