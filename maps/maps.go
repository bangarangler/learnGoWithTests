package maps

import "errors"

type Dictionairy map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionairy) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }
