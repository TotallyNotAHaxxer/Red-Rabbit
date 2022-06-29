package pnglib

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"os"
	"strconv"
	"strings"

	color "main/modg/colors"
	rr6_opts "main/modg/copt"
	warn "main/modules/go-main/0x0001/steg-err"
	"main/modules/go-main/0x0001/utils"
)

const (
	EOFT = "IEND"
)

var (
	m  MetaChunk
	ct string
)

type Header struct {
	Magic_Byte uint64
}

type Chunk struct {
	Chunk_Size uint32
	Chunk_Type uint32
	Chunk_Data []byte
	Chunk_CRC  uint32
}

type MetaChunk struct {
	Chk          Chunk
	Chunk_Offset int64
}

//ProcessImage is the wrapper to parse PNG bytes
func (Image_metachunk *MetaChunk) Run_Functions(image_inject bool, b *bytes.Reader, c *rr6_opts.RR6_options, scan_meta bool) {
	Image_metachunk.validate(b)
	if image_inject {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Modifying Payload         | ", c.Payload)
		m.Chk.Chunk_Data = []byte(c.Payload)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Modifying Type            | ", c.Type)
		m.Chk.Chunk_Type = m.strToInt(c.Type)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Modifying Chunk Size      | ", m.Chk.Chunk_Size)
		m.Chk.Chunk_Size = m.CCS()
		m.Chk.Chunk_CRC = m.CCCRC()
		bm := m.marshalData()
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Modifying data            | ")
		bmb := bm.Bytes()
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Untouched/Unmodded payload |  % X\n", []byte(c.Payload))
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Untouched/Unmodded Byte P  |  % X\n", m.Chk.Chunk_Data)
		utils.J(b, c, bmb)
	}
	if (c.Image_offset != "") && c.Payload_Encode {
		m.Chk.Chunk_Data = utils.XorEncode([]byte(c.Payload), c.Key)
		m.Chk.Chunk_Type = m.strToInt(c.Type)
		m.Chk.Chunk_Size = m.CCS()
		m.Chk.Chunk_CRC = m.CCCRC()
		bm := m.marshalData()
		bmb := bm.Bytes()
		fmt.Printf("<RR6> Payloads and tings: Payload input --> % X\n", []byte(c.Payload))
		fmt.Printf("<RR6> Payloads and tings: Input encoded --> % X\n", m.Chk.Chunk_Data)
		utils.J(b, c, bmb)
	}
	if (c.Image_offset != "") && c.Payload_Decode {
		offset, _ := strconv.ParseInt(c.Image_offset, 10, 64)
		b.Seek(offset, 0)
		m.RC(b)
		origData := m.Chk.Chunk_Data
		m.Chk.Chunk_Data = utils.XorDecode(m.Chk.Chunk_Data, c.Key)
		m.Chk.Chunk_CRC = m.CCCRC()
		bm := m.marshalData()
		bmb := bm.Bytes()
		fmt.Printf("<RR6> Payloads and tings: Payload input --> % X\n", origData)
		fmt.Printf("<RR6> Payloads and tings: Input encoded --> % X\n", m.Chk.Chunk_Data)
		utils.J(b, c, bmb)
	}
	if scan_meta {
		byte_column := 1
		for ct != EOFT {
			Image_metachunk.Get0(b)
			Image_metachunk.RC(b)
			fmt.Printf(" \033[0;31m|>********************************************************************>|\n")
			fmt.Printf(" \033[0;31m|>>> \033[0;35mDiscovered/Found section and chunk number  \033[0;31m%s                      |\n", strconv.Itoa(byte_column))
			fmt.Printf(" \033[0;31m|>********************************************************************>|\n")
			fmt.Printf(" \033[0;35m|>>> + \033[0;35m Chunks Offset     | \033[0;31m %v\n", Image_metachunk.Chunk_Offset)
			fmt.Printf(" |>>> + \033[0;35m Chunks Offset Sp  | \033[0;31m %#02x \n", Image_metachunk.Chunk_Offset)
			fmt.Printf(" |>>> + \033[0;35m Chunks Length     | \033[0;31m %v\n", strconv.Itoa(int(Image_metachunk.Chk.Chunk_Size)))
			fmt.Printf(" |>>> + \033[0;35m Chunks Type       | \033[0;31m %v\n", Image_metachunk.CCTS())
			fmt.Printf(" |>>> + \033[0;35m Chunks Importance | \033[0;31m %v\n", Image_metachunk.checkCritType())
			fmt.Printf(" |>>> + \033[0;35m Chunk CRC         | \033[0;31m %x\n", Image_metachunk.Chk.Chunk_CRC)
			if len(Image_metachunk.Chk.Chunk_Data) >= 1000 {
				fmt.Printf(" |>>> + \033[0;35m Chunk data        | To large to output, choosing not to output\n")
			} else {
				if !c.Suppress {
					fmt.Printf(" |>>> + \033[0;35m Chunk Data        | \033[0;31m %#x\n", Image_metachunk.Chk.Chunk_Data)
				} else if c.Suppress {
					fmt.Printf("Chunk Data: %s\n", "Suppressed")
				}
			}
			fmt.Printf(" |----------------------------------------------------------------------|\n")
			ct = Image_metachunk.CCTS()
			byte_column++
		}
	}
}

