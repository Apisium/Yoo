package yoo

import (
  "math"
  "errors"
  "encoding/binary"
)

type ByteArray struct {
  buf []byte
  posWrite int
  posRead int
  endian binary.ByteOrder
}

var ByteArrayEndian binary.ByteOrder = binary.BigEndian

func CreateByteArray(bytes []byte) *ByteArray {
  var ba *ByteArray
  if len(bytes) > 0 {
    ba = &ByteArray{ buf: bytes }
  } else {
    ba = &ByteArray{}
  }

  ba.endian = binary.BigEndian

  return ba
}

func (this *ByteArray) Length() int {
  return len(this.buf)
}

func (this *ByteArray) Available() int {
  return this.Length() - this.posRead
}

func (this *ByteArray) SetEndian(endian binary.ByteOrder) {
  this.endian = endian
}

func (this *ByteArray) GetEndian() binary.ByteOrder {
  if this.endian == nil {
    return ByteArrayEndian
  }
  return this.endian
}

func (this *ByteArray) SetWritePos(pos int) error {
  if pos > this.Length() {
    this.posWrite = this.Length()
    return errors.New("Buffer end")
  } else {
    this.posWrite = pos
  }
  return nil
}

func (this *ByteArray) SetWriteEnd() {
  this.SetWritePos(this.Length())
}

func (this *ByteArray) GetWritePos() int {
  return this.posWrite
}

func (this *ByteArray) SetReadPos(pos int) error {
  if pos > this.Length() {
    this.posRead = this.Length()
    return errors.New("Buffer end")
  } else {
    this.posRead = pos
  }
  return nil
}

func (this *ByteArray) SetReadEnd() {
  this.SetReadPos(this.Length())
}

func (this *ByteArray) GetReadPos() int {
  return this.posRead
}

func (this *ByteArray) Seek(pos int) error {
  err := this.SetWritePos(pos)
  this.SetReadPos(pos)

  return err
}

func (this *ByteArray) Reset() {
  this.buf = []byte{}
  this.Seek(0)
}

func (this *ByteArray) Bytes() []byte {
  return this.buf
}

func (this *ByteArray) BytesAvailable() []byte {
  return this.buf[this.posRead:]
}

func (this *ByteArray) Read(bytes []byte) (l int, err error) {
  if len(bytes) == 0 {
    return
  }
  if len(bytes) > this.Length() - this.posRead{
    return 0, errors.New("Buffer end")
  }
  l = copy(bytes, this.buf[this.posRead:])
  this.posRead += l

  return l, nil
}

func (this *ByteArray) ReadBytes(bytes []byte, length int, offset int) (l int, err error) {
  return this.Read(bytes[offset:offset + length])
}

func (this *ByteArray) ReadByte() (b byte, err error) {
  bytes := make([]byte, 1)
  _, err = this.ReadBytes(bytes, 1, 0)

  if err == nil{
    b = bytes[0]
  }

  return
}

func (this *ByteArray) ReadLength() (ret int, err error) {
  i := 0
  for _tk, b := true, 0; _tk || (err == nil && b == 255); _tk = false {
    b, err = this.ReadByte()
  }
  ret = int(math.Pow(256, i)) - 1 + b
  return
}

func (this *ByteArray) ReadInt8() (ret int8, err error) {
  err = binary.Read(this, this.endian, &ret)
  return
}

func (this *ByteArray) ReadInt16() (ret int16, err error) {
  err = binary.Read(this, this.endian, &ret)
  return
}

func (this *ByteArray) ReadInt32() (ret int32, err error) {
  err = binary.Read(this, this.endian, &ret)
  return
}

func (this *ByteArray) ReadInt64() (ret int64, err error) {
  err = binary.Read(this, this.endian, &ret)
  return
}

func (this *ByteArray) ReadFloat32() (ret float32, err error) {
  err = binary.Read(this, this.endian, &ret)
  return
}

func (this *ByteArray) ReadFloat64() (ret float64, err error) {
  err = binary.Read(this, this.endian, &ret)
  return
}

func (this *ByteArray) ReadBool() (ret bool, err error) {
  var bb byte
  bb, err = this.ReadByte()
  if err == nil{
    if bb == 1 {
      ret = true
    } else {
      ret = false
    }
  } else {
    ret = false
  }
  return
}

func (this *ByteArray) ReadString(length int) (ret string, err error) {
  bytes := make([]byte, length, length)
  _, err = this.ReadBytes(bytes, length, 0)
  if err == nil{
    ret = string(bytes)
  } else {
    ret = "";
  }
  return
}

func (this *ByteArray) ReadUTF() (ret string, err error) {
  var l int16
  l, err = this.ReadInt16()

  if err != nil{
    return "", err
  }

  return this.ReadString(int(l))
}
