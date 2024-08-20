// Package fs 从文件反序列化任务列表，或将任务序列化到文件
package fs

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"task-tracker/task"
)

var taskFilePath string

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("Get execution patherror: %v", err)
	}

	taskFilePath = filepath.Join(filepath.Dir(ex), "tasks.json")
}

func ReadTaskFromFile() (result []*task.Task, err error) {

	result = make([]*task.Task, 0)

	// 检查文件是否存在
	if _, err := os.Stat(taskFilePath); os.IsNotExist(err) {
		log.Printf("File %s does not exist\n", taskFilePath)
		log.Printf("Create task file %s\n", taskFilePath)
		f, err := os.Create(taskFilePath)
		if err != nil {
			log.Printf("Create task file %s failed.\n", taskFilePath)
			return nil, err
		}
		// 程序函数返回时关闭文件
		defer func(closer io.Closer) {
			err := closer.Close()
			if err != nil {
				log.Printf("Close task file %s failed.\n", taskFilePath)
			}
		}(f)

		if err := json.NewEncoder(f).Encode([]*task.Task{}); err != nil {
			log.Printf("Encode task to file %s failed.\n", taskFilePath)
			return nil, err
		}

		return result, nil
	}

	f, err := os.Open(taskFilePath)
	if err != nil {
		log.Printf("Open task file %s failed.\n", taskFilePath)
		return nil, err
	}

	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Printf("Close task file %s failed.\n", taskFilePath)
		}
	}(f)

	// 读取文件并反序列任务列表
	if err := json.NewDecoder(f).Decode(&result); err != nil {
		log.Printf("Decode task from file %s failed.\n", taskFilePath)
		return nil, err
	}

	return result, nil
}

func WriteTaskToFile(result []*task.Task) error {
	f, err := os.Create(taskFilePath)
	if err != nil {
		log.Printf("Create task file %s failed.\n", taskFilePath)
		return err
	}

	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Printf("Close task file %s failed.\n", taskFilePath)
		}
	}(f)

	if err := json.NewEncoder(f).Encode(result); err != nil {
		log.Printf("Encode task to file %s failed.\n", taskFilePath)
		return err
	}

	return nil
}
