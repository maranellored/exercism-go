package house

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	phrase := subject
	for _, relPhrase := range relPhrases {
		phrase = Embed(phrase, relPhrase)
	}

	return Embed(phrase, nounPhrase)
}

func Song() string {
	s := "This is"
	endPhrase := "that Jack built."

	verses := []string{
		"the house",
		"the malt\nthat lay in",
		"the rat\nthat ate",
		"the cat\nthat killed",
		"the dog\nthat worried",
		"the cow with the crumpled horn\nthat tossed",
		"the maiden all forlorn\nthat milked",
		"the man all tattered and torn\nthat kissed",
		"the priest all shaven and shorn\nthat married",
		"the rooster that crowed in the morn\nthat woke",
		"the farmer sowing his corn\nthat kept",
		"the horse and the hound and the horn\nthat belonged to",
	}

	var song string
	for i, _ := range verses {
		var songLines []string
		for j := i; j >= 0; j-- {
			songLines = append(songLines, verses[j])
		}
		song += Verse(s, songLines, endPhrase)
		if i != (len(verses) - 1) {
			song += "\n\n"
		}
	}

	return song
}
