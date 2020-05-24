package descritor

import (
	"encoding/json"

	"github.com/gocarina/gocsv"
)

// Cargo representa os possíveis cargos pleiteados pelas candidaturas.
type Cargo string

// Cargos pleiteados via candidaturas
const (
	LegMunicipal  = "LM" // Legislativo municipal (a.k.a vereadora)
	ExecMunicipal = "EM" // Executivo municipal (a.k.a prefeita)
)

// Candidatura representa uma
type Candidatura struct {
	Legislatura int    `json:"leg" csv:"leg"`     // Ano eleitoral em que a candidatura foi homologada.
	Cargo       Cargo  `json:"cargo" csv:"cargo"` // Cargo sendo pleiteado pela candidatura.
	UF          string `json:"uf" csv:"uf"`       // Identificador (2 caracteres) de unidade federativa.
}

// ToJSON converte a candidatura para JSON.
func (c *Candidatura) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// ToCSV converte a candidatura para um conjunto de strings
func ToCSV(c []Candidatura) ([]byte, error) {
	b, err := gocsv.MarshalBytes(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// FromJSON cria uma instância de candidatura preenchida com os dados contidos no JSON passado como parâmetro.
func FromJSON(b []byte) (Candidatura, error) {
	c := Candidatura{}
	if err := json.Unmarshal(b, &c); err != nil {
		return c, err
	}
	return c, nil
}
