package collection

func Map[T any, R any](coll []T, iter func(el T) R) []R {
	var newColl []R
	for i := 0; i < len(coll); i++ {
		newColl = append(newColl, iter(coll[i]))
	}
	return newColl
}

func MapErr[T any, R any](coll []T, iter func(el T) (R, error)) ([]R, error) {
	var newColl []R
	for i := 0; i < len(coll); i++ {
		newEl, err := iter(coll[i])
		if err != nil {
			return nil, err
		}
		newColl = append(newColl, newEl)
	}
	return newColl, nil
}

func Each[T any](coll []T, iter func(el T)) {
	for i := 0; i < len(coll); i++ {
		el := coll[i]
		iter(el)
	}
}

func EachErr[T any](coll []T, iter func(el T) error) error {
	for i := 0; i < len(coll); i++ {
		el := coll[i]
		if err := iter(el); err != nil {
			return err
		}
	}
	return nil
}
