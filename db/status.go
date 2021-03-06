package db

import (
	"log"
	"os"
)

// GetStatus :: check that the db instance is running and returns it's details
func GetStatus() (*RunningDBInstanceInfo, error) {
	log.Println("[TRACE] GetStatus ")
	info, err := loadRunningInstanceInfo()
	if err != nil {
		return nil, err
	}

	if info == nil {
		log.Println("[TRACE] GetStatus - loadRunningInstanceInfo returned nil ")
		// we do not have a info file
		return nil, nil
	}

	pidExists, err := pidExists(info.Pid)
	if err != nil {
		return nil, err
	}
	if !pidExists {
		log.Printf("[TRACE] GetStatus - pid %v does not exist\n", info.Pid)
		// nothing to do here
		os.Remove(runningInfoFilePath())
		return nil, nil
	}

	return info, nil
}
