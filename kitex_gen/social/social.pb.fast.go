// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package social

import (
	core "douyin/kitex_gen/core"
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DouyinRelationActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationActionRequest[number], err)
}

func (x *DouyinRelationActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinRelationActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationActionResponse[number], err)
}

func (x *DouyinRelationActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinRelationActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinRelationFollowListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationFollowListRequest[number], err)
}

func (x *DouyinRelationFollowListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.QueryUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationFollowListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationFollowListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationFollowListResponse[number], err)
}

func (x *DouyinRelationFollowListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinRelationFollowListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinRelationFollowListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v core.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *DouyinRelationFollowerListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationFollowerListRequest[number], err)
}

func (x *DouyinRelationFollowerListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.QueryUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationFollowerListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationFollowerListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationFollowerListResponse[number], err)
}

func (x *DouyinRelationFollowerListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinRelationFollowerListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinRelationFollowerListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v core.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *DouyinRelationFriendListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationFriendListRequest[number], err)
}

func (x *DouyinRelationFriendListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationFriendListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinRelationFriendListResponse[number], err)
}

func (x *DouyinRelationFriendListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinRelationFriendListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinRelationFriendListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v FriendUser
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *DouyinMessageActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinMessageActionRequest[number], err)
}

func (x *DouyinMessageActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinMessageActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinMessageActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinMessageActionRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinMessageActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinMessageActionResponse[number], err)
}

func (x *DouyinMessageActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinMessageActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinMessageChatRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinMessageChatRequest[number], err)
}

func (x *DouyinMessageChatRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinMessageChatRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinMessageChatResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinMessageChatResponse[number], err)
}

func (x *DouyinMessageChatResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinMessageChatResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinMessageChatResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Message
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.MessageList = append(x.MessageList, &v)
	return offset, nil
}

func (x *Message) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Message[number], err)
}

func (x *Message) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FromUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CreateTime = &tmp
	return offset, err
}

func (x *FriendUser) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FriendUser[number], err)
}

func (x *FriendUser) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v core.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *FriendUser) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Message = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.MsgType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinRelationActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinRelationActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *DouyinRelationActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ToUserId)
	return offset
}

func (x *DouyinRelationActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.ActionType)
	return offset
}

func (x *DouyinRelationActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinRelationActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinRelationActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.StatusMsg)
	return offset
}

func (x *DouyinRelationFollowListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinRelationFollowListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.QueryUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.QueryUserId)
	return offset
}

func (x *DouyinRelationFollowListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *DouyinRelationFollowListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinRelationFollowListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinRelationFollowListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.StatusMsg)
	return offset
}

func (x *DouyinRelationFollowListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.UserList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.UserList[i])
	}
	return offset
}

func (x *DouyinRelationFollowerListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinRelationFollowerListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.QueryUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.QueryUserId)
	return offset
}

func (x *DouyinRelationFollowerListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *DouyinRelationFollowerListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinRelationFollowerListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinRelationFollowerListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.StatusMsg)
	return offset
}

func (x *DouyinRelationFollowerListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.UserList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.UserList[i])
	}
	return offset
}

func (x *DouyinRelationFriendListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinRelationFriendListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *DouyinRelationFriendListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinRelationFriendListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinRelationFriendListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.StatusMsg)
	return offset
}

func (x *DouyinRelationFriendListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.UserList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.UserList[i])
	}
	return offset
}

func (x *DouyinMessageActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *DouyinMessageActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *DouyinMessageActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ToUserId)
	return offset
}

func (x *DouyinMessageActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.ActionType)
	return offset
}

func (x *DouyinMessageActionRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.Content)
	return offset
}

func (x *DouyinMessageActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinMessageActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinMessageActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.StatusMsg)
	return offset
}

func (x *DouyinMessageChatRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinMessageChatRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *DouyinMessageChatRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ToUserId)
	return offset
}

func (x *DouyinMessageChatResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinMessageChatResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinMessageChatResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.StatusMsg)
	return offset
}

func (x *DouyinMessageChatResponse) fastWriteField3(buf []byte) (offset int) {
	if x.MessageList == nil {
		return offset
	}
	for i := range x.MessageList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.MessageList[i])
	}
	return offset
}

func (x *Message) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *Message) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Message) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ToUserId)
	return offset
}

func (x *Message) fastWriteField3(buf []byte) (offset int) {
	if x.FromUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.FromUserId)
	return offset
}

func (x *Message) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.Content)
	return offset
}

func (x *Message) fastWriteField5(buf []byte) (offset int) {
	if x.CreateTime == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, *x.CreateTime)
	return offset
}

func (x *FriendUser) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FriendUser) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.User)
	return offset
}

