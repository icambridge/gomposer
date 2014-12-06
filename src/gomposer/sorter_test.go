package gomposer


import (
    "testing"
    "reflect"
    "sort"
)

func Test_Version_Sort(t *testing.T) {

    input := []string{"1.2.1","1.2.9","1.2.13","1.2.4"}
    expected := []string{"1.2.1","1.2.4","1.2.9", "1.2.13"}


    sort.Sort(VersionSlice(input))


    if !reflect.DeepEqual(input, expected) {
        t.Errorf("Response body = %v, expected %v", input, expected)
    }

}
