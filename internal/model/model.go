package word

// object returned by this API
type Payload struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

type Usg struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Meaning struct {
	GramGrp   string `json:"gramGrp,omitempty"`
	Def       string `json:"def"`
	Ast       string `json:"ast,omitempty"`
	UsageInfo []Usg  `json:"usg,omitempty"`
}

func (m *Meaning) SenseToMeaning(s Sense) {
	m.GramGrp = s.GramGrp
	m.Def = s.Def
	m.Ast = s.Ast
	for _, usg := range s.Usg {
		m.UsageInfo = append(m.UsageInfo, Usg(usg))
	}
}

func (m *Meaning) Sanitize() {

}
