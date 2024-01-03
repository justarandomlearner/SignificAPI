package word

// object returned by this API
type Payload struct {
	Word     string `json:"word"`
	Meanings []Meaning
}

type Meaning struct {
	GramGrp string `json:"gramGrp,omitempty"`
	Def     string `json:"def"`
	Ast     string `json:"ast,omitempty"`
	Usg     struct {
		Type string `json:"attr"`
		Text string `json:"text"`
	} `json:"usg,omitempty"`
}

func (m *Meaning) SenseToMeaning(s Sense) {
	m.GramGrp = s.GramGrp
	m.Def = s.Def
	m.Ast = s.Ast
	m.Usg.Type = s.Usg.Type
	m.Usg.Text = s.Usg.Text
}
