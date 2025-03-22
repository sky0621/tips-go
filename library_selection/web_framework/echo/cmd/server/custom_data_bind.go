package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type CustomMsg string

func (cm *CustomMsg) UnmarshalJSON(b []byte) error {
	*cm = CustomMsg(fmt.Sprintf("『%s』", string(b)))
	return nil
}

func (cm *CustomMsg) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *cm)), nil
}

type CustomDate time.Time

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}

func (cd *CustomDate) MarshalJSON() ([]byte, error) {
	t := time.Time(*cd)
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

type CustomData struct {
	Msg  CustomMsg  `json:"msg"`
	Date CustomDate `json:"date,omitempty"`
}

func customDataBindSample(c echo.Context) error {
	var cd CustomData
	if err := c.Bind(&cd); err != nil {
		return err
	}
	jMsg, err := cd.Msg.MarshalJSON()
	if err != nil {
		return err
	}
	jDate, err := cd.Date.MarshalJSON()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, string(jMsg)+","+string(jDate))
}
