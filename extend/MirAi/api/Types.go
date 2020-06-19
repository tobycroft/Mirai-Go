package api

func At(target interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type":   "At",
		"target": target,
	}
}

func AtAll() map[string]interface{} {
	return map[string]interface{}{
		"type": "AtAll",
	}
}

func Face(faceId interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type":   "Face",
		"faceId": faceId,
	}
}

func Plain(text interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type": "Plain",
		"text": text,
	}
}

func Image(url interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type": "Image",
		"url":  url,
	}
}

func Xml(xml interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type": "Xml",
		"xml":  xml,
	}
}

func Json(json interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type": "Json",
		"json": json,
	}
}

func App() map[string]interface{} {
	return map[string]interface{}{
		"type":    "App",
		"content": "<>",
	}
}

func Poke() map[string]interface{} {
	return map[string]interface{}{
		"type": "Poke",
		"name": "Poke",
	}
}
