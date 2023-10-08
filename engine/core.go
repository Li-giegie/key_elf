package engine

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/cppdebug/windev"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type TaskManager struct {
	lock  sync.Mutex
	Conf  ConfigurationFile `yaml:"conf"`
	Sleep time.Duration     `yaml:"sleep,omitempty"` //检测间隔时间
	Tasks []*Task           `yaml:"tasks,omitempty"`
	Debug bool              `yaml:"debug"`
	running bool
}

type ConfigurationFile struct {
	name               string
	EnableCheckRefresh bool          `yaml:"enable_check_refresh"`
	Sleep              time.Duration `yaml:"sleep"`
}

var tm *TaskManager

func NewTaskManager(confName string) (*TaskManager, error) {
	f, err := os.Open(confName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buf []byte
	if buf, err = io.ReadAll(f); err != nil {
		return nil, err
	}

	var info os.FileInfo
	if info, err = f.Stat(); err != nil {
		return nil, err
	}
	confHex = string(getMd5(buf))
	confModTime = info.ModTime().UnixNano()

	var tmp TaskManager
	tmp.running = true
	tmp.Conf.name = confName
	err = yaml.Unmarshal(buf, &tmp)
	return &tmp, err
}

func NewKeypressTaskDefault() *TaskManager {
	return &TaskManager{
		Sleep: time.Millisecond * 100,
		Debug: false,
		Conf: ConfigurationFile{
			EnableCheckRefresh: true,
			Sleep:              time.Millisecond * 100,
		},
		Tasks: []*Task{
			&Task{
				Name:        "hello",
				Description: "在cmd中输出hello.",
				Keys: []Key{
					{
						Name: "ctrl",
						Tag:  "left",
					},
					{
						Name: "1",
						Tag:  "left",
					},
				},
				Run: []string{
					"echo hello",
				},
			},
			&Task{
				Name:        "输出world",
				Description: "描述......",
				Keys: []Key{
					{
						Name: "ctrl",
						Tag:  "left",
					},
					{
						Name: "W",
					},
				},
				Run: []string{
					"echo world",
				},
			},
		},
	}
}

func (k *TaskManager) Init() error {
	for id, task := range k.Tasks {
		err := task.Init(task.Keys, id)
		if err != nil {
			return fmt.Errorf("初始化按键错误, 任务名：[%s]: %v\n%s\n", task.Name,err, task.String())
		}
		fmt.Println(task.String())
	}
	return nil
}

func (k *TaskManager) Listen() {
	log.Printf("listen [%x] start ------\n", confHex)
	for k.running {
		for _, task := range k.Tasks {
			//命中检查
			if hitDetection(task.cacheKey, len(task.Keys)) {
				fmt.Println("激活任务:", task.Name)
				//运行命令
				if err := executeCmdS(task.Run); err != nil {
					log.Printf("execute cmd err: task name: [%v] cmd:[%v],%v\n", task.Name, strings.Join(task.Run, " "), err)
				}
				time.Sleep(time.Millisecond*500)
			}
		}
		time.Sleep(k.Sleep)
	}
}

func hitDetection(keys []int, keySize int) bool {
	var n int
	for _, asciiKey := range keys {
		if windev.KeyDownUp(asciiKey) == 1 {
			n++
		}
	}
	if n != 0 && n == keySize {
		return true
	}
	return false
}

func executeCmdS(cmdS []string) error {
	for _, _cmd := range cmdS {
		cmd := exec.Command("cmd", "/C", _cmd)
		err := cmd.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *TaskManager) GenerateYamlFile(name string) error {
	buf, err := yaml.Marshal(k)
	if err != nil {
		return err
	}
	return os.WriteFile(name, buf, 0644)
}

type Task struct {
	id          int      `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description,omitempty"`
	Keys        []Key    `yaml:"keys"`
	Run         []string `yaml:"run"`

	cacheKey []int
}

func (t *Task) Init(keys []Key, i int) error {
	t.cacheKey = make([]int, len(keys))
	t.id = i
	var keyAscii int
	var ok bool
	for _, key := range keys {
		name := strings.ToLower(key.Name)
		tag := strings.ToLower(key.Tag)
		switch name {
		case "mouse":
			switch tag {
			case "left", "right", "center":
				keyAscii, ok = asciiMap[tag+"_"+name]
			default:
				return errors.New("mouse tag not exist ")
			}
		case "win", "shift", "ctrl", "alt":
			switch tag {
			case "left", "right":
				keyAscii, ok = asciiMap[tag+"_"+name]
			case "":
				keyAscii, ok = asciiMap[name]
			default:
				return errors.New("win tag not exist ")
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			switch tag {
			case "left", "right":
				keyAscii, ok = asciiMap[tag+"_"+name]
			case "":
				tag = "left"
				keyAscii, ok = asciiMap[tag+"_"+name]
			default:
				return errors.New("0-9 [number] tag (left|right) not exist ")
			}
		case "select", "print":
			keyAscii, ok = asciiMap[name+"_"]
		default:
			keyAscii, ok = asciiMap[name]
		}
		if !ok {
			return errors.New("key ascii is not exist! ")
		}
		t.cacheKey = append(t.cacheKey, keyAscii)
	}
	return nil
}

func (t *Task) String() string {
	return fmt.Sprintf("taskID: %v, taskName: %v, keys: %v, run: %v", t.id, t.Name, t.Keys, t.Run)
}

type Key struct {
	Name string `yaml:"name,omitempty"`
	Tag  string `yaml:"tag,omitempty"`
}

func (k *Key) String() string {
	return fmt.Sprintf("name:%v,tag:%v", k.Name, k.Tag)
}

var confModTime int64
var confHex string

func ListenConfigChange(name string, sleep time.Duration, actionF func(b []byte)) error {
	var newModTime int64
	var newModHex string
	for {
		info, err := os.Stat(name)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return errors.New("not is a file")
		}
		newModTime = info.ModTime().UnixNano()
		//不一样的话计算hex值
		if confModTime != newModTime {
			confModTime = newModTime
			buf, err := os.ReadFile(name)
			if err != nil {
				return err
			}
			newModHex = getMd5(buf)
			if newModHex != confHex {
				fmt.Println("检测到配置发生变化")
				confHex = newModHex
				actionF(buf)
			}
		}
		time.Sleep(sleep)
	}
}

func findAsciiCode(v int) string {
	for s, i := range asciiMap {
		if v == i {
			return s
		}
	}
	return "not find key"
}

func StartGenerateYamlTemplateFile(name string) error {
	return NewKeypressTaskDefault().GenerateYamlFile(name)
}

func StartTaskManager(name string) (err error) {

	tm, err = NewTaskManager(name)
	if err != nil {
		return fmt.Errorf("启动失败 -1 %v", err)
	}
	if err = tm.Init(); err != nil {
		return fmt.Errorf("启动失败 -2 %v ", err)
	}
	log.Println("启动服务成功------")

	if tm.Conf.EnableCheckRefresh {
		go func() {
			err = ListenConfigChange(name, tm.Conf.Sleep, func(b []byte) {
				_tm, err := NewTaskManager(name)
				if err != nil {
					log.Println("刷新配置失败！-1", err)
					return
				}
				if err = _tm.Init(); err != nil {
					log.Println("刷新配置失败！-2", err)
					return
				}
				tm.running = false
				tm = _tm
				go tm.Listen()
				log.Println("配置刷新成功！")
			})
			if err != nil {
				fmt.Printf("ListenConfigChange err :%v \n",err)
			}
		}()
	}
	tm.Listen()
	return nil
}

func getMd5(b []byte) string {
	m := md5.New()
	m.Write(b)
	return string(m.Sum(nil))
}
