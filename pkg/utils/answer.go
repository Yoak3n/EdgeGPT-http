package utils

import (
	"edgegpt-http/edgegpt"
	"fmt"
)

// Created at 2023/4/10 20:57
// Created by Yoake

func FormatAnswer(answer *edgegpt.Answer) string {
	text := fmt.Sprintf("剩余回复数:%d/%d\n%s", answer.NumUserMessages(), answer.MaxNumUserMessages(), answer.Text())
	return text
}

func OutPutAnswer(text string) string {
	return text
}
