package utils

import (
	"fmt"
	"github.com/Yoak3n/EdgeGPT-http/edgegpt"
)

// Created at 2023/4/10 20:57
// Created by Yoake

func FormatAnswer(answer *edgegpt.Answer) string {
	text := fmt.Sprintf("%s\n(%d/%d)", answer.Text(), answer.NumUserMessages(), answer.MaxNumUserMessages())
	return text
}

//func OutPutAnswer(answer *edgegpt.Answer) string {
//	text := answer.Text()
//	return text
//}
