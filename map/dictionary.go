package dictionary

// custom type - which is a map
type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound    = DictionaryErr("Search Error - no such word in the dictionary")
	ErrWordExist   = DictionaryErr("the word is already existed in dictionary")
	ErrNoWordExist = DictionaryErr("the word is not present in dictionary")
)

// method attached with the type
// search a word in dictionary if found return the meaning else return err
func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	// fmt.Println("d", d[key])
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

// add a word in dictionary, if already there then err
func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]
	// didn't get any keyy
	if !ok {
		d[key] = val
		return nil
	}
	return ErrWordExist
}

// update a word, if not already there then return err
func (d Dictionary) Update(word, updatedMeaning string) error {
	_, ok := d[word]
	if !ok {
		return ErrNoWordExist
	}
	d[word] = updatedMeaning
	return nil
}

// delete a word, if not already there then return err
func (d Dictionary) Delete(word string) error {
	_, ok := d[word]
	if !ok {
		return ErrNoWordExist
	}
	delete(d, word)
	return nil
}
