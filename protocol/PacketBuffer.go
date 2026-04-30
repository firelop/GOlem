package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"

	"github.com/firelop/GOlem/protocol/datatypes"
	"github.com/google/uuid"
)

type PacketBuffer struct {
	*bytes.Buffer
}

func NewPacketBuffer(data []byte) *PacketBuffer {
	return &PacketBuffer{
		Buffer: bytes.NewBuffer(data),
	}
}

// Read methods

func (pb *PacketBuffer) ReadByte() (byte, error) {
	return pb.Buffer.ReadByte()
}

func (pb *PacketBuffer) ReadBool() (bool, error) {
	b, err := pb.ReadByte()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

func (pb *PacketBuffer) ReadVarInt() (int32, error) {
	var value int32
	var bitOffset uint

	for {
		b, err := pb.ReadByte()
		if err != nil {
			if err == io.EOF {
				return 0, io.ErrUnexpectedEOF
			}
			return 0, err
		}
		value |= int32(b&0x7F) << bitOffset

		if (b & 0x80) == 0 {
			break
		}

		bitOffset += 7
		if bitOffset >= 32 {
			return 0, errors.New("VarInt is too big")
		}
	}
	return value, nil
}

func (pb *PacketBuffer) ReadString() (string, error) {
	length, err := pb.ReadVarInt()
	if err != nil {
		return "", err
	}

	data := pb.Next(int(length))
	if len(data) < int(length) {
		return "", io.ErrUnexpectedEOF
	}

	return string(data), nil
}

func (pb *PacketBuffer) ReadUint16() (uint16, error) {
	data := pb.Next(2)
	if len(data) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.BigEndian.Uint16(data)
	return val, nil
}

func (pb *PacketBuffer) ReadInt64() (int64, error) {
	data := pb.Next(8)
	if len(data) < 8 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.BigEndian.Uint64(data)
	return int64(val), nil
}

func (pb *PacketBuffer) ReadUUID() (uuid.UUID, error) {
	data := pb.Next(16)
	if len(data) < 16 {
		return uuid.UUID{}, io.ErrUnexpectedEOF
	}

	var u uuid.UUID
	copy(u[:], data)
	return u, nil
}

func (pb *PacketBuffer) ReadByteArray() ([]byte, error) {
	length, err := pb.ReadVarInt()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("negative length for byte array")
	}

	data := pb.Next(int(length))
	if len(data) < int(length) {
		return nil, io.ErrUnexpectedEOF
	}

	res := make([]byte, length)
	copy(res, data)
	return res, nil
}

func (pb *PacketBuffer) ReadFloat() (float32, error) {
	data := pb.Next(4)
	if len(data) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	return math.Float32frombits(binary.BigEndian.Uint32(data)), nil
}

func (pb *PacketBuffer) ReadDouble() (float64, error) {
	data := pb.Next(8)
	if len(data) < 8 {
		return 0, io.ErrUnexpectedEOF
	}
	return math.Float64frombits(binary.BigEndian.Uint64(data)), nil
}

func (pb *PacketBuffer) ReadInt() (int32, error) {
	data := pb.Next(4)
	if len(data) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	return int32(binary.BigEndian.Uint32(data)), nil
}

// Write methods

func (pb *PacketBuffer) WriteVarInt(value int32) {
	ux := uint32(value)
	for ux >= 0x80 {
		pb.WriteByte(byte(ux) | 0x80)
		ux >>= 7
	}
	pb.WriteByte(byte(ux))
}

func (pb *PacketBuffer) WriteByte(value byte) error {
	return pb.Buffer.WriteByte(value)
}

func (pb *PacketBuffer) WriteBool(value bool) {
	if value {
		pb.WriteByte(0x01)
	} else {
		pb.WriteByte(0x00)
	}
}

func (pb *PacketBuffer) WriteString(value string) {
	pb.WriteVarInt(int32(len(value)))
	pb.Buffer.WriteString(value)
}

func (pb *PacketBuffer) WriteInt32(value int32) {
	var temp [4]byte
	binary.BigEndian.PutUint32(temp[:], uint32(value))
	pb.Write(temp[:])
}

func (pb *PacketBuffer) WriteInt64(value int64) {
	var temp [8]byte
	binary.BigEndian.PutUint64(temp[:], uint64(value))
	pb.Write(temp[:])
}

func (pb *PacketBuffer) WriteUUID(value uuid.UUID) {
	pb.Write(value[:])
}

func (pb *PacketBuffer) WriteByteArray(value []byte) {
	pb.WriteVarInt(int32(len(value)))
	pb.Write(value)
}

func (pb *PacketBuffer) WritePosition(pos datatypes.Position) {
	encode_pos := (int64(pos.X&0x3FFFFFF) << 38) | (int64(pos.Z&0x3FFFFFF) << 12) | int64(pos.Y&0xFFF)
	pb.WriteInt64(encode_pos)
}

func (pb *PacketBuffer) WriteFloat(value float32) {
	var temp [4]byte
	binary.BigEndian.PutUint32(temp[:], math.Float32bits(value))
	pb.Write(temp[:])
}

func (pb *PacketBuffer) WriteDouble(value float64) {
	var temp [8]byte
	binary.BigEndian.PutUint64(temp[:], math.Float64bits(value))
	pb.Write(temp[:])
}

func (pb *PacketBuffer) WriteInt(value int32) {
	var temp [4]byte
	binary.BigEndian.PutUint32(temp[:], uint32(value))
	pb.Write(temp[:])
}

// Generic functions

func WritePrefixedArray[T any](pb *PacketBuffer, slice []T, encoder func(*PacketBuffer, T)) {
	pb.WriteVarInt(int32(len(slice)))
	for _, item := range slice {
		encoder(pb, item)
	}
}

func ReadPrefixedArray[T any](pb *PacketBuffer, decoder func(*PacketBuffer) (T, error)) ([]T, error) {
	length, err := pb.ReadVarInt()
	if err != nil {
		return nil, err
	}

	// On pré-alloue la slice avec la bonne taille pour éviter les appends successifs
	result := make([]T, length)
	for i := 0; i < int(length); i++ {
		val, err := decoder(pb)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}
