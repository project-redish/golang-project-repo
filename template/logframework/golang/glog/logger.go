package logger

import (
	"github.com/golang/glog"
	"flag"
)

const (
	LOG_LEVEL string = "2"
)

func init() {
	flag.Parse()
	glog.Infof("Glog logging framework will be used for loggin purposes . For more info go to https://github.com/golang/glog")
	glog.Infof("The default glog level is %s ", LOG_LEVEL)
	

}
