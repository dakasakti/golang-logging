package logging

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

func TestLogger(t *testing.T) {

	logger.Println("Hello Logger")
}

func TestOutput(t *testing.T) {
	file, _ := os.OpenFile("application-text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	// dikirim ke file bukan ke console lagi
	logger.Info("Hello World")
}

func TestFormatter(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application-json.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Hello World")
}

func TestAddField(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "dakasakti").Info("Hello World")
}

func TestAddFields(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "dakasakti",
	}).Info("Hello World")
}

func TestAlurKerja(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "dakasakti").Info("Hello World")
}

type SampleHook struct{}

func (sh *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (sh *SampleHook) Fire(entry *logrus.Entry) error {

	// implementasi
	fmt.Println(entry.Level, ": Sample Hook => ", entry.Message)

	return nil
}

func TestHook(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.AddHook(&SampleHook{})

	logger.Info("Hook di Print")
}

func TestSingleton(t *testing.T) {
	logrus.Info("Langsung bisa dieksekusi")
}
