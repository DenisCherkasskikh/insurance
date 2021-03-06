// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: insurance/crash_record.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CrashRecord struct {
	Index        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Brand        string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model        string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Owner        string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	LicencePlate string `protobuf:"bytes,5,opt,name=licencePlate,proto3" json:"licencePlate,omitempty"`
	VinNumber    string `protobuf:"bytes,6,opt,name=vinNumber,proto3" json:"vinNumber,omitempty"`
	Mileage      string `protobuf:"bytes,7,opt,name=mileage,proto3" json:"mileage,omitempty"`
	Side         string `protobuf:"bytes,8,opt,name=side,proto3" json:"side,omitempty"`
	Part         string `protobuf:"bytes,9,opt,name=part,proto3" json:"part,omitempty"`
	Payout       string `protobuf:"bytes,10,opt,name=payout,proto3" json:"payout,omitempty"`
}

func (m *CrashRecord) Reset()         { *m = CrashRecord{} }
func (m *CrashRecord) String() string { return proto.CompactTextString(m) }
func (*CrashRecord) ProtoMessage()    {}
func (*CrashRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_88e8e3910adb0d59, []int{0}
}
func (m *CrashRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrashRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrashRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrashRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrashRecord.Merge(m, src)
}
func (m *CrashRecord) XXX_Size() int {
	return m.Size()
}
func (m *CrashRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_CrashRecord.DiscardUnknown(m)
}

var xxx_messageInfo_CrashRecord proto.InternalMessageInfo

func (m *CrashRecord) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *CrashRecord) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CrashRecord) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *CrashRecord) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CrashRecord) GetLicencePlate() string {
	if m != nil {
		return m.LicencePlate
	}
	return ""
}

func (m *CrashRecord) GetVinNumber() string {
	if m != nil {
		return m.VinNumber
	}
	return ""
}

func (m *CrashRecord) GetMileage() string {
	if m != nil {
		return m.Mileage
	}
	return ""
}

func (m *CrashRecord) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *CrashRecord) GetPart() string {
	if m != nil {
		return m.Part
	}
	return ""
}

func (m *CrashRecord) GetPayout() string {
	if m != nil {
		return m.Payout
	}
	return ""
}

func init() {
	proto.RegisterType((*CrashRecord)(nil), "denischerkasskikh.insurance.insurance.CrashRecord")
}

func init() { proto.RegisterFile("insurance/crash_record.proto", fileDescriptor_88e8e3910adb0d59) }

