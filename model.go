package main

type ESG struct {
	Environment Standard `json:"Environment"`
	Social Standard `json:"Social"`
	Governance Standard `json:"Governance"`
	Links Links `json:"Links"`
}

type Standard struct {
	Rating       string `json:"rating"`
	FactorsScore FactorsScore `json:"factorsScore"`
}

type FactorsScore []struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Links []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
