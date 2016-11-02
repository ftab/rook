package api

func (h *Handler) GetRoutes() []Route {
	return []Route{
		{
			"GetStatusDetails",
			"GET",
			"/status",
			h.GetStatusDetails,
		},
		{
			"GetNodes",
			"GET",
			"/node",
			h.GetNodes,
		},
		{
			"GetPools",
			"GET",
			"/pool",
			h.GetPools,
		},
		{
			"CreatePool",
			"POST",
			"/pool",
			h.CreatePool,
		},
		{
			"GetImages",
			"GET",
			"/image",
			h.GetImages,
		},
		{
			"CreateImage",
			"POST",
			"/image",
			h.CreateImage,
		},
		{
			"GetImageMapInfo",
			"GET",
			"/image/mapinfo",
			h.GetImageMapInfo,
		},
		{
			"GetMonitors",
			"GET",
			"/mon",
			h.GetMonitors,
		},
		{
			"GetCrushMap",
			"GET",
			"/crushmap",
			h.GetCrushMap,
		},
		{
			"AddDevice",
			"POST",
			"/device",
			h.AddDevice,
		},
		{
			"RemoveDevice",
			"POST",
			"/device/remove",
			h.RemoveDevice,
		},
	}
}