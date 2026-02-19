package types

func NewMsgCreateClass(creator string, name string, symbol string, uri string, admin string) *MsgCreateClass {
	return &MsgCreateClass{
		Creator: creator,
		Name:    name,
		Symbol:  symbol,
		Uri:     uri,
		Admin:   admin,
	}
}
