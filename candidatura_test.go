package descritor

import (
	"testing"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/matryer/is"
)

func TestCandidatura(t *testing.T) {
	nascimentoAsTime, err := time.Parse("02-01-2006", "29-11-1997")
	if err != nil {
		t.Errorf("want error nil when creating a test date for nascimento, got %q", err)
	}
	s := int64(nascimentoAsTime.Second())
	n := int32(nascimentoAsTime.Nanosecond())
	nascimentoAsTimestamp := &timestamp.Timestamp{Seconds: s, Nanos: n}
	cand := Candidatura{
		Aptdao:                "APTO",
		Legislatura:           2016,
		Cargo:                 "Prefeito",
		UF:                    "AL",
		Municipio:             "Município",
		NumeroUrna:            45555,
		NomeUrna:              "Lelinho do povo",
		Deferimento:           "DEFERIDO",
		TipoAgremiacao:        "PARTIDO ISOLADO",
		NumeroPartido:         45,
		LegendaPartido:        "PSDB",
		NomePartido:           "Partido Socilismo",
		NomeColigacao:         "Todos pela cidade",
		PartidosColigacao:     "PSDB/PT/PSC",
		SituacaoSegundoTurno:  "ELEITO",
		SituacaoPrimeiroTurno: "SEGUNDO TURNO",
		DeclarouBens:          true,
		SequencialCandidato:   "505",
		Candidato: &Candidato{
			Email:           "abuarquemf@gmail.com",
			Nome:            "Aurélio Buarque de Miranda Filho",
			CPF:             "07496470430",
			UF:              "UF",
			Municipio:       "Maceio",
			Nascimento:      nascimentoAsTimestamp,
			TituloEleitoral: "000000000",
			Genero:          "MASCULINO",
			GrauInstrucao:   "ENSINO MEDIO COMPLETO",
			EstadoCivil:     "SOLTEIRO",
			Raca:            "PARDA",
			Ocupacao:        "ENGENHEIRO",
		},
	}
	j, err := cand.ToJSON()
	if err != nil {
		t.Errorf("want error nil when passing struct candidato to JSON, got %q", err)
	}
	is := is.New(t)
	is.NoErr(err)
	is.Equal(`{"leg":2016,"cargo":"Prefeito","uf":"AL","municipio":"Município","numero_urna":45555,"nome_urna":"Lelinho do povo","aptidao":"APTO","deferimento":"DEFERIDO","tipo_agremiacao":"PARTIDO ISOLADO","numero_partio":45,"legenda_partido":"PSDB","nome_partido":"Partido Socilismo","nome_coligacao":"Todos pela cidade","partidos_coligacao":"PSDB/PT/PSC","declarou_bens":true,"situacao_1turno":"SEGUNDO TURNO","situacao_2turno":"ELEITO","sequencial_candidato":"505","candidato":{"uf_origem":"UF","municipio_origem":"Maceio","nascimento":{},"titulo_eleitoral":"000000000","genero":"MASCULINO","grau_instrucao":"ENSINO MEDIO COMPLETO","estado_civil":"SOLTEIRO","raca":"PARDA","ocupacao":"ENGENHEIRO","cpf":"07496470430","nome":"Aurélio Buarque de Miranda Filho","email":"abuarquemf@gmail.com"}}`, string(j))
}

