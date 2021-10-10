package models

import (
	"encoding/csv"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
	"strings"
)

type Operadora struct {
	RawLine            string `json:"raw_line"`
	RegistroANS        string `json:"registro_ans"`
	CNPJ               string `json:"cnpj"`
	RazaoSocial        string `json:"razao_social"`
	NomeFantasia       string `json:"nome_fantasia"`
	Modalidade         string `json:"modalidade"`
	Logradouro         string `json:"logradouro"`
	Numero             string `json:"numero"`
	Complemento        string `json:"complemento"`
	Bairro             string `json:"bairro"`
	Cidade             string `json:"cidade"`
	UF                 string `json:"uf"`
	CEP                string `json:"cep"`
	DDD                string `json:"ddd"`
	Telefone           string `json:"telefone"`
	Fax                string `json:"fax"`
	EnderecoEletronico string `json:"endereco_eletronico"`
	Representante      string `json:"representante"`
	CargoRepresentante string `json:"cargo_representante"`
	DataRegistroANS    string `json:"data_registro_ans"`
}

type Operadoras []*Operadora

func (oo *Operadoras) LoadFromCSV(f *os.File) {
	r := csv.NewReader(transform.NewReader(f, charmap.ISO8859_1.NewDecoder()))
	r.Comma = ';'
	r.FieldsPerRecord = -1
	r.LazyQuotes = true

	for i := 0; i < 2; i++ {
		_, _ = r.Read()
	}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		*oo = append(*oo, &Operadora{
			RawLine:            strings.Join(line, ";"),
			RegistroANS:        line[0],
			CNPJ:               line[1],
			RazaoSocial:        line[2],
			NomeFantasia:       line[3],
			Modalidade:         line[4],
			Logradouro:         line[5],
			Numero:             line[6],
			Complemento:        line[7],
			Bairro:             line[8],
			Cidade:             line[9],
			UF:                 line[10],
			CEP:                line[11],
			DDD:                line[12],
			Telefone:           line[13],
			Fax:                line[14],
			EnderecoEletronico: line[15],
			Representante:      line[16],
			CargoRepresentante: line[17],
			DataRegistroANS:    line[18],
		})
	}
}

func (oo *Operadoras) Filter(str string) Operadoras {
	var result Operadoras
	for i := 0; i < len(*oo); i++ {
		if strings.Contains(strings.ToLower((*oo)[i].RawLine), strings.ToLower(str)) {
			result = append(result, (*oo)[i])
		}
	}
	return result
}
