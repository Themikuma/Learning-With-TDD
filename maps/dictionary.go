package maps

// Dictionary a repository of known words.
type Dictionary map[string]string

const (
	// ErrNotFound is returned when the word is not defined in the dictionary
	ErrNotFound = DictionaryErr("word definition not found")
	//ErrDuplicateKey flags that a word has already been defined by the dictionary
	ErrDuplicateKey = DictionaryErr("duplicate entry detected")
	//ErrWordDoesNotExist is returned when an attempt to update a missing entry has been made
	ErrWordDoesNotExist = DictionaryErr("cannot update missing entry")
)

//DictionaryErr provides a way to access all errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search finds the definition of a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add a word and it's definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrDuplicateKey
	default:
		return err
	}
	return nil
}

//Update an already existing entry with a new definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return err
}
