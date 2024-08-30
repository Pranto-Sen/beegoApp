package routers

import (
	"beegoApp/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/voting", &controllers.MainController{})
	beego.Router("/breed", &controllers.MainController{})
	beego.Router("/favourite", &controllers.MainController{})

	beego.Router("/api/cat", &controllers.VotingController{}, "get:GetCatImages")
	beego.Router("/api/breeds", &controllers.BreedController{}, "get:GetBreeds")
	beego.Router("/api/favourites", &controllers.FavouriteController{}, "post:AddFavourite;get:GetFavourites")
	beego.Router("/api/favourites/:id", &controllers.FavouriteController{}, "delete:DeleteFavourite")
}