func (mc *MetaChunk) marshalData() *bytes.Buffer {
	bytesMSB := new(bytes.Buffer)
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Chunk_Size); err != nil {
		warn.Return_error(err, "<RR6> Stego Module: Could not write the data ", color.REDHB)
		os.Exit(0)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Chunk_Type); err != nil {
		warn.Return_error(err, "<RR6> Stego Module: Could not write the data ", color.REDHB)
		os.Exit(0)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Chunk_Data); err != nil {
		warn.Return_error(err, "<RR6> Stego Module: Could not write the data ", color.REDHB)
		os.Exit(0)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Chunk_CRC); err != nil {
		warn.Return_error(err, "<RR6> Stego Module: Could not write the data ", color.REDHB)
		os.Exit(0)
	}
	return bytesMSB
}

func (mc *MetaChunk) RC(b *bytes.Reader) {
	mc.RCS(b)
	mc.RCT(b)
	mc.RCB(b, mc.Chk.Chunk_Size)
	mc.RCCRC(b)
}

func (mc *MetaChunk) RCS(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Chunk_Size); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got EOF??? Could not write the binary data", color.BLKHB, 229)
		fmt.Println(color.REDHB, "<RR6> Stego Module: Error -> ", err)
		fmt.Println(color.RET_RED)
		os.Exit(0)
	}
}

func (mc *MetaChunk) RCT(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Chunk_Type); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got EOF??? Could not write the binary data", color.BLKHB, 229)
		fmt.Println(color.REDHB, "<RR6> Stego Module: Error -> ", err)
		fmt.Println(color.RET_RED)
		os.Exit(0)

	}
}

func (mc *MetaChunk) RCB(b *bytes.Reader, cLen uint32) {
	mc.Chk.Chunk_Data = make([]byte, cLen)
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Chunk_Data); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got EOF?? Could not read chunk data, NOT going to attempt to retry", color.REDHB, 229)
		fmt.Println(color.RET_RED)
		os.Exit(0)
	}
}

func (mc *MetaChunk) checkCritType() string {
	fChar := string([]rune(mc.CCTS())[0])
	if fChar == strings.ToUpper(fChar) {
		return "Critical"
	}
	return "Ancillary"
}

func (mc *MetaChunk) validate(b *bytes.Reader) {
	var header Header
	if err := binary.Read(b, binary.BigEndian, &header.Magic_Byte); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got EOF?? Could not read Magic byte, NOT going to attempt to retry", color.BLKHB, 229)
		fmt.Println(color.REDHB, "<RR6> Stego Module: Error -> ", err)
		fmt.Println(color.RET_RED)
		os.Exit(0)
	}
	bArr := make([]byte, 8)
	binary.BigEndian.PutUint64(bArr, header.Magic_Byte)
	if string(bArr[1:4]) != "PNG" {
		warn.Return_Warnings("<RR6> Stego Module: Could not verify the PNG header, the values didnt match up", color.BLKHB, 299)
		fmt.Println(color.RET_RED)
		os.Exit(0)
	}
}

func (mc *MetaChunk) CCS() uint32 {
	return uint32(len(mc.Chk.Chunk_Data))
}

func (metac *MetaChunk) CCCRC() uint32 {
	bytesMSB := new(bytes.Buffer)
	if err := binary.Write(bytesMSB, binary.BigEndian, metac.Chk.Chunk_Type); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got e == 0x001??? Could not write the binary data", color.BLKHB, 229)
		fmt.Println(color.REDHB, "<RR6> Stego Module: Error -> ", err)
		fmt.Println(color.RET_RED)
		os.Exit(0)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, metac.Chk.Chunk_Data); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got e == 0x001??? Could not write the binary data", color.BLKHB, 229)
		fmt.Println(color.REDHB, "<RR6> Stego Module: Error -> ", err)
		fmt.Println(color.RET_RED)
		os.Exit(0)

	}
	return crc32.ChecksumIEEE(bytesMSB.Bytes())
}

func (mc *MetaChunk) strToInt(s string) uint32 {
	t := []byte(s)
	return binary.BigEndian.Uint32(t)
}

func (mc *MetaChunk) RCCRC(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Chunk_CRC); err != nil {
		warn.Return_Warnings("<RR6> Stego Module: Got EOF??? Could not read the binary data in the CRC chunk", color.BLKHB, 229)
		fmt.Println(color.REDHB, "<RR6> Stego Module: Error -> ", err)
		fmt.Println(color.RET_RED)
		os.Exit(0)

	}
}

func (mc *MetaChunk) Get0(b *bytes.Reader) {
	offset, _ := b.Seek(0, 1)
	mc.Chunk_Offset = offset
}

func (mc *MetaChunk) CCTS() string {
	h := fmt.Sprintf("%x", mc.Chk.Chunk_Type)
	decoded, _ := hex.DecodeString(h)
	result := string(decoded)
	return result
}
