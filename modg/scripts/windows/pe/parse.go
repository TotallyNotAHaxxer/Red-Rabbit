package parse

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	constants "main/modg/constants"
	system "main/modg/system-runscript"
)

func errors(err error, msg string, exit int) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(exit)
	}
}

func Parser(filename string) (string, error) {
	system.File_inf(filename)
	o, e := os.Open(filename)
	errors(e, "<RR6> Windows Module: Could not open filename", 1)
	pf, e := pe.NewFile(o)
	errors(e, "<RR6> Windows Module: Could not make a new PE file", 1)
	defer o.Close()
	defer pf.Close() // close under defer, without closing this will create a vulnerability
	dh := make([]byte, 96)
	so := make([]byte, 4)
	_, e = o.Read(dh)
	errors(e, " <RR6> Windows Module: Could not inspect or read DOS header", 1)
	peoffset := int64(binary.LittleEndian.Uint32(dh[0x3c:]))
	o.ReadAt(so[:], peoffset)
	sr := io.NewSectionReader(o, 0, 1<<63-1)
	_, e = sr.Seek(peoffset+4, io.SeekStart)
	errors(e, "<RR6> Windows Module: Could not seek offset", 1)
	binary.Read(sr, binary.LittleEndian, &pf.FileHeader)
	switch pf.FileHeader.SizeOfOptionalHeader {
	case constants.SizeofOptionalHeader32:
		binary.Read(sr, binary.LittleEndian, &constants.X32)
	case constants.SizeofOptionalHeader64:
		binary.Read(sr, binary.LittleEndian, &constants.X64)
	}
	fmt.Printf(" |>**********************************************************************<|\n")
	fmt.Printf(" |>>> Discovered section BASIC SECTION START INFO                         |\n")
	fmt.Printf(" |>**********************************************************************<|\n")
	fmt.Printf(" |>>> +  \033[32mMagic Byte       \033[31m| %s%s\n", string(dh[0]), string(dh[1]))
	fmt.Printf(" |>>> +  \033[32mLFANEW Value     \033[31m|%s\n", string(so))
	fmt.Printf(" |>**********************************************************************<|\n")
	fmt.Printf(" |>>> Discovered section COFF File Header                                 |\n")
	fmt.Printf(" |>**********************************************************************<|\n")
	fmt.Printf(" |>>> +  \033[32mMachine Architecture        \033[31m|%#x\n", pf.FileHeader.Machine)
	fmt.Printf(" |>>> +  \033[32mNumber of Sections          \033[31m|%#x\n", pf.FileHeader.NumberOfSections)
	fmt.Printf(" |>>> +  \033[32mSize of Optional Header     \033[31m|%#x\n", pf.FileHeader.SizeOfOptionalHeader)
	fmt.Printf(" |************************************************************************|")
	fmt.Print("\n\n")
	fmt.Printf(" |>**********************************************************************<|\n")
	fmt.Printf(" |>>> Discovered section offsets                                          |\n")
	fmt.Printf(" |>**********************************************************************<|\n")
	fmt.Printf(" |>>> + \033[32mNumber of Sections Field Offset   \033[31m|%#x\n", peoffset+6)
	fmt.Printf(" |>>> + \033[32mSection Table Offset              \033[31m|%#x\n", peoffset+0xF8)
	fmt.Printf(" |*************************************************************************|\n")
	fmt.Print("\n")
	fmt.Printf(" |---------------------------------------------------------\n")
	fmt.Printf(" |>>> Discovered Section - \033[32mOptional Headers\033[31m\n")
	fmt.Printf(" |********************************************************************************************\n")
	fmt.Printf(" |>>> + \033[32mEntry Point                \033[31m| %#x\n", constants.X32.AddressOfEntryPoint)
	fmt.Printf(" |>>> + \033[32mImage Base                 \033[31m| %#x\n", constants.X32.ImageBase)
	fmt.Printf(" |>>> + \033[32mSize of image              \033[31m| %#x\n", constants.X32.SizeOfImage)
	fmt.Printf(" |>>> + \033[32mSection alignment          \033[31m| %#x\n", constants.X32.SectionAlignment)
	fmt.Printf(" |>>> + \033[32mFile attachment            \033[31m| %#x\n", constants.X32.FileAlignment)
	fmt.Printf(" |>>> + \033[32mFile Characteristics       \033[31m| %#x\n", pf.FileHeader.Characteristics)
	fmt.Printf(" |>>> + \033[32mSize of headers            \033[31m| %#x\n", constants.X32.SizeOfHeaders)
	fmt.Printf(" |>>> + \033[32mChecksum                   \033[31m| %#x\n", constants.X32.CheckSum)
	fmt.Printf(" |>>> + \033[32mMachine                    \033[31m| %#x\n", pf.FileHeader.Machine)
	fmt.Printf(" |>>> + \033[32mSubsystem                  \033[31m| %#x\n", constants.X32.Subsystem)
	fmt.Printf(" |>>> + \033[32mDLL Characteristics        \033[31m| %#x\n", constants.X32.DllCharacteristics)
	fmt.Printf(" |********************************************************************************************\n")
	for ix, d := range constants.X32.DataDirectory {
		fmt.Printf(" |---------------------------------------------------------\n")
		fmt.Printf(" |>>> Discovered Section - \033[32m%s\033[31m\n", constants.Windows_data_directories[ix])
		fmt.Printf(" |*******************************************************************************************\n")
		fmt.Printf(" |>>> + \033[32mImage Virtual Address \033[31m| %#x\n", d.VirtualAddress)
		fmt.Printf(" |>>> + \033[32mImage Size            \033[31m| %#x\n", d.Size)
		fmt.Printf(" |*******************************************************************************************\n")
	}
	for _, sn := range pf.Sections {
		fmt.Printf("|---------------------------------------------------------------------\n")
		fmt.Printf("|>>> Discovered Section %s\n", sn.Name)
		fmt.Printf("| \n")
		fmt.Printf("|**********************************************************************\n")
		fmt.Printf("|>>> + \033[32mSection Characteristics              \033[31m| %#x\n", sn.Characteristics)
		fmt.Printf("|>>> + \033[32mSection Virtual Size                 \033[31m| %#x\n", sn.VirtualSize)
		fmt.Printf("|>>> + \033[32mSection Virtual Offset               \033[31m| %#x\n", sn.VirtualAddress)
		fmt.Printf("|>>> + \033[32mSection Raw Size                     \033[31m| %#x\n", sn.Size)
		fmt.Printf("|>>> + \033[32mSection Raw Offset to Data           \033[31m|%#x\n", sn.Offset)
		fmt.Printf("|>>> + \033[32mSection Append Offset (Next Section) \033[31m| %#x\n", sn.Offset+sn.Size)
	}
	return "", nil

}