func TestToCSV(t *testing.T) {
	nascimentoAsTime, err := time.Parse("02-01-2006", "29-11-1997")
	if err != nil {
		t.Errorf("want error nil when creating a test date for nascimento, got %q", err)
	}
	s := int64(nascimentoAsTime.Second())
	n := int32(nascimentoAsTime.Nanosecond())
	nascimentoAsTimestamp := &timestamp.Timestamp{Seconds: s, Nanos: n}
	c := []Candidatura{
		{
			Aptdao:                "APTO",
			Legislatura:           2016,
			Cargo:                 "Prefeito",
			UF:                    "AL",
			Municipio:             "Município",
			NumeroUrna:            45555,
			NomeUrna:              "Lelinho do povo",
			Deferimento:           "DEFERIDO",
			TipoAgremiacao:        "PARTIDO ISOLADO",
			NumeroPartido:         45,
			LegendaPartido:        "PSDB",
			NomePartido:           "Partido Socilismo",
			NomeColigacao:         "Todos pela cidade",
			PartidosColigacao:     "PSDB/PT/PSC",
			DeclarouBens:          true,
			SituacaoSegundoTurno:  "ELEITO",
			SituacaoPrimeiroTurno: "SEGUNDO TURNO",
			SequencialCandidato:   "505",
			Candidato: &Candidato{
				Email:           "abuarquemf@gmail.com",
				Nome:            "Aurélio Buarque de Miranda Filho",
				CPF:             "07496470430",
				UF:              "UF",
				Municipio:       "Maceio",
				Nascimento:      nascimentoAsTimestamp,
				TituloEleitoral: "000000000",
				Genero:          "MASCULINO",
				GrauInstrucao:   "ENSINO MEDIO COMPLETO",
				EstadoCivil:     "SOLTEIRO",
				Raca:            "PARDA",
				Ocupacao:        "ENGENHEIRO",
			},
		},
		{
			Aptdao:                "APTO",
			Legislatura:           2016,
			Cargo:                 "Prefeito",
			UF:                    "AL",
			Municipio:             "Município",
			NumeroUrna:            45555,
			NomeUrna:              "Lelinho do povo",
			Deferimento:           "DEFERIDO",
			TipoAgremiacao:        "PARTIDO ISOLADO",
			NumeroPartido:         45,
			LegendaPartido:        "PSDB",
			NomePartido:           "Partido Socilismo",
			NomeColigacao:         "Todos pela cidade",
			PartidosColigacao:     "PSDB/PT/PSC",
			DeclarouBens:          true,
			SituacaoSegundoTurno:  "",
			SituacaoPrimeiroTurno: "ELEITO",
			SequencialCandidato:   "505",
			Candidato: &Candidato{
				Email:           "abuarquemf@gmail.com",
				Nome:            "Aurélio Buarque de Miranda Filho",
				CPF:             "07496470430",
				UF:              "UF",
				Municipio:       "Maceio",
				Nascimento:      nascimentoAsTimestamp,
				TituloEleitoral: "000000000",
				Genero:          "MASCULINO",
				GrauInstrucao:   "ENSINO MEDIO COMPLETO",
				EstadoCivil:     "SOLTEIRO",
				Raca:            "PARDA",
				Ocupacao:        "ENGENHEIRO",
			},
		},
	}
	b, err := ToCSV(c)
	if err != nil {
		t.Errorf("want error nil when passing struct candidato to CSV, got %q", err)
	}
	is := is.New(t)
	is.NoErr(err)
	is.Equal("leg,cargo,uf,municipio,numero_urna,nome_urna,aptidao,deferimento,tipo_agremiacao,numero_partido,legenda_partido,nome_partido,nome_coligacao,partidos_coligacao,declarou_bens,situacao_1turno,situacao_2turno,sequencial_candidato,uf_origem,municipio_origem,Seconds,Nanos,nascimento,titulo_eleitoral,genero,grau_instrucao,estado_civil,raca,ocupacao,cpf,nome,email\n2016,Prefeito,AL,Município,45555,Lelinho do povo,APTO,DEFERIDO,PARTIDO ISOLADO,45,PSDB,Partido Socilismo,Todos pela cidade,PSDB/PT/PSC,true,SEGUNDO TURNO,ELEITO,505,UF,Maceio,0,0,,000000000,MASCULINO,ENSINO MEDIO COMPLETO,SOLTEIRO,PARDA,ENGENHEIRO,07496470430,Aurélio Buarque de Miranda Filho,abuarquemf@gmail.com\n2016,Prefeito,AL,Município,45555,Lelinho do povo,APTO,DEFERIDO,PARTIDO ISOLADO,45,PSDB,Partido Socilismo,Todos pela cidade,PSDB/PT/PSC,true,ELEITO,,505,UF,Maceio,0,0,,000000000,MASCULINO,ENSINO MEDIO COMPLETO,SOLTEIRO,PARDA,ENGENHEIRO,07496470430,Aurélio Buarque de Miranda Filho,abuarquemf@gmail.com\n", string(b))
}
