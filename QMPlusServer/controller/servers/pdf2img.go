
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
	"gin-vue-admin/init/initlog"
	"github.com/nfnt/resize"
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
		DelLocalFile(name)
	}
	return jpgList,nil
}

 func DelLocalFile(filePath string){
	 // del file
	fmt.Printf("del pdf path:%v\n", filePath)
	delCmd := "rm -f " + filePath
	exec_shell(delCmd)
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

func MergeImage(soruceImage, targeImage,outputPath string) string {
	img_file, err := os.Open(soruceImage)
	if err != nil {
		log.L.Info("MergeImage open err: ",err)
		return ""
	}
	defer img_file.Close()
	img, err := jpeg.Decode(img_file)
	if err != nil {
		log.L.Info("把图片解码为结构体时出错 err: ",err)
		return ""
	}

	wmb_file, err := os.Open(targeImage)
	if err != nil {
		log.L.Info("打开水印图片出错 err: ",err)
		return ""
	}
	defer wmb_file.Close()
	wmb_img, err := png.Decode(wmb_file)
	if err != nil {
		log.L.Info("把水印图片解码为结构体时出错 err: ",err)
		return ""
	}

	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-500, img.Bounds().Dy()-wmb_img.Bounds().Dy()-500)
	b := img.Bounds()
	m := image.NewRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新图片new.jpg,并设置图片质量
	mergedFile := fmt.Sprint("%s_%s",soruceImage,targeImage)
	imgw, err := os.Create(mergedFile)
	if err != nil {
		log.L.Info("Create merged file err: ",err)
		return ""
	}
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()

	fmt.Println("jpg merged success")
	return mergedFile
}

func exec_shell(scmd string){
	if scmd == "" {
		log.L.Info("exec_shell cmd empty")
		return
	}
	cmd := exec.Command("/bin/bash", "-c", scmd)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("exec_shell err:%v\n", err)
	}
	//resp := string(bytes)
	//logrus.Info(resp)
}

func ImgShrink(sourceImg string) string {
	var shrinkImg string
    if len(sourceImg) == 0 {
		log.L.Info("ImgShrink param empty")
		return ""
	}
    file, err := os.Open(sourceImg)
    if err != nil {
		log.L.Info("ImgShrink open err:%v",err)
		return ""
	}
    file.Close()

    img, err := png.Decode(file)
    if err != nil {
        log.L.Info("ImgShrink Decode err:%v",err)
		return ""
	}
	
    m := resize.Resize(200, 0, img, resize.Lanczos3)
    out, err := os.Create(shrinkImg)
    if err != nil {
        log.L.Info("ImgShrink Create err:%v",err)
		return ""
    }
    defer out.Close()
	png.Encode(out, m)
	return shrinkImg
}