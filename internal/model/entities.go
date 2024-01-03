package word

// maps the JSON returned by https://api.dicionario-aberto.net/word/
type Result struct {
	Deleted      int    `json:"deleted"`
	LastRevision int    `json:"last_revision"`
	Deletor      string `json:"deletor"`
	DerivedFrom  string `json:"derived_from"`
	Timestamp    string `json:"timestamp"`
	Word         string `json:"word"`
	RevisionID   int    `json:"revision_id"`
	Normalized   string `json:"normalized"`
	Sense        int    `json:"sense"`
	Creator      string `json:"creator"`
	WordID       int    `json:"word_id"`
	Moderator    string `json:"moderator"`
	XML          string `json:"xml"`
}

// represents an entry found in the returned XML
type Entry struct {
	ID   string `xml:"id,attr"`
	Form struct {
		Orth string `xml:"orth"`
	} `xml:"form"`
	Senses []Sense `xml:"sense"`
	Etym   struct {
		Orig string `xml:"orig,attr"`
	} `xml:"etym"`
}

// represents each sense an entry can have
type Sense struct {
	GramGrp string `xml:"gramGrp"`
	Def     string `xml:"def"`
	Ast     string `xml:"ast,attr,omitempty"`
	Usg     struct {
		Type string `xml:"type,attr"`
		Text string `xml:",chardata"`
	} `xml:"usg"`
}
