package comment

import (
	"fmt"
	"github.com/golang/freetype"
	imgtype "github.com/shamsher31/goimgtype"
	"golang.org/x/image/font"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Fontfile  string
	Fontsize  float64
	Fontdgi   float64
	Fontspace float64
	Fontcolor string
	Startx    int
	Starty    int
	Outputdir string
}

type Comment struct {
	filename string
	content  []string
	filetype string
	config   Config
}

var (
	FONTHINTING = font.HintingFull
	FONTGROUND  = image.White
)

func (c *Comment) AddText() error {
	//解析字体
	fontbytes, err := ioutil.ReadFile(c.config.Fontfile)
	if err != nil {
		log.Println("font read error:", err)
		return err
	}
	f, err := freetype.ParseFont(fontbytes)
	if err != nil {
		log.Println("font parse error:", err)
		return err
	}
	log.Println("font parse finish")

	//初始化一张图片,生成原图
	imgback, err := os.Open(c.filename)
	if err != nil {
		log.Println("image open error:", err)
		return err
	}
	defer imgback.Close()
	img, err := c.DecodeImage(imgback)
	if err != nil {
		log.Println("image open error:", err)
		return err
	}
	bounds := img.Bounds()
	rgba := image.NewNRGBA(bounds)
	draw.Draw(rgba, rgba.Bounds(), img, image.ZP, draw.Src)
	log.Println("background image init finish")

	//在图片上面添加文字
	//设置字体
	//设置大小
	//设置边界
	//设置背景底图
	//设置背景图
	//设置提示
	if c.config.Fontcolor != "white" {
		FONTGROUND = image.Black
	}
	context := freetype.NewContext()
	context.SetDPI(c.config.Fontdgi)
	context.SetFont(f)
	context.SetFontSize(c.config.Fontsize)
	context.SetClip(rgba.Bounds())
	context.SetDst(rgba)
	context.SetSrc(FONTGROUND)
	context.SetHinting(FONTHINTING)
	log.Println("image and font set finish")

	// 画文字
	pt := freetype.Pt(c.config.Startx, c.config.Starty+int(context.PointToFixed(c.config.Fontsize)>>6))
	for _, s := range c.content {
		_, err = context.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return err
		}
		pt.Y += context.PointToFixed(c.config.Fontsize * c.config.Fontspace)
	}
	log.Println("draw font finish")

	filename := fmt.Sprintf("%s/output.%s", c.config.Outputdir, c.filetype)
	imgout, err := os.Create(filename)
	defer imgout.Close()

	if err != nil {
		log.Println("output image create error:", err)
		return err
	}

	return c.EncodeImage(imgout, rgba)
}

func (c *Comment) GetFileClass() {

	imagetype, err := imgtype.Get(c.filename)
	if err != nil {
		c.filetype = "unkown"
		return
	}
	log.Println(imagetype)
	switch imagetype {
	case `image/jpeg`:
		c.filetype = "jpg"
		return
	case `image/png`:
		c.filetype = "png"
		return
	case `image/gif`:
		c.filetype = "gif"
		return
	default:
		c.filetype = "unkown"
	}
}

func StartAddText(filename string, content []string, config Config) error {
	log.Println("start add text")
	c := &Comment{
		filename: filename,
		content:  content,
		filetype: "unkown",
		config:   config,
	}
	c.GetFileClass()
	return c.AddText()
}
