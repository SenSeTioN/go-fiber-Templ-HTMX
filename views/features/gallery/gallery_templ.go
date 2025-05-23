// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package gallery

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "sensetion/go-fiber/views/entities/card"

func Gallery() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = GalleryCSS().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"swiper mySwiper\"><div class=\"swiper-wrapper\"><div class=\"swiper-slide\"><img src=\"/public/images/bgs/unsplash_2ZOHXRaZfFQ.png\" loading=\"lazy\"><div class=\"swiper-lazy-preloader\"></div></div><div class=\"swiper-slide\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = card.CardXL(card.CardXLProps{
			Text:    "123",
			SubText: "sdadsa",
			BGSrc:   "/public/images/bgs/unsplash_2ZOHXRaZfFQ.png",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div><div class=\"swiper-slide\">Slide 3</div><div class=\"swiper-slide\">Slide 4</div><div class=\"swiper-slide\">Slide 5</div><div class=\"swiper-slide\">Slide 6</div><div class=\"swiper-slide\">Slide 7</div><div class=\"swiper-slide\">Slide 8</div><div class=\"swiper-slide\">Slide 9</div></div><div class=\"swiper-button-next\"></div><div class=\"swiper-button-prev\"></div><div class=\"swiper-pagination\"></div><div class=\"autoplay-progress\"><svg viewBox=\"0 0 48 48\"><circle cx=\"24\" cy=\"24\" r=\"20\"></circle></svg> <span></span></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func GalleryCSS() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<style>\r\n\t\t.swiper {\r\n\t\t\tmax-width: 750px;\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 455px;\r\n\t\t\tborder-radius: var(--border-radius-m);\r\n\t\t\toverflow: hidden;\r\n\t\t}\r\n\t\t.autoplay-progress {\r\n      position: absolute;\r\n      left: 12px;\r\n      bottom: 30px;\r\n      z-index: 10;\r\n      width: 45px;\r\n      height: 45px;\r\n      display: flex;\r\n      align-items: center;\r\n      justify-content: center;\r\n      font-weight: bold;\r\n      color: var(--color-orange);\r\n    }\r\n\r\n    .autoplay-progress svg {\r\n      --progress: 0;\r\n      position: absolute;\r\n      left: 0;\r\n      top: 0px;\r\n      z-index: 10;\r\n      width: 100%;\r\n      height: 100%;\r\n      stroke-width: 4px;\r\n      stroke: var(--color-orange);\r\n      fill: none;\r\n      stroke-dashoffset: calc(125.6px * (1 - var(--progress)));\r\n      stroke-dasharray: 125.6;\r\n      transform: rotate(-90deg);\r\n    }\r\n\r\n\t\t.swiper-slide {\r\n\t\t\tbackground-color: var(--color-light-grey-pr);\r\n\t\t}\r\n\r\n\t\t.swiper-pagination{\r\n\t\t\tpadding: 0 20px 6px 0 ;\r\n\t\t}\r\n\r\n\t\t.swiper-horizontal>.swiper-pagination-bullets.swiper-pagination-bullets-dynamic, .swiper-pagination-horizontal.swiper-pagination-bullets.swiper-pagination-bullets-dynamic {\r\n\t\t\tleft: 80%;\r\n\t\t\ttransform: translateX(-12%);\r\n\t\t}\r\n\r\n\t\t.swiper-pagination-bullet-active-prev {\r\n\t\t\topacity: 0.75;\r\n\t\t}\r\n\r\n\t\t.swiper-pagination-bullet-active-next {\r\n\t\t\topacity: 0.75;\r\n\t\t}\r\n\r\n\t\t.swiper-pagination-bullet{\r\n\t\t\tbackground-color: var(--color-white);\r\n\t\t\twidth: 10px;\r\n\t\t\theight: 10px;\r\n\t\t\topacity: 0.75;\r\n\t\t}\r\n\r\n\t\t.swiper-pagination-bullet-active{\r\n\t\t\tbackground-color: var(--color-white);\r\n\t\t\tborder-radius: var(--border-radius-s);\r\n\t\t\twidth: 27px;\r\n\t\t\theight: 10px;\r\n\t\t\topacity: 1;\r\n\t\t}\r\n\r\n\t\t.swiper-slide img {\r\n\t\t\tobject-fit: cover;\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 100%;\r\n\t\t}\r\n\r\n\t\t.swiper-button-next{\r\n\t\t\tcolor: var(--color-light-grey);\r\n\t\t\tmax-width: 40px;\r\n\t\t\tmax-height: 40px;\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 100%;\r\n\t\t\tborder-radius: var(--border-radius-m);\r\n\t\t\tbackground-color: var(--bg-light-grey);\r\n\t\t}\r\n\r\n\t\t.swiper-button-next::after{\r\n\t\t\tfont-size: 16px;\r\n\t\t\tfont-weight: 900;\r\n\t\t}\r\n\r\n\t\t.swiper-button-prev{\r\n\t\t\tcolor: var(--color-light-grey);\r\n\t\t\tmax-width: 40px;\r\n\t\t\tmax-height: 40px;\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 100%;\r\n\t\t\tborder-radius: var(--border-radius-m);\r\n\t\t\tbackground-color: var(--bg-light-grey);\r\n\t\t}\r\n\r\n\t\t.swiper-button-prev::after{\r\n\t\t\tfont-size: 18px;\r\n\t\t\tfont-weight: 900;\r\n\t\t}\r\n\r\n\t</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
