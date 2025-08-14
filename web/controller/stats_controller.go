
package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/alireza0/x-ui/web/service"
    "github.com/alireza0/x-ui/database/model"
)

type StatsController struct {
    userService    service.UserService
    settingService service.SettingService
}

func NewStatsController(userService service.UserService, settingService service.SettingService) *StatsController {
    return &StatsController{
        userService:    userService,
        settingService: settingService,
    }
}

func (c *StatsController) GetUserStats(ctx *gin.Context) {
    token := ctx.Param("token")
    
    // بررسی فعال بودن قابلیت
    if !c.settingService.GetSettings().EnableUserTracking {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "این قابلیت غیرفعال است"})
        return
    }

    // یافتن کاربر بر اساس توکن
    user, err := c.userService.GetUserByToken(token)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
        return
    }

    // برگرداندن اطلاعات کاربر
    ctx.JSON(http.StatusOK, gin.H{
        "upload":   user.Up,
        "download": user.Down,
        "total":    user.Total,
        "expiry":   user.Expiry,
    })
}
