// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/services/events/v1/content.proto
// DO NOT EDIT!

package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/containerd/containerd/protobuf/plugin"

import github_com_opencontainers_go_digest "github.com/opencontainers/go-digest"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContentDelete struct {
	Digest github_com_opencontainers_go_digest.Digest `protobuf:"bytes,1,opt,name=digest,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"digest"`
}

func (m *ContentDelete) Reset()                    { *m = ContentDelete{} }
func (*ContentDelete) ProtoMessage()               {}
func (*ContentDelete) Descriptor() ([]byte, []int) { return fileDescriptorContent, []int{0} }

func init() {
	proto.RegisterType((*ContentDelete)(nil), "containerd.services.events.v1.ContentDelete")
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContentDelete) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	case "digest":
		return string(m.Digest), len(m.Digest) > 0
	}
	return "", false
}
func (m *ContentDelete) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentDelete) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Digest) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContent(dAtA, i, uint64(len(m.Digest)))
		i += copy(dAtA[i:], m.Digest)
	}
	return i, nil
}

func encodeFixed64Content(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Content(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintContent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ContentDelete) Size() (n int) {
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	return n
}

func sovContent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContent(x uint64) (n int) {
	return sovContent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ContentDelete) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContentDelete{`,
		`Digest:` + fmt.Sprintf("%v", this.Digest) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ContentDelete) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContentDelete: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentDelete: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContent
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
func skipContent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContent
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
					return 0, ErrIntOverflowContent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthContent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipContent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthContent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContent   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/events/v1/content.proto", fileDescriptorContent)
}

var fileDescriptorContent = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x26, 0x16, 0x64, 0xea, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x16, 0xeb,
	0xa7, 0x96, 0xa5, 0xe6, 0x95, 0x14, 0xeb, 0x97, 0x19, 0x82, 0x55, 0xa4, 0xe6, 0x95, 0xe8, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0x22, 0x34, 0xe8, 0xc1, 0x14, 0xeb, 0x41, 0x14, 0xeb, 0x95,
	0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x55, 0xea, 0x83, 0x58, 0x10, 0x4d, 0x52, 0x0e,
	0x04, 0xed, 0x06, 0xab, 0x4b, 0x2a, 0x4d, 0xd3, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0xd3, 0x4f,
	0xcb, 0x4c, 0xcd, 0x49, 0x29, 0x48, 0x2c, 0xc9, 0x80, 0x98, 0xa0, 0x14, 0xcd, 0xc5, 0xeb, 0x0c,
	0x71, 0x87, 0x4b, 0x6a, 0x4e, 0x6a, 0x49, 0xaa, 0x90, 0x17, 0x17, 0x5b, 0x4a, 0x66, 0x7a, 0x6a,
	0x71, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xd1, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee,
	0xc9, 0x6b, 0x21, 0x59, 0x95, 0x5f, 0x90, 0x9a, 0x07, 0xb7, 0xa3, 0x58, 0x3f, 0x3d, 0x5f, 0x17,
	0xa2, 0x45, 0xcf, 0x05, 0x4c, 0x05, 0x41, 0x4d, 0x70, 0x8a, 0x39, 0xf1, 0x50, 0x8e, 0xe1, 0xc6,
	0x43, 0x39, 0x86, 0x86, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x82, 0x2f, 0x72, 0x8c, 0x51, 0x76, 0x64, 0x06, 0x9c, 0x35, 0x84, 0x95,
	0xc4, 0x06, 0xf6, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x51, 0xce, 0xec, 0x89, 0x81, 0x01,
	0x00, 0x00,
}
