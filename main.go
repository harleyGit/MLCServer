/*
 * @title: ε­¦δΉ ζ΅θ―
 * @Author: gang.huang
 * @Date: 2021-08-20 22:51:45
 * @LastEditTime: 2021-10-08 10:08:55
 * @FilePath: /GoDemo/main.go
 */

package main

import (
	"bytes"
	"container/list"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"reflect"
	"runtime"
	"sync/atomic"
)

func main() {
	fmt.Printf("<=============== π π π ===============> \n\n")

	//fmt.Println("π welcome to Go Lang! π ")

	// init_array()
	// point_flag()
	//chargeValue()
	//pointTest1()
	//lowVar()
	// section_test()
	//multiplication_table()
	// list_delete()
	// paramTranslate()
	// testAnonymousFunction()
	// testAnnoymousFunction1()
	// testFuncImplInterface()
	// testFuncImplInterface1()
	// testClosure1_1()
	// testVariableParameters("hammer", " mom", " and", " hawk")
	// testDef()
	// testError("π£ β ιθ――ζ΅θ―")
	// fmt.Println(testError1(1, 0))
	// testError2()
	// panic("π£ β  ε΄©ζΊ ")
	// testPanic()
	// testPanic1()

	/*
		var version int = 1
		cmd := testStruct(
			"version",
			&version,
			"show version",
		)
		fmt.Println("εε°εη»ζδ½ cmdοΌ", cmd)
	*/
	// testFuncMethod()
	// testFuncMethod1()
	// testStruct1()
	// testStruct2()
	// testStruct3()
	// testInterface1()
	// testInterface2()
	// testInterface3()
	// testInterface4()
	// testInterface5()
	// testInterface6()

	// testGoroutine1()

	testLock1()

	fmt.Printf("\n\n<=============== π π π ===============> ")

}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
/**
 * @description: εΊεε·ηζε¨
 */

var (
	// εΊεε·
	seq int64
)

func testLock1() {
	//ηζ10δΈͺεΉΆεεΊεε·
	for i := 0; i < 10; i++ {
		go GenID()
	}
	fmt.Println(GenID())
}
func GenID() int64 {
	// ε°θ―εε­ηε’ε εΊεε·
	// δ½Ώη¨εε­ζδ½ε½ζ°atomic.Add Int64()ε―Ήseq()ε½ζ°ε 1ζδ½γ
	// δΈθΏθΏιζζζ²‘ζδ½Ώη¨atomic.Add Int64()ηθΏεεΌδ½δΈΊGen ID()ε½ζ°ηθΏεεΌοΌε ζ­€δΌι ζδΈδΈͺη«ζι?ι’
	// atomic.AddInt64(&seq, 1)
	// return seq

	// ε°θ―εε­ηε’ε εΊεε·
	return atomic.AddInt64(&seq, 1)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func testGoroutine1() {
	fmt.Println("CPUηΊΏη¨ζ°ι: ", runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU())
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
/**
 * @description: ηΆζζ₯ε£
 */
type State interface {
	// θ·εηΆζεε­
	Name() string
	// θ―₯ηΆζζ―ε¦εθ?ΈεηΆζθ½¬η§»
	EnableSameTransit() bool
	// εεΊηΆζεΌε§ζΆ
	OnBegin()
	// εεΊηΆζη»ζζΆ
	OnEnd()
	// ε€ζ­θ½ε¦θ½¬η§»ε°ζδΈͺηΆζ
	CanTransitTo(name string) bool
}

// δ»ηΆζε?δΎθ·εηΆζε
func StateName(s State) string {
	if s == nil {
		return "none"
	}
	// δ½Ώη¨εε°θ·εηΆζηεη§°
	return reflect.TypeOf(s).Elem().Name()
}

/**
 * @description: ηΆζεΊζ¬δΏ‘ζ―
 */
// ηΆζηεΊη‘δΏ‘ζ―ει»θ?€ε?η°
type StateInfo struct {
	// ηΆζε
	name string
}

// ηΆζε
func (s *StateInfo) Name() string {
	return s.name
}

// ζδΎη»ει¨θ?Ύη½?εε­
// setName()ζΉζ³ηι¦ε­ζ―ε°εοΌθ‘¨η€ΊθΏδΈͺζΉζ³εͺθ½ε¨εεεθ’«θ°η¨γ
// θΏιζδ»¬εΈζsetName()δΈθ½θ’«δ½Ώη¨θε¨ηΆζεε§εειζδΏ?ζΉεη§°οΌθζ―ιθΏει’ζε°ηηΆζη?‘ηε¨θͺε¨θ΅εΌ
func (s *StateInfo) setName(name string) {
	s.name = name
}

// εθ?ΈεηΆζθ½¬η§»
func (s *StateInfo) EnableSameTransit() bool {
	return false
}

// ι»θ?€ε°ηΆζεΌε―ζΆε?η°
func (s *StateInfo) OnBegin() {

}

// ι»θ?€ε°ηΆζη»ζζΆε?η°
func (s *StateInfo) OnEnd() {}

// ι»θ?€ε―δ»₯θ½¬η§»ε°δ»»δ½ηΆζ
func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}

/**
 * @description: ηΆζη?‘ηε¨
 */

type StateManager struct {
	// ε·²η»ζ·»ε ηηΆζ
	// ε£°ζδΈδΈͺδ»₯ηΆζεδΈΊι?οΌδ»₯Stateζ₯ε£δΈΊεΌηmap
	stateByName map[string]State
	// ηΆζζΉεζΆηεθ°
	OnChange func(from, to State)
	// ε½εηΆζ
	curr State
}

// ζ·»ε δΈδΈͺηΆζε°η?‘ηε¨δΈ­
func (sm *StateManager) Add(s State) {
	// θ·εηΆζηεη§°
	name := StateName(s)
	// ε°sθ½¬ζ’δΈΊθ½θ?Ύη½?εε­ηζ₯ε£οΌηΆεθ°η¨θ―₯ζ₯ε£
	// ε°sοΌStateζ₯ε£οΌιθΏη±»εζ­θ¨θ½¬ζ’δΈΊεΈ¦ζset Name()ζΉζ³(name string)ηζ₯ε£γ
	// ζ₯ηθ°η¨θΏδΈͺζ₯ε£ηset Name()ζΉζ³θ?Ύη½?ηΆζηεη§°γδ½Ώη¨θ―₯ζΉζ³ε―δ»₯εΏ«ιθ°η¨δΈδΈͺζ₯ε£ε?η°ηεΆδ»ζΉζ³
	s.(interface {
		setName(name string)
	}).setName(name)
	// ζ Ήζ?ηΆζεθ·εε·²η»ζ·»ε ηηΆζοΌζ£ζ₯θ―₯ηΆζζ―ε¦ε­ε¨
	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}
	// ζ Ήζ?εε­δΏε­ε°mapδΈ­
	sm.stateByName[name] = s
}

