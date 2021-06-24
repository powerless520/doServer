package otherUtil

import (
	"strings"
)

type MapStrValSorter []mapStrValSorterItem
type MapStrValSorterWuZhuJun []mapStrValSorterItem
type mapStrValSorterItem struct {
	Key string
	Val string
}

func NewMapStrValSorter(m map[string]string) MapStrValSorter {
	ms := make(MapStrValSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, mapStrValSorterItem{k, v})
	}
	return ms
}
func NewMapStrValSorterWuZhuJun(m map[string]string) MapStrValSorterWuZhuJun {
	ms := make(MapStrValSorterWuZhuJun, 0, len(m))
	for k, v := range m {
		ms = append(ms, mapStrValSorterItem{k, v})
	}
	return ms
}
func (ms MapStrValSorter) Len() int {
	return len(ms)
}
func (ms MapStrValSorter) Less(i, j int) bool {
	return strings.ToLower(ms[i].Val) < strings.ToLower(ms[j].Val)
}
func (ms MapStrValSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms MapStrValSorterWuZhuJun) Len() int {
	return len(ms)
}
func (ms MapStrValSorterWuZhuJun) Less(i, j int) bool {
	return ms[i].Val < ms[j].Val
}
func (ms MapStrValSorterWuZhuJun) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
