//package main
//
//import (
//	"bytes"
//	"encoding/hex"
//	"fmt"
//	"io"
//	"math/rand"
//	"net"
//	"time"
//
//	"github.com/bradfitz/http2"
//	"github.com/tatsuhiro-t/go-http2-hpack"
//)
//
//func main() { //52
//	srcip := "192.168.2.10"
//	srcport := 34480
//	dstip := "192.168.2.2"
//	dstport := 29509
//	target := "192.168.2.2:29509"
//	data := `{"supiOrSuci":"suci-0-208-93-0000-0-0-0000000003","servingNetworkName":"5G:mnc093.mcc208.3gppnetwork.org"}`
//	//调用Free5gcFuzz进行发包
//	var headerkey [9]string
//	var headervalue [9]string
//	headerkey = [9]string{":authority", ":method", ":path", ":scheme", "content-type", "accept", "user-agent", "content-length", "accpet-encoding"}
//	headervalue = [9]string{target, "POST", "/nausf-auth/v1/ue-authentications", "http", "application/json", "application/3gppHal+json,application/problem+json", "OpenAPI-Generator/1.0.0/go", "106", "gzip"}
//	Free5gcFuzz(srcip, srcport, dstip, dstport, data, headerkey[:], headervalue[:])
//}
//
////func Free5gcFuzz(srcip string, srcport int, dstip string, dstport int, data string, headerkey []string, headervalue []string) error {
////	conn := Dial(srcip, srcport, dstip, dstport)
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
////func Dial(srcip string, srcport int, dstip string, dstport int) net.Conn {
////	var localaddr net.TCPAddr
////	var remoteaddr net.TCPAddr
////	localaddr.IP = net.ParseIP(srcip)
////	localaddr.Port = srcport
////	remoteaddr.IP = net.ParseIP(dstip)
////	remoteaddr.Port = dstport
////	if localaddr.IP == nil || remoteaddr.IP == nil {
////		fmt.Println("error")
////	}
////	conn, err := net.DialTCP("tcp", &localaddr, &remoteaddr)
////	if err != nil {
////		panic(err)
////	}
////	return conn
//}
//
////变异算法
//
///*
//hpack
//
//const (
//	FrameData         FrameType = 0x0
//	FrameHeaders      FrameType = 0x1
//	FramePriority     FrameType = 0x2
//	FrameRSTStream    FrameType = 0x3
//	FrameSettings     FrameType = 0x4
//	FramePushPromise  FrameType = 0x5
//	FramePing         FrameType = 0x6
//	FrameGoAway       FrameType = 0x7
//	FrameWindowUpdate FrameType = 0x8
//	FrameContinuation FrameType = 0x9
//)
//
//const (
//	// Data Frame
//	FlagDataEndStream Flags = 0x1
//	FlagDataPadded    Flags = 0x8
//
//	// Headers Frame
//	FlagHeadersEndStream  Flags = 0x1
//	FlagHeadersEndHeaders Flags = 0x4
//	FlagHeadersPadded     Flags = 0x8
//	FlagHeadersPriority   Flags = 0x20
//
//	// Settings Frame
//	FlagSettingsAck Flags = 0x1
//
//	// Ping Frame
//	FlagPingAck Flags = 0x1
//
//	// Continuation Frame
//	FlagContinuationEndHeaders Flags = 0x4
//
//	FlagPushPromiseEndHeaders Flags = 0x4
//	FlagPushPromisePadded     Flags = 0x8
//)
//
//*/
//
//// MutationRate is the rate of mutation
//var MutationRate = 0.005
//
//// PopSize is the size of the population
//var PopSize = 500
//
//func generate(source string, resString chan string) {
//	start := time.Now()
//	rand.Seed(time.Now().UTC().UnixNano())
//
//	target := []byte(source)
//	population := createPopulation(target)
//
//	found := false
//	generation := 0
//	for !found {
//		generation++
//		bestOrganism := getBest(population)
//		// fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestOrganism.DNA), bestOrganism.Fitness)
//		// fmt.Println(string(bestOrganism.DNA))
//		resString <- string(bestOrganism.DNA)
//
//		if bytes.Compare(bestOrganism.DNA, target) == 0 {
//			found = true
//		} else {
//			maxFitness := bestOrganism.Fitness
//			pool := createPool(population, target, maxFitness)
//			population = naturalSelection(pool, population, target)
//		}
//
//	}
//	elapsed := time.Since(start)
//	fmt.Printf("\nTime taken: %s\n", elapsed)
//}
//
//// Organism for this genetic algorithm
//type Organism struct {
//	DNA     []byte
//	Fitness float64
//}
//
//// creates a Organism
//func createOrganism(target []byte) (organism Organism) {
//	ba := make([]byte, len(target))
//	for i := 0; i < len(target); i++ {
//		ba[i] = byte(rand.Intn(95) + 32)
//	}
//	organism = Organism{
//		DNA:     ba,
//		Fitness: 0,
//	}
//	organism.calcFitness(target)
//	return
//}
//
//// creates the initial population
//func createPopulation(target []byte) (population []Organism) {
//	population = make([]Organism, PopSize)
//	for i := 0; i < PopSize; i++ {
//		population[i] = createOrganism(target)
//	}
//	return
//}
//
//// calculates the fitness of the Organism
//func (d *Organism) calcFitness(target []byte) {
//	score := 0
//	for i := 0; i < len(d.DNA); i++ {
//		if d.DNA[i] == target[i] {
//			score++
//		}
//	}
//	d.Fitness = float64(score) / float64(len(d.DNA))
//	return
//}
//
//// create the breeding pool that creates the next generation
//func createPool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
//	pool = make([]Organism, 0)
//	// create a pool for next generation
//	for i := 0; i < len(population); i++ {
//		population[i].calcFitness(target)
//		num := int((population[i].Fitness / maxFitness) * 100)
//		for n := 0; n < num; n++ {
//			pool = append(pool, population[i])
//		}
//	}
//	return
//}
//
//// perform natural selection to create the next generation
//func naturalSelection(pool []Organism, population []Organism, target []byte) []Organism {
//	next := make([]Organism, len(population))
//
//	for i := 0; i < len(population); i++ {
//		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
//		a := pool[r1]
//		b := pool[r2]
//
//		child := crossover(a, b)
//		child.mutate()
//		child.calcFitness(target)
//
//		next[i] = child
//	}
//	return next
//}
//
//// crosses over 2 Organisms
//func crossover(d1 Organism, d2 Organism) Organism {
//	child := Organism{
//		DNA:     make([]byte, len(d1.DNA)),
//		Fitness: 0,
//	}
//	mid := rand.Intn(len(d1.DNA))
//	for i := 0; i < len(d1.DNA); i++ {
//		if i > mid {
//			child.DNA[i] = d1.DNA[i]
//		} else {
//			child.DNA[i] = d2.DNA[i]
//		}
//
//	}
//	return child
//}
//
//// mutate the Organism
//func (d *Organism) mutate() {
//	for i := 0; i < len(d.DNA); i++ {
//		if rand.Float64() < MutationRate {
//			d.DNA[i] = byte(rand.Intn(95) + 32)
//		}
//	}
//}
//
//// Get the best organism
//func getBest(population []Organism) Organism {
//	best := 0.0
//	index := 0
//	for i := 0; i < len(population); i++ {
//		if population[i].Fitness > best {
//			index = i
//			best = population[i].Fitness
//		}
//	}
//	return population[index]
//}
