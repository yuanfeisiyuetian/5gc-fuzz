//
//package main
//
//func main(){
//	target := "192.168.2.7:29503"
//	data := `"{` + `"servingNetworkName":` + `"5G:mnc093.mcc208.3gppnetwork.org",` + `"ausfInstanceId":` + `"0b75b04a-6374-47da-beaf-7815b2aeed67"` + `}"`
//	headerParam := make(map[string]string)
//	headerParam[":authority"] = target
//	headerParam[":method"] = "POST"
//	headerParam[":path"] = "/nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000017/security-information/generate-auth-data"
//	headerParam[":scheme"] = "http"
//	headerParam["accept"] = "application/3gppHal+json,application/problem+json"
//	headerParam["user-agent"] = "OpenAPI-Generator/1.0.0/go"
//	headerParam["content-type"] = "application/json"
//	headerParam["accpet-encoding"] = "gzip"
//	var headerkey [8]string
//	var headervalue [8]string
//	headerkey = [8]string{":authority", ":method", ":path", ":scheme", "accept", "user-agent", "content-type", "accpet-encoding"}
//	headervalue = [8]string{target, "POST", "/nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000017/security-information/generate-auth-data", "http", "application/3gppHal+json,application/problem+json", "OpenAPI-Generator/1.0.0/go", "application/json", "gzip"}
//	Free5gcFuzz(target, data, headerkey[:], headervalue[:])
//}
////func Free5gcFuzz(target, data string, headerkey []string, headervalue []string) error {
////	conn := Dial(target)
////	io.WriteString(conn, http2.ClientPreface)
////
////	framer := http2.NewFramer(conn, conn)
////	//setting
////	settingData, _ := hex.DecodeString("000200000000000400400000000600a00000")
////	framer.WriteRawFrame(http2.FrameType(4), http2.Flags(0), 0, settingData)
////	//window_update
////	windowData, _ := hex.DecodeString("40000000")
////	framer.WriteRawFrame(http2.FrameType(8), http2.Flags(0), 0, windowData)
////
////	//header
////	headers := []*hpack.Header{}
////	for i := 0; i < len(headerkey); i++ {
////		headers = append(headers, hpack.NewHeader(headerkey[i], headervalue[i], false))
////	}
////	// for key, value := range headerParam {
////	// 	headers = append(headers, hpack.NewHeader(key, value, false))
////	// }
////	encoded := &bytes.Buffer{}
////	enc := hpack.NewEncoder(hpack.DEFAULT_HEADER_TABLE_SIZE)
////	enc.Encode(encoded, headers)
////	framer.WriteRawFrame(http2.FrameType(1), http2.Flags(4), 3, encoded.Bytes())
////
////	//发送一个data
////	if len(data) != 0 {
////		framer.WriteRawFrame(http2.FrameType(0), http2.Flags(1), 3, []byte(data))
////	}
////	return nil
////}
////
////func Dial(host string) net.Conn {
////	conn, err := net.Dial("tcp", host)
////	if err != nil {
////		panic(err)
////	}
////	return conn
////}
//
