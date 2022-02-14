package maps

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrEmpytValue = DictionaryErr("cannot add empty value")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string("error: " + e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if value == "" {
		return ErrEmpytValue
	}

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key string, newValue string) error {
	_, err := d.Search(key)

	if newValue == "" {
		return ErrEmpytValue
	}

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[key] = newValue
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
