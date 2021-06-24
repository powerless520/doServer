package csvUtil

import "bytes"

type CsvUtil struct {
}

func (self CsvUtil) ExportCsvFile() {

}

func (self CsvUtil) ExportCsvBuffer(head []string, data [][]string) (buffer_str *bytes.Buffer) {
	head_len := len(head)
	buffer_str = bytes.NewBufferString("\xEF\xBB\xBF")

	if head_len > 0 {
		for index, li := range head {
			ds := ","
			if index == head_len-1 {
				ds = "\n"
			}
			buffer_str.WriteString(li + ds)
		}
	}
	if len(data) <= 0 {
		return
	}
	for _, list := range data {
		list_len := len(list)
		for index, li := range list {
			ds := ","
			if index == list_len-1 {
				ds = "\n"
			}
			buffer_str.WriteString(li + ds)
		}
	}
	return buffer_str
}

func ExportNetLine(sep string, head []string, data []map[string]string) (buffer_str *bytes.Buffer) {
	head_len := len(head)
	buffer_str = bytes.NewBufferString("\xEF\xBB\xBF")

	if head_len > 0 {
		for index, li := range head {
			ds := sep
			if index == head_len-1 {
				ds = "\n"
			}
			buffer_str.WriteString(li + ds)
		}
	}
	if len(data) < 1 {
		return
	}
	for _, list := range data {
		list_len := len(list)
		i := 0
		for _, key := range head {
			ds := sep
			if i == list_len-1 {
				ds = "\n"
			}
			buffer_str.WriteString(list[key] + ds)
			i++
		}
	}
	return buffer_str
}
