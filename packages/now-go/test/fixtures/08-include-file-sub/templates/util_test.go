package templates

import(
 "testing"
)

func TestReadFileContent(t *testing.T) {
 want := "foobar from file"
 got, err := FileContent()

 if err != nil {
   t.Errorf(err.Error())
 }
 if want != got {
  t.Errorf("want \"%s\" but is not what I got \"%s\"", want, got)
 }
}
