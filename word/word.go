package word

// Word type to store individual word data
type Word struct {
	value        string
	similarWords []*Word
	isACareer    bool
	isAnInterest bool
}

// Link will connect
func (w Word) Link(aWord Word) bool {
	if w.isSimilar(aWord) {
		return false
	}
	w.similarWords = append(w.similarWords, &aWord)
	aWord.similarWords = append(aWord.similarWords, &w)
	return true
}

func (w Word) isSimilar(targetWord Word) bool {
	for _, currentWord := range w.similarWords {
		if &(*currentWord) == &targetWord {
			return true
		}
	}

	return false
}

// GetValue returns the words value.
func (w Word) GetValue() string {
	return w.value
}

// SetValue sets the value of Word.
func (w Word) SetValue(value string) {
	w.value = value
}

// New returns a new Word object.
func New(value string, isACareer bool, isAnInterest bool) Word {
	return Word{
		value:        value,
		isACareer:    isACareer,
		isAnInterest: isAnInterest,
		similarWords: make([]*Word, 0),
	}
}
