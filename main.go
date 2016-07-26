package locale
import(
	"strings"
	"strconv"
	"sort"
)

type pair struct{
	v string
	w float64
}

type pairs []pair

func (p pairs) Len()int{
	return len(p)
}
func (p pairs) Less(i,j int) bool{
	return p[j].w < p[i].w
}
func (p pairs) Swap(i,j int){
	p[i], p[j] = p[j], p[i]
}

func GetAll(inp string)[]string{
	languages := []string{
		"English",
		"Arabic",
		"French",
		"German",
		"Japanese",
		"Korean",
		"Mandarin",
		"Portugese",
		"Russian",
		"Spanish",
		"Bengali",
		"Bulgarian",
		"Catalan",
		"Croation",
		"Czech",
		"Danish",
		"Dutch",
		"Estonian",
		"Finnish",
		"Greek",
		"Hebrew",
		"Hindi",
		"Hmong",
		"Hungarian",
		"Indonesian",
		"Italian",
		"Latvian",
		"Lithuanian",
		"Javanese",
		"Norwegian",
		"Persian",
		"Polish",
		"Punjabi",
		"Romanian",
		"Serbian",
		"Slovak",
		"Slovenian",
		"Swedish",
		"Thai",
		"Turkish",
		"Ukrainian",
		"Vietnamese",
	}
	topLanguages := Get(inp)
	for _, v := range topLanguages {
		for i, j := range languages {
			if j == v {
				languages = append(languages[:i], languages[i+1:]...)
			}
		}
	}
	return append(topLanguages,languages...)
}

func Get(inp string) []string {
	i := strings.Split(inp,",")
	b := make([]pair,0)
	for _,v := range i {
		k := strings.TrimSpace(v)
		j := strings.Split(k,";q=")
		if len(j) == 1 { j = append(j,"1.0") }
		f, _ := strconv.ParseFloat(j[1], 64)
		b = append(b,pair{j[0],f})
	}
	sort.Sort(pairs(b))
	result := make([]string,0)
	for _, v := range b{
		var capture string
		switch v.v {
			case "en", "eng":
				capture = "English"
			case "ar", "ara":
				capture = "Arabic"
			case "bn", "ben":
				capture = "Bengali"
			case "bg", "bul":
				capture = "Bulgarian"
			case "ca", "cat":
				capture = "Catalan"
			case "hr", "scr", "hrv":
				capture = "Croation"
			case "cs", "cze", "ces":
				capture = "Czech"
			case "da", "dan":
				capture = "Danish"
			case "nl", "dut", "nid":
				capture = "Dutch"
			case "et", "est":
				capture = "Estonian"
			case "fi", "fin":
				capture = "Finnish"
			case "fr", "fre", "fra":
				capture = "French"
			case "de", "ger", "deu":
				capture = "German"
			case "el", "gre", "ell":
				capture = "Greek"
			case "he", "heb":
				capture = "Hebrew"
			case "hi", "hin":
				capture = "Hindi"
			case "hmv", "mww", "hnj", "hmz", "cqd", "hrm", "hmf":
				capture = "Hmong"
			case "hu", "hun":
				capture = "Hungarian"
			case "id", "ind":
				capture = "Indonesian"
			case "it", "ita":
				capture = "Italian"
			case "ja", "jpn":
				capture = "Japanese"
			case "jv", "jav":
				capture = "Javanese"
			case "lv", "lav":
				capture = "Latvian"
			case "lt", "lit":
				capture = "Lithuanian"
			case "ko", "kor":
				capture = "Korean"
			case "cmn":
				capture = "Mandarin"
			case "no", "nor", "nb", "nob", "nn", "nno":
				capture = "Norwegian"
			case "fa", "per", "fas":
				capture = "Persian"
			case "pl", "pol":
				capture = "Polish"
			case "pt", "por":
				capture = "Portugese"
			case "pa", "pan":
				capture = "Punjabi"
			case "ro", "rum","ron":
				capture = "Romanian"
			case "ru", "rus":
				capture = "Russian"
			case "sr", "scc", "srp":
				capture = "Serbian"
			case "sk", "slo", "slk":
				capture = "Slovak"
			case "sl", "slv":
				capture = "Slovenian"
			case "es", "spa":
				capture = "Spanish"
			case "sv", "swe":
				capture = "Swedish"
			case "th", "tha":
				capture = "Thai"
			case "tr", "tur":
				capture = "Turkish"
			case "uk","ukr":
				capture = "Ukrainian"
			case "vie", "vi":
				capture = "Vietnamese"
		}
		if capture != "" {
			result = append(result,capture)
		}
	}
	return result
}