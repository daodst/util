package log

import (
	"errors"
	"github.com/sirupsen/logrus"
)

type LogModule string

//
var logModules = make(map[string]*logrus.Logger)

//
var moduleLogLevels = make(map[LogModule]logrus.Level)

// 
var defaultLogLevel = logrus.InfoLevel


//
var Log *logrus.Logger = func() *logrus.Logger{
	log1 := logrus.New()
	log1.SetFormatter(diyTextFormatter()) //
	return log1
}()


//,panic
func RegisterModule(name string,level logrus.Level) LogModule{
	if CheckModuleExists(LogModule(name)) {
		panic("log module " + name + " only exists")
	}
	moduleLog := logrus.New() //log
	moduleLog.SetLevel(level) //
	moduleLog.SetFormatter(diyTextFormatter())
	logModules[name] = moduleLog
	moduleLogLevels[LogModule(name)] = level
	return LogModule(name)
}


//
func ResetAllModuleLevel(level logrus.Level) (err error){
	defaultLogLevel = level
	for moduleName,_ := range moduleLogLevels{
		err = SetModuleLevel(moduleName,level)
		if err != nil{
			return err
		}
	}
	return err
}

//
func SetModuleLevel(module LogModule,level logrus.Level) error{
	if !CheckModuleExists(module){
		return errors.New("log module "+string(module)+" not have registerd")
	}
	logModules[string(module)].SetLevel(level)
	moduleLogLevels[module] = level
	return nil
}

//
func CheckModuleExists(module LogModule) bool{
	if _,ok := logModules[string(module)];ok{
		return true
	}else{
		return false
	}
}


//
func GetDefaultLogLevel() logrus.Level{
	return defaultLogLevel
}

//
func GetModelLevels() map[LogModule]logrus.Level{
	return moduleLogLevels
}
