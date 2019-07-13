package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const DATE_FORMAT = "2006-01-02"

type LEVEL byte

const (
	TRACE LEVEL = iota
	INFO
	WARN
	ERROR
	OFF
)

var fileLog *FileLogger

type LoggerConf struct {
	FileDir  string `yaml:"fileDir"`
	FileName string `yaml:"fileName"`
	Prefix   string `yaml:"prefix"`
	Level    string `yaml:"level"`
}

type FileLogger struct {
	fileDir  string
	fileName string
	prefix   string
	date     *time.Time
	logFile  *os.File
	lg       *log.Logger
	logLevel LEVEL
	mu       *sync.RWMutex
	logChan  chan string
}

func InitLogger(conf *LoggerConf) (err error) {
	f := &FileLogger{
		fileDir:  conf.FileDir,
		fileName: conf.FileName,
		prefix:   conf.Prefix,
		mu:       new(sync.RWMutex),
		logChan:  make(chan string, 5000),
	}

	if strings.EqualFold(conf.Level, "OFF") {
		f.logLevel = OFF
	} else if strings.EqualFold(conf.Level, "TRACE") {
		f.logLevel = TRACE
	} else if strings.EqualFold(conf.Level, "WARN") {
		f.logLevel = WARN
	} else if strings.EqualFold(conf.Level, "ERROR") {
		f.logLevel = ERROR
	} else {
		f.logLevel = INFO
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	f.date = &t

	if f.isMustSplit() {
		if err = f.split(); err != nil {
			return
		}
	} else {
		f.isExistOrCreate()

		logFile := filepath.Join(f.fileDir, f.fileName)

		f.logFile, err = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return
		}

		f.lg = log.New(f.logFile, f.prefix, log.LstdFlags|log.Lmicroseconds)
	}

	go f.logWriter()
	go f.fileMonitor()

	fileLog = f
	return
}

// 日志文件是否必须分割
func (f FileLogger) isMustSplit() bool {
	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	return t.After(*f.date)
}

// 日志文件是否存在，不存在则创建
func (f FileLogger) isExistOrCreate() {
	_, err := os.Stat(f.fileDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(f.fileDir, 0755)
	}
}

func (f *FileLogger) split() (err error) {
	logFile := filepath.Join(f.fileDir, f.fileName)
	logFileBak := logFile + "." + f.date.Format(DATE_FORMAT)

	if f.logFile != nil {
		f.logFile.Close()
	}

	err = os.Rename(logFile, logFileBak)
	if err != nil {
		return
	}

	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	f.date = &t

	f.logFile, err = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	f.lg = log.New(f.logFile, f.prefix, log.LstdFlags|log.Lmicroseconds)
	return
}

// 日志写入
func (f *FileLogger) logWriter() {
	defer func() { recover() }()
	
	for {
		select {
		case str := <-f.logChan:
			f.mu.RLock()
			defer f.mu.RUnlock()

			f.lg.Output(2, str)
		}
	}
}

// 日志分割监控
func (f *FileLogger) fileMonitor() {
	defer func() { recover() }()
	
	timer := time.NewTicker(300 * time.Second)
	for {
		select {
		case <-timer.C:
			if f.isMustSplit() {
				f.mu.Lock()
				defer f.mu.Unlock()

				if err := f.split(); err != nil {
					Error("Log split error: %v\n", err)
				}
			}
		}
	}
}

// 关闭日志
func CloseLogger() {
	if fileLog != nil {
		close(fileLog.logChan)
		fileLog.lg = nil
		fileLog.logFile.Close()
	}
}

// 日志写入

func Printf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf(format, v...)
}

func Print(v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprint(v...)
}

func Println(v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintln(v...)
}

func Trace(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if fileLog.logLevel <= TRACE {
		fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf("[TRACE] "+format, v...)
	}
}

func Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if fileLog.logLevel <= INFO {
		fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf("[INFO] "+format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if fileLog.logLevel <= WARN {
		fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf("[WARN] "+format, v...)
	}
}

func Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if fileLog.logLevel <= ERROR {
		fileLog.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf("[ERROR] "+format, v...)
	}
}