// ζ Ήζ?εε­θ·εζε?ηΆζ
func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

// εε§εηΆζη?‘ηε¨
func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

/**
 * @description: ε¨ηΆζι΄θ½¬η§»
 */
// ηΆζζ²‘ζζΎε°ηιθ――
var ErrStateNotFound = errors.New("state not found")

// η¦ζ­’ε¨εηΆζι΄θ½¬η§»
var ErrForbidSameStateTransit = errors.New("forbid same state transit")

// δΈθ½θ½¬η§»ε°ζε?ηΆζ
var ErrCannotTransitToState = errors.New("cannot transit to state")

// θ·εε½εηηΆζ
func (sm *StateManager) CurrState() State {
	return sm.curr
}

// ε½εηΆζθ½ε¦θ½¬η§»ε°η?ζ ηΆζ
func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	// ηΈεηηΆζδΈη¨θ½¬ζ’
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	// δ½Ώη¨ε½εηΆζοΌζ£ζ₯θ½ε¦θ½¬η§»ε°ζε?εε­ηηΆζ
	return sm.curr.CanTransitTo(name)
}

// θ½¬η§»ε°ζε?ηΆζ
func (sm *StateManager) Transit(name string) error {
	// θ·εη?ζ ηΆζ
	next := sm.Get(name)
	// η?ζ δΈε­ε¨
	if next == nil {
		return ErrStateNotFound
	}
	// θ?°ε½θ½¬η§»εηηΆζ
	pre := sm.curr
	// ε½εζηΆζ
	if sm.curr != nil {
		// ηΈεηηΆζδΈη¨θ½¬ζ’
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}
		// δΈθ½θ½¬η§»ε°η?ζ ηΆζ
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		// η»ζε½εηΆζ
		sm.curr.OnEnd()
	}
	// ε°ε½εηΆζεζ’δΈΊθ¦θ½¬η§»ε°ηη?ζ ηΆζ
	sm.curr = next
	// θ°η¨ζ°ηΆζηεΌε§
	sm.curr.OnBegin()
	// ιη₯εθ°
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}
	return nil
}

/**
 * @description: θͺε?δΉηΆζε?η°ηΆζζ₯ε£
 */

// ι²η½?ηΆζ
type IdleState struct {
	StateInfo // δ½Ώη¨State Infoε?η°εΊη‘ζ₯ε£
}

// ιζ°ε?η°ηΆζεΌε§
func (i *IdleState) OnBegin() {
	fmt.Println("Idle State begin")
}

// ιζ°ε?η°ηΆζη»ζ
func (i *IdleState) OnEnd() {
	fmt.Println("Idle State end")
}

// η§»ε¨ηΆζ
type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("Move State begin")
}

// εθ?Έη§»ε¨ηΆζδΊηΈθ½¬ζ’
func (m *MoveState) EnableSameTransit() bool {
	return true
}

// θ·³θ·ηΆζ
type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("Jump State begin")
} // θ·³θ·ηΆζδΈθ½θ½¬η§»ε°η§»ε¨ηΆζ
func (j *JumpState) CanTransitTo(name string) bool {
	return name != "Move State"
}

func testInterface6() {
	// ε?δΎεδΈδΈͺηΆζη?‘ηε¨
	sm := NewStateManager()
	// εεΊηΆζθ½¬η§»ηιη₯
	sm.OnChange = func(from, to State) {
		// ζε°ηΆζθ½¬η§»ηζ΅ε
		fmt.Printf("%s ---> %s\n\n", StateName(from), StateName(to))
	}
	// ζ·»ε 3δΈͺηΆζ
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))
	// ε¨δΈεηΆζι΄θ½¬η§»
	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")
}

// ε°θ£θ½¬η§»ηΆζεθΎεΊζ₯εΏ
func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}
func testInterface5() {
	printType(1024)
	printType("pig")
	printType(true)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// ε?δΉι£θ‘ε¨η©ζ₯ε£
type Flyer interface {
	Fly()
}

// ε?δΉθ‘θ΅°ε¨η©ζ₯ε£
type Walker interface {
	Walk()
}

// ε?δΉιΈη±»
type bird struct{}

// ε?η°ι£θ‘ε¨η©ζ₯ε£
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// δΈΊιΈζ·»ε Walk()ζΉζ³οΌε?η°θ‘θ΅°ε¨η©ζ₯ε£
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// ε?δΉηͺ
type pig struct{}

// δΈΊηͺζ·»ε Walk()ζΉζ³οΌε?η°θ‘θ΅°ε¨η©ζ₯ε£
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}
func testInterface3() {
	// εε»Ίε¨η©ηεε­ε°ε?δΎηζ ε°
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// ιεζ ε°
	for name, obj := range animals {
		// δ½Ώη¨η±»εζ­θ¨θ·εΎfοΌη±»εδΈΊFlyerεis Flyerηζ­θ¨ζεηε€ε?
		f, isFlyer := obj.(Flyer)
		// ε€ζ­ε―Ήθ±‘ζ―ε¦δΈΊθ‘θ΅°ε¨η©
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s is Flyer: %v is Walker: %v\n", name, isFlyer, isWalker)
		// ε¦ζζ―ι£θ‘ε¨η©εθ°η¨ι£θ‘ε¨η©ζ₯ε£
		if isFlyer {
			f.Fly()
		}
		// ε¦ζζ―θ‘θ΅°ε¨η©εθ°η¨θ‘θ΅°ε¨η©ζ₯ε£
		if isWalker {
			w.Walk()
		}
	}
}

