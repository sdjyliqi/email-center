package center

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DusanKasan/parsemail"
	"io/ioutil"
	"testing"
)

func TestTaskCenter(t *testing.T) {
	//TaskCenter()
	contents, err := ioutil.ReadFile("D:\\gowork\\src\\email-center\\data\\发票类\\异常\\20210713000936__yqsidg@ukkl.com__202107130009361898199814091494693.eml")
	//reader.Read(contents)
	emailReader := bytes.NewReader(contents)
	fmt.Println(emailReader)
	email, err := parsemail.Parse(emailReader) // returns Email struct and error
	fmt.Println("====", email.Subject)
	b, err := json.Marshal(email)
	t.Log(email.Subject, string(b), err)
}
