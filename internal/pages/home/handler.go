package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

type MenuTab struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}

	h.registerRoutes()
}

func (h *HomeHandler) registerRoutes() {
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	menuTabs := []MenuTab{
		{Id: 1, Name: "Еда"},
		{Id: 2, Name: "Животные"},
		{Id: 3, Name: "Машины"},
		{Id: 4, Name: "Спорт"},
		{Id: 5, Name: "Музыка"},
		{Id: 6, Name: "Технологии"},
		{Id: 7, Name: "Прочее"},
	}

	return c.Render("home", fiber.Map{"MenuTabs": menuTabs})
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}
