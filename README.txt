
Jar Serialize for cookies

just add a file "json.go"
add:
  func (j *Jar) MarshalJSON() ([]byte, error)
  func Clone(o *Options, json []byte) (*Jar, error)

example:
  cookiejarOptions := cookiejar.Options{
    PublicSuffixList: publicsuffix.List,
  }
  jar, _ := cookiejar.New(&cookiejarOptions)
  client := &http.Client{Jar: jar}

  client.Get("http://bing.com")

  buf, err := json.Marshal(client.Jar)
  // ioutil.WriteFile("cookies.json", buf, 0644)
  jarClone, err := cookiejar.Clone(&cookiejarOptions, buf)

License:
  as the source code of GoLang
