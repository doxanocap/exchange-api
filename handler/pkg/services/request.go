package services

//
//func (dsp *Dispatcher) Get() []byte {
//	resp, err := http.Get(dsp.Url)
//	if err != nil {
//		log.Fatalf("An Error Occured %v", err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	return body
//}
//
//func (dsp *Dispatcher) Post(reqBody []byte) []byte {
//	resp, err := http.Post(
//		dsp.Url,
//		"application/json",
//		bytes.NewBufferString(string(reqBody)))
//	if err != nil {
//		log.Fatalf("An Error Occured %v", err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	return body
//}
