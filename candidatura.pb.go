// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candidatura.proto

package descritor

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Candidatura struct {
	Legislatura           int64      `protobuf:"varint,1,opt,name=Legislatura,proto3" json:"leg" csv:"leg"`
	Cargo                 string     `protobuf:"bytes,2,opt,name=Cargo,proto3" json:"cargo" csv:"cargo"`
	UF                    string     `protobuf:"bytes,3,opt,name=UF,proto3" json:"uf" csv:"uf"`
	Municipio             string     `protobuf:"bytes,4,opt,name=Municipio,proto3" json:"municipio" csv:"municipio"`
	NumeroUrna            int64      `protobuf:"varint,5,opt,name=NumeroUrna,proto3" json:"numero_urna" csv:"numero_urna"`
	NomeUrna              string     `protobuf:"bytes,6,opt,name=NomeUrna,proto3" json:"nome_urna" csv:"nome_urna"`
	Aptdao                string     `protobuf:"bytes,7,opt,name=Aptdao,proto3" json:"aptidao" csv:"aptidao"`
	Deferimento           string     `protobuf:"bytes,8,opt,name=Deferimento,proto3" json:"deferimento" csv:"deferimento"`
	TipoAgremiacao        string     `protobuf:"bytes,9,opt,name=TipoAgremiacao,proto3" json:"tipo_agremiacao" csv:"tipo_agremiacao"`
	NumeroPartido         int64      `protobuf:"varint,10,opt,name=NumeroPartido,proto3" json:"numero_partio" csv:"numero_partido"`
	LegendaPartido        string     `protobuf:"bytes,11,opt,name=LegendaPartido,proto3" json:"legenda_partido" csv:"legenda_partido"`
	NomePartido           string     `protobuf:"bytes,12,opt,name=NomePartido,proto3" json:"nome_partido" csv:"nome_partido"`
	NomeColigacao         string     `protobuf:"bytes,13,opt,name=NomeColigacao,proto3" json:"nome_coligacao" csv:"nome_coligacao"`
	PartidosColigacao     string     `protobuf:"bytes,14,opt,name=PartidosColigacao,proto3" json:"partidos_coligacao" csv:"partidos_coligacao"`
	DeclarouBens          bool       `protobuf:"varint,15,opt,name=DeclarouBens,proto3" json:"declarou_bens" csv:"declarou_bens"`
	SituacaoPrimeiroTurno string     `protobuf:"bytes,16,opt,name=SituacaoPrimeiroTurno,proto3" json:"situacao_1turno" csv:"situacao_1turno"`
	SituacaoSegundoTurno  string     `protobuf:"bytes,17,opt,name=SituacaoSegundoTurno,proto3" json:"situacao_2turno" csv:"situacao_2turno"`
	SequencialCandidato   string     `protobuf:"bytes,18,opt,name=SequencialCandidato,proto3" json:"sequencial_candidato" csv:"sequencial_candidato"`
	Candidato             *Candidato `protobuf:"bytes,19,opt,name=candidato,proto3" json:"candidato,omitempty" csv:"-"`
	XXX_NoUnkeyedLiteral  struct{}   `json:"-" csv:"-"`
	XXX_unrecognized      []byte     `json:"-" csv:"-"`
	XXX_sizecache         int32      `json:"-" csv:"-"`
}

func (m *Candidatura) Reset()         { *m = Candidatura{} }
func (m *Candidatura) String() string { return proto.CompactTextString(m) }
func (*Candidatura) ProtoMessage()    {}
func (*Candidatura) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4fb5b387980952, []int{0}
}

func (m *Candidatura) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidatura.Unmarshal(m, b)
}
func (m *Candidatura) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidatura.Marshal(b, m, deterministic)
}
func (m *Candidatura) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidatura.Merge(m, src)
}
func (m *Candidatura) XXX_Size() int {
	return xxx_messageInfo_Candidatura.Size(m)
}
func (m *Candidatura) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidatura.DiscardUnknown(m)
}

var xxx_messageInfo_Candidatura proto.InternalMessageInfo

func (m *Candidatura) GetLegislatura() int64 {
	if m != nil {
		return m.Legislatura
	}
	return 0
}

