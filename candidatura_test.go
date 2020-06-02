package descritor

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestCandidatura(t *testing.T) {
	is := is.New(t)
	data, err := time.Parse("02-01-2006", "02-10-2016")
	if err != nil {
		t.Errorf("want error nil when creating a test date, got %q", err)
	}
	nascimento, err := time.Parse("02-01-2006", "29-11-1997")
	if err != nil {
		t.Errorf("want error nil when creating a test date for nascimento, got %q", err)
	}
	cand := Candidatura{
		Legislatura:       2016,
		Cargo:             ExecMunicipal,
		UF:                "AL",
		CPF:               "07496470430",
		Turno:             1,
		Descricacao:       "Eleições Municipais 2016",
		Data:              data,
		Municipio:         "Município",
		NumeroUrna:        45555,
		Nome:              "Aurélio Buarque de Miranda Filho",
		NomeUrna:          "Lelinho do povo",
		Email:             "abuarquemf@gmail.com",
		Situacao:          "APTO",
		SituacaoCandidato: "DEFERIDO",
		TipoAgremiacao:    "PARTIDO ISOLADO",
		NumeroPartido:     45,
		LegendaPartido:    "PSDB",
		NomePartido:       "Partido Socilismo",
		NomeColigacao:     "Todos pela cidade",
		PartidosColigacao: "PSDB/PT/PSC",
		UFOrigem:          "UF",
		MunicipioOrigem:   "Maceio",
		Nascimento:        nascimento,
		TituloEleitoral:   "000000000",
		Gênero:            "MASCULINO",
		GrauInstrucao:     "ENSINO MEDIO COMPLETO",
		EstadoCivil:       "SOLTEIRO",
		Raca:              "PARDA",
		OcupacaoCandidato: "ENGENHEIRO",
		SituacaoFinal:     "ELEITO",
	}
	j, err := cand.ToJSON()
	is.NoErr(err)
	is.Equal(`{"leg":2016,"cargo":"EM","uf":"AL","cpf":"07496470430","turno":1,"descricao":"Eleições Municipais 2016","data":"2016-10-02T00:00:00Z","municipio":"Município","numero_urna":45555,"nome":"Aurélio Buarque de Miranda Filho","nome_urna":"Lelinho do povo","email":"abuarquemf@gmail.com","situacao":"APTO","situacao_candidato":"DEFERIDO","tipo_agremiacao":"PARTIDO ISOLADO","numero_partio":45,"legenda_partido":"PSDB","nome_partido":"Partido Socilismo","nome_coligacao":"Todos pela cidade","partidos_coligacao":"PSDB/PT/PSC","uf_origem":"UF","municipio_origem":"Maceio","nascimento":"1997-11-29T00:00:00Z","titulo_eleitoral":"000000000","genero":"MASCULINO","grau_instrucao":"ENSINO MEDIO COMPLETO","estado_civil":"SOLTEIRO","raca":"PARDA","ocupacao":"ENGENHEIRO","situacao_final":"ELEITO"}`, string(j))

}

