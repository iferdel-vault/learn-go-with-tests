package main

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("the word already exists in the dictionary")
	ErrWordDoesNotExists = DictionaryErr("cannot update word since it does not exists in the dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if _, ok := d[word]; !ok {
		return ErrWordDoesNotExists
	}

	d[word] = definition
	return nil
}

// delete a value that is not in the map has no effect
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