func testInterface4() {
	p1 := new(pig)

	var a Walker = p1
	p2 := a.(*pig)

	fmt.Printf("p1=%p p2=%p", p1, p2)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
/**
 * @description: ζ₯εΏε―Ήε€ζ₯ε£
 */
// ε£°ζζ₯εΏεε₯ε¨ζ₯ε£
type LogWriter interface {
	Write(data interface{}) error
}

// ζ₯εΏε¨
type Logger struct {
	// θΏδΈͺζ₯εΏε¨η¨ε°ηζ₯εΏεε₯ε¨
	writerList []LogWriter
}

// ζ³¨εδΈδΈͺζ₯εΏεε₯ε¨
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// ε°δΈδΈͺdataη±»εηζ°ζ?εε₯ζ₯εΏ
func (l *Logger) Log(data interface{}) {
	// ιεζζζ³¨εηεε₯ε¨
	for _, writer := range l.writerList {
		// ε°ζ₯εΏθΎεΊε°ζ―δΈδΈͺεε₯ε¨δΈ­
		writer.Write(data)
	}
}

// εε»Ίζ₯εΏε¨ηε?δΎοΏΌ
func NewLogger() *Logger {
	return &Logger{}
}

// ε£°ζζδ»Άεε₯ε¨
/**
 * @description: ζδ»Άεε₯ε¨
 * ζδ»Άεε₯ε¨ηεθ½ζ―ζ Ήζ?δΈδΈͺζδ»Άεεε»Ίζ₯εΏζδ»ΆοΌfile WriterηSet FileζΉζ³οΌγ
 * ε¨ζζ₯εΏεε₯ζΆοΌε°ζ₯εΏεε₯ζδ»ΆδΈ­γ
 */
// ε£°ζζδ»Άεε₯ε¨οΌε¨η»ζδ½δΈ­δΏε­δΈδΈͺζδ»Άε₯ζοΌδ»₯ζΉδΎΏζ―ζ¬‘εε₯ζΆζδ½
type fileWriter struct {
	file *os.File
}

// θ?Ύη½?ζδ»Άεε₯ε¨εε₯ηζδ»Άε
func (f *fileWriter) SetFile(filename string) (err error) {
	// ε¦ζζδ»Άε·²η»ζεΌοΌε³ι­εδΈδΈͺζδ»Ά
	// θθε°SetFile()ζΉζ³ε―δ»₯θ’«ε€ζ¬‘θ°η¨οΌε½ζ°ε―ιε₯ζ§οΌ
	// εθ?ΎδΉεε·²η»θ°η¨θΏSet File()εεζ¬‘θ°η¨οΌζ­€ζΆηf.fileδΈδΈΊη©ΊοΌε°±ιθ¦ε³ι­δΉεηζδ»ΆοΌιζ°εε»Ίζ°ηζδ»Άγ
	if f.file != nil {
		f.file.Close()
	}
	// εε»ΊδΈδΈͺζδ»ΆεΉΆδΏε­ζδ»Άε₯ζ
	f.file, err = os.Create(filename)
	// ε¦ζεε»ΊηθΏη¨εΊη°ιθ――οΌεθΏειθ――
	return err
}

// ε?η°LogWriterηWrite()ζΉζ³
func (f *fileWriter) Write(data interface{}) error {
	// ε¦ζζδ»Άζ²‘ζεε€ε₯½οΌζδ»Άε₯ζδΈΊnil
	// ζ­€ζΆδ½Ώη¨errorsεηNew()ε½ζ°θΏεδΈδΈͺιθ――ε―Ήθ±‘οΌεε«δΈδΈͺε­η¬¦δΈ²βfile not createdβ
	if f.file == nil {
		// ζ₯εΏζδ»Άζ²‘ζεε€ε₯½
		return errors.New("file not created")
	}
	// ε°ζ°ζ?εΊεεδΈΊε­η¬¦δΈ²
	// δ½Ώη¨fmt.Sprintfε°dataθ½¬ζ’δΈΊε­η¬¦δΈ²οΌθΏιδ½Ώη¨ηζ ΌεΌεεζ°ζ―β%vβοΌζζζ―ε°dataζεΆζ¬ζ₯ηεΌθ½¬ζ’δΈΊε­η¬¦δΈ²
	str := fmt.Sprintf("%v\n", data)
	// ιθΏf.fileηWrite()ζΉζ³οΌε°strε­η¬¦δΈ²θ½¬ζ’δΈΊ[]byteε­θζ°η»οΌεεε₯ε°ζδ»ΆδΈ­γε¦ζεηιθ――οΌεθΏε
	_, err := f.file.Write([]byte(str))
	return err
}

// εε»Ίζδ»Άεε₯ε¨ε?δΎ
func newFileWriter() *fileWriter {
	return &fileWriter{}
}

/**ε½δ»€θ‘εε₯
 * @description:
 */
// ε½δ»€θ‘εε₯ε¨
type consoleWriter struct{}

// ε?η°LogWriterηWrite()ζΉζ³
func (f *consoleWriter) Write(data interface{}) error {
	// ε°ζ°ζ?εΊεεδΈΊε­η¬¦δΈ²
	str := fmt.Sprintf("%v\n", data)
	// ε°ζ°ζ?δ»₯ε­θζ°η»εε₯ε½δ»€θ‘δΈ­
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// εε»Ίε½δ»€θ‘εε₯ε¨ε?δΎ
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

/** δ½Ώη¨ζ₯εΏ
 * @description:
 */

// εε»Ίζ₯εΏε¨
func testInterface2() {
	// εε»Ίζ₯εΏε¨
	l := NewLogger()
	// εε»Ίε½δ»€θ‘εε₯ε¨
	cw := newConsoleWriter()
	// ζ³¨εε½δ»€θ‘εε₯ε¨ε°ζ₯εΏε¨δΈ­
	l.RegisterWriter(cw)
	// εε»Ίζδ»Άεε₯ε¨
	fw := newFileWriter()
	// θ?Ύη½?ζδ»Άε
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	// ζ³¨εζδ»Άεε₯ε¨ε°ζ₯εΏε¨δΈ­
	l.RegisterWriter(fw)

	// εδΈδΈͺζ₯εΏ
	l.Log("hello")
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//ε?δΉδΈδΈͺζ°ζ?εε₯ε¨
type DataWriter interface {
	// interface{}η±»εηdataοΌθΏεδΈδΈͺerrorη»ζθ‘¨η€Ίε―θ½εηηιθ――
	WriteData(data interface{}) error
}

// ε?δΉζδ»Άη»ζοΌη¨δΊε?η°DataWriter
type file struct{}

// ε?η°DataWriterζ₯ε£ηWriteData()ζΉζ³
func (d *file) WriteData(data interface{}) error {
	// ζ¨‘ζεε₯ζ°ζ?
	fmt.Println("Write Data:", data)
	return nil
}

/** ζ₯ε£ηζΉζ³δΈε?η°ζ₯ε£ηη±»εζΉζ³ζ ΌεΌδΈθ΄
 * @description:
 */
func testInterface1() {
	// ε?δΎεfileθ΅εΌη»fοΌfηη±»εδΈΊ*file
	f := new(file)
	// ε£°ζδΈδΈͺDataWriterηζ₯ε£
	var writer DataWriter
	// ε°ζ₯ε£θ΅εΌfοΌδΉε°±ζ―οΌfileη±»ε
	writer = f
	// δ½Ώη¨DataWriterζ₯ε£θΏθ‘ζ°ζ?εε₯
	writer.WriteData("data")
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

// θ½¦θ½?
type Wheel struct {
	Size int
}

// θ½¦
type Car struct {
	Wheel
	// εΌζ
	Engine struct {
		Power int    // εη
		Type  string // η±»ε
	}
}

func testStruct3() {
	c := Car{
		// εε§εθ½?ε­οΌεε§εη»ζδ½εε΅οΌ
		Wheel: Wheel{
			Size: 18,
		},
		// εε§εεΌζοΌεε§εεε΅εΏεη»ζδ½οΌ
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// ε―ι£θ‘η
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// ε―θ‘θ΅°η
type Walkable struct{}

func (f *Walkable) Walk() {
	fmt.Println("can calk")
}

// δΊΊη±»
type Human struct {
	Walkable
	// δΊΊη±»θ½θ‘θ΅°
}

// ιΈη±»
type Bird struct {
	Walkable
	// ιΈη±»θ½θ‘θ΅°
	Flying
	// ιΈη±»θ½ι£θ‘
}

func testStruct2() {
	// ε?δΎειΈη±»
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()
	// ε?δΎεδΊΊη±»
	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
type BasicColor struct {
	R, G, B float32
}
type Color struct {
	// η»ζδ½εε΅
	BasicColor
	Alpha float32
}

/**
 * @description: ε£°ζη»ζδ½εε΅
 */
func testStruct1() {
	// ε?δΎεδΈδΈͺε?ζ΄ι’θ²η»ζδ½
	var c Color
	c.R = 1
	c.G = 1
	c.B = 0
	c.Alpha = 1
	fmt.Printf("%+v", c)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// ε?δΎεδΈδΈͺιθΏε­η¬¦δΈ²ζ ε°ε½ζ°εηηmap
// εε»ΊδΈδΈͺmapε?δΎοΌθΏδΈͺmapιθΏδΊδ»ΆεοΌstringοΌε³θεθ°εθ‘¨οΌ[]func(interface{}οΌοΌ
// εδΈδΈͺδΊδ»Άεη§°ε―θ½ε­ε¨ε€δΈͺδΊδ»Άεθ°οΌε ζ­€δ½Ώη¨εθ°εθ‘¨δΏε­γεθ°ηε½ζ°ε£°ζδΈΊfunc(interface{})
var eventByName = make(map[string][]func(interface{}))

/**δΊδ»Άζ³¨ε
 * @description: ζ³¨εδΊδ»ΆοΌζδΎδΊδ»Άεεεθ°ε½ζ°
 */
func RegisterEvent(name string, callback func(interface{})) {
	// ιθΏεε­ζ₯ζΎδΊδ»Άεθ‘¨
	list := eventByName[name]
	// ε¨εθ‘¨εηδΈ­ζ·»ε ε½ζ°
	// δΈΊεδΈδΈͺδΊδ»Άεη§°ε¨ε·²η»ζ³¨εηδΊδ»Άεθ°ηεθ‘¨δΈ­εζ·»ε δΈδΈͺεθ°ε½ζ°
	list = append(list, callback)
	// δΏε­δΏ?ζΉηδΊδ»Άεθ‘¨εη
	eventByName[name] = list
}

/**δΊδ»Άθ°η¨
 * @description:θ°η¨δΊδ»Ά
 */
func CallEvent(name string, param interface{}) {
	// ιθΏεε­ζΎε°δΊδ»Άεθ‘¨
	list := eventByName[name]
	// ιεθΏδΈͺδΊδ»Άηζζεθ°
	for _, callback := range list {
		// δΌ ε₯εζ°θ°η¨εθ°
		callback(param)
	}
}

/**δ½Ώη¨δΊδ»Άη³»η»
 * @description:
 */
// ε£°ζθ§θ²ηη»ζδ½
type Actor struct{}

// δΈΊθ§θ²ζ·»ε δΈδΈͺδΊδ»Άε€ηε½ζ°
// ζ₯ζparamεζ°οΌη±»εδΈΊinterface{}οΌδΈδΊδ»Άη³»η»ηε½ζ°οΌfunc(interface{})οΌη­ΎεδΈθ΄
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// ε¨ε±δΊδ»Ά
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}
func testFuncMethod1() {
	// ε?δΎεδΈδΈͺθ§θ²
	a := new(Actor)
	// ζ³¨εεδΈΊOn Skillηεθ°
	RegisterEvent("On Skill", a.OnEvent) // εζ¬‘ε¨OnSkillδΈζ³¨εε¨ε±δΊδ»Ά
	RegisterEvent("On Skill", GlobalEvent)
	// θ°η¨δΊδ»ΆοΌζζζ³¨εηεεε½ζ°ι½δΌθ’«θ°η¨
	CallEvent("On Skill", 100)
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
//ε£°ζδΈδΈͺη»ζδ½
type class struct{}

// η»η»ζδ½ζ·»ε Do()ζΉζ³
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

// ζ?ιε½ζ°ηDo()ζΉζ³
func funcDo(v int) {
	fmt.Println("call function do:", v)
}

/**
 * @description:ζΉζ³εε½ζ°ηη»δΈθ°η¨
 */
func testFuncMethod() {
	// ε£°ζδΈδΈͺε½ζ°εθ°
	var delegate func(int)
	// εε»Ίη»ζδ½ε?δΎ
	c := new(class)
	// ε°εθ°θ?ΎδΈΊcηDoζΉζ³
	delegate = c.Do
	// θ°η¨
	delegate(100)
	// ε°εθ°θ?ΎδΈΊζ?ιε½ζ°
	delegate = funcDo
	// θ°η¨3
	delegate(100)
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
type Command struct {
	Name    string // ζδ»€εη§°
	Var     *int   // ζδ»€η»ε?ηει
	Comment string // ζδ»€ηθ§£ι
}

/**
 * @description: εε°εη»ζδ½ε?δΎε
 * @param {string} name
 * @param {*int} varref
 * @param {string} comment ζθΏ°
 * @return {*}
 */
func testStruct(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// ε΄©ζΊζΆιθ¦δΌ ιηδΈδΈζδΏ‘ζ―
type panicContext struct {
	function string // ζε¨ε½ζ°
}

// δΏζ€ζΉεΌεθ?ΈδΈδΈͺε½ζ°
func ProtectRun(entry func()) {
	// δ½Ώη¨deferε°ι­εε»ΆθΏζ§θ‘οΌε½panicθ§¦εε΄©ζΊζΆοΌProtectRun()ε½ζ°ε°η»ζθΏθ‘οΌζ­€ζΆdeferεηι­εε°δΌεηθ°η¨
	defer func() {
		// εηε?ζΊζΆοΌθ·εpanicδΌ ιηδΈδΈζεΉΆζε°
		// recover()θ·εε°panicδΌ ε₯ηεζ°
		err := recover()
		switch err.(type) {
		case runtime.Error: // ε¦ζιθ――ζ―ζRuntimeε±ζεΊηθΏθ‘ζΆιθ――οΌε¦η©Ίζιθ?Ώι?γι€ζ°δΈΊ0η­ζε΅οΌζε°θΏθ‘ζΆιθ――
			fmt.Println("runtime error:", err)
		default: // ιθΏθ‘ζΆιθ――
			fmt.Println("error:", err)
		}
	}()
	entry()
}

/**
 * @description: ε?ζΊε€η
 * @param {*}
 * @return {*}
 */
func testPanic1() {
	fmt.Println("θΏθ‘ε")
	// εθ?ΈδΈζ?΅ζε¨θ§¦εηιθ――
	ProtectRun(func() {
		fmt.Println("ζε¨ε?ζΊε")
		// δ½Ώη¨panicδΌ ιδΈδΈζ
		// δ½Ώη¨panicζε¨θ§¦εδΈδΈͺιθ――οΌεΉΆε°δΈδΈͺη»ζδ½ιεΈ¦δΏ‘ζ―δΌ ιθΏε»οΌζ­€ζΆοΌrecoverε°±δΌθ·εε°θΏδΈͺη»ζδ½δΏ‘ζ―οΌεΉΆζε°εΊζ₯
		panic(&panicContext{
			"ζε¨θ§¦εpanic",
		})
		fmt.Println("ζε¨ε?ζΊε")
	})
	// ζζι ζη©Ίζιθ?Ώι?ιθ――
	ProtectRun(func() {
		fmt.Println("θ΅εΌε?ζΊε")
		var a *int
		// ζ¨‘ζδ»£η δΈ­η©Ίζιθ΅εΌι ζηιθ――οΌζ­€ζΆδΌη±Runtimeε±ζεΊιθ――οΌθ’«ProtectRun()ε½ζ°ηrecover()ε½ζ°ζθ·ε°
		*a = 1
		fmt.Println("θ΅εΌε?ζΊε")
	})
	fmt.Println("θΏθ‘ε")
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func testPanic() {
	defer fmt.Println("π£ ε?ζΊεθ¦εηδΊζ1 ")
	defer fmt.Println("β ε?ζΊεθ¦εηδΊζ2 ")

	panic("ε?ζΊ")
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// ε£°ζδΈδΈͺθ§£ζιθ――
type ParseError struct {
	Filename string // ζδ»Άε
	Line     int    // θ‘ε·
}

// ε?η°errorζ₯ε£οΌθΏειθ――ζθΏ°
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

/**
 * @description: θͺε?δΉError
 * @param {*}
 * @return {*}
 */
func testError2() {
	var e error
	// εε»ΊδΈδΈͺιθ――ε?δΎοΌεε«ζδ»Άεεθ‘ε·
	e = &ParseError{"main.go", 1}

	// ιθΏerrorζ₯ε£ζ₯ηιθ――ζθΏ°
	fmt.Println(e.Error())

	// ζ Ήζ?ιθ――ζ₯ε£ηε·δ½η±»εοΌθ·εθ―¦η»ηιθ――δΏ‘ζ―
	switch detail := e.(type) {
	case *ParseError: // θΏζ―δΈδΈͺθ§£ζιθ――
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default: // εΆδ»η±»εηιθ――
		fmt.Println("other error")
	}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

/**
 * @description:ι€ζ³ιθ――ζ΅θ―
 * @param {*}
 * @return {*}
 */
// ε?δΉι€ζ°δΈΊ0ηιθ――
var errDivisionByZero = errors.New("division by zero")

func testError1(dividend, divisor int) (int, error) {
	fmt.Printf("<=============== π π π ===============> \n\n")

	// ε€ζ­ι€ζ°δΈΊ0ηζε΅εΉΆθΏε
	if divisor == 0 {

		return 0, errDivisionByZero
	}
	// ζ­£εΈΈθ?‘η?οΌζε°η©Ίιθ――
	return dividend / divisor, nil

}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// ιθ――ε­η¬¦δΈ²
type errorString struct {
	s string
}

// θΏεεηδ½η§ιθ――
// ε?η°errorζ₯ε£ηError()ζΉζ³οΌθ―₯ζΉζ³θΏεζεδΈ­ηιθ――ζθΏ°
func (e *errorString) Error() string {
	return e.s
}

/**
 * @description: ιθ――
 * @param {*}
 * @return {*}
 */
func testError(text string) {
	fmt.Printf("<=============== π π π ===============> \n\n")
	fmt.Print(&errorString{text})
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

/**ε»ΆθΏθ―­ζ³
 * @description:
 * @param {*}
 * @return {*}
 */
func testDef() {
	fmt.Printf("<=============== π π π ===============> \n\n")

	filename := "/Users/harleyhuang/Documents/GitHub/Go/GoDemo/main.go"

	f, err := os.Open(filename)
	if err != nil {
		return
	}
	// ε»ΆθΏθ°η¨CloseοΌζ­€ζΆCloseδΈδΌθ’«θ°η¨
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// deferζΊεΆθ§¦εοΌθ°η¨Closeε³ι­ζδ»Ά
		return
	}
	size := info.Size()
	// deferζΊεΆθ§¦εοΌθ°η¨Closeε³ι­ζδ»Ά
	fmt.Println("ζδ»ΆsizeοΌ", size)
}

/**
 * @description: ε―εεζ°
 * @param {*}
 * @return {*}
 */
func testVariableParameters(slist ...string) {
	fmt.Printf("<=============== π π π ===============> \n\n")

	// ε?δΉδΈδΈͺε­θηΌε²οΌεΏ«ιε°θΏζ₯ε­η¬¦δΈ²
	var b bytes.Buffer
	// ιεε―εεζ°εθ‘¨slistοΌη±»εδΈΊ[]string
	for _, s := range slist {
		// ε°ιεεΊηε­η¬¦δΈ²θΏη»­εε₯ε­θζ°η»
		b.WriteString(s)
	}
	// ε°θΏζ₯ε₯½ηε­θζ°η»θ½¬ζ’δΈΊε­η¬¦δΈ²εΉΆθΎεΊ
	fmt.Printf(b.String())

}

/**
 * @description: ι­εηθ?°εΏζεΊ
 * @param {*}
 * @return {*}
 */
func testClosure1_1() {
	fmt.Printf("<=============== π π π ===============> \n\n")

	// εε»ΊδΈδΈͺη΄―ε ε¨οΌεε§εΌδΈΊ1οΌ
	// θΏεηaccumulatorζ―η±»εδΈΊfunc() intηε½ζ°ειγ
	accumulator := testClosure1(1)
	// θ°η¨accumulator()ζΆοΌεΌε§ζ§θ‘func() int{}εΏεε½ζ°ι»θΎοΌη΄ε°θΏεη±»ε εΌ
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// ζε°η΄―ε ε¨ηε½ζ°ε°ε
	fmt.Printf("%p\n", accumulator)
	// εε»ΊδΈδΈͺη΄―ε ε¨οΌεε§εΌδΈΊ
	accumulator2 := testClosure1(10)
	// η΄―ε 1εΉΆζε°
	fmt.Println(accumulator2())
	// ζε°η΄―ε ε¨ηε½ζ°ε°ε
	fmt.Printf("%p\n", accumulator2)

}

/**
 * @description: η΄―ε ε¨ηζε½ζ°οΌθΏδΈͺε½ζ°θΎεΊδΈδΈͺεε§εΌοΌθ°η¨ζΆθΏεδΈδΈͺδΈΊεε§εΌεε»Ίηι­εε½ζ°
 * @param {*}
 * @return {*}
 */
func testClosure1(value int) func() int {

	// θΏεδΈδΈͺι­εε½ζ°οΌζ―ζ¬‘θΏεδΌεε»ΊδΈδΈͺζ°ηε½ζ°ε?δΎ
	return func() int {
		// ε―ΉεΌη¨ηtestClosure1εζ°ειθΏθ‘η΄―ε οΌ
		// ζ³¨ζvalueδΈζ―θ¦θΏεηεΏεε½ζ°ε?δΉηοΌδ½ζ―θ’«θΏδΈͺεΏεε½ζ°εΌη¨οΌζδ»₯ε½’ζι­εγ
		value++
		// θΏεδΈδΈͺη΄―ε εΌ
		return value
	}
}

//ε½ζ°ε?δΉδΈΊη±»ε
type FuncCaller func(interface{})

// ε?η°InvokerηCall
func (f FuncCaller) Call(p interface{}) {
	// θ°η¨f()ε½ζ°ζ¬δ½
	f(p)
}

/**
 * @description: ε½ζ°ε?η°ζ₯ε£
 * @param {*}
 * @return {*}
 */
func testFuncImplInterface1() {
	fmt.Printf("<=============== π π π ===============> \n\n")

	// ε£°ζζ₯ε£ει
	var invoker Invoker
	// ε°εΏεε½ζ°θ½¬δΈΊFunc Callerη±»εοΌεθ΅εΌη»ζ₯ε£
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	// δ½Ώη¨ζ₯ε£θ°η¨Func Caller.CallοΌει¨δΌθ°η¨ε½ζ°ζ¬δ½
	invoker.Call("π ε½ζ°ζ₯ε£ hello")
}

// θ°η¨ε¨ζ₯ε£
type Invoker interface {
	// ιθ¦ε?η°δΈδΈͺCall()ζΉζ³
	Call(interface{})
}

type Struct struct{}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

/**
 * @description: η»ζδ½ε?η°ζ₯ε£
 * @param {*}
 * @return {*}
 */
func testFuncImplInterface() {
	fmt.Printf("<=============== π π π ===============> \n\n")

	// ε£°ζζ₯ε£ει
	var invoker Invoker
	// ε?δΎεη»ζδ½
	s := new(Struct)
	// ε°ε?δΎεηη»ζδ½θ΅εΌε°ζ₯ε£
	invoker = s
	// δ½Ώη¨ζ₯ε£θ°η¨ε?δΎεη»ζδ½ηζΉζ³Struct.Call
	invoker.Call("π hello ε½ζ°ε?η°ζ₯ε£")
}

/**
 * @description: εΏεε½ζ°ε°θ£
 * @param {*}
 * @return {*}
 */
func testAnnoymousFunction1() {

	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	// ε?δΉε½δ»€θ‘skillParamοΌδ»ε½δ»€θ‘θΎε₯βskillε―δ»₯ε°η©Ίζ Όεηε­η¬¦δΈ²δΌ ε₯skill Paramζιει
	var skillParam = flag.String("skill", "", "skill to perform")

	// θ§£ζε½δ»€θ‘εζ°οΌθ§£ζε?ζεοΌskillParamζιειε°ζεε½δ»€θ‘δΌ ε₯ηεΌ
	flag.Parse()

	// ε?δΉδΈδΈͺδ»ε­η¬¦δΈ²ζ ε°ε°func()ηmapοΌηΆεε‘«εθΏδΈͺmap
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		}}

	// skillParamζ―δΈδΈͺ*stringη±»εηζιειοΌδ½Ώη¨*skill Paramθ·εε°ε½δ»€θ‘δΌ θΏζ₯ηεΌοΌεΉΆε¨mapδΈ­ζ₯ζΎε―ΉεΊε½δ»€θ‘εζ°ζε?ηε­η¬¦δΈ²ηε½ζ°
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

//ιεεηηζ―δΈͺεη΄ οΌιθΏη»ε?ε½ζ°θΏθ‘εη΄ θ?Ώι?
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

/**
 * @description: εΏεε½ζ°
 * @param {*}
 * @return {*}
 */
func testAnonymousFunction() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	// δ½Ώη¨εΏεε½ζ°ζε°εηεε?Ή18
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	func(data int) {
		fmt.Println("hello", data)
	}(100)

	// ε°εΏεε½ζ°δ½δΏε­ε°f()
	f := func(data int) {
		fmt.Println("hello", data)
	}
	// δ½Ώη¨f()θ°η¨
	f(100)

}

// η¨δΊζ΅θ―εΌδΌ ιζζηη»ζδ½
type Data struct {
	// ζ΅θ―εηε¨εζ°δΌ ιδΈ­ηζζ
	complax []int

	instance InnerData
	// ε?δΎειηinner Data
	ptr *InnerData
	// ε°ptrε£°ζδΈΊInner Dataηζιη±»ε
}

// δ»£θ‘¨εη§η»ζδ½ε­ζ?΅
type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	// θΎεΊεζ°ηζεζε΅
	// δ½Ώη¨ζ ΌεΌεηβ%+vβε¨θ―θΎεΊinειηθ―¦η»η»ζοΌδ»₯δΎΏθ§ε―Dataη»ζε¨δΌ ιεεηει¨ζ°εΌηεεζε΅
	fmt.Printf("in Func value: %+v\n", inFunc)
	// ζε°inFuncηζιοΌε¨θ?‘η?ζΊδΈ­οΌζ₯ζηΈεε°εδΈη±»εηΈεηειοΌθ‘¨η€Ίηζ―εδΈεεε­εΊε
	fmt.Printf("in Func ptr: %p\n", &inFunc)
	return inFunc
}

/**
 * @description: εΌδΌ ιηζ΅θ―ε½ζ°
 * @param {*}
 * @return {*}
 */
func paramTranslate() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	in := Data{
		//εη
		complax: []int{1, 2, 3},
		// η»ζδ½
		instance: InnerData{
			5,
		},
		// ζι
		ptr: &InnerData{1},
	}
	// θΎε₯η»ζηζεζε΅
	fmt.Printf("in value: %+v\n", in)
	// θΎε₯η»ζηζιε°ε
	fmt.Printf("in ptr: %p\n", &in)
	// δΌ ε₯η»ζδ½οΌθΏεεη±»εηη»ζδ½
	out := passByValue(in)
	// θΎεΊη»ζηζεζε΅
	fmt.Printf("out value: %+v\n", out)
	// θΎεΊη»ζηζιε°ε
	fmt.Printf("out ptr: %p\n", &out)
}

// εθ‘¨ε ι€
func list_delete() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	l := list.New()

	// ε°Ύι¨ζ·»ε 
	l.PushBack("canon")
	// ε€΄ι¨ζ·»ε 
	l.PushFront(67)
	// ε°Ύι¨ζ·»ε εδΏε­εη΄ ε₯ζ
	element := l.PushBack("fist")
	// ε¨fistδΉεζ·»ε high
	l.InsertAfter("high", element)
	// ε¨fistδΉεζ·»ε noon
	l.InsertBefore("noon", element)
	// δ½Ώη¨
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

// δΉδΉδΉζ³θ‘¨οΌ
func multiplication_table() {

	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	// ιεοΌε³ε?ε€ηη¬¬ε θ‘
	for y := 1; y <= 9; y++ {
		// ιεοΌε³ε?θΏδΈθ‘ζε€ε°ε
		for x := 1; x <= y; x++ {
			fmt.Printf("%dοΌ%d=%d ", x, y, x*y)
		}
		// ζε¨ηζεθ½¦
		fmt.Println()
	}
}

// εη
func section_test() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	var a = [4]int{10, 20, 30, 40}

	fmt.Println(a, "\n", a[1:3])
}

// εε§εζ°η»
func init_array() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	var member = [...]string{"ζΉε―ι¦¨", "ζ―δΈͺ", "ζΎε±η²Ύπ³", "ζΉζΊε?Έ", "ζ―δΈͺ", "θ°η?ι¬Όπ", "ζδΊ", "ζ―δΈͺ", "ε°θθΉπ¦οΈ"}

	for k, v := range member {
		fmt.Println(k, v)

	}

	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	v := scene["route2"]

	fmt.Println(v)

	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}

// δ½Ώη¨ζιειθ·εε½δ»€θ‘ηθΎε₯δΏ‘ζ―
func point_flag() {
	// ε?δΉε½δ»€θ‘εζ°
	/*
	* 3δΈͺεζ°εε«ε¦δΈοΌ
	* εζ°εη§°οΌε¨η»εΊη¨θΎε₯εζ°ζΆοΌδ½Ώη¨θΏδΈͺεη§°
	* εζ°εΌηι»θ?€εΌοΌδΈflagζδ½Ώη¨ηε½ζ°εε»Ίειη±»εε―ΉεΊοΌStringε―ΉεΊε­η¬¦δΈ²γIntε―ΉεΊζ΄εγBoolε―ΉεΊεΈε°εη­
	* εζ°θ―΄ζοΌδ½Ώη¨-helpζΆοΌδΌεΊη°ε¨θ―΄ζδΈ­
	 */
	var mode = flag.String("mode", "π π", "process mode")

	// θ§£ζε½δ»€θ‘εζ°
	flag.Parse()

	fmt.Println(*mode)
}

// ε½ζ°ηδΊ€ζ’
func chargeValue() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	// εε€δΈ€δΈͺειοΌθ΅εΌ1ε2
	x, y := 1, 2
	// δΊ€ζ’ειεΌ
	swap(&x, &y)

	fmt.Println(x, y)

}

func swap(a, b *int) {
	// εaζιηεΌοΌθ΅η»δΈ΄ζΆει
	t := *a

	// εbζιηεΌοΌθ΅η»aζιζεηει
	// ζ³¨ζοΌζ­€ζΆβ*aβηζζδΈζ―εaζιηεΌοΌθζ―βaζεηειβ
	*a = *b
	// ε°aζιηεΌθ΅η»bζιζεηει
	*b = t
}

// δ»ζιθ·εζιζεηεΌ
func pointTest1() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	var house string = "π  ζΏε± 366ββ26-404"

	ptr := &house

	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)

	value := *ptr

	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)
}

func pointTest0() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	var cat int = 1
	var str string = "banana"

	fmt.Printf("%p %p ", &cat, &str)
}

// η­ειε£°ζεΉΆεε§ε
func lowVar() {
	fmt.Printf("\n\n <=============== π π π ===============> \n\n")

	// Goθ―­θ¨ηζ¨ε―Όε£°ζεζ³οΌηΌθ―ε¨δΌθͺε¨ζ Ήζ?ε³εΌη±»εζ¨ζ­εΊε·¦εΌηε―ΉεΊη±»ε
	// ζ³¨ζοΌη±δΊδ½Ώη¨δΊβ:=βοΌθδΈζ―θ΅εΌηβ=βοΌε ζ­€ζ¨ε―Όε£°ζεζ³ηε·¦εΌειεΏι‘»ζ―ζ²‘ζε?δΉθΏηειγθ₯ε?δΉθΏοΌε°δΌεηηΌθ―ιθ――γ
	hp := 10

	// ζ³¨ζοΌε¨ε€δΈͺη­ειε£°ζεθ΅εΌδΈ­οΌθ³ε°ζδΈδΈͺζ°ε£°ζηειεΊη°ε¨ε·¦εΌδΈ­οΌε³δΎΏεΆδ»ειεε―θ½ζ―ιε€ε£°ζηοΌηΌθ―ε¨δΉδΈδΌζ₯ι
	conn, err := net.Dial("tcp", "127.0.0.1: 8080")
	conn2, err := net.Dial("tcp", "127.0.0.1: 8080")

	fmt.Printf("hp: %d, conn: %s, err: %s, conn2: %s", hp, conn, err, conn2)

}
