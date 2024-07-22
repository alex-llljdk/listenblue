package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"zjhx.com/pkg/cons"
	"zjhx.com/record/api/internal/svc"
	"zjhx.com/record/api/internal/types"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordByIdLogic {
	return &RecordByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordByIdLogic) RecordById(req *types.RecordByIdRequest) (resp *types.RecordByIdRespone, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.RecordRpc.GetDetailRecordById(l.ctx, &record.DetailRecordReq{
		Id: uint64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	if res.DataType <= 1 {
		var keywordResponse map[string][]string
		var summaryResponse types.AnserRespone
		var scanningTmp map[string][]types.ScanningRespone
		var reviewTmp map[string][]types.ReviewResponse
		var transResponse []types.TransResponse
		var innerText types.InnerTextRespone
		//解析关键词
		fmt.Println(1)
		err = json.Unmarshal([]byte(res.Keyword), &keywordResponse)
		if err != nil {
			return nil, err
		}
		var keyMap []map[string]string
		for _, v := range keywordResponse["answer"] {
			keyMap = append(keyMap, map[string]string{
				"key": v,
			})
		}
		fmt.Println(2)
		//解析全文概括
		err = json.Unmarshal([]byte(res.Summary), &summaryResponse)
		if err != nil {
			return nil, err
		}
		//章节速览
		err = json.Unmarshal([]byte(res.Scanning), &scanningTmp)
		if err != nil {
			return nil, err
		}
		fmt.Println(scanningTmp)
		fmt.Println(3)
		var scanningMap []map[string]string
		for _, v := range scanningTmp["answer"] {
			scanningMap = append(scanningMap, map[string]string{
				"title": v.Title,
				"desc":  v.Desc,
			})
		}
		err = json.Unmarshal([]byte(res.Review), &reviewTmp)
		if err != nil {
			fmt.Println("err--------------", err)
			return nil, err
		}
		fmt.Println(4)
		var reviewMap []map[string]string
		for _, v := range reviewTmp["answer"] {
			reviewMap = append(reviewMap, map[string]string{
				"ques": v.Question,
				"ans":  v.Answer,
			})
		}
		err = json.Unmarshal([]byte(res.TransText), &transResponse)
		if err != nil {
			return nil, err
		}
		fmt.Println(5)
		var transMap []map[string]string
		for _, v := range transResponse {
			transMap = append(transMap, map[string]string{
				"timestamp": v.TimeStamp,
				"content":   v.Content,
			})
		}

		err = json.Unmarshal([]byte(res.InnerText), &innerText)
		if err != nil {
			return nil, err
		}
		fmt.Println(6)
		innerTextMap := map[string]string{
			"recordInfo":      innerText.RecordInfo,
			"source_language": innerText.SourceLanguage,
			"dest_language":   innerText.DestLanguage,
			"content":         innerText.Content,
			"savedTime":       innerText.SaveTime,
		}

		return &types.RecordByIdRespone{
			Code:      cons.ServiceOKCode,
			Msg:       cons.ServiceOKMsg,
			Id:        res.Id,
			Path:      res.Path,
			Title:     res.Title,
			UserId:    res.UserId,
			DataType:  res.DataType,
			Keyword:   keyMap,
			Summary:   summaryResponse.Answer,
			Scanning:  scanningMap,
			InnerText: innerTextMap,
			Review:    reviewMap,
			TransText: transMap,
		}, nil

	} else { //pdf文档解析
		//对得到的结果进行解析
		//对得到的结果进行解析
		type titleResp struct {
			Answer string `json:"answer"`
		}
		var tResp titleResp
		err = json.Unmarshal([]byte(res.Title), &tResp)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		type summaryResp struct {
			Answer string `json:"answer"`
		}
		var summResp summaryResp
		err = json.Unmarshal([]byte(res.Summary), &summResp)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		var tmp_keyword []map[string]string
		var tmp_scanning []map[string]string

		// tmp_scanning 文档速度解析
		type Page1 struct {
			Index string `json:"index"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
		}
		type Answer1 struct {
			Pages []Page1 `json:"answer"`
		}
		var answer1 Answer1
		err := json.Unmarshal([]byte(res.Scanning), &answer1)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		for _, page := range answer1.Pages {
			tmp_scanning = append(tmp_scanning, map[string]string{
				"index": page.Index,
				"title": page.Title,
				"desc":  page.Desc,
			})
		}

		// keyword 关键要点解析
		type Page2 struct {
			Index string `json:"index"`
			Desc  string `json:"desc"`
		}
		type Answer2 struct {
			Pages []Page2 `json:"answer"`
		}
		var answer2 Answer2
		err = json.Unmarshal([]byte(res.Keyword), &answer2)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		for _, page := range answer2.Pages {
			tmp_keyword = append(tmp_keyword, map[string]string{
				"index": page.Index,
				"desc":  page.Desc,
			})
		}

		return &types.RecordByIdRespone{
			Code:      cons.ServiceOKCode,
			Msg:       cons.ServiceOKMsg,
			Id:        res.Id,
			Path:      res.Path,
			Title:     tResp.Answer,
			UserId:    res.UserId,
			DataType:  res.DataType,
			Keyword:   tmp_keyword,
			Summary:   summResp.Answer,
			Scanning:  tmp_scanning,
			InnerText: nil,
			Review:    nil,
			TransText: nil,
		}, nil
	}
}
