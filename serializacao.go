package descritor

import (
	"encoding/json"

	"github.com/gocarina/gocsv"
)

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
func FromJSON(b []byte) (*Candidatura, error) {
	c := &Candidatura{}
	if err := json.Unmarshal(b, c); err != nil {
		return c, err
	}
	return c, nil
}
