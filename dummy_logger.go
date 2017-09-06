package logi

type DummyLogger struct{}

func (l *DummyLogger) Debug(v ...interface{}) {

}

func (l *DummyLogger) Debugf(format string, v ...interface{}) {

}

func (l *DummyLogger) Info(v ...interface{}) {

}

func (l *DummyLogger) Infof(format string, v ...interface{}) {
}

func (l *DummyLogger) Warn(v ...interface{}) {
}

func (l *DummyLogger) Warnf(format string, v ...interface{}) {
}

func (l *DummyLogger) Error(v ...interface{}) {
}

func (l *DummyLogger) Errorf(format string, v ...interface{}) {
}

func (l *DummyLogger) Fatal(v ...interface{}) {
}

func (l *DummyLogger) Fatalf(format string, v ...interface{}) {
}

func (l *DummyLogger) Panic(v ...interface{}) {
}

func (l *DummyLogger) Panicf(format string, v ...interface{}) {
}
