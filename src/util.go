package main

import "golang.org/x/text/encoding/simplifiedchinese"

/**

 */
func UTF8ToGBK(text string) (string,error){
	dst := make([]byte,len(text)*2)
	tr :=simplifiedchinese.GB18030.NewDecoder()
	nDst ,_,err:=tr.Transform(dst,[]byte(text),true)
	if err != nil{
		return text,err
	}
	return string(dst[:nDst]),nil
}
