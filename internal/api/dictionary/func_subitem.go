package dictionary

import (
	"campus-recruiting-api/configs"
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/dictionary"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type subItemRequest struct {
	Code string `json:"code" uri:"code" binding:"required"` // 字典标识
}

type subItemResponse struct {
	List []dictionary.SubItemData `json:"list"`
}

// SubItem 根据字典标识获取字典子项
// @Summary 根据字典标识获取字典子项
// @Description 根据字典标识获取字典子项
// @Tags API.dictionary
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body subItemRequest true "请求信息"
// @Success 200 {object} subItemResponse
// @Failure 400 {object} code.Failure
// @Router /api/dictionary/subItem/{code} [get]
func (h *handler) SubItem() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(subItemRequest)
		res := new(subItemResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		// 从Redis获取缓存数据
		if h.cache.Exists(req.Code) {
			cacheData, err := h.cache.Get(req.Code)
			if err != nil {
				h.logger.Error("Redis：获取字典项失败", zap.Error(err))
			} else {
				var data []dictionary.SubItemData
				err = json.Unmarshal([]byte(cacheData), &data)
				if err != nil {
					h.logger.Error("Redis：反序列化数据失败", zap.Error(err))
				} else {
					res.List = data
					ctx.Payload(res)
					return
				}
			}
		}

		// 从数据库获取数据
		data, err := h.dictionaryService.SubItem(ctx, req.Code)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.DictionarySubItemError,
				code.Text(code.DictionarySubItemError)).WithError(err),
			)
			return
		}

		// 缓存至Redis
		jsonData, err := json.Marshal(data)
		if err != nil {
			h.logger.Error("Redis：序列化数据失败", zap.Error(err))
		} else {
			err = h.cache.Set(req.Code, string(jsonData), configs.NotExpired)
			if err != nil {
				h.logger.Error("Redis：设置字典项失败", zap.Error(err))
			}
		}

		res.List = data
		ctx.Payload(res)
	}
}
