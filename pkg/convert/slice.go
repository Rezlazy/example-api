package convert

func Slice[T, K any](inputItems []T, converter func(T) K) []K {
	items := make([]K, 0, len(inputItems))
	for _, inputItem := range inputItems {
		item := converter(inputItem)
		items = append(items, item)
	}
	return items
}
