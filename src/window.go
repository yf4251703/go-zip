package main

import (
	"github.com/lxn/walk"
	"fmt"
	"github.com/lxn/walk/declarative"
	"archive/zip"
	"os"
	"io"
)

type Window interface {
	// 展示窗体接口
	ShowWindow()
}

// 创建压缩 解压缩界面类
type ComWindow struct {
	Window
	*walk.MainWindow
}

type LabWindow struct {
	Window
}

// 创建界面类
func Show(Window_type string) {
	var Win Window
	switch Window_type {
	case "main_window":
		Win = &ComWindow{}
	case "lab_window":
		Win = &LabWindow{}
	default:
		fmt.Println("参数传递错误!!")
	}
	Win.ShowWindow()
}
// 执行成功后提示
var lable *walk.Label
// 提示信息
var text string

func (comWindow *ComWindow) ShowWindow() {
	pathWindow := new(ComWindow)
	// 选择解压文件  文本框
	var unZipEdit *walk.LineEdit
	// 解压后文件  文本框
	var saveUnZipEdit *walk.LineEdit
	// 选择解压文件  按钮
	var unZipButn *walk.PushButton
	// 选择解压文件  按钮
	var saveUnZipButn *walk.PushButton

	//开始解压按钮
	var startUnZipButn *walk.PushButton

	//开始压缩按钮
	var startZipButn *walk.PushButton
	// 选择解压文件  文本框
	var zipEdit *walk.LineEdit
	// 解压后文件  文本框
	var saveZipEdit *walk.LineEdit
	// 选择解压文件  按钮
	var zipButn *walk.PushButton
	// 选择解压文件  按钮
	var saveZipButn *walk.PushButton


	err := declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow,     // 关联主窗体
		Title:    "文件压缩",                     // 标题
		MinSize:  declarative.Size{480, 230}, //指定窗体宽度跟高度
		// 布局
		Layout: declarative.HBox{}, // 水平布局 垂直布局VBox
		Children: []declarative.Widget{
			//左边区域
			declarative.Composite{
				Layout: declarative.Grid{
					Columns: 2,
					Spacing: 10,
				}, // 左边区域分为2列布局  间距10
				Children: []declarative.Widget{
					// 文本框
					declarative.LineEdit{
						AssignTo:    &unZipEdit, // 将创建好的文本框与变量关联 可以根据该变量获得文本值
						ToolTipText: "请选择需要解压文件路径",
					},
					declarative.PushButton{ // 选择解压文件按钮
						AssignTo: &unZipButn,
						Text: "选择解压文件",
						OnClicked: func() {
							// 点击按钮  将文件选择器选择的文件路径赋值给文本框
							unZipEdit.SetText(pathWindow.OpenFileMananger())
						},
					},
					// 创建解压后 文件存放的文本框
					declarative.LineEdit{
						AssignTo:    &saveUnZipEdit,
						ToolTipText: "请输入解压后文件的路径",
					},
					declarative.PushButton{ // 选择解压文件按钮
						AssignTo: &saveUnZipButn,
						Text: "选择解压位置",
						OnClicked: func() {
							// 点击按钮  将文件选择器选择的文件路径赋值给文本框
							saveUnZipEdit.SetText(pathWindow.OpenDirManager())
						},
					},
					// 创建需要压缩的文件的文本框
					declarative.LineEdit{
						AssignTo:    &zipEdit,
						ToolTipText: "请选择需要压缩文件路径",
					},

					declarative.PushButton{ // 选择解压文件按钮
						AssignTo: &zipButn,
						Text: "选择压缩文件",
						OnClicked: func() {
							// 点击按钮  将文件选择器选择的文件路径赋值给文本框
							zipEdit.SetText(pathWindow.OpenFileMananger())
						},
					},
					// 创建解压后 文件存放的文本框
					declarative.LineEdit{
						AssignTo:    &saveZipEdit,
						ToolTipText: "请输入解压后文件的路径",
					},
					declarative.PushButton{ // 选择解压文件按钮
						AssignTo: &saveZipButn,
						Text: "选择压缩位置",
						OnClicked: func() {
							// 点击按钮  将文件选择器选择的文件路径赋值给文本框
							saveZipEdit.SetText(pathWindow.OpenDirManager())

						},
					},
					declarative.Label{
						AssignTo:&lable,
						Text:"",
					},
				},
			},
			//右边区域
			declarative.Composite{
				Layout: declarative.Grid{
					Rows:2,
					Spacing:50,
				},
				Children:[]declarative.Widget{
					declarative.PushButton{
						AssignTo:&startUnZipButn,
						Text:"开始解压",
						OnClicked: func() {
							// 解压文件 传递压缩文件路径 跟解压后文件路径
							pathWindow.StartToUnZip(unZipEdit.Text(),saveUnZipEdit.Text())
							text="文件保存成功"
							Show("lab_window")
						},
					},

					declarative.PushButton{
						AssignTo:&startZipButn,
						Text:"开始压缩",
						OnClicked: func() {
							// 压缩文件 传递压缩文件路径 跟解压后文件路径
							pathWindow.StartToZip(zipEdit.Text(),saveZipEdit.Text())
							text="文件压缩成功"
							Show("lab_window")
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		fmt.Println(err)
	}

	pathWindow.SetX(650) // x坐标
	pathWindow.SetY(350) // y 坐标
	pathWindow.Run()     // 运行窗口，才能将创建的窗体给用户展示出来
}

// 单开文件选择对话框
func (mv *ComWindow) OpenFileMananger() (fileParh string) {
	// 1 创建文件对话框对象
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "所有文档(*.*)|*.*|文本文档(*.txt)|*.txt|图片文件(*.jpg)|*.jpg"
	// 2 打开文件对话框
	_, err := dlg.ShowOpen(mv) // 如果单击对话框中的”打开“按钮 返回true 否则返回false
	if err != nil {
		fmt.Println(err)
	}
	// 3 获取选中的文件
	fileParh = dlg.FilePath
	return fileParh

}

// 打开浏览文件窗口
func (mv *ComWindow) OpenDirManager() (fileParh string) {
	// 创建对话框窗口
	dlg := new(walk.FileDialog)
	dlg.Title = "浏览文件"
	// 打开窗口
	_, err := dlg.ShowBrowseFolder(mv)
	if err != nil {
		fmt.Println(err)
	}
	// 获取选中的路径 并且返回
	fileParh = dlg.FilePath
	return fileParh
}


// 实现文件解压操作
func (mv *ComWindow) StartToUnZip(file string ,saveFile string)  {
	// 1：获取第一个文本框中的 需要解压文件的路径 并读取文件的内容
		reader,err:=zip.OpenReader(file)
		if err !=nil{
			fmt.Println(err)
		}
		defer reader.Close()
	// 2：循环遍历压缩包里面的文件
	for  _,file:=range reader.File{
		// 打开从压缩文件中的内容
		rc,err:=file.Open()
		if err != nil{
			fmt.Println(err)
		}
		defer rc.Close()
		// 获得当前文件名
		newName:=saveFile+file.Name
		newName,err =UTF8ToGBK(newName)
		if err!=nil{
			fmt.Println(err)
		}
		// 判断是否为文件夹
		if file.FileInfo().IsDir(){
			// 创建文件夹
			err:=os.MkdirAll(newName,os.ModePerm)
			if err != nil{
				fmt.Println(err)
			}
		}else{
			// 是否为文件
			f,err:=os.Create(newName)
			if err!=nil{
				fmt.Println(err)
			}
			defer f.Close()
			_,err1:=io.Copy(f,rc)
			if err1 !=nil{
				fmt.Println(err1)
			}
		}
	}
	// 3：判断是否是文件夹   如果是文件夹 就创建
	// 4: 如果是文件 则创建

}

// 实现文件压缩
func (mv *ComWindow) StartToZip(filePath string,savePath string)  {
		// 1 获取压缩文件 创建zip文件
		d,err:=os.Create(savePath)
		if err != nil{
			fmt.Println(err)
		}
		defer d.Close()
		// 2 获取压缩位置路径 打开该文件
		file,err:=os.Open(filePath)
		if err != nil{
			fmt.Println(err)
		}
		defer file.Close()
		// 3 将压缩文件写入压缩包中
			// 3.1 获取要压缩文件的信息
			info,err :=file.Stat()
			if err != nil{
				fmt.Println(err)
			}
			header,err:=zip.FileInfoHeader(info)
			if err !=nil{
				fmt.Println(err)
			}
			// 3.2 将要压缩的文件写入压缩包中
			w:=zip.NewWriter(d) // 根据创建的压缩包创建了一个writer 的指针 通过该指针可以对压缩包进行操作
			defer w.Close()
			writer,err:=w.CreateHeader(header)
			if err != nil{
				fmt.Println(err)
			}

			io.Copy(writer,file)

}

// 将提示信息打印
func (lab *LabWindow) ShowWindow(){
		lable.SetText(text)
}