func (x *FriendUser) fastWriteField2(buf []byte) (offset int) {
	if x.Message == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.Message)
	return offset
}

func (x *FriendUser) fastWriteField3(buf []byte) (offset int) {
	if x.MsgType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.MsgType)
	return offset
}

func (x *DouyinRelationActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinRelationActionRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *DouyinRelationActionRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ToUserId)
	return n
}

func (x *DouyinRelationActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.ActionType)
	return n
}

func (x *DouyinRelationActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinRelationActionResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinRelationActionResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.StatusMsg)
	return n
}

func (x *DouyinRelationFollowListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinRelationFollowListRequest) sizeField1() (n int) {
	if x.QueryUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.QueryUserId)
	return n
}

func (x *DouyinRelationFollowListRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *DouyinRelationFollowListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinRelationFollowListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinRelationFollowListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.StatusMsg)
	return n
}

func (x *DouyinRelationFollowListResponse) sizeField3() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.UserList {
		n += fastpb.SizeMessage(3, x.UserList[i])
	}
	return n
}

func (x *DouyinRelationFollowerListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinRelationFollowerListRequest) sizeField1() (n int) {
	if x.QueryUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.QueryUserId)
	return n
}

func (x *DouyinRelationFollowerListRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *DouyinRelationFollowerListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinRelationFollowerListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinRelationFollowerListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.StatusMsg)
	return n
}

func (x *DouyinRelationFollowerListResponse) sizeField3() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.UserList {
		n += fastpb.SizeMessage(3, x.UserList[i])
	}
	return n
}

func (x *DouyinRelationFriendListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinRelationFriendListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *DouyinRelationFriendListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinRelationFriendListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinRelationFriendListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.StatusMsg)
	return n
}

func (x *DouyinRelationFriendListResponse) sizeField3() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.UserList {
		n += fastpb.SizeMessage(3, x.UserList[i])
	}
	return n
}

func (x *DouyinMessageActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *DouyinMessageActionRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *DouyinMessageActionRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ToUserId)
	return n
}

func (x *DouyinMessageActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.ActionType)
	return n
}

func (x *DouyinMessageActionRequest) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.Content)
	return n
}

func (x *DouyinMessageActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinMessageActionResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinMessageActionResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.StatusMsg)
	return n
}

func (x *DouyinMessageChatRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinMessageChatRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *DouyinMessageChatRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ToUserId)
	return n
}

func (x *DouyinMessageChatResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinMessageChatResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinMessageChatResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.StatusMsg)
	return n
}

func (x *DouyinMessageChatResponse) sizeField3() (n int) {
	if x.MessageList == nil {
		return n
	}
	for i := range x.MessageList {
		n += fastpb.SizeMessage(3, x.MessageList[i])
	}
	return n
}

func (x *Message) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *Message) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Message) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ToUserId)
	return n
}

func (x *Message) sizeField3() (n int) {
	if x.FromUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.FromUserId)
	return n
}

func (x *Message) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.Content)
	return n
}

func (x *Message) sizeField5() (n int) {
	if x.CreateTime == nil {
		return n
	}
	n += fastpb.SizeString(5, *x.CreateTime)
	return n
}

func (x *FriendUser) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FriendUser) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.User)
	return n
}

func (x *FriendUser) sizeField2() (n int) {
	if x.Message == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.Message)
	return n
}

func (x *FriendUser) sizeField3() (n int) {
	if x.MsgType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.MsgType)
	return n
}

var fieldIDToName_DouyinRelationActionRequest = map[int32]string{
	1: "UserId",
	2: "ToUserId",
	3: "ActionType",
}

var fieldIDToName_DouyinRelationActionResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_DouyinRelationFollowListRequest = map[int32]string{
	1: "QueryUserId",
	2: "UserId",
}

var fieldIDToName_DouyinRelationFollowListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserList",
}

var fieldIDToName_DouyinRelationFollowerListRequest = map[int32]string{
	1: "QueryUserId",
	2: "UserId",
}

var fieldIDToName_DouyinRelationFollowerListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserList",
}

var fieldIDToName_DouyinRelationFriendListRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_DouyinRelationFriendListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserList",
}

var fieldIDToName_DouyinMessageActionRequest = map[int32]string{
	1: "UserId",
	2: "ToUserId",
	3: "ActionType",
	4: "Content",
}

var fieldIDToName_DouyinMessageActionResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_DouyinMessageChatRequest = map[int32]string{
	1: "UserId",
	2: "ToUserId",
}

var fieldIDToName_DouyinMessageChatResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "MessageList",
}

var fieldIDToName_Message = map[int32]string{
	1: "Id",
	2: "ToUserId",
	3: "FromUserId",
	4: "Content",
	5: "CreateTime",
}

var fieldIDToName_FriendUser = map[int32]string{
	1: "User",
	2: "Message",
	3: "MsgType",
}

