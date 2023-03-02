package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"go-test/pkg/logger"
	"log"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

type Test struct {
	a int32
	b float32
	c string
}

type A struct {
	Test
	d string
}

type Dog interface {
	bark() string
}

type Huskie struct {
	age  int
	name string
}

func (h Huskie) bark() string {
	return "Ha!!"
}

func (h Huskie) jump() {
	fmt.Println("Huskie Jump")
}

type Sibainu struct {
	name string
	Dog
}

type RealDog struct{}

func (s Sibainu) bark() string {
	return s.Dog.bark()
}
func (r RealDog) bark() string {
	return "Si"
}

func main() {
	// testSlice()
	// testSliceWhile()
	// testStruct()
	// testMem()
	// testInterface()
	// normalLog()
	// fmt.Println("---------------")
	// zapLog()
	// wrappedLogger()
	// r := guidTrans("0b8be8cfb61b485c9f2c8d3eddf538e6")
	// r := guidTrans("95ac48c5-b966-47b2-8323-5e0d82c60ca9")
	// goRoutineCh()
	// testStorage()
	// slicePointerTest()
	// fmt.Println(int(math.Ceil(float64(11) / float64(2))))

	// myTT()
	bit()
}

func testSlice() {
	s := []int{1, 2, 3}
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
	fmt.Println(s, len(s), cap(s))

	if a := s[1]; a == 2 {
		fmt.Printf("a is %d\n", a)
	}
}

func testSliceWhile() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Start Loop")
	batchNum := 3
	for {
		// if len(s) == 0 {
		// 	break
		// }
		if len(s) > batchNum {
			fmt.Printf("sssss: %v\n", s[:batchNum])
			s = s[batchNum:]
		} else {
			fmt.Printf("sssss: %v\n", s)
			break
		}
		// fmt.Printf("rrrrr: %v\n", s)
	}
	fmt.Println("Loop over")

}

func testStruct() {
	// t1 := new(Test)
	// fmt.Println(t1)
	t1 := Test{a: 1}
	fmt.Printf("t1: %v\n", t1)

	t2 := &Test{a: 1}
	fmt.Printf("t2: %v\n", t2)

	structInput(t1, t2)
	fmt.Println(t1, t2)

	a := A{Test{1, 2, "a"}, "b"}
	fmt.Println(a)
}

func structInput(i1 Test, i2 *Test) {
	i1.a = 3
	i2.a = 3
	fmt.Println(i1, i2)
}

func testMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

func testInterface() {
	// var d Dog = Huskie{3, "Haha"}
	// fmt.Println("dog", d, d.bark())
	// (d.(Huskie)).jump()

	var d Dog = Sibainu{"Siba", RealDog{}}
	fmt.Println("dog", d, d.bark())
	// (d.(Huskie)).jump()
}

func normalLog() {
	log.Println("normal log")
}

func zapLog() {
	// baseLogger, _ := zap.NewProduction()

	// defer baseLogger.Sync()
	// logger := baseLogger.Sugar()

	f, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	config := zap.NewDevelopmentConfig()
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.AddSync(os.Stderr),
			zap.DebugLevel,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config.EncoderConfig),
			zapcore.AddSync(f),
			zap.DebugLevel,
		),
	)
	baseLogger := zap.New(core)
	defer baseLogger.Sync()

	logger := baseLogger.Sugar()
	// logger.Infof("Check order. id: %d, name: %v", 123, "Fruit")
	// logger.Infow("Check order.", "id", 123, "name", "Fruit")
	logger.Infoln("zap log")
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		// zap.String("url", url),
		// zap.Int("attempt", 3),
		// zap.Duration("backoff", time.Second),
		"attempt ", 3,
		"backoff ", time.Second,
	)
}

func wrappedLogger() {
	l := logger.New()
	l.Infof("my logger!!!!")
}

func myTask(id int, isRand bool) {
	fmt.Printf("[Task %d] START!\n", id)
	for i := 0; i < 10; i++ {
		randNum := 1
		if isRand {
			dice, _ := rand.Int(rand.Reader, big.NewInt(3))
			randNum = int(dice.Int64()) + 1
		}
		time.Sleep(time.Duration(randNum) * time.Second)
		fmt.Printf("[Task %d] Proc: %d\n", id, i+1)
	}
	fmt.Printf("[Task %d] END!\n", id)
}

func goRoutineNoWait() {
	// old way, no waiting
	fmt.Println("Test Start.")
	go myTask(100, true)
	go myTask(200, false)
	time.Sleep(15 * time.Second)
	fmt.Println("Test Done.")
}

func goRoutineWQ() {
	// use wait queue and wrapped to a task queue

	fmt.Println("Test Start.")
	var wq WQueue = &WaitQueue{}
	wq.Init()
	wq.Add(myTask, false)
	wq.Add(myTask, false)
	wq.Run()
	fmt.Println("Test Done.")
}

func goRoutineCh() {
	// use channel

	fmt.Println("Test Start.")
	c := make(chan string)

	// anonymous goroutine
	go func(c chan string) {
		time.Sleep(3 * time.Second)
		fmt.Println("Hello " + <-c + "!")
		c <- "hh"
		// fmt.Println("Hello !")
	}(c)

	// c <- "John"
	<-c
	fmt.Println("Test Done.")
}

func guidTrans(id string) string {
	result := ""
	if strings.Contains(id, "-") {
		result = strings.ReplaceAll(id, "-", "")
	} else {
		// 95ac48c5-b966-47b2-8323-5e0d82c60ca9
		result = id[:8] + "-" + id[8:12] + "-" + id[12:16] + "-" + id[16:20] + "-" + id[20:]
	}

	fmt.Println(result)
	return result
}

func testStorage() {
	ctx := context.Background()
	bucket1, err := blob.OpenBucket(ctx, "mem://")
	if err != nil {
		panic("aaaa")
	}
	defer bucket1.Close()
	err = bucket1.WriteAll(ctx, "my-key", []byte("hello world"), nil)
	if err != nil {
		panic("rrrr")
	}
	data, err := bucket1.ReadAll(ctx, "my-key")
	if err != nil {
		panic("rrrr")
	}
	fmt.Println(string(data))

	bucket2, err := blob.OpenBucket(ctx, "file://./data")
	if err != nil {
		fmt.Println(err)
		panic("bbb")
	}
	defer bucket2.Close()
	b, err := bucket2.ReadAll(ctx, "test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	w, err := bucket2.NewWriter(ctx, "foo.txt", nil)
	if err != nil {
		panic("rrrr")
	}
	w.Write([]byte("it's a test\n"))
	w.Close()
}

func slicePointerTest() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		// sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
		// fmt.Printf("hhhhh111 - %v", sh)
	}(slice)
	// func(slice *[]string){
	//     *slice = append(*slice, "a")
	//     (*slice)[0]="b";
	//     (*slice)[1]="b";
	//     fmt.Print(*slice)
	// }(&slice)
	// sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Print(slice)
	// fmt.Printf("hhhhh000 - %v", sh)
}

func myTT() {
	str := "x-txsp-send-time"
	ts, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("ttttttt: %d, %d\n", ts, time.Now().Unix() - ts)
}

func bit() {
	a := 1
	b := 2
	c := 4
	d := a | b
	fmt.Println(d)
	fmt.Println(d & c)
	fmt.Println(d & a)
	fmt.Println(d & b)
}