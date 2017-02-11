package cookiejar_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/yuanlixg/cookiejar"
	"golang.org/x/net/publicsuffix"
)

func ExampleClone_nil_nil() {
	jar, err := cookiejar.Clone(nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf, err := json.MarshalIndent(jar, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
	// Output:
	// {
	//   "Entries": {},
	//   "NextSeqNum": 0
	// }
}

func ExampleClone_o_nil() {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.Clone(&cookiejarOptions, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf, err := json.MarshalIndent(jar, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
	// Output:
	// {
	//   "Entries": {},
	//   "NextSeqNum": 0
	// }
}

func ExampleClone() {
	client := &http.Client{}

	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&cookiejarOptions)

	// apply it to default client, work ?
	client.Jar = jar

	client.Get("https://baidu.com")

	buf, err := json.Marshal(client.Jar)
	if err != nil {
		fmt.Println(err)
		return
	}

	jarClone, err := cookiejar.Clone(&cookiejarOptions, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf, err = json.MarshalIndent(jarClone, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ss := strings.Split(string(buf), "\n")
	// only `head -3` && `tail -1`, intermediate is unpredictable
	if len(ss) >= 4 {
		ss = append(ss[:3], ss[len(ss)-1])
	}
	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i])
	}
	// Output:
	// {
	//   "Entries": {
	//     "baidu.com": {
	// }
}
