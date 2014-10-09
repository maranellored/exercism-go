package foodchain

import (
	"strings"
)

const (
	TestVersion = 1
)

func getVerse(index int) string {

	var ref = []string{``,

		`I know an old lady who swallowed a fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a spider.
It wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a bird.
How absurd to swallow a bird!
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a cat.
Imagine that, to swallow a cat!
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a dog.
What a hog, to swallow a dog!
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a goat.
Just opened her throat and swallowed a goat!
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a cow.
I don't know how she swallowed a cow!
She swallowed the cow to catch the goat.
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

		`I know an old lady who swallowed a horse.
She's dead, of course!`,
	}

	if index > 8 || index < 1 {
		return ""
	}

	return ref[index]

}

func Verse(index int) string {
	return getVerse(index)
}

func Verses(a, b int) string {
	verseA := getVerse(a)
	verseB := getVerse(b)

	stringSlice := []string{verseA, verseB}
	return strings.Join(stringSlice, "\n\n")
}

func Song() string {
	var stringSlice []string
	for i := 1; i <= 8; i++ {
		verse := getVerse(i)
		stringSlice = append(stringSlice, verse)
	}
	return strings.Join(stringSlice, "\n\n")
}
