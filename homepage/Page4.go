package homepage

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func Page4(w fyne.Window) fyne.CanvasObject {
	return NewJoystick(150, 20, color.RGBA{R: 255, A: 255}, func(x, y float32) {
		// 这里可以处理摇杆的回调
		println("Joystick position:", x, y)
	})
}

// Joystick 结构体定义了摇杆控件的基本属性
type Joystick struct {
	widget.BaseWidget // 继承自 Fyne 的 BaseWidget

	centerX, centerY float32            // 摇杆中心点的坐标
	radius           float32            // 摇杆的半径
	knobRadius       float32            // 摇杆柄的半径
	knobX, knobY     float32            // 摇杆柄的当前坐标
	knobColor        color.Color        // 摇杆柄的颜色
	callback         func(x, y float32) // 回调函数，用于处理摇杆的位置变化
}

// NewJoystick 创建一个新的摇杆控件
func NewJoystick(radius float32, knobRadius float32, knobColor color.Color, callback func(x, y float32)) *Joystick {
	j := &Joystick{
		radius:     radius,     // 设置摇杆的半径
		knobRadius: knobRadius, // 设置摇杆柄的半径
		knobColor:  knobColor,  // 设置摇杆柄的颜色
		callback:   callback,   // 设置回调函数
	}
	j.ExtendBaseWidget(j) // 扩展 BaseWidget
	return j
}

// CreateRenderer 创建摇杆控件的渲染器
func (j *Joystick) CreateRenderer() fyne.WidgetRenderer {
	background := canvas.NewCircle(color.Gray{128}) // 创建背景圆
	knob := canvas.NewCircle(j.knobColor)           // 创建摇杆柄
	return &joystickRenderer{
		joystick:   j,          // 摇杆控件
		background: background, // 背景圆
		knob:       knob,       // 摇杆柄
	}
}

// joystickRenderer 结构体定义了摇杆控件的渲染器
type joystickRenderer struct {
	joystick   *Joystick      // 摇杆控件
	background *canvas.Circle // 背景圆
	knob       *canvas.Circle // 摇杆柄
}

// MinSize 返回摇杆控件的最小尺寸
func (r *joystickRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.joystick.radius*2, r.joystick.radius*2)
}

// Layout 布局摇杆控件的子控件
func (r *joystickRenderer) Layout(size fyne.Size) {
	r.joystick.centerX = size.Width / 2                                                                      // 设置摇杆中心点的 X 坐标
	r.joystick.centerY = size.Height / 2                                                                     // 设置摇杆中心点的 Y 坐标
	r.background.Resize(size)                                                                                // 调整背景圆的大小
	r.knob.Resize(fyne.NewSize(r.joystick.knobRadius*2, r.joystick.knobRadius*2))                            // 调整摇杆柄的大小
	r.knob.Move(fyne.NewPos(r.joystick.knobX-r.joystick.knobRadius, r.joystick.knobY-r.joystick.knobRadius)) // 移动摇杆柄
}

// Refresh 刷新摇杆控件的显示
func (r *joystickRenderer) Refresh() {
	r.knob.Move(fyne.NewPos(r.joystick.knobX-r.joystick.knobRadius, r.joystick.knobY-r.joystick.knobRadius)) // 移动摇杆柄
	canvas.Refresh(r.joystick)                                                                               // 刷新摇杆控件
}

// Objects 返回摇杆控件的子控件
func (r *joystickRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.knob}
}

// Destroy 销毁摇杆控件的渲染器
func (r *joystickRenderer) Destroy() {}

// Dragged 处理拖动事件
func (j *Joystick) Dragged(e *fyne.DragEvent) {
	x := e.Position.X - j.centerX             // 计算摇杆柄相对于中心点的 X 坐标
	y := e.Position.Y - j.centerY             // 计算摇杆柄相对于中心点的 Y 坐标
	distance := math.Sqrt(float64(x*x + y*y)) // 计算摇杆柄到中心点的距离

	// 如果距离超过摇杆的半径，则限制摇杆柄的位置
	if distance > float64(j.radius) {
		ratio := float64(j.radius) / distance
		x *= float32(ratio)
		y *= float32(ratio)
	}

	j.knobX = j.centerX + x // 更新摇杆柄的 X 坐标
	j.knobY = j.centerY + y // 更新摇杆柄的 Y 坐标

	// 调用回调函数，传递摇杆柄的相对位置
	if j.callback != nil {
		j.callback(x/j.radius, y/j.radius)
	}

	j.Refresh() // 刷新摇杆控件
}

// DragEnd 处理拖动结束事件
func (j *Joystick) DragEnd() {
	j.knobX = j.centerX // 将摇杆柄重置到中心点的 X 坐标
	j.knobY = j.centerY // 将摇杆柄重置到中心点的 Y 坐标
	j.Refresh()         // 刷新摇杆控件
}