func TestToCSV(t *testing.T) {
	data, err := time.Parse("02-01-2006", "02-10-2016")
	if err != nil {
		t.Errorf("want error nil when creating a test date, got %q", err)
	}
	nascimento, err := time.Parse("02-01-2006", "29-11-1997")
	if err != nil {
		t.Errorf("want error nil when creating a test date for nascimento, got %q", err)
	}
	c := []Candidatura{
		{
			Legislatura:       2016,
			Cargo:             ExecMunicipal,
			UF:                "AL",
			CPF:               "07496470430",
			Turno:             1,
			Descricacao:       "Eleições Municipais 2016",
			Data:              data,
			Municipio:         "Município",
			NumeroUrna:        45555,
			Nome:              "Aurélio Buarque de Miranda Filho",
			NomeUrna:          "Lelinho do povo",
			Email:             "abuarquemf@gmail.com",
			Situacao:          "APTO",
			SituacaoCandidato: "DEFERIDO",
			TipoAgremiacao:    "PARTIDO ISOLADO",
			NumeroPartido:     45,
			LegendaPartido:    "PSDB",
			NomePartido:       "Partido Socialismo",
			NomeColigacao:     "Todos pela cidade",
			PartidosColigacao: "PSDB/PT/PSC",
			UFOrigem:          "UF",
			MunicipioOrigem:   "Maceio",
			Nascimento:        nascimento,
			TituloEleitoral:   "000000000",
			Gênero:            "MASCULINO",
			GrauInstrucao:     "ENSINO MEDIO COMPLETO",
			EstadoCivil:       "SOLTEIRO",
			Raca:              "PARDA",
			OcupacaoCandidato: "ENGENHEIRO",
			SituacaoFinal:     "ELEITO",
		},
		{
			Legislatura:       2016,
			Cargo:             ExecMunicipal,
			UF:                "AL",
			CPF:               "00000000000",
			Turno:             1,
			Descricacao:       "Eleições Municipais 2016",
			Data:              data,
			Municipio:         "Município",
			NumeroUrna:        45555,
			Nome:              "Marcela Moreira",
			NomeUrna:          "Marcela guerreira",
			Email:             "mmoreira@gmail.com",
			Situacao:          "APTO",
			SituacaoCandidato: "INDEFERIDO",
			TipoAgremiacao:    "PARTIDO ISOLADO",
			NumeroPartido:     45,
			LegendaPartido:    "PSDB",
			NomePartido:       "Partido Socialismo",
			NomeColigacao:     "Todos em nome da cidade",
			PartidosColigacao: "PSTU/PSB/PSL",
			UFOrigem:          "UF",
			MunicipioOrigem:   "Maceio",
			Nascimento:        nascimento,
			TituloEleitoral:   "123456789",
			Gênero:            "FEMININO",
			GrauInstrucao:     "ENSINO MEDIO INCOMPLETO",
			EstadoCivil:       "CASADO",
			Raca:              "NEGRO",
			OcupacaoCandidato: "ARTISTA",
			SituacaoFinal:     "SUPLENTE",
		},
	}
	b, _ := ToCSV(c)
	is := is.New(t)
	is.NoErr(err)
	is.Equal("leg,cargo,uf,cpf,turno,desricao,data,municipio,numero_urna,nome,nome_urna,email,situacao,situacao_candidato,tipo_agremiacao,numero_partido,legenda_partido,nome_partido,nome_coligacao,partidos_coligacao,uf_origem,municipio_origem,nascimento,titulo_eleitoral,genero,grau_instrucao,estado_civil,raca,ocupacao,situacao_final\n2016,EM,AL,07496470430,1,Eleições Municipais 2016,2016-10-02T00:00:00Z,Município,45555,Aurélio Buarque de Miranda Filho,Lelinho do povo,abuarquemf@gmail.com,APTO,DEFERIDO,PARTIDO ISOLADO,45,PSDB,Partido Socialismo,Todos pela cidade,PSDB/PT/PSC,UF,Maceio,1997-11-29T00:00:00Z,000000000,MASCULINO,ENSINO MEDIO COMPLETO,SOLTEIRO,PARDA,ENGENHEIRO,ELEITO\n2016,EM,AL,00000000000,1,Eleições Municipais 2016,2016-10-02T00:00:00Z,Município,45555,Marcela Moreira,Marcela guerreira,mmoreira@gmail.com,APTO,INDEFERIDO,PARTIDO ISOLADO,45,PSDB,Partido Socialismo,Todos em nome da cidade,PSTU/PSB/PSL,UF,Maceio,1997-11-29T00:00:00Z,123456789,FEMININO,ENSINO MEDIO INCOMPLETO,CASADO,NEGRO,ARTISTA,SUPLENTE\n", string(b))
}
