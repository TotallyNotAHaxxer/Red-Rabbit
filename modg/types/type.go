package types

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

type Result struct {
	Result_Sink string
	Url         string
}
