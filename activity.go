package readFile

import (
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-lib/logger"
    "io/ioutil"
    "path/filepath"
    "os"

)

const (
    ivField    = "filePath"
    fileContent    = "fileContent"
    fileLocation = "fileLocation"
    fileName="fileName"
    fileSize="fileSize"
)

var activityLog = logger.GetLogger("tibco-activity-fileRead")

type DemoActivity struct {
    metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
    return &DemoActivity{metadata: metadata}
}

func (a *DemoActivity) Metadata() *activity.Metadata {
    return a.metadata
}
func (a *DemoActivity) Eval(context activity.Context) (done bool, err error) {
    activityLog.Info("Executing DemoActivity activity")
    //Read Inputs

    if context.GetInput(ivField) == nil {
        // return error to the engine

        return false, activity.NewError("File path is not provided", "readFile-4001", nil)
    }
    field1v := context.GetInput(ivField).(string)



    data, err := ioutil.ReadFile(field1v)

    if err != nil {

        return false, activity.NewError("Error while reading file (inaccurate file path)", "readFile-4001", nil)
    }
    dir, file := filepath.Split(field1v)
    fileInfo, err :=os.Stat(field1v)


    //Set output
    context.SetOutput(fileContent, string(data))
    context.SetOutput(fileName, file)
    context.SetOutput(fileLocation, dir)
    context.SetOutput(fileSize, fileInfo.Size())

    return true, nil
}