func (m *Candidatura) GetCargo() string {
	if m != nil {
		return m.Cargo
	}
	return ""
}

func (m *Candidatura) GetUF() string {
	if m != nil {
		return m.UF
	}
	return ""
}

func (m *Candidatura) GetMunicipio() string {
	if m != nil {
		return m.Municipio
	}
	return ""
}

func (m *Candidatura) GetNumeroUrna() int64 {
	if m != nil {
		return m.NumeroUrna
	}
	return 0
}

func (m *Candidatura) GetNomeUrna() string {
	if m != nil {
		return m.NomeUrna
	}
	return ""
}

func (m *Candidatura) GetAptdao() string {
	if m != nil {
		return m.Aptdao
	}
	return ""
}

func (m *Candidatura) GetDeferimento() string {
	if m != nil {
		return m.Deferimento
	}
	return ""
}

func (m *Candidatura) GetTipoAgremiacao() string {
	if m != nil {
		return m.TipoAgremiacao
	}
	return ""
}

func (m *Candidatura) GetNumeroPartido() int64 {
	if m != nil {
		return m.NumeroPartido
	}
	return 0
}

func (m *Candidatura) GetLegendaPartido() string {
	if m != nil {
		return m.LegendaPartido
	}
	return ""
}

func (m *Candidatura) GetNomePartido() string {
	if m != nil {
		return m.NomePartido
	}
	return ""
}

func (m *Candidatura) GetNomeColigacao() string {
	if m != nil {
		return m.NomeColigacao
	}
	return ""
}

func (m *Candidatura) GetPartidosColigacao() string {
	if m != nil {
		return m.PartidosColigacao
	}
	return ""
}

func (m *Candidatura) GetDeclarouBens() bool {
	if m != nil {
		return m.DeclarouBens
	}
	return false
}

func (m *Candidatura) GetSituacaoPrimeiroTurno() string {
	if m != nil {
		return m.SituacaoPrimeiroTurno
	}
	return ""
}

func (m *Candidatura) GetSituacaoSegundoTurno() string {
	if m != nil {
		return m.SituacaoSegundoTurno
	}
	return ""
}

func (m *Candidatura) GetSequencialCandidato() string {
	if m != nil {
		return m.SequencialCandidato
	}
	return ""
}

func (m *Candidatura) GetCandidato() *Candidato {
	if m != nil {
		return m.Candidato
	}
	return nil
}

type Candidato struct {
	UF                   string               `protobuf:"bytes,1,opt,name=UF,proto3" json:"uf_origem" csv:"uf_origem"`
	Municipio            string               `protobuf:"bytes,2,opt,name=Municipio,proto3" json:"municipio_origem" csv:"municipio_origem"`
	Nascimento           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Nascimento,proto3" json:"nascimento" csv:"nascimento"`
	TituloEleitoral      string               `protobuf:"bytes,4,opt,name=TituloEleitoral,proto3" json:"titulo_eleitoral" csv:"titulo_eleitoral"`
	Genero               string               `protobuf:"bytes,5,opt,name=Genero,proto3" json:"genero" csv:"genero"`
	GrauInstrucao        string               `protobuf:"bytes,6,opt,name=GrauInstrucao,proto3" json:"grau_instrucao" csv:"grau_instrucao"`
	EstadoCivil          string               `protobuf:"bytes,7,opt,name=EstadoCivil,proto3" json:"estado_civil" csv:"estado_civil"`
	Raca                 string               `protobuf:"bytes,8,opt,name=Raca,proto3" json:"raca" csv:"raca"`
	Ocupacao             string               `protobuf:"bytes,9,opt,name=Ocupacao,proto3" json:"ocupacao" csv:"ocupacao"`
	CPF                  string               `protobuf:"bytes,10,opt,name=CPF,proto3" json:"cpf" csv:"cpf"`
	Nome                 string               `protobuf:"bytes,11,opt,name=Nome,proto3" json:"nome" csv:"nome"`
	Email                string               `protobuf:"bytes,12,opt,name=Email,proto3" json:"email" csv:"email"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" csv:"-"`
	XXX_unrecognized     []byte               `json:"-" csv:"-"`
	XXX_sizecache        int32                `json:"-" csv:"-"`
}

