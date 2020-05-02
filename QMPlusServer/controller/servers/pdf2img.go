
//package main
package servers

//pdf2img m2.pdf aabbcc tmp2
//unidoc/pdf/model/writer.go +371  s := "Unlicensed UniDoc - Get a license on https://unidoc.io"
//https://github.com/unidoc/unidoc

//2. go get gopkg.in/gographics/imagick.v2/imagick
// sudo apt install libmagic-dev libmagickwand-dev
import (
	"fmt"
	"os/exec"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	//"math"
	"os"

	pdf "github.com/unidoc/unidoc/pdf/model"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func SplitPdf(inputPath,tmppdfPath, outputPath string) ([]string,error) {
	//create result path
	createPathCmd := "mkdir -p " + outputPath
	exec_shell(createPathCmd)

	var jpgList []string
	f, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("source file :%s not exist err:%v\n", inputPath,err)
		return nil,err
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		fmt.Printf("pdf split reader err:%v\n", err)
		return nil,err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		fmt.Printf("pdf check is encrypted,please check err:%v\n", err)
		return nil,err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return nil,err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		fmt.Printf("pdf get pages err:%v\n", err)
		return nil,err
	}

	filnameList := []string{}
	for i := 0; i < numPages; i++ {
		name := fmt.Sprintf("%s_%d.pdf",tmppdfPath, i)
		filnameList = append(filnameList, name)
	}

	for item, name := range filnameList {
		pdfWriter := pdf.NewPdfWriter()
		page, err := pdfReader.GetPage(item + 1)
		if err != nil {
			return nil,err
		}
		err = pdfWriter.AddPage(page)
		if err != nil {
			return nil,err
		}

		fWrite, err := os.Create(name)
		if err != nil {
			return nil,err
		}
		defer fWrite.Close()

		err = pdfWriter.Write(fWrite)
		if err != nil {
			return nil,err
		}

		//pdf Convert jpg
		picName := outputPath + "//"
		picName += name[0 : len(name)-3]
		picName += "jpg"
		jpgList = append(jpgList,picName)
		err = ConvertPdfToJpg(name, picName)
		if err != nil {
			fmt.Printf("convertPdfToJpg err:%v\n", err)
			return nil,err
		}
		//del pdf file
		fmt.Printf("del pdf path:%v\n", name)
		delCmd := "rm -f " + name
		exec_shell(delCmd)
	}
	return jpgList,nil
}

func ConvertPdfToJpg(pdfName string, imageName string) error {
	// Setup
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Must be *before* ReadImageFile
	// Make sure our image is high quality
	if err := mw.SetResolution(300, 300); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	// Must be *after* ReadImageFile
	// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
		return err
	}

	// Set any compression (100 = max quality)
	if err := mw.SetCompressionQuality(95); err != nil {
		return err
	}

	// Select only first page of pdf
	mw.SetIteratorIndex(0)

	// Convert into JPG
	if err := mw.SetFormat("jpg"); err != nil {
		return err
	}

	// Save File
	return mw.WriteImage(imageName)
}
/*
const MaxWidth float64 = 600

func fixSize(img1W, img2W int) (new1W, new2W int) {
	var ( //为了方便计算，将两个图片的宽转为 float64
		img1Width, img2Width = float64(img1W), float64(img2W)
		ratio1, ratio2       float64
	)

	minWidth := math.Min(img1Width, img2Width) // 取出两张图片中宽度最小的为基准

	if minWidth > 600 { // 如果最小宽度大于600，那么两张图片都需要进行缩放
		ratio1 = MaxWidth / img1Width // 图片1的缩放比例
		ratio2 = MaxWidth / img2Width // 图片2的缩放比例

		// 原宽度 * 比例 = 新宽度
		return int(img1Width * ratio1), int(img2Width * ratio2)
	}

	// 如果最小宽度小于600，那么需要将较大的图片缩放，使得两张图片的宽度一致
	if minWidth == img1Width {
		ratio2 = minWidth / img2Width // 图片2的缩放比例
		return img1W, int(img2Width * ratio2)
	}

	ratio1 = minWidth / img1Width // 图片1的缩放比例
	return int(img1Width * ratio1), img2W
}
*/
func MergeImageEx(soruceImage, targeImage string) {
	file, err := os.Create("dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open(soruceImage)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	file2, err := os.Open(targeImage)
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()
	img2, _ := jpeg.Decode(file2)

	jpg := image.NewRGBA(image.Rect(200, 200, 500, 500))

	draw.Draw(jpg, jpg.Bounds(), img2, img2.Bounds().Min, draw.Over)                   //首先将一个图片信息存入jpg
	draw.Draw(jpg, jpg.Bounds(), img, img.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over) //将另外一张图片信息存入jpg

	// draw.DrawMask(jpg, jpg.Bounds(), img, img.Bounds().Min, img2, img2.Bounds().Min, draw.Src) // 利用这种方法不能够将两个图片直接合成？目前尚不知道原因。

	jpeg.Encode(file, jpg, nil)

}

func MergeImage3(soruceImage, targeImage string) {
	img_file, err := os.Open(soruceImage)
	if err != nil {
		fmt.Println("打开图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer img_file.Close()
	img, err := jpeg.Decode(img_file)
	if err != nil {
		fmt.Println("把图片解码为结构体时出错")
		fmt.Println(img)
		os.Exit(-1)
	}

	//水印,用的是我自己支付宝的二维码
	wmb_file, err := os.Open(targeImage)
	if err != nil {
		fmt.Println("打开水印图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer wmb_file.Close()
	wmb_img, err := png.Decode(wmb_file)
	if err != nil {
		fmt.Println("把水印图片解码为结构体时出错")
		fmt.Println(err)
		os.Exit(-1)
	}

	//把水印写在右下角，并向0坐标偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-500, img.Bounds().Dy()-wmb_img.Bounds().Dy()-500)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()

	fmt.Println("jpg merged success")
}

func exec_shell(scmd string){
	if scmd == "" {
		fmt.Printf("exec_shell cmd empty\n")
	}
	cmd := exec.Command("/bin/bash", "-c", scmd)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("exec_shell err:%v\n", err)
	}
	//resp := string(bytes)
	//logrus.Info(resp)
}
