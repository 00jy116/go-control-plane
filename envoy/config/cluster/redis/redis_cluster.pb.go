// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/cluster/redis/redis_cluster.proto

package v2

import (
	fmt "fmt"
	io "io"
	math "math"
	time "time"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RedisClusterConfig struct {
	// Interval between successive topology refresh requests. If not set, this defaults to 5s.
	ClusterRefreshRate *time.Duration `protobuf:"bytes,1,opt,name=cluster_refresh_rate,json=clusterRefreshRate,proto3,stdduration" json:"cluster_refresh_rate,omitempty"`
	// Timeout for topology refresh request. If not set, this defaults to 3s.
	ClusterRefreshTimeout *time.Duration `protobuf:"bytes,2,opt,name=cluster_refresh_timeout,json=clusterRefreshTimeout,proto3,stdduration" json:"cluster_refresh_timeout,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *RedisClusterConfig) Reset()         { *m = RedisClusterConfig{} }
func (m *RedisClusterConfig) String() string { return proto.CompactTextString(m) }
func (*RedisClusterConfig) ProtoMessage()    {}
func (*RedisClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6593a6ec218c02, []int{0}
}
func (m *RedisClusterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedisClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedisClusterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedisClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisClusterConfig.Merge(m, src)
}
func (m *RedisClusterConfig) XXX_Size() int {
	return m.Size()
}
func (m *RedisClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RedisClusterConfig proto.InternalMessageInfo

func (m *RedisClusterConfig) GetClusterRefreshRate() *time.Duration {
	if m != nil {
		return m.ClusterRefreshRate
	}
	return nil
}

func (m *RedisClusterConfig) GetClusterRefreshTimeout() *time.Duration {
	if m != nil {
		return m.ClusterRefreshTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RedisClusterConfig)(nil), "envoy.config.cluster.redis.RedisClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/redis/redis_cluster.proto", fileDescriptor_6d6593a6ec218c02)
}

var fileDescriptor_6d6593a6ec218c02 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0x86, 0x90, 0xf1, 0x50, 0x31, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0x29, 0xb0, 0x7a, 0x3d, 0x88, 0x7a, 0x3d, 0x98, 0x1c, 0x58, 0xa5, 0x94, 0x5c, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x65, 0x52, 0x69, 0x9a, 0x7e, 0x4a, 0x69, 0x51, 0x62,
	0x49, 0x66, 0x7e, 0x1e, 0x44, 0xaf, 0x94, 0x78, 0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0xaa,
	0x3e, 0x8c, 0x01, 0x95, 0x10, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8,
	0xd2, 0x15, 0x46, 0x2e, 0xa1, 0x20, 0x90, 0xc1, 0xce, 0x10, 0x5b, 0x9c, 0xc1, 0x76, 0x0a, 0x45,
	0x73, 0x89, 0x40, 0xad, 0x8d, 0x2f, 0x4a, 0x4d, 0x2b, 0x4a, 0x2d, 0xce, 0x88, 0x2f, 0x4a, 0x2c,
	0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd4, 0x83, 0x38, 0x42, 0x0f, 0xe6, 0x08,
	0x3d, 0x17, 0xa8, 0x23, 0x9c, 0xf8, 0x66, 0xdc, 0x97, 0x67, 0xdc, 0xf5, 0xf2, 0x00, 0x33, 0xeb,
	0x2a, 0x46, 0x26, 0x2d, 0x86, 0x20, 0x21, 0xa8, 0x31, 0x41, 0x10, 0x53, 0x82, 0x12, 0x4b, 0x52,
	0x85, 0x12, 0xb9, 0xc4, 0xd1, 0x0d, 0x2f, 0xc9, 0xcc, 0x4d, 0xcd, 0x2f, 0x2d, 0x91, 0x60, 0x22,
	0xd5, 0x7c, 0x51, 0x54, 0xf3, 0x43, 0x20, 0xe6, 0x38, 0xf9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x5c, 0x1a, 0x99, 0xf9, 0x90, 0x28, 0x28, 0x28, 0xca,
	0xaf, 0xa8, 0xd4, 0xc3, 0x1d, 0xba, 0x4e, 0x82, 0xc8, 0x61, 0x11, 0x00, 0xb2, 0x3d, 0x80, 0x31,
	0x8a, 0xa9, 0xcc, 0x28, 0x89, 0x0d, 0xec, 0x14, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x57, 0xeb, 0xe7, 0xcb, 0x01, 0x00, 0x00,
}

func (m *RedisClusterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedisClusterConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClusterRefreshRate != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRedisCluster(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ClusterRefreshRate)))
		n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.ClusterRefreshRate, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ClusterRefreshTimeout != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRedisCluster(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ClusterRefreshTimeout)))
		n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.ClusterRefreshTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRedisCluster(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RedisClusterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterRefreshRate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ClusterRefreshRate)
		n += 1 + l + sovRedisCluster(uint64(l))
	}
	if m.ClusterRefreshTimeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ClusterRefreshTimeout)
		n += 1 + l + sovRedisCluster(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRedisCluster(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRedisCluster(x uint64) (n int) {
	return sovRedisCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RedisClusterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedisCluster
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
			return fmt.Errorf("proto: RedisClusterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedisClusterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterRefreshRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRedisCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterRefreshRate == nil {
				m.ClusterRefreshRate = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.ClusterRefreshRate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterRefreshTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRedisCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterRefreshTimeout == nil {
				m.ClusterRefreshTimeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.ClusterRefreshTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRedisCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRedisCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRedisCluster
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
					return 0, ErrIntOverflowRedisCluster
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
					return 0, ErrIntOverflowRedisCluster
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
				return 0, ErrInvalidLengthRedisCluster
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRedisCluster
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRedisCluster
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
				next, err := skipRedisCluster(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRedisCluster
				}
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
	ErrInvalidLengthRedisCluster = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRedisCluster   = fmt.Errorf("proto: integer overflow")
)