func (m *Candidato) Reset()         { *m = Candidato{} }
func (m *Candidato) String() string { return proto.CompactTextString(m) }
func (*Candidato) ProtoMessage()    {}
func (*Candidato) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4fb5b387980952, []int{1}
}

func (m *Candidato) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidato.Unmarshal(m, b)
}
func (m *Candidato) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidato.Marshal(b, m, deterministic)
}
func (m *Candidato) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidato.Merge(m, src)
}
func (m *Candidato) XXX_Size() int {
	return xxx_messageInfo_Candidato.Size(m)
}
func (m *Candidato) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidato.DiscardUnknown(m)
}

var xxx_messageInfo_Candidato proto.InternalMessageInfo

func (m *Candidato) GetUF() string {
	if m != nil {
		return m.UF
	}
	return ""
}

func (m *Candidato) GetMunicipio() string {
	if m != nil {
		return m.Municipio
	}
	return ""
}

func (m *Candidato) GetNascimento() *timestamp.Timestamp {
	if m != nil {
		return m.Nascimento
	}
	return nil
}

func (m *Candidato) GetTituloEleitoral() string {
	if m != nil {
		return m.TituloEleitoral
	}
	return ""
}

func (m *Candidato) GetGenero() string {
	if m != nil {
		return m.Genero
	}
	return ""
}

func (m *Candidato) GetGrauInstrucao() string {
	if m != nil {
		return m.GrauInstrucao
	}
	return ""
}

func (m *Candidato) GetEstadoCivil() string {
	if m != nil {
		return m.EstadoCivil
	}
	return ""
}

func (m *Candidato) GetRaca() string {
	if m != nil {
		return m.Raca
	}
	return ""
}

func (m *Candidato) GetOcupacao() string {
	if m != nil {
		return m.Ocupacao
	}
	return ""
}

func (m *Candidato) GetCPF() string {
	if m != nil {
		return m.CPF
	}
	return ""
}

func (m *Candidato) GetNome() string {
	if m != nil {
		return m.Nome
	}
	return ""
}