var fileDescriptor_88e8e3910adb0d59 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd1, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x06, 0xf0, 0xa6, 0xb7, 0x7f, 0x6e, 0x0c, 0x93, 0x85, 0x90, 0x87, 0xca, 0x42, 0x95, 0x90,
	0x98, 0x92, 0x81, 0x85, 0x99, 0x32, 0x23, 0x94, 0x91, 0x05, 0x39, 0xf6, 0x51, 0x63, 0x25, 0xb1,
	0x23, 0xdb, 0x81, 0x76, 0xe4, 0x0d, 0x78, 0x2c, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x11, 0x64, 0x07,
	0x48, 0xd9, 0xbe, 0xef, 0xe7, 0x73, 0x96, 0x63, 0xb4, 0x92, 0xca, 0xb6, 0x86, 0x29, 0x0e, 0x29,
	0x37, 0xcc, 0x16, 0x4f, 0x06, 0xb8, 0x36, 0x22, 0x69, 0x8c, 0x76, 0x1a, 0x5f, 0x0a, 0x50, 0xd2,
	0xf2, 0x02, 0x4c, 0xc9, 0xac, 0x2d, 0x65, 0x59, 0x24, 0xbf, 0xf3, 0x63, 0x5a, 0xbf, 0x4e, 0xd1,
	0xc9, 0xc6, 0x6f, 0x67, 0x61, 0x19, 0x9f, 0xa1, 0xb9, 0x54, 0x02, 0x76, 0x24, 0xba, 0x88, 0xae,
	0xe2, 0x6c, 0x28, 0x5e, 0x73, 0xc3, 0x94, 0x20, 0xd3, 0x41, 0x43, 0xf1, 0x5a, 0x6b, 0x01, 0x15,
	0xf9, 0x37, 0x68, 0x28, 0x5e, 0xf5, 0x8b, 0x02, 0x43, 0x66, 0x83, 0x86, 0x82, 0xd7, 0xe8, 0xb4,
	0x92, 0x1c, 0x14, 0x87, 0x87, 0x8a, 0x39, 0x20, 0xf3, 0xf0, 0xf8, 0xc7, 0xf0, 0x0a, 0xc5, 0xcf,
	0x52, 0xdd, 0xb7, 0x75, 0x0e, 0x86, 0x2c, 0xc2, 0xc0, 0x08, 0x98, 0xa0, 0x65, 0x2d, 0x2b, 0x60,
	0x5b, 0x20, 0xcb, 0xf0, 0xf6, 0x53, 0x31, 0x46, 0x33, 0x2b, 0x05, 0x90, 0xff, 0x81, 0x43, 0xf6,
	0xd6, 0x30, 0xe3, 0x48, 0x3c, 0x98, 0xcf, 0xf8, 0x1c, 0x2d, 0x1a, 0xb6, 0xd7, 0xad, 0x23, 0x28,
	0xe8, 0x77, 0xbb, 0xcd, 0xde, 0x3b, 0x1a, 0x1d, 0x3a, 0x1a, 0x7d, 0x76, 0x34, 0x7a, 0xeb, 0xe9,
	0xe4, 0xd0, 0xd3, 0xc9, 0x47, 0x4f, 0x27, 0x8f, 0x37, 0x5b, 0xe9, 0x8a, 0x36, 0x4f, 0xb8, 0xae,
	0xd3, 0x3b, 0x7f, 0xcf, 0xcd, 0xd1, 0x3d, 0xd3, 0xf1, 0xfe, 0xbb, 0xa3, 0xec, 0xf6, 0x0d, 0xd8,
	0x7c, 0x11, 0x7e, 0xe1, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x79, 0x7a, 0x25, 0xa5, 0x01,
	0x00, 0x00,
}

func (m *CrashRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrashRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrashRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payout) > 0 {
		i -= len(m.Payout)
		copy(dAtA[i:], m.Payout)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Payout)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Part) > 0 {
		i -= len(m.Part)
		copy(dAtA[i:], m.Part)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Part)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Side) > 0 {
		i -= len(m.Side)
		copy(dAtA[i:], m.Side)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Side)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Mileage) > 0 {
		i -= len(m.Mileage)
		copy(dAtA[i:], m.Mileage)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Mileage)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.VinNumber) > 0 {
		i -= len(m.VinNumber)
		copy(dAtA[i:], m.VinNumber)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.VinNumber)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LicencePlate) > 0 {
		i -= len(m.LicencePlate)
		copy(dAtA[i:], m.LicencePlate)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.LicencePlate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Model) > 0 {
		i -= len(m.Model)
		copy(dAtA[i:], m.Model)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Model)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Brand) > 0 {
		i -= len(m.Brand)
		copy(dAtA[i:], m.Brand)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Brand)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCrashRecord(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCrashRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovCrashRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CrashRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Brand)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Model)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.LicencePlate)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.VinNumber)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Mileage)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Side)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Part)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	l = len(m.Payout)
	if l > 0 {
		n += 1 + l + sovCrashRecord(uint64(l))
	}
	return n
}

func sovCrashRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrashRecord(x uint64) (n int) {
	return sovCrashRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CrashRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrashRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CrashRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrashRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Brand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Brand = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Model = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LicencePlate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LicencePlate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VinNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VinNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mileage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mileage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Side", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Side = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Part", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Part = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrashRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrashRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrashRecord
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCrashRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrashRecord
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCrashRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCrashRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCrashRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCrashRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCrashRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrashRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCrashRecord = fmt.Errorf("proto: unexpected end of group")
)
