#按键精灵
### 用于定义一组全局按键，当定义的按键被触发后在终端执行一些操作，由配置决定
### 支持windows，其他平台没有做测试不详
#### 生成配置
```
用于生成一个默认的配置文件
.\keyelf.exe create config [配置文件名称]
```

#### 启动
##### 默认程序会在当前路径下查找conf.yaml
```
用于指定配置文件启动
.\keyelf.exe -c [配置文件路径]
```


#### 配置文件，任务hello将会在用户按下左边的ctrl+左边的数字1在终端输出hello：
```yaml
conf:
  enable_check_refresh: true    #是否开启热更新
  sleep: 100ms                  #配置文件热更新检查间隔时间，时间可以稍微长一点
sleep: 100ms                    #按键是否触发间隔时间
tasks:
  - name: hello                 #任务名字
    description: 在cmd中输出hello. #任务描述介绍
    keys:                       #按键列表
      - name: ctrl              #按键值
        tag: left               #左边的ctrl
      - name: "1"               #数字1
        tag: left               #左边的数字生效
    run:
      - echo hello              #按键触发后要执行的命令，这里是在命令行中输出了hello
  - name: 输出world             #多任务，根据需要定义足够的任务
    description: 描述......
    keys:
      - name: ctrl
        tag: left
      - name: W
    run:
      - echo world
debug: false
```


### 支持的按键列表
```
const mouse_left = 0x01   // 鼠标左键
const mouse_right = 0x02  // 鼠标右键
const mouse_center = 0x04 // 鼠标中键

const backspace = 0x08      // BackSpace
const tab = 0x09            // Tab
const num_lock_close = 0x0C // Num Lock关闭时的数字区 5
const enter = 0x0D          // Enter
const shift = 0x10          // Shift/左右shift都可以
const ctrl = 0x11           // Ctrl/左右Ctrl都可以
const alt = 0x12            // Alt/左右Alt都可以
const pause = 0x13          /* Pause*/
const caps_lock = 0x14      // Caps Lock
const esc = 0x1B            // Esc
const space = 0x20          // Space
const pgup = 0x21           // PgUp
const pgdn = 0x22           // PgDn
const end = 0x23            // End
const home = 0x24           // Home
const left = 0x25           // 方向左
const up = 0x26             // 方向上
const right = 0x27          // 方向右
const down = 0x28           // 方向下
const select_ = 0x29        /* Select*/
const print_ = 0x2A         /* Print*/
const execute = 0x2B        /* Execute*/
const print_screen = 0x2C   /* Print Screen键(抓屏)*/
const insert = 0x2D         // Insert
const delete = 0x2E         // Delete
const help = 0x2F           /* Help*/

const left_0 = 0x30 // '0'
const left_1 = 0x31 // '1'
const left_2 = 0x32 // '2'
const left_3 = 0x33 // '3'
const left_4 = 0x34 // '4'
const left_5 = 0x35 // '5'
const left_6 = 0x36 // '6'
const left_7 = 0x37 // '7'
const left_8 = 0x38 // '8'
const left_9 = 0x39 // '9'

const A = 0x41 // 'A'
const B = 0x42 // 'B'
const C = 0x43 // 'C'
const D = 0x44 // 'D'
const E = 0x45 // 'E'
const F = 0x46 // 'F'
const G = 0x47 // 'G'
const H = 0x48 // 'H'
const I = 0x49 // 'I'
const J = 0x4A // 'J'
const K = 0x4B // 'K'
const L = 0x4C // 'L'
const M = 0x4D // 'M'
const N = 0x4E // 'N'
const O = 0x4F // 'O'
const P = 0x50 // 'P'
const Q = 0x51 // 'Q'
const R = 0x52 // 'R'
const S = 0x53 // 'S'
const T = 0x54 // 'T'
const U = 0x55 // 'U'
const V = 0x56 // 'V'
const W = 0x57 // 'W'
const X = 0x58 // 'X'
const Y = 0x59 // 'Y'
const Z = 0x5A // 'Z'

const left_win = 0x5B  // 左win键
const right_win = 0x5C // 右win键

const right_0 = 0x60  // 数字区 0
const right_1 = 0x61  // 数字区 1
const right_2 = 0x62  // 数字区 2
const right_3 = 0x63  // 数字区 3
const right_4 = 0x64  // 数字区 4
const right_5 = 0x65  // 数字区 5
const right_6 = 0x66  // 数字区 6
const right_7 = 0x67  // 数字区 7
const right_8 = 0x68  // 数字区 8
const right_9 = 0x69  // 数字区 9
const multiply = 0x6A // 数字区 *
const add = 0x6B      // 数字区 +

const sepatator = 0x6C /* Sepatator*/
const subtract = 0x6D  // 数字区 -
const decimal = 0x6E   // 数字区 小数点
const divide = 0x6F    // 数字区 /

const F1 = 0x70  // F1
const F2 = 0x71  // F2
const F3 = 0x72  // F3
const F4 = 0x73  // F4
const F5 = 0x74  // F5
const F6 = 0x75  // F6
const F7 = 0x76  // F7
const F8 = 0x77  // F8
const F9 = 0x78  // F9
const F10 = 0x79 // F10
const F11 = 0x7A // F11
const F12 = 0x7B // F12
const F13 = 0x7C // F13
const F14 = 0x7D // F14
const F15 = 0x7E // F15
const F16 = 0x7F // F16
const F17 = 0x80 // F17
const F18 = 0x81 // F18
const F19 = 0x82 // F19
const F20 = 0x83 // F20
const F21 = 0x84 // F21
const F22 = 0x85 // F22
const F23 = 0x86 // F23
const F24 = 0x87 // F24

const num_lock = 0x90    // Num Lock
const scroll_lock = 0x91 /* Scroll Lock*/
const left_shift = 0xA0  // 左shift
const right_shift = 0xA1 // 右shift
const left_ctrl = 0xA2   // 左Ctrl
const right_ctrl = 0xA3  // 右Ctrl
const left_alt = 0xA4    // 左Alt
const right_alt = 0xA5   // 右Alt
const OEM_1 = 0xBA       // ';:'
const OEM_PLUS = 0xBB    // '=+'
const OEM_COMMA = 0xBC   // ',<'
const OEM_MINUS = 0xBD   // '-_'
const OEM_PERIOD = 0xBE  // '.>'
const OEM_2 = 0xBF       // '/?'
const OEM_3 = 0xC0       // '`~'
const OEM_4 = 0xDB       // '[{'
const OEM_5 = 0xDC       // '\|'
const OEM_6 = 0xDD       // ']}'
const OEM_7 = 0xDE       // ''"'

```