func (m *Candidato) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CandidaturasDeCidade struct {
	Group                map[string]*Candidatura `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CandidaturasDeCidade) Reset()         { *m = CandidaturasDeCidade{} }
func (m *CandidaturasDeCidade) String() string { return proto.CompactTextString(m) }
func (*CandidaturasDeCidade) ProtoMessage()    {}
func (*CandidaturasDeCidade) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4fb5b387980952, []int{2}
}

func (m *CandidaturasDeCidade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CandidaturasDeCidade.Unmarshal(m, b)
}
func (m *CandidaturasDeCidade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CandidaturasDeCidade.Marshal(b, m, deterministic)
}
func (m *CandidaturasDeCidade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandidaturasDeCidade.Merge(m, src)
}
func (m *CandidaturasDeCidade) XXX_Size() int {
	return xxx_messageInfo_CandidaturasDeCidade.Size(m)
}
func (m *CandidaturasDeCidade) XXX_DiscardUnknown() {
	xxx_messageInfo_CandidaturasDeCidade.DiscardUnknown(m)
}

var xxx_messageInfo_CandidaturasDeCidade proto.InternalMessageInfo

func (m *CandidaturasDeCidade) GetGroup() map[string]*Candidatura {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*Candidatura)(nil), "descritor.Candidatura")
	proto.RegisterType((*Candidato)(nil), "descritor.Candidato")
	proto.RegisterType((*CandidaturasDeCidade)(nil), "descritor.CandidaturasDeCidade")
	proto.RegisterMapType((map[string]*Candidatura)(nil), "descritor.CandidaturasDeCidade.GroupEntry")
}

func init() { proto.RegisterFile("candidatura.proto", fileDescriptor_bc4fb5b387980952) }

var fileDescriptor_bc4fb5b387980952 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdf, 0x6a, 0xdb, 0x4a,
	0x10, 0xc6, 0x91, 0x1d, 0xe7, 0x44, 0xe3, 0xfc, 0xdd, 0xf8, 0x84, 0xc5, 0x1c, 0x4e, 0x8d, 0x29,
	0xc5, 0x94, 0xa0, 0x14, 0xb7, 0x17, 0x25, 0x57, 0x4d, 0x1d, 0x27, 0x14, 0xd2, 0xd4, 0x28, 0xce,
	0x03, 0x6c, 0xa4, 0x89, 0x58, 0x2a, 0x69, 0xd5, 0xd5, 0x6e, 0x20, 0xaf, 0xd4, 0x37, 0xe9, 0xb3,
	0xf4, 0x25, 0xca, 0x8e, 0x64, 0x5b, 0x4e, 0x7d, 0xb7, 0xf3, 0xcd, 0x37, 0xb3, 0x83, 0xf6, 0x37,
	0x82, 0xa3, 0x48, 0xe4, 0xb1, 0x8c, 0x85, 0xb1, 0x5a, 0x04, 0x85, 0x56, 0x46, 0x31, 0x3f, 0xc6,
	0x32, 0xd2, 0xd2, 0x28, 0xdd, 0x7f, 0x95, 0x28, 0x95, 0xa4, 0x78, 0x46, 0x89, 0x07, 0xfb, 0x78,
	0x66, 0x64, 0x86, 0xa5, 0x11, 0x59, 0x51, 0x79, 0x87, 0xbf, 0x3a, 0xd0, 0x9d, 0xac, 0x3a, 0xb0,
	0x01, 0x74, 0x6f, 0x30, 0x91, 0x65, 0x4a, 0x21, 0xf7, 0x06, 0xde, 0xa8, 0x1d, 0x36, 0x25, 0xd6,
	0x83, 0xce, 0x44, 0xe8, 0x44, 0xf1, 0xd6, 0xc0, 0x1b, 0xf9, 0x61, 0x15, 0xb0, 0x7d, 0x68, 0xdd,
	0x5f, 0xf1, 0x36, 0x49, 0xad, 0xfb, 0x2b, 0xf6, 0x1f, 0xf8, 0x5f, 0x6d, 0x2e, 0x23, 0x59, 0x48,
	0xc5, 0xb7, 0x48, 0x5e, 0x09, 0xec, 0x7f, 0x80, 0x5b, 0x9b, 0xa1, 0x56, 0xf7, 0x3a, 0x17, 0xbc,
	0x43, 0x97, 0x34, 0x14, 0xd6, 0x87, 0x9d, 0x5b, 0x95, 0x21, 0x65, 0xb7, 0xa9, 0x78, 0x19, 0xb3,
	0x13, 0xd8, 0xbe, 0x28, 0x4c, 0x2c, 0x14, 0xff, 0x87, 0x32, 0x75, 0xe4, 0x26, 0xbf, 0xc4, 0x47,
	0xd4, 0x32, 0xc3, 0xdc, 0x28, 0xbe, 0x43, 0xc9, 0xa6, 0xc4, 0xde, 0xc0, 0xfe, 0x5c, 0x16, 0xea,
	0x22, 0xd1, 0x98, 0x49, 0x11, 0x09, 0xc5, 0x7d, 0x32, 0xbd, 0x50, 0xd9, 0x6b, 0xd8, 0xab, 0x66,
	0x99, 0x09, 0x6d, 0x64, 0xac, 0x38, 0xd0, 0x80, 0xeb, 0xa2, 0xeb, 0x76, 0x83, 0x09, 0xe6, 0xb1,
	0x58, 0xd8, 0xba, 0x55, 0xb7, 0x75, 0xd5, 0xcd, 0xe5, 0x66, 0x5f, 0x98, 0x76, 0xab, 0xb9, 0x1a,
	0x12, 0xdd, 0xa7, 0x32, 0x9c, 0xa8, 0x54, 0x26, 0x34, 0xd6, 0x1e, 0x79, 0xd6, 0x45, 0x76, 0x0a,
	0x47, 0x75, 0x41, 0xb9, 0x72, 0xee, 0x93, 0xf3, 0xef, 0x04, 0x1b, 0xc2, 0xee, 0x25, 0x46, 0xa9,
	0xd0, 0xca, 0x7e, 0xc6, 0xbc, 0xe4, 0x07, 0x03, 0x6f, 0xb4, 0x13, 0xae, 0x69, 0xec, 0x03, 0xfc,
	0x7b, 0x27, 0x8d, 0x75, 0xfe, 0x99, 0xfb, 0x46, 0x52, 0xab, 0xb9, 0xd5, 0xb9, 0xe2, 0x87, 0xd4,
	0x75, 0x73, 0x92, 0x8d, 0xa1, 0xb7, 0x48, 0xdc, 0x61, 0x62, 0xf3, 0xb8, 0x2e, 0x3a, 0xa2, 0xa2,
	0x8d, 0x39, 0xf6, 0x0e, 0x8e, 0xef, 0xf0, 0x87, 0xc5, 0x3c, 0x92, 0x22, 0x5d, 0xe0, 0xa6, 0x38,
	0xa3, 0x92, 0x4d, 0x29, 0x36, 0x06, 0x3f, 0x5a, 0xfa, 0x8e, 0x07, 0xde, 0xa8, 0x3b, 0xee, 0x05,
	0x4b, 0xae, 0x83, 0xa5, 0x31, 0x5c, 0xd9, 0x86, 0xbf, 0x5b, 0xe0, 0xaf, 0x3a, 0x54, 0x44, 0x7a,
	0x9b, 0x89, 0x6c, 0xbd, 0x24, 0xf2, 0x1c, 0xe0, 0x56, 0x94, 0x51, 0x0d, 0x4f, 0x9b, 0x2e, 0xec,
	0x07, 0xd5, 0xf6, 0x04, 0x8b, 0xed, 0x09, 0xe6, 0x8b, 0xed, 0x09, 0x1b, 0x6e, 0x36, 0x82, 0x83,
	0xb9, 0x34, 0x36, 0x55, 0xd3, 0x14, 0xdd, 0x74, 0x22, 0xad, 0x89, 0x7f, 0x29, 0x3b, 0x76, 0xaf,
	0x31, 0x47, 0xad, 0x88, 0x79, 0x3f, 0xac, 0x23, 0x47, 0xc0, 0xb5, 0x16, 0xf6, 0x4b, 0x5e, 0x1a,
	0x6d, 0xdd, 0xbb, 0x56, 0xd0, 0xaf, 0x8b, 0x8e, 0xa4, 0x69, 0x69, 0x44, 0xac, 0x26, 0xf2, 0x49,
	0xa6, 0x35, 0xfe, 0x4d, 0x89, 0x31, 0xd8, 0x0a, 0x45, 0x24, 0x6a, 0xf8, 0xe9, 0xec, 0x76, 0xe9,
	0x5b, 0x64, 0x8b, 0x06, 0xef, 0xcb, 0x98, 0x1d, 0x42, 0x7b, 0x32, 0xbb, 0x22, 0xbe, 0xfd, 0xd0,
	0x1d, 0x5d, 0x07, 0x87, 0x5d, 0xcd, 0x32, 0x9d, 0xdd, 0xc6, 0x4f, 0x33, 0x21, 0xd3, 0x9a, 0xdd,
	0x2a, 0x18, 0xfe, 0xf4, 0xa0, 0xd7, 0xf8, 0x73, 0x94, 0x97, 0x38, 0x91, 0xb1, 0x88, 0x91, 0x7d,
	0x82, 0x4e, 0xa2, 0x95, 0x2d, 0xb8, 0x37, 0x68, 0x8f, 0xba, 0xe3, 0xb7, 0x1b, 0x9e, 0xad, 0xe9,
	0x0f, 0xae, 0x9d, 0x79, 0x9a, 0x1b, 0xfd, 0x1c, 0x56, 0x85, 0xfd, 0x19, 0xc0, 0x4a, 0x74, 0x43,
	0x7e, 0xc7, 0xe7, 0xfa, 0x25, 0xdd, 0x91, 0x9d, 0x42, 0xe7, 0x49, 0xa4, 0x16, 0xe9, 0x19, 0xbb,
	0xe3, 0x93, 0xcd, 0x37, 0x84, 0x95, 0xe9, 0xbc, 0xf5, 0xd1, 0x7b, 0xd8, 0xa6, 0x27, 0x7c, 0xff,
	0x27, 0x00, 0x00, 0xff, 0xff, 0xff, 0x18, 0xc9, 0xd0, 0x2e, 0x05, 0x00, 0x00,
}
