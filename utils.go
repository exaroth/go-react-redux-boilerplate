package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var defaultJSONHeaders = map[string]string{
	"Content-Type": "application/json",
}

// Set response headers for the response.
func SetHeaders(w http.ResponseWriter, headers map[string]string) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	return
}

func SendJSONResponse(w http.ResponseWriter, data interface{}, headers map[string]string) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err != nil {
		return err
	}
	if headers != nil {
		SetHeaders(w, headers)
	}
	SetHeaders(w, defaultJSONHeaders)
	return nil
}

func SendJSONError(w http.ResponseWriter, err error) {
	e := SendJSONResponse(w, map[string]string{
		"error": err.Error(),
	}, nil)
	if e != nil {
		Log.Error("", e, nil)
		panic(e)
	}
	return
}

func RenderTemplate(w http.ResponseWriter, tplName string, ctx map[string]interface{}, headers map[string]string) error {
	tpl := Config.GetTemplate(tplName)
	if tpl == nil {
		return fmt.Errorf("Template %s not found", tplName)
	}
	err := tpl.Execute(w, ctx)
	if err != nil {
		return err
	}
	SetHeaders(w, headers)
	return nil
}
