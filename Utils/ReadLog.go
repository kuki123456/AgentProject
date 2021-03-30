package Utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)
var (
	UPLOAD,
	ViewEvent,
	StateChangeEvent,
	NetworkEvent,
	LaunchEvent,
	LagEvent,
	JSErrorEvent,
	H5Event,
	CustomMetricEvent,
	CustomLogEvent,
	CustomEventEvent,
	CrashEvent,
	ActionEvent  int
)
func ReadLine(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func handler(line string) {

		if strings.Contains(line, "UPLOAD") {
			UPLOAD += 1
		}
		if strings.Contains(line, "ViewEvent") {
			ViewEvent += 1

		}
		if strings.Contains(line, "StateChangeEvent") {
			StateChangeEvent += 1

		}
		if strings.Contains(line, "NetworkEvent") {
			NetworkEvent += 1
		}
		if strings.Contains(line, "LaunchEvent") {
			LaunchEvent += 1

		}
		if strings.Contains(line,"LagEvent"){
			LagEvent+=1
		}
		if strings.Contains(line,"JSErrorEvent"){
			JSErrorEvent+=1
		}
		if strings.Contains(line, "H5Event") {
			H5Event += 1

		}
		if strings.Contains(line, "CustomMetricEvent") {
			CustomMetricEvent += 1

		}
		if strings.Contains(line, "CustomLogEvent") {
			CustomLogEvent += 1


	}
		if strings.Contains(line,"CustomEventEvent"){
			CustomEventEvent+=1
		}
		if strings.Contains(line,"CrashEvent"){
			CrashEvent+=1
		}
		if strings.Contains(line,"ActionEvent"){
			ActionEvent+=1
		}
}
func InitCount(){
	UPLOAD=0
		ViewEvent=0
		StateChangeEvent=0
		NetworkEvent=0
		LaunchEvent=0
		LagEvent=0
		JSErrorEvent=0
		H5Event=0
		CustomMetricEvent=0
		CustomLogEvent=0
		CustomEventEvent=0
		CrashEvent=0
		ActionEvent =0
}